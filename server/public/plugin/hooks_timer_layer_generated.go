// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

// Code generated by "make pluginapi"
// DO NOT EDIT

package plugin

import (
	"io"
	"net/http"
	timePkg "time"

	"github.com/mattermost/mattermost/server/public/model"
)

type hooksTimerLayer struct {
	pluginID  string
	hooksImpl Hooks
	metrics   metricsInterface
}

func (hooks *hooksTimerLayer) recordTime(startTime timePkg.Time, name string, success bool) {
	if hooks.metrics != nil {
		elapsedTime := float64(timePkg.Since(startTime)) / float64(timePkg.Second)
		hooks.metrics.ObservePluginHookDuration(hooks.pluginID, name, success, elapsedTime)
	}
}

func (hooks *hooksTimerLayer) OnActivate() error {
	startTime := timePkg.Now()
	_returnsA := hooks.hooksImpl.OnActivate()
	hooks.recordTime(startTime, "OnActivate", _returnsA == nil)
	return _returnsA
}

func (hooks *hooksTimerLayer) Implemented() ([]string, error) {
	startTime := timePkg.Now()
	_returnsA, _returnsB := hooks.hooksImpl.Implemented()
	hooks.recordTime(startTime, "Implemented", _returnsB == nil)
	return _returnsA, _returnsB
}

func (hooks *hooksTimerLayer) OnDeactivate() error {
	startTime := timePkg.Now()
	_returnsA := hooks.hooksImpl.OnDeactivate()
	hooks.recordTime(startTime, "OnDeactivate", _returnsA == nil)
	return _returnsA
}

func (hooks *hooksTimerLayer) OnConfigurationChange() error {
	startTime := timePkg.Now()
	_returnsA := hooks.hooksImpl.OnConfigurationChange()
	hooks.recordTime(startTime, "OnConfigurationChange", _returnsA == nil)
	return _returnsA
}

func (hooks *hooksTimerLayer) ServeHTTP(c *Context, w http.ResponseWriter, r *http.Request) {
	startTime := timePkg.Now()
	hooks.hooksImpl.ServeHTTP(c, w, r)
	hooks.recordTime(startTime, "ServeHTTP", true)
}

func (hooks *hooksTimerLayer) ExecuteCommand(c *Context, args *model.CommandArgs) (*model.CommandResponse, *model.AppError) {
	startTime := timePkg.Now()
	_returnsA, _returnsB := hooks.hooksImpl.ExecuteCommand(c, args)
	hooks.recordTime(startTime, "ExecuteCommand", _returnsB == nil)
	return _returnsA, _returnsB
}

func (hooks *hooksTimerLayer) UserHasBeenCreated(c *Context, user *model.User) {
	startTime := timePkg.Now()
	hooks.hooksImpl.UserHasBeenCreated(c, user)
	hooks.recordTime(startTime, "UserHasBeenCreated", true)
}

func (hooks *hooksTimerLayer) UserWillLogIn(c *Context, user *model.User) string {
	startTime := timePkg.Now()
	_returnsA := hooks.hooksImpl.UserWillLogIn(c, user)
	hooks.recordTime(startTime, "UserWillLogIn", true)
	return _returnsA
}

func (hooks *hooksTimerLayer) UserHasLoggedIn(c *Context, user *model.User) {
	startTime := timePkg.Now()
	hooks.hooksImpl.UserHasLoggedIn(c, user)
	hooks.recordTime(startTime, "UserHasLoggedIn", true)
}

func (hooks *hooksTimerLayer) MessageWillBePosted(c *Context, post *model.Post) (*model.Post, string) {
	startTime := timePkg.Now()
	_returnsA, _returnsB := hooks.hooksImpl.MessageWillBePosted(c, post)
	hooks.recordTime(startTime, "MessageWillBePosted", true)
	return _returnsA, _returnsB
}

func (hooks *hooksTimerLayer) MessageWillBeUpdated(c *Context, newPost, oldPost *model.Post) (*model.Post, string) {
	startTime := timePkg.Now()
	_returnsA, _returnsB := hooks.hooksImpl.MessageWillBeUpdated(c, newPost, oldPost)
	hooks.recordTime(startTime, "MessageWillBeUpdated", true)
	return _returnsA, _returnsB
}

func (hooks *hooksTimerLayer) MessageHasBeenPosted(c *Context, post *model.Post) {
	startTime := timePkg.Now()
	hooks.hooksImpl.MessageHasBeenPosted(c, post)
	hooks.recordTime(startTime, "MessageHasBeenPosted", true)
}

func (hooks *hooksTimerLayer) MessageHasBeenUpdated(c *Context, newPost, oldPost *model.Post) {
	startTime := timePkg.Now()
	hooks.hooksImpl.MessageHasBeenUpdated(c, newPost, oldPost)
	hooks.recordTime(startTime, "MessageHasBeenUpdated", true)
}

func (hooks *hooksTimerLayer) ChannelHasBeenCreated(c *Context, channel *model.Channel) {
	startTime := timePkg.Now()
	hooks.hooksImpl.ChannelHasBeenCreated(c, channel)
	hooks.recordTime(startTime, "ChannelHasBeenCreated", true)
}

func (hooks *hooksTimerLayer) UserHasJoinedChannel(c *Context, channelMember *model.ChannelMember, actor *model.User) {
	startTime := timePkg.Now()
	hooks.hooksImpl.UserHasJoinedChannel(c, channelMember, actor)
	hooks.recordTime(startTime, "UserHasJoinedChannel", true)
}

func (hooks *hooksTimerLayer) UserHasLeftChannel(c *Context, channelMember *model.ChannelMember, actor *model.User) {
	startTime := timePkg.Now()
	hooks.hooksImpl.UserHasLeftChannel(c, channelMember, actor)
	hooks.recordTime(startTime, "UserHasLeftChannel", true)
}

func (hooks *hooksTimerLayer) UserHasJoinedTeam(c *Context, teamMember *model.TeamMember, actor *model.User) {
	startTime := timePkg.Now()
	hooks.hooksImpl.UserHasJoinedTeam(c, teamMember, actor)
	hooks.recordTime(startTime, "UserHasJoinedTeam", true)
}

func (hooks *hooksTimerLayer) UserHasLeftTeam(c *Context, teamMember *model.TeamMember, actor *model.User) {
	startTime := timePkg.Now()
	hooks.hooksImpl.UserHasLeftTeam(c, teamMember, actor)
	hooks.recordTime(startTime, "UserHasLeftTeam", true)
}

func (hooks *hooksTimerLayer) FileWillBeUploaded(c *Context, info *model.FileInfo, file io.Reader, output io.Writer) (*model.FileInfo, string) {
	startTime := timePkg.Now()
	_returnsA, _returnsB := hooks.hooksImpl.FileWillBeUploaded(c, info, file, output)
	hooks.recordTime(startTime, "FileWillBeUploaded", true)
	return _returnsA, _returnsB
}

func (hooks *hooksTimerLayer) ReactionHasBeenAdded(c *Context, reaction *model.Reaction) {
	startTime := timePkg.Now()
	hooks.hooksImpl.ReactionHasBeenAdded(c, reaction)
	hooks.recordTime(startTime, "ReactionHasBeenAdded", true)
}

func (hooks *hooksTimerLayer) ReactionHasBeenRemoved(c *Context, reaction *model.Reaction) {
	startTime := timePkg.Now()
	hooks.hooksImpl.ReactionHasBeenRemoved(c, reaction)
	hooks.recordTime(startTime, "ReactionHasBeenRemoved", true)
}

func (hooks *hooksTimerLayer) OnPluginClusterEvent(c *Context, ev model.PluginClusterEvent) {
	startTime := timePkg.Now()
	hooks.hooksImpl.OnPluginClusterEvent(c, ev)
	hooks.recordTime(startTime, "OnPluginClusterEvent", true)
}

func (hooks *hooksTimerLayer) OnWebSocketConnect(webConnID, userID string) {
	startTime := timePkg.Now()
	hooks.hooksImpl.OnWebSocketConnect(webConnID, userID)
	hooks.recordTime(startTime, "OnWebSocketConnect", true)
}

func (hooks *hooksTimerLayer) OnWebSocketDisconnect(webConnID, userID string) {
	startTime := timePkg.Now()
	hooks.hooksImpl.OnWebSocketDisconnect(webConnID, userID)
	hooks.recordTime(startTime, "OnWebSocketDisconnect", true)
}

func (hooks *hooksTimerLayer) WebSocketMessageHasBeenPosted(webConnID, userID string, req *model.WebSocketRequest) {
	startTime := timePkg.Now()
	hooks.hooksImpl.WebSocketMessageHasBeenPosted(webConnID, userID, req)
	hooks.recordTime(startTime, "WebSocketMessageHasBeenPosted", true)
}

func (hooks *hooksTimerLayer) RunDataRetention(nowTime, batchSize int64) (int64, error) {
	startTime := timePkg.Now()
	_returnsA, _returnsB := hooks.hooksImpl.RunDataRetention(nowTime, batchSize)
	hooks.recordTime(startTime, "RunDataRetention", _returnsB == nil)
	return _returnsA, _returnsB
}

func (hooks *hooksTimerLayer) OnInstall(c *Context, event model.OnInstallEvent) error {
	startTime := timePkg.Now()
	_returnsA := hooks.hooksImpl.OnInstall(c, event)
	hooks.recordTime(startTime, "OnInstall", _returnsA == nil)
	return _returnsA
}

func (hooks *hooksTimerLayer) OnSendDailyTelemetry() {
	startTime := timePkg.Now()
	hooks.hooksImpl.OnSendDailyTelemetry()
	hooks.recordTime(startTime, "OnSendDailyTelemetry", true)
}

func (hooks *hooksTimerLayer) OnCloudLimitsUpdated(limits *model.ProductLimits) {
	startTime := timePkg.Now()
	hooks.hooksImpl.OnCloudLimitsUpdated(limits)
	hooks.recordTime(startTime, "OnCloudLimitsUpdated", true)
}

func (hooks *hooksTimerLayer) ConfigurationWillBeSaved(newCfg *model.Config) (*model.Config, error) {
	startTime := timePkg.Now()
	_returnsA, _returnsB := hooks.hooksImpl.ConfigurationWillBeSaved(newCfg)
	hooks.recordTime(startTime, "ConfigurationWillBeSaved", _returnsB == nil)
	return _returnsA, _returnsB
}

func (hooks *hooksTimerLayer) NotificationWillBePushed(pushNotification *model.PushNotification, userID string) (*model.PushNotification, string) {
	startTime := timePkg.Now()
	_returnsA, _returnsB := hooks.hooksImpl.NotificationWillBePushed(pushNotification, userID)
	hooks.recordTime(startTime, "NotificationWillBePushed", true)
	return _returnsA, _returnsB
}
