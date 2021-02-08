// Code generated by mockery v1.0.0. DO NOT EDIT.

// Regenerate this file using `make store-mocks`.

package mocks

import (
	model "github.com/mattermost/mattermost-server/v5/model"
	store "github.com/mattermost/mattermost-server/v5/store"
	mock "github.com/stretchr/testify/mock"
)

// SharedChannelStore is an autogenerated mock type for the SharedChannelStore type
type SharedChannelStore struct {
	mock.Mock
}

// Delete provides a mock function with given fields: channelId
func (_m *SharedChannelStore) Delete(channelId string) (bool, error) {
	ret := _m.Called(channelId)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(channelId)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(channelId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteRemote provides a mock function with given fields: remoteId
func (_m *SharedChannelStore) DeleteRemote(remoteId string) (bool, error) {
	ret := _m.Called(remoteId)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(remoteId)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(remoteId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: channelId
func (_m *SharedChannelStore) Get(channelId string) (*model.SharedChannel, error) {
	ret := _m.Called(channelId)

	var r0 *model.SharedChannel
	if rf, ok := ret.Get(0).(func(string) *model.SharedChannel); ok {
		r0 = rf(channelId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.SharedChannel)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(channelId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields: offset, limit, opts
func (_m *SharedChannelStore) GetAll(offset int, limit int, opts store.SharedChannelFilterOpts) ([]*model.SharedChannel, error) {
	ret := _m.Called(offset, limit, opts)

	var r0 []*model.SharedChannel
	if rf, ok := ret.Get(0).(func(int, int, store.SharedChannelFilterOpts) []*model.SharedChannel); ok {
		r0 = rf(offset, limit, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.SharedChannel)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int, store.SharedChannelFilterOpts) error); ok {
		r1 = rf(offset, limit, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllCount provides a mock function with given fields: opts
func (_m *SharedChannelStore) GetAllCount(opts store.SharedChannelFilterOpts) (int64, error) {
	ret := _m.Called(opts)

	var r0 int64
	if rf, ok := ret.Get(0).(func(store.SharedChannelFilterOpts) int64); ok {
		r0 = rf(opts)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(store.SharedChannelFilterOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAttachment provides a mock function with given fields: fileId, remoteId
func (_m *SharedChannelStore) GetAttachment(fileId string, remoteId string) (*model.SharedChannelAttachment, error) {
	ret := _m.Called(fileId, remoteId)

	var r0 *model.SharedChannelAttachment
	if rf, ok := ret.Get(0).(func(string, string) *model.SharedChannelAttachment); ok {
		r0 = rf(fileId, remoteId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.SharedChannelAttachment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(fileId, remoteId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRemote provides a mock function with given fields: id
func (_m *SharedChannelStore) GetRemote(id string) (*model.SharedChannelRemote, error) {
	ret := _m.Called(id)

	var r0 *model.SharedChannelRemote
	if rf, ok := ret.Get(0).(func(string) *model.SharedChannelRemote); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.SharedChannelRemote)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRemoteByIds provides a mock function with given fields: channelId, remoteId
func (_m *SharedChannelStore) GetRemoteByIds(channelId string, remoteId string) (*model.SharedChannelRemote, error) {
	ret := _m.Called(channelId, remoteId)

	var r0 *model.SharedChannelRemote
	if rf, ok := ret.Get(0).(func(string, string) *model.SharedChannelRemote); ok {
		r0 = rf(channelId, remoteId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.SharedChannelRemote)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(channelId, remoteId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRemotes provides a mock function with given fields: channelId
func (_m *SharedChannelStore) GetRemotes(channelId string) ([]*model.SharedChannelRemote, error) {
	ret := _m.Called(channelId)

	var r0 []*model.SharedChannelRemote
	if rf, ok := ret.Get(0).(func(string) []*model.SharedChannelRemote); ok {
		r0 = rf(channelId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.SharedChannelRemote)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(channelId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRemotesStatus provides a mock function with given fields: channelId
func (_m *SharedChannelStore) GetRemotesStatus(channelId string) ([]*model.SharedChannelRemoteStatus, error) {
	ret := _m.Called(channelId)

	var r0 []*model.SharedChannelRemoteStatus
	if rf, ok := ret.Get(0).(func(string) []*model.SharedChannelRemoteStatus); ok {
		r0 = rf(channelId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.SharedChannelRemoteStatus)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(channelId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUser provides a mock function with given fields: userId, remoteId
func (_m *SharedChannelStore) GetUser(userId string, remoteId string) (*model.SharedChannelUser, error) {
	ret := _m.Called(userId, remoteId)

	var r0 *model.SharedChannelUser
	if rf, ok := ret.Get(0).(func(string, string) *model.SharedChannelUser); ok {
		r0 = rf(userId, remoteId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.SharedChannelUser)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(userId, remoteId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HasChannel provides a mock function with given fields: channelID
func (_m *SharedChannelStore) HasChannel(channelID string) (bool, error) {
	ret := _m.Called(channelID)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(channelID)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(channelID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HasRemote provides a mock function with given fields: channelID, remoteId
func (_m *SharedChannelStore) HasRemote(channelID string, remoteId string) (bool, error) {
	ret := _m.Called(channelID, remoteId)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, string) bool); ok {
		r0 = rf(channelID, remoteId)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(channelID, remoteId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: sc
func (_m *SharedChannelStore) Save(sc *model.SharedChannel) (*model.SharedChannel, error) {
	ret := _m.Called(sc)

	var r0 *model.SharedChannel
	if rf, ok := ret.Get(0).(func(*model.SharedChannel) *model.SharedChannel); ok {
		r0 = rf(sc)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.SharedChannel)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.SharedChannel) error); ok {
		r1 = rf(sc)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveAttachment provides a mock function with given fields: remote
func (_m *SharedChannelStore) SaveAttachment(remote *model.SharedChannelAttachment) (*model.SharedChannelAttachment, error) {
	ret := _m.Called(remote)

	var r0 *model.SharedChannelAttachment
	if rf, ok := ret.Get(0).(func(*model.SharedChannelAttachment) *model.SharedChannelAttachment); ok {
		r0 = rf(remote)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.SharedChannelAttachment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.SharedChannelAttachment) error); ok {
		r1 = rf(remote)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveRemote provides a mock function with given fields: remote
func (_m *SharedChannelStore) SaveRemote(remote *model.SharedChannelRemote) (*model.SharedChannelRemote, error) {
	ret := _m.Called(remote)

	var r0 *model.SharedChannelRemote
	if rf, ok := ret.Get(0).(func(*model.SharedChannelRemote) *model.SharedChannelRemote); ok {
		r0 = rf(remote)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.SharedChannelRemote)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.SharedChannelRemote) error); ok {
		r1 = rf(remote)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveUser provides a mock function with given fields: remote
func (_m *SharedChannelStore) SaveUser(remote *model.SharedChannelUser) (*model.SharedChannelUser, error) {
	ret := _m.Called(remote)

	var r0 *model.SharedChannelUser
	if rf, ok := ret.Get(0).(func(*model.SharedChannelUser) *model.SharedChannelUser); ok {
		r0 = rf(remote)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.SharedChannelUser)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.SharedChannelUser) error); ok {
		r1 = rf(remote)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: sc
func (_m *SharedChannelStore) Update(sc *model.SharedChannel) (*model.SharedChannel, error) {
	ret := _m.Called(sc)

	var r0 *model.SharedChannel
	if rf, ok := ret.Get(0).(func(*model.SharedChannel) *model.SharedChannel); ok {
		r0 = rf(sc)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.SharedChannel)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.SharedChannel) error); ok {
		r1 = rf(sc)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateAttachmentLastSyncAt provides a mock function with given fields: id, syncTime
func (_m *SharedChannelStore) UpdateAttachmentLastSyncAt(id string, syncTime int64) error {
	ret := _m.Called(id, syncTime)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, int64) error); ok {
		r0 = rf(id, syncTime)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateRemote provides a mock function with given fields: remote
func (_m *SharedChannelStore) UpdateRemote(remote *model.SharedChannelRemote) (*model.SharedChannelRemote, error) {
	ret := _m.Called(remote)

	var r0 *model.SharedChannelRemote
	if rf, ok := ret.Get(0).(func(*model.SharedChannelRemote) *model.SharedChannelRemote); ok {
		r0 = rf(remote)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.SharedChannelRemote)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.SharedChannelRemote) error); ok {
		r1 = rf(remote)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateRemoteLastSyncAt provides a mock function with given fields: id, syncTime
func (_m *SharedChannelStore) UpdateRemoteLastSyncAt(id string, syncTime int64) error {
	ret := _m.Called(id, syncTime)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, int64) error); ok {
		r0 = rf(id, syncTime)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateUserLastSyncAt provides a mock function with given fields: id, syncTime
func (_m *SharedChannelStore) UpdateUserLastSyncAt(id string, syncTime int64) error {
	ret := _m.Called(id, syncTime)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, int64) error); ok {
		r0 = rf(id, syncTime)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpsertAttachment provides a mock function with given fields: remote
func (_m *SharedChannelStore) UpsertAttachment(remote *model.SharedChannelAttachment) (string, error) {
	ret := _m.Called(remote)

	var r0 string
	if rf, ok := ret.Get(0).(func(*model.SharedChannelAttachment) string); ok {
		r0 = rf(remote)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.SharedChannelAttachment) error); ok {
		r1 = rf(remote)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
