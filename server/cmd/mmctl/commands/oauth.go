// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package commands

import (
	"context"
	"fmt"

	"github.com/mattermost/mattermost/server/public/model"

	"github.com/mattermost/mattermost/server/v8/cmd/mmctl/client"
	"github.com/mattermost/mattermost/server/v8/cmd/mmctl/printer"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var OAuthCmd = &cobra.Command{
	Use:   "oauth",
	Short: "Management of OAuth2 apps",
}

var ListOAuthAppsCmd = &cobra.Command{
	Use:     "list",
	Short:   "List OAuth2 apps",
	Long:    "list all OAuth2 apps",
	Example: "  oauth list",
	RunE:    withClient(listOAuthAppsCmdF),
}

func listOAuthAppsCmdF(c client.Client, command *cobra.Command, args []string) error {
	// Fetch all apps with a very large limit so we get them all.
	apps, _, err := c.GetOAuthApps(command.Context(), 0, 100000000)
	if err != nil {
		return errors.Wrap(err, "Failed to fetch oauth2 apps")
	}

	userIds := make([]string, len(apps))
	for i := range apps {
		userIds[i] = apps[i].CreatorId
	}

	users, _, err := c.GetUsersByIds(context.Background(), userIds)
	if err != nil {
		return errors.Wrap(err, "Failed to fetch users for oauth2 apps")
	}

	usersByID := map[string]*model.User{}
	for _, user := range users {
		usersByID[user.Id] = user
	}

	for _, app := range apps {
		ownerName := app.CreatorId
		if owner, ok := usersByID[app.CreatorId]; ok {
			ownerName = owner.Username
		}
		printer.PrintT(fmt.Sprintf("{{.Id}}: {{.Name}} (Created by %s)", ownerName), app)
	}

	return nil
}

func init() {
	OAuthCmd.AddCommand(
		ListOAuthAppsCmd,
	)

	RootCmd.AddCommand(OAuthCmd)
}
