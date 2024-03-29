// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/aws-sdk-go/service/sagemakerfeaturestoreruntime/sagemakerfeaturestoreruntimeiface (interfaces: SageMakerFeatureStoreRuntimeAPI)

// Package sagemakerfeaturestoreruntimemock is a generated GoMock package.
package sagemakerfeaturestoreruntimemock

import (
	context "context"
	reflect "reflect"

	request "github.com/aws/aws-sdk-go/aws/request"
	sagemakerfeaturestoreruntime "github.com/aws/aws-sdk-go/service/sagemakerfeaturestoreruntime"
	gomock "github.com/golang/mock/gomock"
)

// MockSageMakerFeatureStoreRuntimeAPI is a mock of SageMakerFeatureStoreRuntimeAPI interface.
type MockSageMakerFeatureStoreRuntimeAPI struct {
	ctrl     *gomock.Controller
	recorder *MockSageMakerFeatureStoreRuntimeAPIMockRecorder
}

// MockSageMakerFeatureStoreRuntimeAPIMockRecorder is the mock recorder for MockSageMakerFeatureStoreRuntimeAPI.
type MockSageMakerFeatureStoreRuntimeAPIMockRecorder struct {
	mock *MockSageMakerFeatureStoreRuntimeAPI
}

// NewMockSageMakerFeatureStoreRuntimeAPI creates a new mock instance.
func NewMockSageMakerFeatureStoreRuntimeAPI(ctrl *gomock.Controller) *MockSageMakerFeatureStoreRuntimeAPI {
	mock := &MockSageMakerFeatureStoreRuntimeAPI{ctrl: ctrl}
	mock.recorder = &MockSageMakerFeatureStoreRuntimeAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSageMakerFeatureStoreRuntimeAPI) EXPECT() *MockSageMakerFeatureStoreRuntimeAPIMockRecorder {
	return m.recorder
}

// BatchGetRecord mocks base method.
func (m *MockSageMakerFeatureStoreRuntimeAPI) BatchGetRecord(arg0 *sagemakerfeaturestoreruntime.BatchGetRecordInput) (*sagemakerfeaturestoreruntime.BatchGetRecordOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BatchGetRecord", arg0)
	ret0, _ := ret[0].(*sagemakerfeaturestoreruntime.BatchGetRecordOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BatchGetRecord indicates an expected call of BatchGetRecord.
func (mr *MockSageMakerFeatureStoreRuntimeAPIMockRecorder) BatchGetRecord(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchGetRecord", reflect.TypeOf((*MockSageMakerFeatureStoreRuntimeAPI)(nil).BatchGetRecord), arg0)
}

// BatchGetRecordRequest mocks base method.
func (m *MockSageMakerFeatureStoreRuntimeAPI) BatchGetRecordRequest(arg0 *sagemakerfeaturestoreruntime.BatchGetRecordInput) (*request.Request, *sagemakerfeaturestoreruntime.BatchGetRecordOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BatchGetRecordRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*sagemakerfeaturestoreruntime.BatchGetRecordOutput)
	return ret0, ret1
}

// BatchGetRecordRequest indicates an expected call of BatchGetRecordRequest.
func (mr *MockSageMakerFeatureStoreRuntimeAPIMockRecorder) BatchGetRecordRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchGetRecordRequest", reflect.TypeOf((*MockSageMakerFeatureStoreRuntimeAPI)(nil).BatchGetRecordRequest), arg0)
}

// BatchGetRecordWithContext mocks base method.
func (m *MockSageMakerFeatureStoreRuntimeAPI) BatchGetRecordWithContext(arg0 context.Context, arg1 *sagemakerfeaturestoreruntime.BatchGetRecordInput, arg2 ...request.Option) (*sagemakerfeaturestoreruntime.BatchGetRecordOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BatchGetRecordWithContext", varargs...)
	ret0, _ := ret[0].(*sagemakerfeaturestoreruntime.BatchGetRecordOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BatchGetRecordWithContext indicates an expected call of BatchGetRecordWithContext.
func (mr *MockSageMakerFeatureStoreRuntimeAPIMockRecorder) BatchGetRecordWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchGetRecordWithContext", reflect.TypeOf((*MockSageMakerFeatureStoreRuntimeAPI)(nil).BatchGetRecordWithContext), varargs...)
}

// DeleteRecord mocks base method.
func (m *MockSageMakerFeatureStoreRuntimeAPI) DeleteRecord(arg0 *sagemakerfeaturestoreruntime.DeleteRecordInput) (*sagemakerfeaturestoreruntime.DeleteRecordOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRecord", arg0)
	ret0, _ := ret[0].(*sagemakerfeaturestoreruntime.DeleteRecordOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteRecord indicates an expected call of DeleteRecord.
func (mr *MockSageMakerFeatureStoreRuntimeAPIMockRecorder) DeleteRecord(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRecord", reflect.TypeOf((*MockSageMakerFeatureStoreRuntimeAPI)(nil).DeleteRecord), arg0)
}

// DeleteRecordRequest mocks base method.
func (m *MockSageMakerFeatureStoreRuntimeAPI) DeleteRecordRequest(arg0 *sagemakerfeaturestoreruntime.DeleteRecordInput) (*request.Request, *sagemakerfeaturestoreruntime.DeleteRecordOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRecordRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*sagemakerfeaturestoreruntime.DeleteRecordOutput)
	return ret0, ret1
}

// DeleteRecordRequest indicates an expected call of DeleteRecordRequest.
func (mr *MockSageMakerFeatureStoreRuntimeAPIMockRecorder) DeleteRecordRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRecordRequest", reflect.TypeOf((*MockSageMakerFeatureStoreRuntimeAPI)(nil).DeleteRecordRequest), arg0)
}

// DeleteRecordWithContext mocks base method.
func (m *MockSageMakerFeatureStoreRuntimeAPI) DeleteRecordWithContext(arg0 context.Context, arg1 *sagemakerfeaturestoreruntime.DeleteRecordInput, arg2 ...request.Option) (*sagemakerfeaturestoreruntime.DeleteRecordOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteRecordWithContext", varargs...)
	ret0, _ := ret[0].(*sagemakerfeaturestoreruntime.DeleteRecordOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteRecordWithContext indicates an expected call of DeleteRecordWithContext.
func (mr *MockSageMakerFeatureStoreRuntimeAPIMockRecorder) DeleteRecordWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRecordWithContext", reflect.TypeOf((*MockSageMakerFeatureStoreRuntimeAPI)(nil).DeleteRecordWithContext), varargs...)
}

// GetRecord mocks base method.
func (m *MockSageMakerFeatureStoreRuntimeAPI) GetRecord(arg0 *sagemakerfeaturestoreruntime.GetRecordInput) (*sagemakerfeaturestoreruntime.GetRecordOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRecord", arg0)
	ret0, _ := ret[0].(*sagemakerfeaturestoreruntime.GetRecordOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRecord indicates an expected call of GetRecord.
func (mr *MockSageMakerFeatureStoreRuntimeAPIMockRecorder) GetRecord(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRecord", reflect.TypeOf((*MockSageMakerFeatureStoreRuntimeAPI)(nil).GetRecord), arg0)
}

// GetRecordRequest mocks base method.
func (m *MockSageMakerFeatureStoreRuntimeAPI) GetRecordRequest(arg0 *sagemakerfeaturestoreruntime.GetRecordInput) (*request.Request, *sagemakerfeaturestoreruntime.GetRecordOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRecordRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*sagemakerfeaturestoreruntime.GetRecordOutput)
	return ret0, ret1
}

// GetRecordRequest indicates an expected call of GetRecordRequest.
func (mr *MockSageMakerFeatureStoreRuntimeAPIMockRecorder) GetRecordRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRecordRequest", reflect.TypeOf((*MockSageMakerFeatureStoreRuntimeAPI)(nil).GetRecordRequest), arg0)
}

// GetRecordWithContext mocks base method.
func (m *MockSageMakerFeatureStoreRuntimeAPI) GetRecordWithContext(arg0 context.Context, arg1 *sagemakerfeaturestoreruntime.GetRecordInput, arg2 ...request.Option) (*sagemakerfeaturestoreruntime.GetRecordOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetRecordWithContext", varargs...)
	ret0, _ := ret[0].(*sagemakerfeaturestoreruntime.GetRecordOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRecordWithContext indicates an expected call of GetRecordWithContext.
func (mr *MockSageMakerFeatureStoreRuntimeAPIMockRecorder) GetRecordWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRecordWithContext", reflect.TypeOf((*MockSageMakerFeatureStoreRuntimeAPI)(nil).GetRecordWithContext), varargs...)
}

// PutRecord mocks base method.
func (m *MockSageMakerFeatureStoreRuntimeAPI) PutRecord(arg0 *sagemakerfeaturestoreruntime.PutRecordInput) (*sagemakerfeaturestoreruntime.PutRecordOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutRecord", arg0)
	ret0, _ := ret[0].(*sagemakerfeaturestoreruntime.PutRecordOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutRecord indicates an expected call of PutRecord.
func (mr *MockSageMakerFeatureStoreRuntimeAPIMockRecorder) PutRecord(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutRecord", reflect.TypeOf((*MockSageMakerFeatureStoreRuntimeAPI)(nil).PutRecord), arg0)
}

// PutRecordRequest mocks base method.
func (m *MockSageMakerFeatureStoreRuntimeAPI) PutRecordRequest(arg0 *sagemakerfeaturestoreruntime.PutRecordInput) (*request.Request, *sagemakerfeaturestoreruntime.PutRecordOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutRecordRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*sagemakerfeaturestoreruntime.PutRecordOutput)
	return ret0, ret1
}

// PutRecordRequest indicates an expected call of PutRecordRequest.
func (mr *MockSageMakerFeatureStoreRuntimeAPIMockRecorder) PutRecordRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutRecordRequest", reflect.TypeOf((*MockSageMakerFeatureStoreRuntimeAPI)(nil).PutRecordRequest), arg0)
}

// PutRecordWithContext mocks base method.
func (m *MockSageMakerFeatureStoreRuntimeAPI) PutRecordWithContext(arg0 context.Context, arg1 *sagemakerfeaturestoreruntime.PutRecordInput, arg2 ...request.Option) (*sagemakerfeaturestoreruntime.PutRecordOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PutRecordWithContext", varargs...)
	ret0, _ := ret[0].(*sagemakerfeaturestoreruntime.PutRecordOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutRecordWithContext indicates an expected call of PutRecordWithContext.
func (mr *MockSageMakerFeatureStoreRuntimeAPIMockRecorder) PutRecordWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutRecordWithContext", reflect.TypeOf((*MockSageMakerFeatureStoreRuntimeAPI)(nil).PutRecordWithContext), varargs...)
}
