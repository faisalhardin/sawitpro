// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/faisalhardin/sawitpro/internal/entity/interfaces (interfaces: EstateHandler)

// Package mocks is a generated GoMock package.
package mocks

import (
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockEstateHandler is a mock of EstateHandler interface.
type MockEstateHandler struct {
	ctrl     *gomock.Controller
	recorder *MockEstateHandlerMockRecorder
}

// MockEstateHandlerMockRecorder is the mock recorder for MockEstateHandler.
type MockEstateHandlerMockRecorder struct {
	mock *MockEstateHandler
}

// NewMockEstateHandler creates a new mock instance.
func NewMockEstateHandler(ctrl *gomock.Controller) *MockEstateHandler {
	mock := &MockEstateHandler{ctrl: ctrl}
	mock.recorder = &MockEstateHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEstateHandler) EXPECT() *MockEstateHandlerMockRecorder {
	return m.recorder
}

// GetDronePlan mocks base method.
func (m *MockEstateHandler) GetDronePlan(arg0 http.ResponseWriter, arg1 *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GetDronePlan", arg0, arg1)
}

// GetDronePlan indicates an expected call of GetDronePlan.
func (mr *MockEstateHandlerMockRecorder) GetDronePlan(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDronePlan", reflect.TypeOf((*MockEstateHandler)(nil).GetDronePlan), arg0, arg1)
}

// GetEstateStats mocks base method.
func (m *MockEstateHandler) GetEstateStats(arg0 http.ResponseWriter, arg1 *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GetEstateStats", arg0, arg1)
}

// GetEstateStats indicates an expected call of GetEstateStats.
func (mr *MockEstateHandlerMockRecorder) GetEstateStats(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEstateStats", reflect.TypeOf((*MockEstateHandler)(nil).GetEstateStats), arg0, arg1)
}

// InsertEstate mocks base method.
func (m *MockEstateHandler) InsertEstate(arg0 http.ResponseWriter, arg1 *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "InsertEstate", arg0, arg1)
}

// InsertEstate indicates an expected call of InsertEstate.
func (mr *MockEstateHandlerMockRecorder) InsertEstate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertEstate", reflect.TypeOf((*MockEstateHandler)(nil).InsertEstate), arg0, arg1)
}

// InsertTree mocks base method.
func (m *MockEstateHandler) InsertTree(arg0 http.ResponseWriter, arg1 *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "InsertTree", arg0, arg1)
}

// InsertTree indicates an expected call of InsertTree.
func (mr *MockEstateHandlerMockRecorder) InsertTree(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertTree", reflect.TypeOf((*MockEstateHandler)(nil).InsertTree), arg0, arg1)
}
