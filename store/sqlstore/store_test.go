// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package sqlstore

import (
	"fmt"
	"os"
	"sync"
	"testing"

	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/mattermost/mattermost-server/v5/store"
	"github.com/mattermost/mattermost-server/v5/store/searchtest"
	"github.com/mattermost/mattermost-server/v5/store/storetest"
)

type storeType struct {
	Name        string
	SqlSettings *model.SqlSettings
	SqlSupplier *SqlSupplier
	Store       store.Store
}

var storeTypes []*storeType
var storeTypesBench []*storeType

func StoreTest(t *testing.T, f func(*testing.T, store.Store)) {
	defer func() {
		if err := recover(); err != nil {
			tearDownStores()
			panic(err)
		}
	}()
	for _, st := range storeTypes {
		st := st
		t.Run(st.Name, func(t *testing.T) {
			if testing.Short() {
				t.SkipNow()
			}
			f(t, st.Store)
		})
	}
}

func StoreTestWithSearchTestEngine(t *testing.T, f func(*testing.T, store.Store, *searchtest.SearchTestEngine)) {
	defer func() {
		if err := recover(); err != nil {
			tearDownStores()
			panic(err)
		}
	}()

	for _, st := range storeTypes {
		st := st
		searchTestEngine := &searchtest.SearchTestEngine{
			Driver: *st.SqlSettings.DriverName,
		}

		t.Run(st.Name, func(t *testing.T) { f(t, st.Store, searchTestEngine) })
	}
}

func StoreTestWithSqlSupplier(t *testing.T, f func(*testing.T, store.Store, storetest.SqlSupplier)) {
	defer func() {
		if err := recover(); err != nil {
			tearDownStores()
			panic(err)
		}
	}()
	for _, st := range storeTypes {
		st := st
		t.Run(st.Name, func(t *testing.T) {
			if testing.Short() {
				t.SkipNow()
			}
			f(t, st.Store, st.SqlSupplier)
		})
	}
}

func StoreBenchWithSqlSupplier(b *testing.B, f func(*testing.B, store.Store)) {
	for _, st := range storeTypesBench {
		st := st
		fmt.Println("Running benchmark on: ", *st.SqlSettings.DataSource)
		b.Run(st.Name, func(b *testing.B) {
			f(b, st.Store)
		})
	}
}

func initStoresForBench() {
	driver := model.DATABASE_DRIVER_MYSQL
	dataSource := "mmuser:mostest@tcp(localhost:3306)/mattermost_test?charset=utf8mb4,utf8&readTimeout=30s&writeTimeout=30s" // XXX Point me to the right DB
	settings := &model.SqlSettings{
		DriverName:                  &driver,
		DataSource:                  &dataSource,
		DataSourceReplicas:          []string{},
		DataSourceSearchReplicas:    []string{},
		MaxIdleConns:                new(int),
		ConnMaxLifetimeMilliseconds: new(int),
		MaxOpenConns:                new(int),
		Trace:                       model.NewBool(false),
		AtRestEncryptKey:            model.NewString(model.NewRandomString(32)),
		QueryTimeout:                new(int),
	}
	*settings.MaxIdleConns = 10
	*settings.ConnMaxLifetimeMilliseconds = 3600000
	*settings.MaxOpenConns = 100
	*settings.QueryTimeout = 60

	storeTypesBench = append(storeTypesBench, &storeType{
		Name:        "MySQL",
		SqlSettings: settings,
	})

	for _, st := range storeTypesBench {
		st := st
		st.SqlSupplier = NewSqlSupplier(*st.SqlSettings, nil)
		st.Store = st.SqlSupplier
	}
}

func initStores() {
	if testing.Short() {
		return
	}
	// In CI, we already run the entire test suite for both mysql and postgres in parallel.
	// So we just run the tests for the current database set.
	if os.Getenv("IS_CI") == "true" {
		switch os.Getenv("MM_SQLSETTINGS_DRIVERNAME") {
		case "mysql":
			storeTypes = append(storeTypes, &storeType{
				Name:        "MySQL",
				SqlSettings: storetest.MakeSqlSettings(model.DATABASE_DRIVER_MYSQL),
			})
		case "postgres":
			storeTypes = append(storeTypes, &storeType{
				Name:        "PostgreSQL",
				SqlSettings: storetest.MakeSqlSettings(model.DATABASE_DRIVER_POSTGRES),
			})
		}
	} else {
		storeTypes = append(storeTypes, &storeType{
			Name:        "MySQL",
			SqlSettings: storetest.MakeSqlSettings(model.DATABASE_DRIVER_MYSQL),
		})
		storeTypes = append(storeTypes, &storeType{
			Name:        "PostgreSQL",
			SqlSettings: storetest.MakeSqlSettings(model.DATABASE_DRIVER_POSTGRES),
		})
	}

	defer func() {
		if err := recover(); err != nil {
			tearDownStores()
			panic(err)
		}
	}()
	var wg sync.WaitGroup
	for _, st := range storeTypes {
		st := st
		wg.Add(1)
		go func() {
			defer wg.Done()
			st.SqlSupplier = NewSqlSupplier(*st.SqlSettings, nil)
			st.Store = st.SqlSupplier
			st.Store.DropAllTables()
			st.Store.MarkSystemRanUnitTests()
		}()
	}
	wg.Wait()
}

var tearDownStoresOnce sync.Once

func tearDownStores() {
	if testing.Short() {
		return
	}
	tearDownStoresOnce.Do(func() {
		var wg sync.WaitGroup
		wg.Add(len(storeTypes))
		for _, st := range storeTypes {
			st := st
			go func() {
				if st.Store != nil {
					st.Store.Close()
				}
				if st.SqlSettings != nil {
					storetest.CleanupSqlSettings(st.SqlSettings)
				}
				wg.Done()
			}()
		}
		wg.Wait()
	})
}
