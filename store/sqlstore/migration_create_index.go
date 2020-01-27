// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package sqlstore

import (
	"context"
	"database/sql"
	"errors"
	"strings"

	"github.com/mattermost/mattermost-server/v5/model"
)

// CreateIndex - asynchronous migration that adds an index to table
type CreateIndex struct {
	name      string
	table     string
	columns   []string
	indexType string
	unique    bool
}

// NewCreateIndex creates a migration that adds an index
func NewCreateIndex(indexName string, tableName string, columnNames []string, indexType string, unique bool) *CreateIndex {
	return &CreateIndex{
		name:      indexName,
		table:     tableName,
		columns:   columnNames,
		indexType: indexType,
		unique:    unique,
	}
}

// Name returns name of the migration, should be unique
func (m *CreateIndex) Name() string {
	return "add_index_" + m.name
}

type pgIndexData struct {
	Name    string `db:"relname"`
	IsValid bool   `db:"indisvalid"`
}

func checkPostgreSQLIndex(ss SqlStore, name string) (*pgIndexData, error) {
	idxData := pgIndexData{}
	err := ss.GetMaster().SelectOne(&idxData, "SELECT relname, indisvalid FROM pg_class, pg_index WHERE pg_index.indexrelid = pg_class.oid AND pg_class.relname = $1", name)
	// ErrNoRows means there is no index
	if err == sql.ErrNoRows {
		return nil, nil
	}
	// other error means something went wrong
	if err != nil {
		return nil, err
	}
	return &idxData, nil
}

// GetStatus returns if the migration should be executed or not
func (m *CreateIndex) GetStatus(ss SqlStore) (asyncMigrationStatus, error) {
	if ss.DriverName() == model.DATABASE_DRIVER_POSTGRES {
		if m.indexType == INDEX_TYPE_FULL_TEXT && len(m.columns) != 1 {
			return failed, errors.New("Unable to create multi column full text index")
		}
		idxData, err := checkPostgreSQLIndex(ss, m.name)
		// other error means something went wrong
		if err != nil {
			return unknown, err
		}
		// check if the index is invalid, this can happen if create index concurrently was aborted
		// in that case we have to drop this index and create it again
		// this may block if there is a long running query on the table
		if idxData != nil && idxData.IsValid != true {
			_, err = ss.GetMaster().ExecNoTimeout("DROP INDEX " + m.name)
			if err != nil {
				return unknown, err
			}
		}
		return run, nil
	} else if ss.DriverName() == model.DATABASE_DRIVER_MYSQL {
		if m.indexType == INDEX_TYPE_FULL_TEXT {
			return failed, errors.New("Unable to create full text index concurrently")
		}
		count, err := ss.GetMaster().SelectInt("SELECT COUNT(0) AS index_exists FROM information_schema.statistics WHERE TABLE_SCHEMA = DATABASE() and table_name = ? AND index_name = ?", m.table, m.name)
		if err != nil {
			return unknown, err
		}
		if count > 0 {
			return skip, nil
		}
	}
	return run, nil
}

// Execute runs the migration
// Explicit connection is passed so that all queries run in a single session
func (m *CreateIndex) Execute(ctx context.Context, ss SqlStore, conn *sql.Conn) (asyncMigrationStatus, error) {
	uniqueStr := ""
	if m.unique {
		uniqueStr = "UNIQUE "
	}

	if ss.DriverName() == model.DATABASE_DRIVER_POSTGRES {
		// because of retries we check for invalid index here too
		// but I'm not sure if it's really necessary
		idxData, err := checkPostgreSQLIndex(ss, m.name)
		if err != nil {
			return unknown, err
		}
		// check if the index is invalid
		// in that case we have to drop this index and create it again
		if idxData != nil && idxData.IsValid != true {
			_, err = conn.ExecContext(ctx, "DROP INDEX "+m.name)
			if err != nil {
				return unknown, err
			}
		}
		query := ""
		if m.indexType == INDEX_TYPE_FULL_TEXT {
			columnName := m.columns[0]
			postgresColumnNames := convertMySQLFullTextColumnsToPostgres(columnName)
			query = "CREATE INDEX CONCURRENTLY " + m.name + " ON " + m.table + " USING gin(to_tsvector('english', " + postgresColumnNames + "))"
		} else {
			query = "CREATE " + uniqueStr + "INDEX CONCURRENTLY " + m.name + " ON " + m.table + " (" + strings.Join(m.columns, ", ") + ")"
		}

		_, err = conn.ExecContext(ctx, query)
		if err != nil {
			return failed, err
		}
	} else if ss.DriverName() == model.DATABASE_DRIVER_MYSQL {
		_, err := conn.ExecContext(ctx, "CREATE  "+uniqueStr+" INDEX "+m.name+" ON "+m.table+" ("+strings.Join(m.columns, ", ")+")")
		if err != nil {
			return failed, err
		}
	}
	return complete, nil
}
