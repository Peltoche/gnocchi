// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package phonenumbers

import (
	context "context"

	contacts "github.com/Peltoche/gnocchi/internal/service/contacts"

	mock "github.com/stretchr/testify/mock"

	sqlstorage "github.com/Peltoche/gnocchi/internal/tools/sqlstorage"

	uuid "github.com/Peltoche/gnocchi/internal/tools/uuid"
)

// MockService is an autogenerated mock type for the Service type
type MockService struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, cmd
func (_m *MockService) Create(ctx context.Context, cmd *CreateCmd) (*Phone, error) {
	ret := _m.Called(ctx, cmd)

	var r0 *Phone
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *CreateCmd) (*Phone, error)); ok {
		return rf(ctx, cmd)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *CreateCmd) *Phone); ok {
		r0 = rf(ctx, cmd)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Phone)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *CreateCmd) error); ok {
		r1 = rf(ctx, cmd)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteContactPhone provides a mock function with given fields: ctx, contact, phoneID
func (_m *MockService) DeleteContactPhone(ctx context.Context, contact *contacts.Contact, phoneID uuid.UUID) error {
	ret := _m.Called(ctx, contact, phoneID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *contacts.Contact, uuid.UUID) error); ok {
		r0 = rf(ctx, contact, phoneID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllForContact provides a mock function with given fields: ctx, contact, cmd
func (_m *MockService) GetAllForContact(ctx context.Context, contact *contacts.Contact, cmd *sqlstorage.PaginateCmd) ([]Phone, error) {
	ret := _m.Called(ctx, contact, cmd)

	var r0 []Phone
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *contacts.Contact, *sqlstorage.PaginateCmd) ([]Phone, error)); ok {
		return rf(ctx, contact, cmd)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *contacts.Contact, *sqlstorage.PaginateCmd) []Phone); ok {
		r0 = rf(ctx, contact, cmd)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]Phone)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *contacts.Contact, *sqlstorage.PaginateCmd) error); ok {
		r1 = rf(ctx, contact, cmd)
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