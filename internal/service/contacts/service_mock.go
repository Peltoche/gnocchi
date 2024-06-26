// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package contacts

import (
	context "context"

	sqlstorage "github.com/Peltoche/gnocchi/internal/tools/sqlstorage"
	mock "github.com/stretchr/testify/mock"

	uuid "github.com/Peltoche/gnocchi/internal/tools/uuid"
)

// MockService is an autogenerated mock type for the Service type
type MockService struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, cmd
func (_m *MockService) Create(ctx context.Context, cmd *CreateCmd) (*Contact, error) {
	ret := _m.Called(ctx, cmd)

	var r0 *Contact
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *CreateCmd) (*Contact, error)); ok {
		return rf(ctx, cmd)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *CreateCmd) *Contact); ok {
		r0 = rf(ctx, cmd)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Contact)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *CreateCmd) error); ok {
		r1 = rf(ctx, cmd)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, contact
func (_m *MockService) Delete(ctx context.Context, contact *Contact) error {
	ret := _m.Called(ctx, contact)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *Contact) error); ok {
		r0 = rf(ctx, contact)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EditName provides a mock function with given fields: ctx, cmd
func (_m *MockService) EditName(ctx context.Context, cmd *EditNameCmd) (*Contact, error) {
	ret := _m.Called(ctx, cmd)

	var r0 *Contact
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *EditNameCmd) (*Contact, error)); ok {
		return rf(ctx, cmd)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *EditNameCmd) *Contact); ok {
		r0 = rf(ctx, cmd)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Contact)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *EditNameCmd) error); ok {
		r1 = rf(ctx, cmd)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields: ctx, paginateCmd
func (_m *MockService) GetAll(ctx context.Context, paginateCmd *sqlstorage.PaginateCmd) ([]Contact, error) {
	ret := _m.Called(ctx, paginateCmd)

	var r0 []Contact
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *sqlstorage.PaginateCmd) ([]Contact, error)); ok {
		return rf(ctx, paginateCmd)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *sqlstorage.PaginateCmd) []Contact); ok {
		r0 = rf(ctx, paginateCmd)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]Contact)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *sqlstorage.PaginateCmd) error); ok {
		r1 = rf(ctx, paginateCmd)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *MockService) GetByID(ctx context.Context, id uuid.UUID) (*Contact, error) {
	ret := _m.Called(ctx, id)

	var r0 *Contact
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) (*Contact, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) *Contact); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Contact)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewMockService creates a new instance of MockService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockService {
	mock := &MockService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
