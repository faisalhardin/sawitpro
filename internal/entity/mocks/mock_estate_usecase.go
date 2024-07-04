// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/faisalhardin/sawitpro/internal/entity/interfaces (interfaces: EstateUsecase)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	model "github.com/faisalhardin/sawitpro/internal/entity/model"
	gomock "github.com/golang/mock/gomock"
)

// MockEstateUsecase is a mock of EstateUsecase interface.
type MockEstateUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockEstateUsecaseMockRecorder
}

// MockEstateUsecaseMockRecorder is the mock recorder for MockEstateUsecase.
type MockEstateUsecaseMockRecorder struct {
	mock *MockEstateUsecase
}

// NewMockEstateUsecase creates a new mock instance.
func NewMockEstateUsecase(ctrl *gomock.Controller) *MockEstateUsecase {
	mock := &MockEstateUsecase{ctrl: ctrl}
	mock.recorder = &MockEstateUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEstateUsecase) EXPECT() *MockEstateUsecaseMockRecorder {
	return m.recorder
}

// GetDronePlanByEstateUUID mocks base method.
func (m *MockEstateUsecase) GetDronePlanByEstateUUID(arg0 context.Context, arg1 string) (model.EstateDronePlanResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDronePlanByEstateUUID", arg0, arg1)
	ret0, _ := ret[0].(model.EstateDronePlanResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDronePlanByEstateUUID indicates an expected call of GetDronePlanByEstateUUID.
func (mr *MockEstateUsecaseMockRecorder) GetDronePlanByEstateUUID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDronePlanByEstateUUID", reflect.TypeOf((*MockEstateUsecase)(nil).GetDronePlanByEstateUUID), arg0, arg1)
}

// GetEstateStatsByUUID mocks base method.
func (m *MockEstateUsecase) GetEstateStatsByUUID(arg0 context.Context, arg1 string) (model.EstateStats, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEstateStatsByUUID", arg0, arg1)
	ret0, _ := ret[0].(model.EstateStats)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEstateStatsByUUID indicates an expected call of GetEstateStatsByUUID.
func (mr *MockEstateUsecaseMockRecorder) GetEstateStatsByUUID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEstateStatsByUUID", reflect.TypeOf((*MockEstateUsecase)(nil).GetEstateStatsByUUID), arg0, arg1)
}

// InsertEstate mocks base method.
func (m *MockEstateUsecase) InsertEstate(arg0 context.Context, arg1 model.InsertEstateRequest) (model.InsertEstateResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertEstate", arg0, arg1)
	ret0, _ := ret[0].(model.InsertEstateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertEstate indicates an expected call of InsertEstate.
func (mr *MockEstateUsecaseMockRecorder) InsertEstate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertEstate", reflect.TypeOf((*MockEstateUsecase)(nil).InsertEstate), arg0, arg1)
}

// InsertNewTree mocks base method.
func (m *MockEstateUsecase) InsertNewTree(arg0 context.Context, arg1 model.InsertNewTreeRequest) (model.InsertNewTreeResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertNewTree", arg0, arg1)
	ret0, _ := ret[0].(model.InsertNewTreeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertNewTree indicates an expected call of InsertNewTree.
func (mr *MockEstateUsecaseMockRecorder) InsertNewTree(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertNewTree", reflect.TypeOf((*MockEstateUsecase)(nil).InsertNewTree), arg0, arg1)
}
