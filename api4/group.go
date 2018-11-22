// Copyright (c) 2018-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package api4

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/mattermost/mattermost-server/model"
)

const (
	groupMemberActionCreate = iota
	groupMemberActionDelete
)

func (api *API) InitGroup() {
	// GET /api/v4/groups/:group_id
	api.BaseRoutes.Groups.Handle("/{group_id:[A-Za-z0-9]+}",
		api.ApiSessionRequiredTrustRequester(getGroup)).Methods("GET")

	// PUT /api/v4/groups/:group_id/patch
	api.BaseRoutes.Groups.Handle("/{group_id:[A-Za-z0-9]+}/patch",
		api.ApiSessionRequired(patchGroup)).Methods("PUT")

	// POST /api/v4/teams/:team_id/link
	// POST /api/v4/channels/:channel_id/link
	api.BaseRoutes.Groups.Handle("/{group_id:[A-Za-z0-9]+}/{syncable_type:teams|channels}/{syncable_id:[A-Za-z0-9]+}/link",
		api.ApiSessionRequired(linkGroupSyncable)).Methods("POST")

	// DELETE /api/v4/teams/:team_id/link
	// DELETE /api/v4/channels/:channel_id/link
	api.BaseRoutes.Groups.Handle("/{group_id:[A-Za-z0-9]+}/{syncable_type:teams|channels}/{syncable_id:[A-Za-z0-9]+}/link",
		api.ApiSessionRequired(unlinkGroupSyncable)).Methods("DELETE")

	// GET /api/v4/teams/:team_id
	// GET /api/v4/channels/:channel_id
	api.BaseRoutes.Groups.Handle("/{group_id:[A-Za-z0-9]+}/{syncable_type:teams|channels}/{syncable_id:[A-Za-z0-9]+}",
		api.ApiSessionRequired(getGroupSyncable)).Methods("GET")

	// GET /api/v4/teams
	// GET /api/v4/channels
	api.BaseRoutes.Groups.Handle("/{group_id:[A-Za-z0-9]+}/{syncable_type:teams|channels}",
		api.ApiSessionRequired(getGroupSyncables)).Methods("GET")

	// PUT /api/v4/teams/:team_id/patch
	// PUT /api/v4/channels/:channel_id/patch
	api.BaseRoutes.Groups.Handle("/{group_id:[A-Za-z0-9]+}/{syncable_type:teams|channels}/{syncable_id:[A-Za-z0-9]+}/patch",
		api.ApiSessionRequired(patchGroupSyncable)).Methods("PUT")

	// GET /api/v4/groups/:group_id/members?page=0&per_page=100
	api.BaseRoutes.Groups.Handle("/{group_id:[A-Za-z0-9]+}/members",
		api.ApiSessionRequiredTrustRequester(getGroupMembers)).Methods("GET")
}

func getGroup(c *Context, w http.ResponseWriter, r *http.Request) {
	c.RequireGroupId()
	if c.Err != nil {
		return
	}

	if c.App.License() == nil || !*c.App.License().Features.LDAPGroups {
		c.Err = model.NewAppError("Api4.getGroup", "api.groups.license.error", nil, "", http.StatusNotImplemented)
		return
	}

	if !c.App.SessionHasPermissionTo(c.Session, model.PERMISSION_MANAGE_SYSTEM) {
		c.SetPermissionError(model.PERMISSION_MANAGE_SYSTEM)
		return
	}

	group, err := c.App.GetGroup(c.Params.GroupId)
	if err != nil {
		c.Err = err
		return
	}

	b, marshalErr := json.Marshal(group)
	if marshalErr != nil {
		c.Err = model.NewAppError("Api4.getGroup", "api.group.marshal_error", nil, marshalErr.Error(),
			http.StatusNotImplemented)
		return
	}

	w.Write(b)
}

func patchGroup(c *Context, w http.ResponseWriter, r *http.Request) {
	c.RequireGroupId()
	if c.Err != nil {
		return
	}

	groupPatch := model.GroupPatchFromJson(r.Body)
	if groupPatch == nil {
		c.SetInvalidParam("group")
		return
	}

	if c.App.License() == nil || !*c.App.License().Features.LDAPGroups {
		c.Err = model.NewAppError("Api4.patchGroup", "api.group.license.error", nil, "", http.StatusNotImplemented)
		return
	}

	if !c.App.SessionHasPermissionTo(c.Session, model.PERMISSION_MANAGE_SYSTEM) {
		c.SetPermissionError(model.PERMISSION_MANAGE_SYSTEM)
		return
	}

	group, err := c.App.GetGroup(c.Params.GroupId)
	if err != nil {
		c.Err = err
		return
	}

	group.Patch(groupPatch)

	group, err = c.App.UpdateGroup(group)
	if err != nil {
		c.Err = err
		return
	}

	b, marshalErr := json.Marshal(group)
	if marshalErr != nil {
		c.Err = model.NewAppError("Api4.patchGroup", "api.group.marshal_error", nil, marshalErr.Error(),
			http.StatusNotImplemented)
		return
	}

	w.Write(b)
}

func linkGroupSyncable(c *Context, w http.ResponseWriter, r *http.Request) {
	c.RequireGroupId()
	if c.Err != nil {
		return
	}

	c.RequireSyncableId()
	if c.Err != nil {
		return
	}
	syncableID := c.Params.SyncableId

	c.RequireSyncableType()
	if c.Err != nil {
		return
	}
	syncableType := c.Params.SyncableType

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		c.Err = model.NewAppError("Api4.createGroupSyncable", "api.group.io_error", nil, err.Error(),
			http.StatusNotImplemented)
	}

	var patch *model.GroupSyncablePatch
	err = json.Unmarshal(body, &patch)
	if err != nil || patch == nil {
		c.SetInvalidParam(fmt.Sprintf("Group%s", syncableType.String()))
		return
	}

	if c.App.License() == nil || !*c.App.License().Features.LDAPGroups {
		c.Err = model.NewAppError("Api4.createGroupSyncable", "api.group.license.error", nil, "",
			http.StatusNotImplemented)
		return
	}

	if !c.App.SessionHasPermissionTo(c.Session, model.PERMISSION_MANAGE_SYSTEM) {
		c.SetPermissionError(model.PERMISSION_MANAGE_SYSTEM)
		return
	}

	groupSyncable, appErr := c.App.GetGroupSyncable(c.Params.GroupId, syncableID, syncableType)
	if appErr != nil {
		if appErr.Id != "store.sql_group.no_rows" {
			c.Err = appErr
			return
		}
	}

	if groupSyncable == nil || groupSyncable.DeleteAt == 0 {
		groupSyncable = &model.GroupSyncable{
			GroupId:    c.Params.GroupId,
			SyncableId: syncableID,
			Type:       syncableType,
		}
		groupSyncable.Patch(patch)
		groupSyncable, appErr = c.App.CreateGroupSyncable(groupSyncable)
		if appErr != nil {
			c.Err = appErr
			return
		}
	} else {
		groupSyncable.DeleteAt = 0
		groupSyncable.Patch(patch)
		groupSyncable, appErr = c.App.UpdateGroupSyncable(groupSyncable)
		if appErr != nil {
			c.Err = appErr
			return
		}
	}

	w.WriteHeader(http.StatusCreated)

	var marshalErr error
	b, marshalErr := json.Marshal(groupSyncable)
	if marshalErr != nil {
		c.Err = model.NewAppError("Api4.createGroupSyncable", "api.group.marshal_error", nil,
			marshalErr.Error(), http.StatusNotImplemented)
		return
	}

	w.Write(b)
}

func getGroupSyncable(c *Context, w http.ResponseWriter, r *http.Request) {
	c.RequireGroupId()
	if c.Err != nil {
		return
	}

	c.RequireSyncableId()
	if c.Err != nil {
		return
	}
	syncableID := c.Params.SyncableId

	c.RequireSyncableType()
	if c.Err != nil {
		return
	}
	syncableType := c.Params.SyncableType

	if c.App.License() == nil || !*c.App.License().Features.LDAPGroups {
		c.Err = model.NewAppError("Api4.getGroupSyncable", "api.groups.license.error", nil, "", http.StatusNotImplemented)
		return
	}

	if !c.App.SessionHasPermissionTo(c.Session, model.PERMISSION_MANAGE_SYSTEM) {
		c.SetPermissionError(model.PERMISSION_MANAGE_SYSTEM)
		return
	}

	groupSyncable, err := c.App.GetGroupSyncable(c.Params.GroupId, syncableID, syncableType)
	if err != nil {
		c.Err = err
		return
	}

	b, marshalErr := json.Marshal(groupSyncable)
	if marshalErr != nil {
		c.Err = model.NewAppError("Api4.getGroupSyncable", "api.group.marshal_error", nil, marshalErr.Error(),
			http.StatusNotImplemented)
		return
	}

	w.Write(b)
}

func getGroupSyncables(c *Context, w http.ResponseWriter, r *http.Request) {
	c.RequireGroupId()
	if c.Err != nil {
		return
	}

	c.RequireSyncableType()
	if c.Err != nil {
		return
	}
	syncableType := c.Params.SyncableType

	if c.App.License() == nil || !*c.App.License().Features.LDAPGroups {
		c.Err = model.NewAppError("Api4.getGroupSyncables", "api.group.license.error", nil, "", http.StatusNotImplemented)
		return
	}

	if !c.App.SessionHasPermissionTo(c.Session, model.PERMISSION_MANAGE_SYSTEM) {
		c.SetPermissionError(model.PERMISSION_MANAGE_SYSTEM)
		return
	}

	groupSyncables, err := c.App.GetGroupSyncables(c.Params.GroupId, syncableType)
	if err != nil {
		c.Err = err
		return
	}

	if len(groupSyncables) == 0 {
		w.Write([]byte("[]"))
		return
	}

	b, marshalErr := json.Marshal(groupSyncables)
	if marshalErr != nil {
		c.Err = model.NewAppError("Api4.getGroupSyncables", "api.group.marshal_error", nil, marshalErr.Error(),
			http.StatusNotImplemented)
		return
	}

	w.Write(b)
}

func patchGroupSyncable(c *Context, w http.ResponseWriter, r *http.Request) {
	c.RequireGroupId()
	if c.Err != nil {
		return
	}

	c.RequireSyncableId()
	if c.Err != nil {
		return
	}
	syncableID := c.Params.SyncableId

	c.RequireSyncableType()
	if c.Err != nil {
		return
	}
	syncableType := c.Params.SyncableType

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		c.Err = model.NewAppError("Api4.patchGroupSyncable", "api.group.io_error", nil, err.Error(),
			http.StatusNotImplemented)
	}

	var patch *model.GroupSyncablePatch
	err = json.Unmarshal(body, &patch)
	if err != nil || patch == nil {
		c.SetInvalidParam(fmt.Sprintf("Group[%s]Patch", syncableType.String()))
		return
	}

	if c.App.License() == nil || !*c.App.License().Features.LDAPGroups {
		c.Err = model.NewAppError("Api4.patchGroupSyncable", "api.group.license.error", nil, "",
			http.StatusNotImplemented)
		return
	}

	if !c.App.SessionHasPermissionTo(c.Session, model.PERMISSION_MANAGE_SYSTEM) {
		c.SetPermissionError(model.PERMISSION_MANAGE_SYSTEM)
		return
	}

	var appErr *model.AppError
	groupSyncable, appErr := c.App.GetGroupSyncable(c.Params.GroupId, syncableID, syncableType)
	if appErr != nil {
		c.Err = appErr
		return
	}

	groupSyncable.Patch(patch)

	groupSyncable, appErr = c.App.UpdateGroupSyncable(groupSyncable)
	if appErr != nil {
		c.Err = appErr
		return
	}

	b, marshalErr := json.Marshal(groupSyncable)
	if marshalErr != nil {
		c.Err = model.NewAppError("Api4.patchGroupSyncable", "api.group.marshal_error", nil, marshalErr.Error(),
			http.StatusNotImplemented)
		return
	}

	w.Write(b)
}

func unlinkGroupSyncable(c *Context, w http.ResponseWriter, r *http.Request) {
	c.RequireGroupId()
	if c.Err != nil {
		return
	}

	c.RequireSyncableId()
	if c.Err != nil {
		return
	}
	syncableID := c.Params.SyncableId

	c.RequireSyncableType()
	if c.Err != nil {
		return
	}
	syncableType := c.Params.SyncableType

	if c.App.License() == nil || !*c.App.License().Features.LDAPGroups {
		c.Err = model.NewAppError("Api4.unlinkGroupSyncable", "api.group.license.error", nil, "",
			http.StatusNotImplemented)
		return
	}

	if !c.App.SessionHasPermissionTo(c.Session, model.PERMISSION_MANAGE_SYSTEM) {
		c.SetPermissionError(model.PERMISSION_MANAGE_SYSTEM)
		return
	}

	_, err := c.App.DeleteGroupSyncable(c.Params.GroupId, syncableID, syncableType)
	if err != nil {
		c.Err = err
		return
	}

	w.WriteHeader(http.StatusNoContent)
	w.Write(nil)
}

func getGroupMembers(c *Context, w http.ResponseWriter, r *http.Request) {
	if c.App.License() == nil || !*c.App.License().Features.LDAPGroups {
		c.Err = model.NewAppError("Api4.getGroupMembers", "api.group.license.error", nil, "", http.StatusNotImplemented)
		return
	}

	if !c.App.SessionHasPermissionTo(c.Session, model.PERMISSION_MANAGE_SYSTEM) {
		c.SetPermissionError(model.PERMISSION_MANAGE_SYSTEM)
		return
	}

	c.RequireGroupId()
	if c.Err != nil {
		return
	}

	members, count, err := c.App.GetGroupMemberUsersPage(c.Params.GroupId, c.Params.Page, c.Params.PerPage)
	if err != nil {
		c.Err = err
		return
	}

	b, marshalErr := json.Marshal(struct {
		Members []*model.User `json:"members"`
		Count   int           `json:"total_member_count"`
	}{
		Members: members,
		Count:   count,
	})
	if marshalErr != nil {
		c.Err = model.NewAppError("Api4.getGroupMembers", "api.group.marshal_error", nil, marshalErr.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(b)
}
