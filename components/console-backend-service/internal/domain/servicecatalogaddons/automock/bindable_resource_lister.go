// Code generated by mockery v1.0.0. DO NOT EDIT.

package automock

import gqlschema "github.com/kyma-project/kyma/components/console-backend-service/internal/gqlschema"
import mock "github.com/stretchr/testify/mock"

// bindableResourceLister is an autogenerated mock type for the bindableResourceLister type
type bindableResourceLister struct {
	mock.Mock
}

// ListResources provides a mock function with given fields: namespace
func (_m *bindableResourceLister) ListResources(namespace string) ([]gqlschema.BindableResourcesOutputItem, error) {
	ret := _m.Called(namespace)

	var r0 []gqlschema.BindableResourcesOutputItem
	if rf, ok := ret.Get(0).(func(string) []gqlschema.BindableResourcesOutputItem); ok {
		r0 = rf(namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]gqlschema.BindableResourcesOutputItem)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(namespace)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}