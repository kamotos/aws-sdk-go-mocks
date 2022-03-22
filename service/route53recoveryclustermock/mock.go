// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/aws-sdk-go/service/route53recoverycluster/route53recoveryclusteriface (interfaces: Route53RecoveryClusterAPI)

// Package route53recoveryclustermock is a generated GoMock package.
package route53recoveryclustermock

import (
	context "context"
	reflect "reflect"

	request "github.com/aws/aws-sdk-go/aws/request"
	route53recoverycluster "github.com/aws/aws-sdk-go/service/route53recoverycluster"
	gomock "github.com/golang/mock/gomock"
)

// MockRoute53RecoveryClusterAPI is a mock of Route53RecoveryClusterAPI interface.
type MockRoute53RecoveryClusterAPI struct {
	ctrl     *gomock.Controller
	recorder *MockRoute53RecoveryClusterAPIMockRecorder
}

// MockRoute53RecoveryClusterAPIMockRecorder is the mock recorder for MockRoute53RecoveryClusterAPI.
type MockRoute53RecoveryClusterAPIMockRecorder struct {
	mock *MockRoute53RecoveryClusterAPI
}

// NewMockRoute53RecoveryClusterAPI creates a new mock instance.
func NewMockRoute53RecoveryClusterAPI(ctrl *gomock.Controller) *MockRoute53RecoveryClusterAPI {
	mock := &MockRoute53RecoveryClusterAPI{ctrl: ctrl}
	mock.recorder = &MockRoute53RecoveryClusterAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRoute53RecoveryClusterAPI) EXPECT() *MockRoute53RecoveryClusterAPIMockRecorder {
	return m.recorder
}

// GetRoutingControlState mocks base method.
func (m *MockRoute53RecoveryClusterAPI) GetRoutingControlState(arg0 *route53recoverycluster.GetRoutingControlStateInput) (*route53recoverycluster.GetRoutingControlStateOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRoutingControlState", arg0)
	ret0, _ := ret[0].(*route53recoverycluster.GetRoutingControlStateOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRoutingControlState indicates an expected call of GetRoutingControlState.
func (mr *MockRoute53RecoveryClusterAPIMockRecorder) GetRoutingControlState(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRoutingControlState", reflect.TypeOf((*MockRoute53RecoveryClusterAPI)(nil).GetRoutingControlState), arg0)
}

// GetRoutingControlStateRequest mocks base method.
func (m *MockRoute53RecoveryClusterAPI) GetRoutingControlStateRequest(arg0 *route53recoverycluster.GetRoutingControlStateInput) (*request.Request, *route53recoverycluster.GetRoutingControlStateOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRoutingControlStateRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*route53recoverycluster.GetRoutingControlStateOutput)
	return ret0, ret1
}

// GetRoutingControlStateRequest indicates an expected call of GetRoutingControlStateRequest.
func (mr *MockRoute53RecoveryClusterAPIMockRecorder) GetRoutingControlStateRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRoutingControlStateRequest", reflect.TypeOf((*MockRoute53RecoveryClusterAPI)(nil).GetRoutingControlStateRequest), arg0)
}

// GetRoutingControlStateWithContext mocks base method.
func (m *MockRoute53RecoveryClusterAPI) GetRoutingControlStateWithContext(arg0 context.Context, arg1 *route53recoverycluster.GetRoutingControlStateInput, arg2 ...request.Option) (*route53recoverycluster.GetRoutingControlStateOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetRoutingControlStateWithContext", varargs...)
	ret0, _ := ret[0].(*route53recoverycluster.GetRoutingControlStateOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRoutingControlStateWithContext indicates an expected call of GetRoutingControlStateWithContext.
func (mr *MockRoute53RecoveryClusterAPIMockRecorder) GetRoutingControlStateWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRoutingControlStateWithContext", reflect.TypeOf((*MockRoute53RecoveryClusterAPI)(nil).GetRoutingControlStateWithContext), varargs...)
}

// UpdateRoutingControlState mocks base method.
func (m *MockRoute53RecoveryClusterAPI) UpdateRoutingControlState(arg0 *route53recoverycluster.UpdateRoutingControlStateInput) (*route53recoverycluster.UpdateRoutingControlStateOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRoutingControlState", arg0)
	ret0, _ := ret[0].(*route53recoverycluster.UpdateRoutingControlStateOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateRoutingControlState indicates an expected call of UpdateRoutingControlState.
func (mr *MockRoute53RecoveryClusterAPIMockRecorder) UpdateRoutingControlState(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRoutingControlState", reflect.TypeOf((*MockRoute53RecoveryClusterAPI)(nil).UpdateRoutingControlState), arg0)
}

// UpdateRoutingControlStateRequest mocks base method.
func (m *MockRoute53RecoveryClusterAPI) UpdateRoutingControlStateRequest(arg0 *route53recoverycluster.UpdateRoutingControlStateInput) (*request.Request, *route53recoverycluster.UpdateRoutingControlStateOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRoutingControlStateRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*route53recoverycluster.UpdateRoutingControlStateOutput)
	return ret0, ret1
}

// UpdateRoutingControlStateRequest indicates an expected call of UpdateRoutingControlStateRequest.
func (mr *MockRoute53RecoveryClusterAPIMockRecorder) UpdateRoutingControlStateRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRoutingControlStateRequest", reflect.TypeOf((*MockRoute53RecoveryClusterAPI)(nil).UpdateRoutingControlStateRequest), arg0)
}

// UpdateRoutingControlStateWithContext mocks base method.
func (m *MockRoute53RecoveryClusterAPI) UpdateRoutingControlStateWithContext(arg0 context.Context, arg1 *route53recoverycluster.UpdateRoutingControlStateInput, arg2 ...request.Option) (*route53recoverycluster.UpdateRoutingControlStateOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateRoutingControlStateWithContext", varargs...)
	ret0, _ := ret[0].(*route53recoverycluster.UpdateRoutingControlStateOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateRoutingControlStateWithContext indicates an expected call of UpdateRoutingControlStateWithContext.
func (mr *MockRoute53RecoveryClusterAPIMockRecorder) UpdateRoutingControlStateWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRoutingControlStateWithContext", reflect.TypeOf((*MockRoute53RecoveryClusterAPI)(nil).UpdateRoutingControlStateWithContext), varargs...)
}

// UpdateRoutingControlStates mocks base method.
func (m *MockRoute53RecoveryClusterAPI) UpdateRoutingControlStates(arg0 *route53recoverycluster.UpdateRoutingControlStatesInput) (*route53recoverycluster.UpdateRoutingControlStatesOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRoutingControlStates", arg0)
	ret0, _ := ret[0].(*route53recoverycluster.UpdateRoutingControlStatesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateRoutingControlStates indicates an expected call of UpdateRoutingControlStates.
func (mr *MockRoute53RecoveryClusterAPIMockRecorder) UpdateRoutingControlStates(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRoutingControlStates", reflect.TypeOf((*MockRoute53RecoveryClusterAPI)(nil).UpdateRoutingControlStates), arg0)
}

// UpdateRoutingControlStatesRequest mocks base method.
func (m *MockRoute53RecoveryClusterAPI) UpdateRoutingControlStatesRequest(arg0 *route53recoverycluster.UpdateRoutingControlStatesInput) (*request.Request, *route53recoverycluster.UpdateRoutingControlStatesOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRoutingControlStatesRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*route53recoverycluster.UpdateRoutingControlStatesOutput)
	return ret0, ret1
}

// UpdateRoutingControlStatesRequest indicates an expected call of UpdateRoutingControlStatesRequest.
func (mr *MockRoute53RecoveryClusterAPIMockRecorder) UpdateRoutingControlStatesRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRoutingControlStatesRequest", reflect.TypeOf((*MockRoute53RecoveryClusterAPI)(nil).UpdateRoutingControlStatesRequest), arg0)
}

// UpdateRoutingControlStatesWithContext mocks base method.
func (m *MockRoute53RecoveryClusterAPI) UpdateRoutingControlStatesWithContext(arg0 context.Context, arg1 *route53recoverycluster.UpdateRoutingControlStatesInput, arg2 ...request.Option) (*route53recoverycluster.UpdateRoutingControlStatesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateRoutingControlStatesWithContext", varargs...)
	ret0, _ := ret[0].(*route53recoverycluster.UpdateRoutingControlStatesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateRoutingControlStatesWithContext indicates an expected call of UpdateRoutingControlStatesWithContext.
func (mr *MockRoute53RecoveryClusterAPIMockRecorder) UpdateRoutingControlStatesWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRoutingControlStatesWithContext", reflect.TypeOf((*MockRoute53RecoveryClusterAPI)(nil).UpdateRoutingControlStatesWithContext), varargs...)
}
