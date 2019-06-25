// Copyright (c) 2016-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package model

type UserGetOptions struct {
	// Filters the users in the team
	InTeamId string
	// Filters the users not in the team
	NotInTeamId string
	// Filters the users in the channel
	InChannelId string
	// Filters the users not in the channel
	NotInChannelId string
	// Filters the users group constrained
	GroupConstrained bool
	// Filters the users without a team
	WithoutTeam bool
	// Filters the inactive users
	Inactive bool
	// Filters for the given role
	Role string
	// Sorting option
	Sort string
	// Restrict to search in a list of teams and channels
	ViewRestrictions *ViewUsersRestrictions
	// Page
	Page int
	// Page size
	PerPage int
}

type UserGetByIdsOptions struct {
	// IsAdmin tracks whether or not the request is being made by an administrator. Does nothing when provided by a client.
	IsAdmin bool

	// Since filters the users based on their UpdateAt timestamp.
	Since int64

	// Restrict to search in a list of teams and channels. Does nothing when provided by a client.
	ViewRestrictions *ViewUsersRestrictions
}
