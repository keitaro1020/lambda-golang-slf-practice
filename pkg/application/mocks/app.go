// Code generated by MockGen. DO NOT EDIT.
// Source: ../pkg/application/app.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	domain "github.com/keitaro1020/lambda-golang-slf-practice/pkg/domain"
	reflect "reflect"
)

// MockApp is a mock of App interface
type MockApp struct {
	ctrl     *gomock.Controller
	recorder *MockAppMockRecorder
}

// MockAppMockRecorder is the mock recorder for MockApp
type MockAppMockRecorder struct {
	mock *MockApp
}

// NewMockApp creates a new mock instance
func NewMockApp(ctrl *gomock.Controller) *MockApp {
	mock := &MockApp{ctrl: ctrl}
	mock.recorder = &MockAppMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockApp) EXPECT() *MockAppMockRecorder {
	return m.recorder
}

// SQSWorker mocks base method
func (m *MockApp) SQSWorker(ctx context.Context, message string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SQSWorker", ctx, message)
	ret0, _ := ret[0].(error)
	return ret0
}

// SQSWorker indicates an expected call of SQSWorker
func (mr *MockAppMockRecorder) SQSWorker(ctx, message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SQSWorker", reflect.TypeOf((*MockApp)(nil).SQSWorker), ctx, message)
}

// S3Worker mocks base method
func (m *MockApp) S3Worker(ctx context.Context, bucket, filename string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "S3Worker", ctx, bucket, filename)
	ret0, _ := ret[0].(error)
	return ret0
}

// S3Worker indicates an expected call of S3Worker
func (mr *MockAppMockRecorder) S3Worker(ctx, bucket, filename interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "S3Worker", reflect.TypeOf((*MockApp)(nil).S3Worker), ctx, bucket, filename)
}

// GetCat mocks base method
func (m *MockApp) GetCat(ctx context.Context, id domain.CatID) (*domain.Cat, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCat", ctx, id)
	ret0, _ := ret[0].(*domain.Cat)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCat indicates an expected call of GetCat
func (mr *MockAppMockRecorder) GetCat(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCat", reflect.TypeOf((*MockApp)(nil).GetCat), ctx, id)
}

// GetCats mocks base method
func (m *MockApp) GetCats(ctx context.Context, first int64) (domain.Cats, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCats", ctx, first)
	ret0, _ := ret[0].(domain.Cats)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCats indicates an expected call of GetCats
func (mr *MockAppMockRecorder) GetCats(ctx, first interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCats", reflect.TypeOf((*MockApp)(nil).GetCats), ctx, first)
}
