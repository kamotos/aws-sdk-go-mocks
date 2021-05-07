// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/aws-sdk-go/service/ioteventsdata/ioteventsdataiface (interfaces: IoTEventsDataAPI)

// Package ioteventsdatamock is a generated GoMock package.
package ioteventsdatamock

import (
	context "context"
	reflect "reflect"

	request "github.com/aws/aws-sdk-go/aws/request"
	ioteventsdata "github.com/aws/aws-sdk-go/service/ioteventsdata"
	gomock "github.com/golang/mock/gomock"
)

// MockIoTEventsDataAPI is a mock of IoTEventsDataAPI interface.
type MockIoTEventsDataAPI struct {
	ctrl     *gomock.Controller
	recorder *MockIoTEventsDataAPIMockRecorder
}

// MockIoTEventsDataAPIMockRecorder is the mock recorder for MockIoTEventsDataAPI.
type MockIoTEventsDataAPIMockRecorder struct {
	mock *MockIoTEventsDataAPI
}

// NewMockIoTEventsDataAPI creates a new mock instance.
func NewMockIoTEventsDataAPI(ctrl *gomock.Controller) *MockIoTEventsDataAPI {
	mock := &MockIoTEventsDataAPI{ctrl: ctrl}
	mock.recorder = &MockIoTEventsDataAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIoTEventsDataAPI) EXPECT() *MockIoTEventsDataAPIMockRecorder {
	return m.recorder
}

// BatchPutMessage mocks base method.
func (m *MockIoTEventsDataAPI) BatchPutMessage(arg0 *ioteventsdata.BatchPutMessageInput) (*ioteventsdata.BatchPutMessageOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BatchPutMessage", arg0)
	ret0, _ := ret[0].(*ioteventsdata.BatchPutMessageOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BatchPutMessage indicates an expected call of BatchPutMessage.
func (mr *MockIoTEventsDataAPIMockRecorder) BatchPutMessage(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchPutMessage", reflect.TypeOf((*MockIoTEventsDataAPI)(nil).BatchPutMessage), arg0)
}

// BatchPutMessageRequest mocks base method.
func (m *MockIoTEventsDataAPI) BatchPutMessageRequest(arg0 *ioteventsdata.BatchPutMessageInput) (*request.Request, *ioteventsdata.BatchPutMessageOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BatchPutMessageRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*ioteventsdata.BatchPutMessageOutput)
	return ret0, ret1
}

// BatchPutMessageRequest indicates an expected call of BatchPutMessageRequest.
func (mr *MockIoTEventsDataAPIMockRecorder) BatchPutMessageRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchPutMessageRequest", reflect.TypeOf((*MockIoTEventsDataAPI)(nil).BatchPutMessageRequest), arg0)
}

// BatchPutMessageWithContext mocks base method.
func (m *MockIoTEventsDataAPI) BatchPutMessageWithContext(arg0 context.Context, arg1 *ioteventsdata.BatchPutMessageInput, arg2 ...request.Option) (*ioteventsdata.BatchPutMessageOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BatchPutMessageWithContext", varargs...)
	ret0, _ := ret[0].(*ioteventsdata.BatchPutMessageOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BatchPutMessageWithContext indicates an expected call of BatchPutMessageWithContext.
func (mr *MockIoTEventsDataAPIMockRecorder) BatchPutMessageWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchPutMessageWithContext", reflect.TypeOf((*MockIoTEventsDataAPI)(nil).BatchPutMessageWithContext), varargs...)
}

// BatchUpdateDetector mocks base method.
func (m *MockIoTEventsDataAPI) BatchUpdateDetector(arg0 *ioteventsdata.BatchUpdateDetectorInput) (*ioteventsdata.BatchUpdateDetectorOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BatchUpdateDetector", arg0)
	ret0, _ := ret[0].(*ioteventsdata.BatchUpdateDetectorOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BatchUpdateDetector indicates an expected call of BatchUpdateDetector.
func (mr *MockIoTEventsDataAPIMockRecorder) BatchUpdateDetector(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchUpdateDetector", reflect.TypeOf((*MockIoTEventsDataAPI)(nil).BatchUpdateDetector), arg0)
}

// BatchUpdateDetectorRequest mocks base method.
func (m *MockIoTEventsDataAPI) BatchUpdateDetectorRequest(arg0 *ioteventsdata.BatchUpdateDetectorInput) (*request.Request, *ioteventsdata.BatchUpdateDetectorOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BatchUpdateDetectorRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*ioteventsdata.BatchUpdateDetectorOutput)
	return ret0, ret1
}

// BatchUpdateDetectorRequest indicates an expected call of BatchUpdateDetectorRequest.
func (mr *MockIoTEventsDataAPIMockRecorder) BatchUpdateDetectorRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchUpdateDetectorRequest", reflect.TypeOf((*MockIoTEventsDataAPI)(nil).BatchUpdateDetectorRequest), arg0)
}

// BatchUpdateDetectorWithContext mocks base method.
func (m *MockIoTEventsDataAPI) BatchUpdateDetectorWithContext(arg0 context.Context, arg1 *ioteventsdata.BatchUpdateDetectorInput, arg2 ...request.Option) (*ioteventsdata.BatchUpdateDetectorOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BatchUpdateDetectorWithContext", varargs...)
	ret0, _ := ret[0].(*ioteventsdata.BatchUpdateDetectorOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BatchUpdateDetectorWithContext indicates an expected call of BatchUpdateDetectorWithContext.
func (mr *MockIoTEventsDataAPIMockRecorder) BatchUpdateDetectorWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchUpdateDetectorWithContext", reflect.TypeOf((*MockIoTEventsDataAPI)(nil).BatchUpdateDetectorWithContext), varargs...)
}

// DescribeDetector mocks base method.
func (m *MockIoTEventsDataAPI) DescribeDetector(arg0 *ioteventsdata.DescribeDetectorInput) (*ioteventsdata.DescribeDetectorOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeDetector", arg0)
	ret0, _ := ret[0].(*ioteventsdata.DescribeDetectorOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeDetector indicates an expected call of DescribeDetector.
func (mr *MockIoTEventsDataAPIMockRecorder) DescribeDetector(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeDetector", reflect.TypeOf((*MockIoTEventsDataAPI)(nil).DescribeDetector), arg0)
}

// DescribeDetectorRequest mocks base method.
func (m *MockIoTEventsDataAPI) DescribeDetectorRequest(arg0 *ioteventsdata.DescribeDetectorInput) (*request.Request, *ioteventsdata.DescribeDetectorOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeDetectorRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*ioteventsdata.DescribeDetectorOutput)
	return ret0, ret1
}

// DescribeDetectorRequest indicates an expected call of DescribeDetectorRequest.
func (mr *MockIoTEventsDataAPIMockRecorder) DescribeDetectorRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeDetectorRequest", reflect.TypeOf((*MockIoTEventsDataAPI)(nil).DescribeDetectorRequest), arg0)
}

// DescribeDetectorWithContext mocks base method.
func (m *MockIoTEventsDataAPI) DescribeDetectorWithContext(arg0 context.Context, arg1 *ioteventsdata.DescribeDetectorInput, arg2 ...request.Option) (*ioteventsdata.DescribeDetectorOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeDetectorWithContext", varargs...)
	ret0, _ := ret[0].(*ioteventsdata.DescribeDetectorOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeDetectorWithContext indicates an expected call of DescribeDetectorWithContext.
func (mr *MockIoTEventsDataAPIMockRecorder) DescribeDetectorWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeDetectorWithContext", reflect.TypeOf((*MockIoTEventsDataAPI)(nil).DescribeDetectorWithContext), varargs...)
}

// ListDetectors mocks base method.
func (m *MockIoTEventsDataAPI) ListDetectors(arg0 *ioteventsdata.ListDetectorsInput) (*ioteventsdata.ListDetectorsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListDetectors", arg0)
	ret0, _ := ret[0].(*ioteventsdata.ListDetectorsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListDetectors indicates an expected call of ListDetectors.
func (mr *MockIoTEventsDataAPIMockRecorder) ListDetectors(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDetectors", reflect.TypeOf((*MockIoTEventsDataAPI)(nil).ListDetectors), arg0)
}

// ListDetectorsRequest mocks base method.
func (m *MockIoTEventsDataAPI) ListDetectorsRequest(arg0 *ioteventsdata.ListDetectorsInput) (*request.Request, *ioteventsdata.ListDetectorsOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListDetectorsRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*ioteventsdata.ListDetectorsOutput)
	return ret0, ret1
}

// ListDetectorsRequest indicates an expected call of ListDetectorsRequest.
func (mr *MockIoTEventsDataAPIMockRecorder) ListDetectorsRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDetectorsRequest", reflect.TypeOf((*MockIoTEventsDataAPI)(nil).ListDetectorsRequest), arg0)
}

// ListDetectorsWithContext mocks base method.
func (m *MockIoTEventsDataAPI) ListDetectorsWithContext(arg0 context.Context, arg1 *ioteventsdata.ListDetectorsInput, arg2 ...request.Option) (*ioteventsdata.ListDetectorsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListDetectorsWithContext", varargs...)
	ret0, _ := ret[0].(*ioteventsdata.ListDetectorsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListDetectorsWithContext indicates an expected call of ListDetectorsWithContext.
func (mr *MockIoTEventsDataAPIMockRecorder) ListDetectorsWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDetectorsWithContext", reflect.TypeOf((*MockIoTEventsDataAPI)(nil).ListDetectorsWithContext), varargs...)
}
