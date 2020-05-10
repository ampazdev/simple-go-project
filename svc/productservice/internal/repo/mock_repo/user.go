// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ampazdev/simple-go-project/services/userservices/internal (interfaces: UserRepo)

// Package mock_repo is a generated GoMock package.
package mock_repo

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	entity "github.com/ampazdev/simple-go-project/services/userservices/internal/entity"
	reflect "reflect"
)

// MockUserRepo is a mock of UserRepo interface
type MockUserRepo struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepoMockRecorder
}

// MockUserRepoMockRecorder is the mock recorder for MockUserRepo
type MockUserRepoMockRecorder struct {
	mock *MockUserRepo
}

// NewMockUserRepo creates a new mock instance
func NewMockUserRepo(ctrl *gomock.Controller) *MockUserRepo {
	mock := &MockUserRepo{ctrl: ctrl}
	mock.recorder = &MockUserRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUserRepo) EXPECT() *MockUserRepoMockRecorder {
	return m.recorder
}

// GetAll mocks base method
func (m *MockUserRepo) GetAll(arg0 context.Context) ([]*entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", arg0)
	ret0, _ := ret[0].([]*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll
func (mr *MockUserRepoMockRecorder) GetAll(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockUserRepo)(nil).GetAll), arg0)
}

// GetByID mocks base method
func (m *MockUserRepo) GetByID(arg0 context.Context, arg1 int64) (*entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", arg0, arg1)
	ret0, _ := ret[0].(*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID
func (mr *MockUserRepoMockRecorder) GetByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockUserRepo)(nil).GetByID), arg0, arg1)
}
