// Code generated by MockGen. DO NOT EDIT.
// Source: selector.go

// Package setuptest is a generated GoMock package.
package setuptest

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockChainNetworkExistingSelector is a mock of ChainNetworkExistingSelector interface
type MockChainNetworkExistingSelector struct {
	ctrl     *gomock.Controller
	recorder *MockChainNetworkExistingSelectorMockRecorder
}

// MockChainNetworkExistingSelectorMockRecorder is the mock recorder for MockChainNetworkExistingSelector
type MockChainNetworkExistingSelectorMockRecorder struct {
	mock *MockChainNetworkExistingSelector
}

// NewMockChainNetworkExistingSelector creates a new mock instance
func NewMockChainNetworkExistingSelector(ctrl *gomock.Controller) *MockChainNetworkExistingSelector {
	mock := &MockChainNetworkExistingSelector{ctrl: ctrl}
	mock.recorder = &MockChainNetworkExistingSelectorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockChainNetworkExistingSelector) EXPECT() *MockChainNetworkExistingSelectorMockRecorder {
	return m.recorder
}

// Select mocks base method
func (m *MockChainNetworkExistingSelector) Select(chain, network, receiver string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Select", chain, network, receiver)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Select indicates an expected call of Select
func (mr *MockChainNetworkExistingSelectorMockRecorder) Select(chain, network, receiver interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Select", reflect.TypeOf((*MockChainNetworkExistingSelector)(nil).Select), chain, network, receiver)
}