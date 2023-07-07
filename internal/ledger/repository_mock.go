// Code generated by mockery v2.14.1. DO NOT EDIT.

package ledger

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	unitofwork "github.com/tauki/invoiceexchange/internal/unitofwork"

	uuid "github.com/google/uuid"
)

// MockRepository is an autogenerated mock type for the Repository type
type MockRepository struct {
	mock.Mock
}

// AddLedgerEntry provides a mock function with given fields: ctx, _a1, opts
func (_m *MockRepository) AddLedgerEntry(ctx context.Context, _a1 *Ledger, opts ...unitofwork.Option) error {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *Ledger, ...unitofwork.Option) error); ok {
		r0 = rf(ctx, _a1, opts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetLedgerEntriesByInvoiceID provides a mock function with given fields: ctx, invoiceID, opts
func (_m *MockRepository) GetLedgerEntriesByInvoiceID(ctx context.Context, invoiceID uuid.UUID, opts ...unitofwork.Option) ([]*Ledger, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, invoiceID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []*Ledger
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, ...unitofwork.Option) []*Ledger); ok {
		r0 = rf(ctx, invoiceID, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*Ledger)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID, ...unitofwork.Option) error); ok {
		r1 = rf(ctx, invoiceID, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewMockRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockRepository creates a new instance of MockRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockRepository(t mockConstructorTestingTNewMockRepository) *MockRepository {
	mock := &MockRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}