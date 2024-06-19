// Package mock_database is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	database "github.com/matheusmhmelo/FullCycle-rate-limiter/internal/infra/database"
)

// MockRateInterface is a mock of RateInterface interface.
type MockRateInterface struct {
	ctrl     *gomock.Controller
	recorder *MockRateInterfaceMockRecorder
}

// MockRateInterfaceMockRecorder is the mock recorder for MockRateInterface.
type MockRateInterfaceMockRecorder struct {
	mock *MockRateInterface
}

// NewMockRateInterface creates a new mock instance.
func NewMockRateInterface(ctrl *gomock.Controller) *MockRateInterface {
	mock := &MockRateInterface{ctrl: ctrl}
	mock.recorder = &MockRateInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRateInterface) EXPECT() *MockRateInterfaceMockRecorder {
	return m.recorder
}

// Block mocks base method.
func (m *MockRateInterface) Block(ctx context.Context, limiterType database.LimiterType, val string, ttl time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Block", ctx, limiterType, val, ttl)
	ret0, _ := ret[0].(error)
	return ret0
}

// Block indicates an expected call of Block.
func (mr *MockRateInterfaceMockRecorder) Block(ctx, limiterType, val, ttl interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Block", reflect.TypeOf((*MockRateInterface)(nil).Block), ctx, limiterType, val, ttl)
}

// FindBlocker mocks base method.
func (m *MockRateInterface) FindBlocker(ctx context.Context, limiterType database.LimiterType, val string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindBlocker", ctx, limiterType, val)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindBlocker indicates an expected call of FindBlocker.
func (mr *MockRateInterfaceMockRecorder) FindBlocker(ctx, limiterType, val interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindBlocker", reflect.TypeOf((*MockRateInterface)(nil).FindBlocker), ctx, limiterType, val)
}

// FindRequests mocks base method.
func (m *MockRateInterface) FindRequests(ctx context.Context, limiterType database.LimiterType, val string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindRequests", ctx, limiterType, val)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindRequests indicates an expected call of FindRequests.
func (mr *MockRateInterfaceMockRecorder) FindRequests(ctx, limiterType, val interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindRequests", reflect.TypeOf((*MockRateInterface)(nil).FindRequests), ctx, limiterType, val)
}

// NewRequest mocks base method.
func (m *MockRateInterface) NewRequest(ctx context.Context, limiterType database.LimiterType, val string, ttl time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewRequest", ctx, limiterType, val, ttl)
	ret0, _ := ret[0].(error)
	return ret0
}

// NewRequest indicates an expected call of NewRequest.
func (mr *MockRateInterfaceMockRecorder) NewRequest(ctx, limiterType, val, ttl interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewRequest", reflect.TypeOf((*MockRateInterface)(nil).NewRequest), ctx, limiterType, val, ttl)
}
