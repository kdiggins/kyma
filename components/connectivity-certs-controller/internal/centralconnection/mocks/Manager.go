// Code generated by mockery v1.0.0. DO NOT EDIT.
package mocks

import mock "github.com/stretchr/testify/mock"
import v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
import v1alpha1 "github.com/kyma-project/kyma/components/connectivity-certs-controller/pkg/apis/applicationconnector/v1alpha1"

// Manager is an autogenerated mock type for the Manager type
type Manager struct {
	mock.Mock
}

// Create provides a mock function with given fields: _a0
func (_m *Manager) Create(_a0 *v1alpha1.CentralConnection) (*v1alpha1.CentralConnection, error) {
	ret := _m.Called(_a0)

	var r0 *v1alpha1.CentralConnection
	if rf, ok := ret.Get(0).(func(*v1alpha1.CentralConnection) *v1alpha1.CentralConnection); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1alpha1.CentralConnection)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*v1alpha1.CentralConnection) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: name, options
func (_m *Manager) Get(name string, options v1.GetOptions) (*v1alpha1.CentralConnection, error) {
	ret := _m.Called(name, options)

	var r0 *v1alpha1.CentralConnection
	if rf, ok := ret.Get(0).(func(string, v1.GetOptions) *v1alpha1.CentralConnection); ok {
		r0 = rf(name, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1alpha1.CentralConnection)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, v1.GetOptions) error); ok {
		r1 = rf(name, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: _a0
func (_m *Manager) Update(_a0 *v1alpha1.CentralConnection) (*v1alpha1.CentralConnection, error) {
	ret := _m.Called(_a0)

	var r0 *v1alpha1.CentralConnection
	if rf, ok := ret.Get(0).(func(*v1alpha1.CentralConnection) *v1alpha1.CentralConnection); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1alpha1.CentralConnection)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*v1alpha1.CentralConnection) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
