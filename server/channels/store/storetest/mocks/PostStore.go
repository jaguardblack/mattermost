// Code generated by mockery v2.23.2. DO NOT EDIT.

// Regenerate this file using `make store-mocks`.

package mocks

import (
	context "context"

	model "github.com/mattermost/mattermost/server/public/model"
	mock "github.com/stretchr/testify/mock"

	request "github.com/mattermost/mattermost/server/public/shared/request"

	store "github.com/mattermost/mattermost/server/v8/channels/store"
)

// PostStore is an autogenerated mock type for the PostStore type
type PostStore struct {
	mock.Mock
}

// AnalyticsPostCount provides a mock function with given fields: options
func (_m *PostStore) AnalyticsPostCount(options *model.PostCountOptions) (int64, error) {
	ret := _m.Called(options)

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(*model.PostCountOptions) (int64, error)); ok {
		return rf(options)
	}
	if rf, ok := ret.Get(0).(func(*model.PostCountOptions) int64); ok {
		r0 = rf(options)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(*model.PostCountOptions) error); ok {
		r1 = rf(options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AnalyticsPostCountsByDay provides a mock function with given fields: options
func (_m *PostStore) AnalyticsPostCountsByDay(options *model.AnalyticsPostCountsOptions) (model.AnalyticsRows, error) {
	ret := _m.Called(options)

	var r0 model.AnalyticsRows
	var r1 error
	if rf, ok := ret.Get(0).(func(*model.AnalyticsPostCountsOptions) (model.AnalyticsRows, error)); ok {
		return rf(options)
	}
	if rf, ok := ret.Get(0).(func(*model.AnalyticsPostCountsOptions) model.AnalyticsRows); ok {
		r0 = rf(options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.AnalyticsRows)
		}
	}

	if rf, ok := ret.Get(1).(func(*model.AnalyticsPostCountsOptions) error); ok {
		r1 = rf(options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AnalyticsUserCountsWithPostsByDay provides a mock function with given fields: teamID
func (_m *PostStore) AnalyticsUserCountsWithPostsByDay(teamID string) (model.AnalyticsRows, error) {
	ret := _m.Called(teamID)

	var r0 model.AnalyticsRows
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (model.AnalyticsRows, error)); ok {
		return rf(teamID)
	}
	if rf, ok := ret.Get(0).(func(string) model.AnalyticsRows); ok {
		r0 = rf(teamID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.AnalyticsRows)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(teamID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ClearCaches provides a mock function with given fields:
func (_m *PostStore) ClearCaches() {
	_m.Called()
}

// Delete provides a mock function with given fields: rctx, postID, timestamp, deleteByID
func (_m *PostStore) Delete(rctx request.CTX, postID string, timestamp int64, deleteByID string) error {
	ret := _m.Called(rctx, postID, timestamp, deleteByID)

	var r0 error
	if rf, ok := ret.Get(0).(func(request.CTX, string, int64, string) error); ok {
		r0 = rf(rctx, postID, timestamp, deleteByID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: ctx, id, opts, userID, sanitizeOptions
func (_m *PostStore) Get(ctx context.Context, id string, opts model.GetPostsOptions, userID string, sanitizeOptions map[string]bool) (*model.PostList, error) {
	ret := _m.Called(ctx, id, opts, userID, sanitizeOptions)

	var r0 *model.PostList
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, model.GetPostsOptions, string, map[string]bool) (*model.PostList, error)); ok {
		return rf(ctx, id, opts, userID, sanitizeOptions)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, model.GetPostsOptions, string, map[string]bool) *model.PostList); ok {
		r0 = rf(ctx, id, opts, userID, sanitizeOptions)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.PostList)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, model.GetPostsOptions, string, map[string]bool) error); ok {
		r1 = rf(ctx, id, opts, userID, sanitizeOptions)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDirectPostParentsForExportAfter provides a mock function with given fields: limit, afterID, includeArchivedChannels
func (_m *PostStore) GetDirectPostParentsForExportAfter(limit int, afterID string, includeArchivedChannels bool) ([]*model.DirectPostForExport, error) {
	ret := _m.Called(limit, afterID, includeArchivedChannels)

	var r0 []*model.DirectPostForExport
	var r1 error
	if rf, ok := ret.Get(0).(func(int, string, bool) ([]*model.DirectPostForExport, error)); ok {
		return rf(limit, afterID, includeArchivedChannels)
	}
	if rf, ok := ret.Get(0).(func(int, string, bool) []*model.DirectPostForExport); ok {
		r0 = rf(limit, afterID, includeArchivedChannels)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.DirectPostForExport)
		}
	}

	if rf, ok := ret.Get(1).(func(int, string, bool) error); ok {
		r1 = rf(limit, afterID, includeArchivedChannels)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetEditHistoryForPost provides a mock function with given fields: postId
func (_m *PostStore) GetEditHistoryForPost(postId string) ([]*model.Post, error) {
	ret := _m.Called(postId)

	var r0 []*model.Post
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]*model.Post, error)); ok {
		return rf(postId)
	}
	if rf, ok := ret.Get(0).(func(string) []*model.Post); ok {
		r0 = rf(postId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Post)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(postId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetEtag provides a mock function with given fields: channelID, allowFromCache, collapsedThreads
func (_m *PostStore) GetEtag(channelID string, allowFromCache bool, collapsedThreads bool) string {
	ret := _m.Called(channelID, allowFromCache, collapsedThreads)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, bool, bool) string); ok {
		r0 = rf(channelID, allowFromCache, collapsedThreads)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetFlaggedPosts provides a mock function with given fields: userID, offset, limit
func (_m *PostStore) GetFlaggedPosts(userID string, offset int, limit int) (*model.PostList, error) {
	ret := _m.Called(userID, offset, limit)

	var r0 *model.PostList
	var r1 error
	if rf, ok := ret.Get(0).(func(string, int, int) (*model.PostList, error)); ok {
		return rf(userID, offset, limit)
	}
	if rf, ok := ret.Get(0).(func(string, int, int) *model.PostList); ok {
		r0 = rf(userID, offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.PostList)
		}
	}

	if rf, ok := ret.Get(1).(func(string, int, int) error); ok {
		r1 = rf(userID, offset, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetFlaggedPostsForChannel provides a mock function with given fields: userID, channelID, offset, limit
func (_m *PostStore) GetFlaggedPostsForChannel(userID string, channelID string, offset int, limit int) (*model.PostList, error) {
	ret := _m.Called(userID, channelID, offset, limit)

	var r0 *model.PostList
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, int, int) (*model.PostList, error)); ok {
		return rf(userID, channelID, offset, limit)
	}
	if rf, ok := ret.Get(0).(func(string, string, int, int) *model.PostList); ok {
		r0 = rf(userID, channelID, offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.PostList)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string, int, int) error); ok {
		r1 = rf(userID, channelID, offset, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetFlaggedPostsForTeam provides a mock function with given fields: userID, teamID, offset, limit
func (_m *PostStore) GetFlaggedPostsForTeam(userID string, teamID string, offset int, limit int) (*model.PostList, error) {
	ret := _m.Called(userID, teamID, offset, limit)

	var r0 *model.PostList
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, int, int) (*model.PostList, error)); ok {
		return rf(userID, teamID, offset, limit)
	}
	if rf, ok := ret.Get(0).(func(string, string, int, int) *model.PostList); ok {
		r0 = rf(userID, teamID, offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.PostList)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string, int, int) error); ok {
		r1 = rf(userID, teamID, offset, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMaxPostSize provides a mock function with given fields:
func (_m *PostStore) GetMaxPostSize() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// GetNthRecentPostTime provides a mock function with given fields: n
func (_m *PostStore) GetNthRecentPostTime(n int64) (int64, error) {
	ret := _m.Called(n)

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(int64) (int64, error)); ok {
		return rf(n)
	}
	if rf, ok := ret.Get(0).(func(int64) int64); ok {
		r0 = rf(n)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(n)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOldest provides a mock function with given fields:
func (_m *PostStore) GetOldest() (*model.Post, error) {
	ret := _m.Called()

	var r0 *model.Post
	var r1 error
	if rf, ok := ret.Get(0).(func() (*model.Post, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *model.Post); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Post)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOldestEntityCreationTime provides a mock function with given fields:
func (_m *PostStore) GetOldestEntityCreationTime() (int64, error) {
	ret := _m.Called()

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func() (int64, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetParentsForExportAfter provides a mock function with given fields: limit, afterID, includeArchivedChannels
func (_m *PostStore) GetParentsForExportAfter(limit int, afterID string, includeArchivedChannels bool) ([]*model.PostForExport, error) {
	ret := _m.Called(limit, afterID, includeArchivedChannels)

	var r0 []*model.PostForExport
	var r1 error
	if rf, ok := ret.Get(0).(func(int, string, bool) ([]*model.PostForExport, error)); ok {
		return rf(limit, afterID, includeArchivedChannels)
	}
	if rf, ok := ret.Get(0).(func(int, string, bool) []*model.PostForExport); ok {
		r0 = rf(limit, afterID, includeArchivedChannels)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.PostForExport)
		}
	}

	if rf, ok := ret.Get(1).(func(int, string, bool) error); ok {
		r1 = rf(limit, afterID, includeArchivedChannels)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPostAfterTime provides a mock function with given fields: channelID, timestamp, collapsedThreads
func (_m *PostStore) GetPostAfterTime(channelID string, timestamp int64, collapsedThreads bool) (*model.Post, error) {
	ret := _m.Called(channelID, timestamp, collapsedThreads)

	var r0 *model.Post
	var r1 error
	if rf, ok := ret.Get(0).(func(string, int64, bool) (*model.Post, error)); ok {
		return rf(channelID, timestamp, collapsedThreads)
	}
	if rf, ok := ret.Get(0).(func(string, int64, bool) *model.Post); ok {
		r0 = rf(channelID, timestamp, collapsedThreads)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Post)
		}
	}

	if rf, ok := ret.Get(1).(func(string, int64, bool) error); ok {
		r1 = rf(channelID, timestamp, collapsedThreads)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPostIdAfterTime provides a mock function with given fields: channelID, timestamp, collapsedThreads
func (_m *PostStore) GetPostIdAfterTime(channelID string, timestamp int64, collapsedThreads bool) (string, error) {
	ret := _m.Called(channelID, timestamp, collapsedThreads)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, int64, bool) (string, error)); ok {
		return rf(channelID, timestamp, collapsedThreads)
	}
	if rf, ok := ret.Get(0).(func(string, int64, bool) string); ok {
		r0 = rf(channelID, timestamp, collapsedThreads)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string, int64, bool) error); ok {
		r1 = rf(channelID, timestamp, collapsedThreads)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPostIdBeforeTime provides a mock function with given fields: channelID, timestamp, collapsedThreads
func (_m *PostStore) GetPostIdBeforeTime(channelID string, timestamp int64, collapsedThreads bool) (string, error) {
	ret := _m.Called(channelID, timestamp, collapsedThreads)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, int64, bool) (string, error)); ok {
		return rf(channelID, timestamp, collapsedThreads)
	}
	if rf, ok := ret.Get(0).(func(string, int64, bool) string); ok {
		r0 = rf(channelID, timestamp, collapsedThreads)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string, int64, bool) error); ok {
		r1 = rf(channelID, timestamp, collapsedThreads)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPostReminderMetadata provides a mock function with given fields: postID
func (_m *PostStore) GetPostReminderMetadata(postID string) (*store.PostReminderMetadata, error) {
	ret := _m.Called(postID)

	var r0 *store.PostReminderMetadata
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*store.PostReminderMetadata, error)); ok {
		return rf(postID)
	}
	if rf, ok := ret.Get(0).(func(string) *store.PostReminderMetadata); ok {
		r0 = rf(postID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*store.PostReminderMetadata)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(postID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPostReminders provides a mock function with given fields: now
func (_m *PostStore) GetPostReminders(now int64) ([]*model.PostReminder, error) {
	ret := _m.Called(now)

	var r0 []*model.PostReminder
	var r1 error
	if rf, ok := ret.Get(0).(func(int64) ([]*model.PostReminder, error)); ok {
		return rf(now)
	}
	if rf, ok := ret.Get(0).(func(int64) []*model.PostReminder); ok {
		r0 = rf(now)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.PostReminder)
		}
	}

	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(now)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPosts provides a mock function with given fields: options, allowFromCache, sanitizeOptions
func (_m *PostStore) GetPosts(options model.GetPostsOptions, allowFromCache bool, sanitizeOptions map[string]bool) (*model.PostList, error) {
	ret := _m.Called(options, allowFromCache, sanitizeOptions)

	var r0 *model.PostList
	var r1 error
	if rf, ok := ret.Get(0).(func(model.GetPostsOptions, bool, map[string]bool) (*model.PostList, error)); ok {
		return rf(options, allowFromCache, sanitizeOptions)
	}
	if rf, ok := ret.Get(0).(func(model.GetPostsOptions, bool, map[string]bool) *model.PostList); ok {
		r0 = rf(options, allowFromCache, sanitizeOptions)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.PostList)
		}
	}

	if rf, ok := ret.Get(1).(func(model.GetPostsOptions, bool, map[string]bool) error); ok {
		r1 = rf(options, allowFromCache, sanitizeOptions)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPostsAfter provides a mock function with given fields: options, sanitizeOptions
func (_m *PostStore) GetPostsAfter(options model.GetPostsOptions, sanitizeOptions map[string]bool) (*model.PostList, error) {
	ret := _m.Called(options, sanitizeOptions)

	var r0 *model.PostList
	var r1 error
	if rf, ok := ret.Get(0).(func(model.GetPostsOptions, map[string]bool) (*model.PostList, error)); ok {
		return rf(options, sanitizeOptions)
	}
	if rf, ok := ret.Get(0).(func(model.GetPostsOptions, map[string]bool) *model.PostList); ok {
		r0 = rf(options, sanitizeOptions)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.PostList)
		}
	}

	if rf, ok := ret.Get(1).(func(model.GetPostsOptions, map[string]bool) error); ok {
		r1 = rf(options, sanitizeOptions)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPostsBatchForIndexing provides a mock function with given fields: startTime, startPostID, limit
func (_m *PostStore) GetPostsBatchForIndexing(startTime int64, startPostID string, limit int) ([]*model.PostForIndexing, error) {
	ret := _m.Called(startTime, startPostID, limit)

	var r0 []*model.PostForIndexing
	var r1 error
	if rf, ok := ret.Get(0).(func(int64, string, int) ([]*model.PostForIndexing, error)); ok {
		return rf(startTime, startPostID, limit)
	}
	if rf, ok := ret.Get(0).(func(int64, string, int) []*model.PostForIndexing); ok {
		r0 = rf(startTime, startPostID, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.PostForIndexing)
		}
	}

	if rf, ok := ret.Get(1).(func(int64, string, int) error); ok {
		r1 = rf(startTime, startPostID, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPostsBefore provides a mock function with given fields: options, sanitizeOptions
func (_m *PostStore) GetPostsBefore(options model.GetPostsOptions, sanitizeOptions map[string]bool) (*model.PostList, error) {
	ret := _m.Called(options, sanitizeOptions)

	var r0 *model.PostList
	var r1 error
	if rf, ok := ret.Get(0).(func(model.GetPostsOptions, map[string]bool) (*model.PostList, error)); ok {
		return rf(options, sanitizeOptions)
	}
	if rf, ok := ret.Get(0).(func(model.GetPostsOptions, map[string]bool) *model.PostList); ok {
		r0 = rf(options, sanitizeOptions)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.PostList)
		}
	}

	if rf, ok := ret.Get(1).(func(model.GetPostsOptions, map[string]bool) error); ok {
		r1 = rf(options, sanitizeOptions)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPostsByIds provides a mock function with given fields: postIds
func (_m *PostStore) GetPostsByIds(postIds []string) ([]*model.Post, error) {
	ret := _m.Called(postIds)

	var r0 []*model.Post
	var r1 error
	if rf, ok := ret.Get(0).(func([]string) ([]*model.Post, error)); ok {
		return rf(postIds)
	}
	if rf, ok := ret.Get(0).(func([]string) []*model.Post); ok {
		r0 = rf(postIds)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Post)
		}
	}

	if rf, ok := ret.Get(1).(func([]string) error); ok {
		r1 = rf(postIds)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPostsByThread provides a mock function with given fields: threadID, since
func (_m *PostStore) GetPostsByThread(threadID string, since int64) ([]*model.Post, error) {
	ret := _m.Called(threadID, since)

	var r0 []*model.Post
	var r1 error
	if rf, ok := ret.Get(0).(func(string, int64) ([]*model.Post, error)); ok {
		return rf(threadID, since)
	}
	if rf, ok := ret.Get(0).(func(string, int64) []*model.Post); ok {
		r0 = rf(threadID, since)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Post)
		}
	}

	if rf, ok := ret.Get(1).(func(string, int64) error); ok {
		r1 = rf(threadID, since)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPostsCreatedAt provides a mock function with given fields: channelID, timestamp
func (_m *PostStore) GetPostsCreatedAt(channelID string, timestamp int64) ([]*model.Post, error) {
	ret := _m.Called(channelID, timestamp)

	var r0 []*model.Post
	var r1 error
	if rf, ok := ret.Get(0).(func(string, int64) ([]*model.Post, error)); ok {
		return rf(channelID, timestamp)
	}
	if rf, ok := ret.Get(0).(func(string, int64) []*model.Post); ok {
		r0 = rf(channelID, timestamp)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Post)
		}
	}

	if rf, ok := ret.Get(1).(func(string, int64) error); ok {
		r1 = rf(channelID, timestamp)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPostsSince provides a mock function with given fields: options, allowFromCache, sanitizeOptions
func (_m *PostStore) GetPostsSince(options model.GetPostsSinceOptions, allowFromCache bool, sanitizeOptions map[string]bool) (*model.PostList, error) {
	ret := _m.Called(options, allowFromCache, sanitizeOptions)

	var r0 *model.PostList
	var r1 error
	if rf, ok := ret.Get(0).(func(model.GetPostsSinceOptions, bool, map[string]bool) (*model.PostList, error)); ok {
		return rf(options, allowFromCache, sanitizeOptions)
	}
	if rf, ok := ret.Get(0).(func(model.GetPostsSinceOptions, bool, map[string]bool) *model.PostList); ok {
		r0 = rf(options, allowFromCache, sanitizeOptions)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.PostList)
		}
	}

	if rf, ok := ret.Get(1).(func(model.GetPostsSinceOptions, bool, map[string]bool) error); ok {
		r1 = rf(options, allowFromCache, sanitizeOptions)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPostsSinceForSync provides a mock function with given fields: options, cursor, limit
func (_m *PostStore) GetPostsSinceForSync(options model.GetPostsSinceForSyncOptions, cursor model.GetPostsSinceForSyncCursor, limit int) ([]*model.Post, model.GetPostsSinceForSyncCursor, error) {
	ret := _m.Called(options, cursor, limit)

	var r0 []*model.Post
	var r1 model.GetPostsSinceForSyncCursor
	var r2 error
	if rf, ok := ret.Get(0).(func(model.GetPostsSinceForSyncOptions, model.GetPostsSinceForSyncCursor, int) ([]*model.Post, model.GetPostsSinceForSyncCursor, error)); ok {
		return rf(options, cursor, limit)
	}
	if rf, ok := ret.Get(0).(func(model.GetPostsSinceForSyncOptions, model.GetPostsSinceForSyncCursor, int) []*model.Post); ok {
		r0 = rf(options, cursor, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Post)
		}
	}

	if rf, ok := ret.Get(1).(func(model.GetPostsSinceForSyncOptions, model.GetPostsSinceForSyncCursor, int) model.GetPostsSinceForSyncCursor); ok {
		r1 = rf(options, cursor, limit)
	} else {
		r1 = ret.Get(1).(model.GetPostsSinceForSyncCursor)
	}

	if rf, ok := ret.Get(2).(func(model.GetPostsSinceForSyncOptions, model.GetPostsSinceForSyncCursor, int) error); ok {
		r2 = rf(options, cursor, limit)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetRepliesForExport provides a mock function with given fields: parentID
func (_m *PostStore) GetRepliesForExport(parentID string) ([]*model.ReplyForExport, error) {
	ret := _m.Called(parentID)

	var r0 []*model.ReplyForExport
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]*model.ReplyForExport, error)); ok {
		return rf(parentID)
	}
	if rf, ok := ret.Get(0).(func(string) []*model.ReplyForExport); ok {
		r0 = rf(parentID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.ReplyForExport)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(parentID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSingle provides a mock function with given fields: id, inclDeleted
func (_m *PostStore) GetSingle(id string, inclDeleted bool) (*model.Post, error) {
	ret := _m.Called(id, inclDeleted)

	var r0 *model.Post
	var r1 error
	if rf, ok := ret.Get(0).(func(string, bool) (*model.Post, error)); ok {
		return rf(id, inclDeleted)
	}
	if rf, ok := ret.Get(0).(func(string, bool) *model.Post); ok {
		r0 = rf(id, inclDeleted)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Post)
		}
	}

	if rf, ok := ret.Get(1).(func(string, bool) error); ok {
		r1 = rf(id, inclDeleted)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HasAutoResponsePostByUserSince provides a mock function with given fields: options, userId
func (_m *PostStore) HasAutoResponsePostByUserSince(options model.GetPostsSinceOptions, userId string) (bool, error) {
	ret := _m.Called(options, userId)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(model.GetPostsSinceOptions, string) (bool, error)); ok {
		return rf(options, userId)
	}
	if rf, ok := ret.Get(0).(func(model.GetPostsSinceOptions, string) bool); ok {
		r0 = rf(options, userId)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(model.GetPostsSinceOptions, string) error); ok {
		r1 = rf(options, userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InvalidateLastPostTimeCache provides a mock function with given fields: channelID
func (_m *PostStore) InvalidateLastPostTimeCache(channelID string) {
	_m.Called(channelID)
}

// Overwrite provides a mock function with given fields: rctx, post
func (_m *PostStore) Overwrite(rctx request.CTX, post *model.Post) (*model.Post, error) {
	ret := _m.Called(rctx, post)

	var r0 *model.Post
	var r1 error
	if rf, ok := ret.Get(0).(func(request.CTX, *model.Post) (*model.Post, error)); ok {
		return rf(rctx, post)
	}
	if rf, ok := ret.Get(0).(func(request.CTX, *model.Post) *model.Post); ok {
		r0 = rf(rctx, post)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Post)
		}
	}

	if rf, ok := ret.Get(1).(func(request.CTX, *model.Post) error); ok {
		r1 = rf(rctx, post)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OverwriteMultiple provides a mock function with given fields: posts
func (_m *PostStore) OverwriteMultiple(posts []*model.Post) ([]*model.Post, int, error) {
	ret := _m.Called(posts)

	var r0 []*model.Post
	var r1 int
	var r2 error
	if rf, ok := ret.Get(0).(func([]*model.Post) ([]*model.Post, int, error)); ok {
		return rf(posts)
	}
	if rf, ok := ret.Get(0).(func([]*model.Post) []*model.Post); ok {
		r0 = rf(posts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Post)
		}
	}

	if rf, ok := ret.Get(1).(func([]*model.Post) int); ok {
		r1 = rf(posts)
	} else {
		r1 = ret.Get(1).(int)
	}

	if rf, ok := ret.Get(2).(func([]*model.Post) error); ok {
		r2 = rf(posts)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// PermanentDeleteBatch provides a mock function with given fields: endTime, limit
func (_m *PostStore) PermanentDeleteBatch(endTime int64, limit int64) (int64, error) {
	ret := _m.Called(endTime, limit)

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(int64, int64) (int64, error)); ok {
		return rf(endTime, limit)
	}
	if rf, ok := ret.Get(0).(func(int64, int64) int64); ok {
		r0 = rf(endTime, limit)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(int64, int64) error); ok {
		r1 = rf(endTime, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PermanentDeleteBatchForRetentionPolicies provides a mock function with given fields: now, globalPolicyEndTime, limit, cursor
func (_m *PostStore) PermanentDeleteBatchForRetentionPolicies(now int64, globalPolicyEndTime int64, limit int64, cursor model.RetentionPolicyCursor) (int64, model.RetentionPolicyCursor, error) {
	ret := _m.Called(now, globalPolicyEndTime, limit, cursor)

	var r0 int64
	var r1 model.RetentionPolicyCursor
	var r2 error
	if rf, ok := ret.Get(0).(func(int64, int64, int64, model.RetentionPolicyCursor) (int64, model.RetentionPolicyCursor, error)); ok {
		return rf(now, globalPolicyEndTime, limit, cursor)
	}
	if rf, ok := ret.Get(0).(func(int64, int64, int64, model.RetentionPolicyCursor) int64); ok {
		r0 = rf(now, globalPolicyEndTime, limit, cursor)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(int64, int64, int64, model.RetentionPolicyCursor) model.RetentionPolicyCursor); ok {
		r1 = rf(now, globalPolicyEndTime, limit, cursor)
	} else {
		r1 = ret.Get(1).(model.RetentionPolicyCursor)
	}

	if rf, ok := ret.Get(2).(func(int64, int64, int64, model.RetentionPolicyCursor) error); ok {
		r2 = rf(now, globalPolicyEndTime, limit, cursor)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// PermanentDeleteByChannel provides a mock function with given fields: rctx, channelID
func (_m *PostStore) PermanentDeleteByChannel(rctx request.CTX, channelID string) error {
	ret := _m.Called(rctx, channelID)

	var r0 error
	if rf, ok := ret.Get(0).(func(request.CTX, string) error); ok {
		r0 = rf(rctx, channelID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PermanentDeleteByUser provides a mock function with given fields: rctx, userID
func (_m *PostStore) PermanentDeleteByUser(rctx request.CTX, userID string) error {
	ret := _m.Called(rctx, userID)

	var r0 error
	if rf, ok := ret.Get(0).(func(request.CTX, string) error); ok {
		r0 = rf(rctx, userID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Save provides a mock function with given fields: rctx, post
func (_m *PostStore) Save(rctx request.CTX, post *model.Post) (*model.Post, error) {
	ret := _m.Called(rctx, post)

	var r0 *model.Post
	var r1 error
	if rf, ok := ret.Get(0).(func(request.CTX, *model.Post) (*model.Post, error)); ok {
		return rf(rctx, post)
	}
	if rf, ok := ret.Get(0).(func(request.CTX, *model.Post) *model.Post); ok {
		r0 = rf(rctx, post)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Post)
		}
	}

	if rf, ok := ret.Get(1).(func(request.CTX, *model.Post) error); ok {
		r1 = rf(rctx, post)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveMultiple provides a mock function with given fields: posts
func (_m *PostStore) SaveMultiple(posts []*model.Post) ([]*model.Post, int, error) {
	ret := _m.Called(posts)

	var r0 []*model.Post
	var r1 int
	var r2 error
	if rf, ok := ret.Get(0).(func([]*model.Post) ([]*model.Post, int, error)); ok {
		return rf(posts)
	}
	if rf, ok := ret.Get(0).(func([]*model.Post) []*model.Post); ok {
		r0 = rf(posts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Post)
		}
	}

	if rf, ok := ret.Get(1).(func([]*model.Post) int); ok {
		r1 = rf(posts)
	} else {
		r1 = ret.Get(1).(int)
	}

	if rf, ok := ret.Get(2).(func([]*model.Post) error); ok {
		r2 = rf(posts)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Search provides a mock function with given fields: teamID, userID, params
func (_m *PostStore) Search(teamID string, userID string, params *model.SearchParams) (*model.PostList, error) {
	ret := _m.Called(teamID, userID, params)

	var r0 *model.PostList
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, *model.SearchParams) (*model.PostList, error)); ok {
		return rf(teamID, userID, params)
	}
	if rf, ok := ret.Get(0).(func(string, string, *model.SearchParams) *model.PostList); ok {
		r0 = rf(teamID, userID, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.PostList)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string, *model.SearchParams) error); ok {
		r1 = rf(teamID, userID, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SearchPostsForUser provides a mock function with given fields: rctx, paramsList, userID, teamID, page, perPage
func (_m *PostStore) SearchPostsForUser(rctx request.CTX, paramsList []*model.SearchParams, userID string, teamID string, page int, perPage int) (*model.PostSearchResults, error) {
	ret := _m.Called(rctx, paramsList, userID, teamID, page, perPage)

	var r0 *model.PostSearchResults
	var r1 error
	if rf, ok := ret.Get(0).(func(request.CTX, []*model.SearchParams, string, string, int, int) (*model.PostSearchResults, error)); ok {
		return rf(rctx, paramsList, userID, teamID, page, perPage)
	}
	if rf, ok := ret.Get(0).(func(request.CTX, []*model.SearchParams, string, string, int, int) *model.PostSearchResults); ok {
		r0 = rf(rctx, paramsList, userID, teamID, page, perPage)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.PostSearchResults)
		}
	}

	if rf, ok := ret.Get(1).(func(request.CTX, []*model.SearchParams, string, string, int, int) error); ok {
		r1 = rf(rctx, paramsList, userID, teamID, page, perPage)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetPostReminder provides a mock function with given fields: reminder
func (_m *PostStore) SetPostReminder(reminder *model.PostReminder) error {
	ret := _m.Called(reminder)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.PostReminder) error); ok {
		r0 = rf(reminder)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: rctx, newPost, oldPost
func (_m *PostStore) Update(rctx request.CTX, newPost *model.Post, oldPost *model.Post) (*model.Post, error) {
	ret := _m.Called(rctx, newPost, oldPost)

	var r0 *model.Post
	var r1 error
	if rf, ok := ret.Get(0).(func(request.CTX, *model.Post, *model.Post) (*model.Post, error)); ok {
		return rf(rctx, newPost, oldPost)
	}
	if rf, ok := ret.Get(0).(func(request.CTX, *model.Post, *model.Post) *model.Post); ok {
		r0 = rf(rctx, newPost, oldPost)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Post)
		}
	}

	if rf, ok := ret.Get(1).(func(request.CTX, *model.Post, *model.Post) error); ok {
		r1 = rf(rctx, newPost, oldPost)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewPostStore interface {
	mock.TestingT
	Cleanup(func())
}

// NewPostStore creates a new instance of PostStore. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPostStore(t mockConstructorTestingTNewPostStore) *PostStore {
	mock := &PostStore{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
