// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/aws-sdk-go/service/appconfigdata/appconfigdataiface (interfaces: AppConfigDataAPI)

// Package appconfigdatamock is a generated GoMock package.
package appconfigdatamock

import (
	context "context"
	reflect "reflect"

	request "github.com/aws/aws-sdk-go/aws/request"
	appconfigdata "github.com/aws/aws-sdk-go/service/appconfigdata"
	gomock "github.com/golang/mock/gomock"
)

// MockAppConfigDataAPI is a mock of AppConfigDataAPI interface.
type MockAppConfigDataAPI struct {
	ctrl     *gomock.Controller
	recorder *MockAppConfigDataAPIMockRecorder
}

// MockAppConfigDataAPIMockRecorder is the mock recorder for MockAppConfigDataAPI.
type MockAppConfigDataAPIMockRecorder struct {
	mock *MockAppConfigDataAPI
}

// NewMockAppConfigDataAPI creates a new mock instance.
func NewMockAppConfigDataAPI(ctrl *gomock.Controller) *MockAppConfigDataAPI {
	mock := &MockAppConfigDataAPI{ctrl: ctrl}
	mock.recorder = &MockAppConfigDataAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAppConfigDataAPI) EXPECT() *MockAppConfigDataAPIMockRecorder {
	return m.recorder
}

// GetLatestConfiguration mocks base method.
func (m *MockAppConfigDataAPI) GetLatestConfiguration(arg0 *appconfigdata.GetLatestConfigurationInput) (*appconfigdata.GetLatestConfigurationOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLatestConfiguration", arg0)
	ret0, _ := ret[0].(*appconfigdata.GetLatestConfigurationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLatestConfiguration indicates an expected call of GetLatestConfiguration.
func (mr *MockAppConfigDataAPIMockRecorder) GetLatestConfiguration(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLatestConfiguration", reflect.TypeOf((*MockAppConfigDataAPI)(nil).GetLatestConfiguration), arg0)
}

// GetLatestConfigurationRequest mocks base method.
func (m *MockAppConfigDataAPI) GetLatestConfigurationRequest(arg0 *appconfigdata.GetLatestConfigurationInput) (*request.Request, *appconfigdata.GetLatestConfigurationOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLatestConfigurationRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*appconfigdata.GetLatestConfigurationOutput)
	return ret0, ret1
}

// GetLatestConfigurationRequest indicates an expected call of GetLatestConfigurationRequest.
func (mr *MockAppConfigDataAPIMockRecorder) GetLatestConfigurationRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLatestConfigurationRequest", reflect.TypeOf((*MockAppConfigDataAPI)(nil).GetLatestConfigurationRequest), arg0)
}

// GetLatestConfigurationWithContext mocks base method.
func (m *MockAppConfigDataAPI) GetLatestConfigurationWithContext(arg0 context.Context, arg1 *appconfigdata.GetLatestConfigurationInput, arg2 ...request.Option) (*appconfigdata.GetLatestConfigurationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetLatestConfigurationWithContext", varargs...)
	ret0, _ := ret[0].(*appconfigdata.GetLatestConfigurationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLatestConfigurationWithContext indicates an expected call of GetLatestConfigurationWithContext.
func (mr *MockAppConfigDataAPIMockRecorder) GetLatestConfigurationWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLatestConfigurationWithContext", reflect.TypeOf((*MockAppConfigDataAPI)(nil).GetLatestConfigurationWithContext), varargs...)
}

// StartConfigurationSession mocks base method.
func (m *MockAppConfigDataAPI) StartConfigurationSession(arg0 *appconfigdata.StartConfigurationSessionInput) (*appconfigdata.StartConfigurationSessionOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartConfigurationSession", arg0)
	ret0, _ := ret[0].(*appconfigdata.StartConfigurationSessionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StartConfigurationSession indicates an expected call of StartConfigurationSession.
func (mr *MockAppConfigDataAPIMockRecorder) StartConfigurationSession(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartConfigurationSession", reflect.TypeOf((*MockAppConfigDataAPI)(nil).StartConfigurationSession), arg0)
}

// StartConfigurationSessionRequest mocks base method.
func (m *MockAppConfigDataAPI) StartConfigurationSessionRequest(arg0 *appconfigdata.StartConfigurationSessionInput) (*request.Request, *appconfigdata.StartConfigurationSessionOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartConfigurationSessionRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*appconfigdata.StartConfigurationSessionOutput)
	return ret0, ret1
}

// StartConfigurationSessionRequest indicates an expected call of StartConfigurationSessionRequest.
func (mr *MockAppConfigDataAPIMockRecorder) StartConfigurationSessionRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartConfigurationSessionRequest", reflect.TypeOf((*MockAppConfigDataAPI)(nil).StartConfigurationSessionRequest), arg0)
}

// StartConfigurationSessionWithContext mocks base method.
func (m *MockAppConfigDataAPI) StartConfigurationSessionWithContext(arg0 context.Context, arg1 *appconfigdata.StartConfigurationSessionInput, arg2 ...request.Option) (*appconfigdata.StartConfigurationSessionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "StartConfigurationSessionWithContext", varargs...)
	ret0, _ := ret[0].(*appconfigdata.StartConfigurationSessionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StartConfigurationSessionWithContext indicates an expected call of StartConfigurationSessionWithContext.
func (mr *MockAppConfigDataAPIMockRecorder) StartConfigurationSessionWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartConfigurationSessionWithContext", reflect.TypeOf((*MockAppConfigDataAPI)(nil).StartConfigurationSessionWithContext), varargs...)
}
