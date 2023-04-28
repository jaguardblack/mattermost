// Code generated by mockery v2.23.2. DO NOT EDIT.

// Regenerate this file using `make misc-mocks`.

package mocks

import (
	model "github.com/mattermost/mattermost-server/server/v8/model"
	mock "github.com/stretchr/testify/mock"
)

// LicenseValidatorIface is an autogenerated mock type for the LicenseValidatorIface type
type LicenseValidatorIface struct {
	mock.Mock
}

// LicenseFromBytes provides a mock function with given fields: licenseBytes
func (_m *LicenseValidatorIface) LicenseFromBytes(licenseBytes []byte) (*model.License, *model.AppError) {
	ret := _m.Called(licenseBytes)

	var r0 *model.License
	var r1 *model.AppError
	if rf, ok := ret.Get(0).(func([]byte) (*model.License, *model.AppError)); ok {
		return rf(licenseBytes)
	}
	if rf, ok := ret.Get(0).(func([]byte) *model.License); ok {
		r0 = rf(licenseBytes)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.License)
		}
	}

	if rf, ok := ret.Get(1).(func([]byte) *model.AppError); ok {
		r1 = rf(licenseBytes)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// ValidateLicense provides a mock function with given fields: signed
func (_m *LicenseValidatorIface) ValidateLicense(signed []byte) (bool, string) {
	ret := _m.Called(signed)

	var r0 bool
	var r1 string
	if rf, ok := ret.Get(0).(func([]byte) (bool, string)); ok {
		return rf(signed)
	}
	if rf, ok := ret.Get(0).(func([]byte) bool); ok {
		r0 = rf(signed)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func([]byte) string); ok {
		r1 = rf(signed)
	} else {
		r1 = ret.Get(1).(string)
	}

	return r0, r1
}

type mockConstructorTestingTNewLicenseValidatorIface interface {
	mock.TestingT
	Cleanup(func())
}

// NewLicenseValidatorIface creates a new instance of LicenseValidatorIface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewLicenseValidatorIface(t mockConstructorTestingTNewLicenseValidatorIface) *LicenseValidatorIface {
	mock := &LicenseValidatorIface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
