package mocks

import common "github.com/arreyder/compliance-masonry/lib/common"
import mock "github.com/stretchr/testify/mock"

// Workspace is an autogenerated mock type for the Workspace type
type Workspace struct {
	mock.Mock
}

// GetAllComponents provides a mock function with given fields:
func (_m *Workspace) GetAllComponents() []common.Component {
	ret := _m.Called()

	var r0 []common.Component
	if rf, ok := ret.Get(0).(func() []common.Component); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]common.Component)
		}
	}

	return r0
}

// GetAllVerificationsWith provides a mock function with given fields: standardKey, controlKey
func (_m *Workspace) GetAllVerificationsWith(standardKey string, controlKey string) common.Verifications {
	ret := _m.Called(standardKey, controlKey)

	var r0 common.Verifications
	if rf, ok := ret.Get(0).(func(string, string) common.Verifications); ok {
		r0 = rf(standardKey, controlKey)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Verifications)
		}
	}

	return r0
}

// GetCertification provides a mock function with given fields:
func (_m *Workspace) GetCertification() common.Certification {
	ret := _m.Called()

	var r0 common.Certification
	if rf, ok := ret.Get(0).(func() common.Certification); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Certification)
		}
	}

	return r0
}

// GetComponent provides a mock function with given fields: componentKey
func (_m *Workspace) GetComponent(componentKey string) (common.Component, bool) {
	ret := _m.Called(componentKey)

	var r0 common.Component
	if rf, ok := ret.Get(0).(func(string) common.Component); ok {
		r0 = rf(componentKey)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Component)
		}
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(string) bool); ok {
		r1 = rf(componentKey)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// GetStandard provides a mock function with given fields: standardKey
func (_m *Workspace) GetStandard(standardKey string) (common.Standard, bool) {
	ret := _m.Called(standardKey)

	var r0 common.Standard
	if rf, ok := ret.Get(0).(func(string) common.Standard); ok {
		r0 = rf(standardKey)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Standard)
		}
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(string) bool); ok {
		r1 = rf(standardKey)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// LoadCertification provides a mock function with given fields: _a0
func (_m *Workspace) LoadCertification(_a0 string) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LoadComponents provides a mock function with given fields: _a0
func (_m *Workspace) LoadComponents(_a0 string) []error {
	ret := _m.Called(_a0)

	var r0 []error
	if rf, ok := ret.Get(0).(func(string) []error); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]error)
		}
	}

	return r0
}

// LoadStandards provides a mock function with given fields: _a0
func (_m *Workspace) LoadStandards(_a0 string) []error {
	ret := _m.Called(_a0)

	var r0 []error
	if rf, ok := ret.Get(0).(func(string) []error); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]error)
		}
	}

	return r0
}

var _ common.Workspace = (*Workspace)(nil)
