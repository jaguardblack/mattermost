// Code generated by mockery v2.23.2. DO NOT EDIT.

// Regenerate this file using `make einterfaces-mocks`.

package mocks

import (
	model "github.com/mattermost/mattermost/server/public/model"
	mock "github.com/stretchr/testify/mock"
)

// ElasticsearchAggregatorInterface is an autogenerated mock type for the ElasticsearchAggregatorInterface type
type ElasticsearchAggregatorInterface struct {
	mock.Mock
}

// MakeScheduler provides a mock function with given fields:
func (_m *ElasticsearchAggregatorInterface) MakeScheduler() model.Scheduler {
	ret := _m.Called()

	var r0 model.Scheduler
	if rf, ok := ret.Get(0).(func() model.Scheduler); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.Scheduler)
		}
	}

	return r0
}

// MakeWorker provides a mock function with given fields:
func (_m *ElasticsearchAggregatorInterface) MakeWorker() model.Worker {
	ret := _m.Called()

	var r0 model.Worker
	if rf, ok := ret.Get(0).(func() model.Worker); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.Worker)
		}
	}

	return r0
}

type mockConstructorTestingTNewElasticsearchAggregatorInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewElasticsearchAggregatorInterface creates a new instance of ElasticsearchAggregatorInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewElasticsearchAggregatorInterface(t mockConstructorTestingTNewElasticsearchAggregatorInterface) *ElasticsearchAggregatorInterface {
	mock := &ElasticsearchAggregatorInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
