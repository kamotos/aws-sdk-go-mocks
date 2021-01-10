// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/aws-sdk-go/service/iotfleethub/iotfleethubiface (interfaces: IoTFleetHubAPI)

// Package iotfleethubmock is a generated GoMock package.
package iotfleethubmock

import (
	context "context"
	request "github.com/aws/aws-sdk-go/aws/request"
	iotfleethub "github.com/aws/aws-sdk-go/service/iotfleethub"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockIoTFleetHubAPI is a mock of IoTFleetHubAPI interface
type MockIoTFleetHubAPI struct {
	ctrl     *gomock.Controller
	recorder *MockIoTFleetHubAPIMockRecorder
}

// MockIoTFleetHubAPIMockRecorder is the mock recorder for MockIoTFleetHubAPI
type MockIoTFleetHubAPIMockRecorder struct {
	mock *MockIoTFleetHubAPI
}

// NewMockIoTFleetHubAPI creates a new mock instance
func NewMockIoTFleetHubAPI(ctrl *gomock.Controller) *MockIoTFleetHubAPI {
	mock := &MockIoTFleetHubAPI{ctrl: ctrl}
	mock.recorder = &MockIoTFleetHubAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIoTFleetHubAPI) EXPECT() *MockIoTFleetHubAPIMockRecorder {
	return m.recorder
}

// CreateApplication mocks base method
func (m *MockIoTFleetHubAPI) CreateApplication(arg0 *iotfleethub.CreateApplicationInput) (*iotfleethub.CreateApplicationOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateApplication", arg0)
	ret0, _ := ret[0].(*iotfleethub.CreateApplicationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateApplication indicates an expected call of CreateApplication
func (mr *MockIoTFleetHubAPIMockRecorder) CreateApplication(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateApplication", reflect.TypeOf((*MockIoTFleetHubAPI)(nil).CreateApplication), arg0)
}

// CreateApplicationRequest mocks base method
func (m *MockIoTFleetHubAPI) CreateApplicationRequest(arg0 *iotfleethub.CreateApplicationInput) (*request.Request, *iotfleethub.CreateApplicationOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateApplicationRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*iotfleethub.CreateApplicationOutput)
	return ret0, ret1
}

// CreateApplicationRequest indicates an expected call of CreateApplicationRequest
func (mr *MockIoTFleetHubAPIMockRecorder) CreateApplicationRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateApplicationRequest", reflect.TypeOf((*MockIoTFleetHubAPI)(nil).CreateApplicationRequest), arg0)
}

// CreateApplicationWithContext mocks base method
func (m *MockIoTFleetHubAPI) CreateApplicationWithContext(arg0 context.Context, arg1 *iotfleethub.CreateApplicationInput, arg2 ...request.Option) (*iotfleethub.CreateApplicationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateApplicationWithContext", varargs...)
	ret0, _ := ret[0].(*iotfleethub.CreateApplicationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateApplicationWithContext indicates an expected call of CreateApplicationWithContext
func (mr *MockIoTFleetHubAPIMockRecorder) CreateApplicationWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateApplicationWithContext", reflect.TypeOf((*MockIoTFleetHubAPI)(nil).CreateApplicationWithContext), varargs...)
}

// DeleteApplication mocks base method
func (m *MockIoTFleetHubAPI) DeleteApplication(arg0 *iotfleethub.DeleteApplicationInput) (*iotfleethub.DeleteApplicationOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteApplication", arg0)
	ret0, _ := ret[0].(*iotfleethub.DeleteApplicationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteApplication indicates an expected call of DeleteApplication
func (mr *MockIoTFleetHubAPIMockRecorder) DeleteApplication(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteApplication", reflect.TypeOf((*MockIoTFleetHubAPI)(nil).DeleteApplication), arg0)
}

// DeleteApplicationRequest mocks base method
func (m *MockIoTFleetHubAPI) DeleteApplicationRequest(arg0 *iotfleethub.DeleteApplicationInput) (*request.Request, *iotfleethub.DeleteApplicationOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteApplicationRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*iotfleethub.DeleteApplicationOutput)
	return ret0, ret1
}

// DeleteApplicationRequest indicates an expected call of DeleteApplicationRequest
func (mr *MockIoTFleetHubAPIMockRecorder) DeleteApplicationRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteApplicationRequest", reflect.TypeOf((*MockIoTFleetHubAPI)(nil).DeleteApplicationRequest), arg0)
}

// DeleteApplicationWithContext mocks base method
func (m *MockIoTFleetHubAPI) DeleteApplicationWithContext(arg0 context.Context, arg1 *iotfleethub.DeleteApplicationInput, arg2 ...request.Option) (*iotfleethub.DeleteApplicationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteApplicationWithContext", varargs...)
	ret0, _ := ret[0].(*iotfleethub.DeleteApplicationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteApplicationWithContext indicates an expected call of DeleteApplicationWithContext
func (mr *MockIoTFleetHubAPIMockRecorder) DeleteApplicationWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteApplicationWithContext", reflect.TypeOf((*MockIoTFleetHubAPI)(nil).DeleteApplicationWithContext), varargs...)
}

// DescribeApplication mocks base method
func (m *MockIoTFleetHubAPI) DescribeApplication(arg0 *iotfleethub.DescribeApplicationInput) (*iotfleethub.DescribeApplicationOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeApplication", arg0)
	ret0, _ := ret[0].(*iotfleethub.DescribeApplicationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeApplication indicates an expected call of DescribeApplication
func (mr *MockIoTFleetHubAPIMockRecorder) DescribeApplication(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeApplication", reflect.TypeOf((*MockIoTFleetHubAPI)(nil).DescribeApplication), arg0)
}

// DescribeApplicationRequest mocks base method
func (m *MockIoTFleetHubAPI) DescribeApplicationRequest(arg0 *iotfleethub.DescribeApplicationInput) (*request.Request, *iotfleethub.DescribeApplicationOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeApplicationRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*iotfleethub.DescribeApplicationOutput)
	return ret0, ret1
}

// DescribeApplicationRequest indicates an expected call of DescribeApplicationRequest
func (mr *MockIoTFleetHubAPIMockRecorder) DescribeApplicationRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeApplicationRequest", reflect.TypeOf((*MockIoTFleetHubAPI)(nil).DescribeApplicationRequest), arg0)
}

// DescribeApplicationWithContext mocks base method
func (m *MockIoTFleetHubAPI) DescribeApplicationWithContext(arg0 context.Context, arg1 *iotfleethub.DescribeApplicationInput, arg2 ...request.Option) (*iotfleethub.DescribeApplicationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeApplicationWithContext", varargs...)
	ret0, _ := ret[0].(*iotfleethub.DescribeApplicationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeApplicationWithContext indicates an expected call of DescribeApplicationWithContext
func (mr *MockIoTFleetHubAPIMockRecorder) DescribeApplicationWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeApplicationWithContext", reflect.TypeOf((*MockIoTFleetHubAPI)(nil).DescribeApplicationWithContext), varargs...)
}

// ListApplications mocks base method
func (m *MockIoTFleetHubAPI) ListApplications(arg0 *iotfleethub.ListApplicationsInput) (*iotfleethub.ListApplicationsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListApplications", arg0)
	ret0, _ := ret[0].(*iotfleethub.ListApplicationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListApplications indicates an expected call of ListApplications
func (mr *MockIoTFleetHubAPIMockRecorder) ListApplications(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListApplications", reflect.TypeOf((*MockIoTFleetHubAPI)(nil).ListApplications), arg0)
}

// ListApplicationsPages mocks base method
func (m *MockIoTFleetHubAPI) ListApplicationsPages(arg0 *iotfleethub.ListApplicationsInput, arg1 func(*iotfleethub.ListApplicationsOutput, bool) bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListApplicationsPages", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListApplicationsPages indicates an expected call of ListApplicationsPages
func (mr *MockIoTFleetHubAPIMockRecorder) ListApplicationsPages(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListApplicationsPages", reflect.TypeOf((*MockIoTFleetHubAPI)(nil).ListApplicationsPages), arg0, arg1)
}

// ListApplicationsPagesWithContext mocks base method
func (m *MockIoTFleetHubAPI) ListApplicationsPagesWithContext(arg0 context.Context, arg1 *iotfleethub.ListApplicationsInput, arg2 func(*iotfleethub.ListApplicationsOutput, bool) bool, arg3 ...request.Option) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListApplicationsPagesWithContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListApplicationsPagesWithContext indicates an expected call of ListApplicationsPagesWithContext
func (mr *MockIoTFleetHubAPIMockRecorder) ListApplicationsPagesWithContext(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListApplicationsPagesWithContext", reflect.TypeOf((*MockIoTFleetHubAPI)(nil).ListApplicationsPagesWithContext), varargs...)
}

// ListApplicationsRequest mocks base method
func (m *MockIoTFleetHubAPI) ListApplicationsRequest(arg0 *iotfleethub.ListApplicationsInput) (*request.Request, *iotfleethub.ListApplicationsOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListApplicationsRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*iotfleethub.ListApplicationsOutput)
	return ret0, ret1
}

// ListApplicationsRequest indicates an expected call of ListApplicationsRequest
func (mr *MockIoTFleetHubAPIMockRecorder) ListApplicationsRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListApplicationsRequest", reflect.TypeOf((*MockIoTFleetHubAPI)(nil).ListApplicationsRequest), arg0)
}

// ListApplicationsWithContext mocks base method
func (m *MockIoTFleetHubAPI) ListApplicationsWithContext(arg0 context.Context, arg1 *iotfleethub.ListApplicationsInput, arg2 ...request.Option) (*iotfleethub.ListApplicationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListApplicationsWithContext", varargs...)
	ret0, _ := ret[0].(*iotfleethub.ListApplicationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListApplicationsWithContext indicates an expected call of ListApplicationsWithContext
func (mr *MockIoTFleetHubAPIMockRecorder) ListApplicationsWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListApplicationsWithContext", reflect.TypeOf((*MockIoTFleetHubAPI)(nil).ListApplicationsWithContext), varargs...)
}

// ListTagsForResource mocks base method
func (m *MockIoTFleetHubAPI) ListTagsForResource(arg0 *iotfleethub.ListTagsForResourceInput) (*iotfleethub.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTagsForResource", arg0)
	ret0, _ := ret[0].(*iotfleethub.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTagsForResource indicates an expected call of ListTagsForResource
func (mr *MockIoTFleetHubAPIMockRecorder) ListTagsForResource(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResource", reflect.TypeOf((*MockIoTFleetHubAPI)(nil).ListTagsForResource), arg0)
}

// ListTagsForResourceRequest mocks base method
func (m *MockIoTFleetHubAPI) ListTagsForResourceRequest(arg0 *iotfleethub.ListTagsForResourceInput) (*request.Request, *iotfleethub.ListTagsForResourceOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTagsForResourceRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*iotfleethub.ListTagsForResourceOutput)
	return ret0, ret1
}

// ListTagsForResourceRequest indicates an expected call of ListTagsForResourceRequest
func (mr *MockIoTFleetHubAPIMockRecorder) ListTagsForResourceRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResourceRequest", reflect.TypeOf((*MockIoTFleetHubAPI)(nil).ListTagsForResourceRequest), arg0)
}

// ListTagsForResourceWithContext mocks base method
func (m *MockIoTFleetHubAPI) ListTagsForResourceWithContext(arg0 context.Context, arg1 *iotfleethub.ListTagsForResourceInput, arg2 ...request.Option) (*iotfleethub.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTagsForResourceWithContext", varargs...)
	ret0, _ := ret[0].(*iotfleethub.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTagsForResourceWithContext indicates an expected call of ListTagsForResourceWithContext
func (mr *MockIoTFleetHubAPIMockRecorder) ListTagsForResourceWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResourceWithContext", reflect.TypeOf((*MockIoTFleetHubAPI)(nil).ListTagsForResourceWithContext), varargs...)
}

// TagResource mocks base method
func (m *MockIoTFleetHubAPI) TagResource(arg0 *iotfleethub.TagResourceInput) (*iotfleethub.TagResourceOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TagResource", arg0)
	ret0, _ := ret[0].(*iotfleethub.TagResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TagResource indicates an expected call of TagResource
func (mr *MockIoTFleetHubAPIMockRecorder) TagResource(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TagResource", reflect.TypeOf((*MockIoTFleetHubAPI)(nil).TagResource), arg0)
}

// TagResourceRequest mocks base method
func (m *MockIoTFleetHubAPI) TagResourceRequest(arg0 *iotfleethub.TagResourceInput) (*request.Request, *iotfleethub.TagResourceOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TagResourceRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*iotfleethub.TagResourceOutput)
	return ret0, ret1
}

// TagResourceRequest indicates an expected call of TagResourceRequest
func (mr *MockIoTFleetHubAPIMockRecorder) TagResourceRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TagResourceRequest", reflect.TypeOf((*MockIoTFleetHubAPI)(nil).TagResourceRequest), arg0)
}

// TagResourceWithContext mocks base method
func (m *MockIoTFleetHubAPI) TagResourceWithContext(arg0 context.Context, arg1 *iotfleethub.TagResourceInput, arg2 ...request.Option) (*iotfleethub.TagResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "TagResourceWithContext", varargs...)
	ret0, _ := ret[0].(*iotfleethub.TagResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TagResourceWithContext indicates an expected call of TagResourceWithContext
func (mr *MockIoTFleetHubAPIMockRecorder) TagResourceWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TagResourceWithContext", reflect.TypeOf((*MockIoTFleetHubAPI)(nil).TagResourceWithContext), varargs...)
}

// UntagResource mocks base method
func (m *MockIoTFleetHubAPI) UntagResource(arg0 *iotfleethub.UntagResourceInput) (*iotfleethub.UntagResourceOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UntagResource", arg0)
	ret0, _ := ret[0].(*iotfleethub.UntagResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UntagResource indicates an expected call of UntagResource
func (mr *MockIoTFleetHubAPIMockRecorder) UntagResource(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UntagResource", reflect.TypeOf((*MockIoTFleetHubAPI)(nil).UntagResource), arg0)
}

// UntagResourceRequest mocks base method
func (m *MockIoTFleetHubAPI) UntagResourceRequest(arg0 *iotfleethub.UntagResourceInput) (*request.Request, *iotfleethub.UntagResourceOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UntagResourceRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*iotfleethub.UntagResourceOutput)
	return ret0, ret1
}

// UntagResourceRequest indicates an expected call of UntagResourceRequest
func (mr *MockIoTFleetHubAPIMockRecorder) UntagResourceRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UntagResourceRequest", reflect.TypeOf((*MockIoTFleetHubAPI)(nil).UntagResourceRequest), arg0)
}

// UntagResourceWithContext mocks base method
func (m *MockIoTFleetHubAPI) UntagResourceWithContext(arg0 context.Context, arg1 *iotfleethub.UntagResourceInput, arg2 ...request.Option) (*iotfleethub.UntagResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UntagResourceWithContext", varargs...)
	ret0, _ := ret[0].(*iotfleethub.UntagResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UntagResourceWithContext indicates an expected call of UntagResourceWithContext
func (mr *MockIoTFleetHubAPIMockRecorder) UntagResourceWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UntagResourceWithContext", reflect.TypeOf((*MockIoTFleetHubAPI)(nil).UntagResourceWithContext), varargs...)
}

// UpdateApplication mocks base method
func (m *MockIoTFleetHubAPI) UpdateApplication(arg0 *iotfleethub.UpdateApplicationInput) (*iotfleethub.UpdateApplicationOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateApplication", arg0)
	ret0, _ := ret[0].(*iotfleethub.UpdateApplicationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateApplication indicates an expected call of UpdateApplication
func (mr *MockIoTFleetHubAPIMockRecorder) UpdateApplication(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateApplication", reflect.TypeOf((*MockIoTFleetHubAPI)(nil).UpdateApplication), arg0)
}

// UpdateApplicationRequest mocks base method
func (m *MockIoTFleetHubAPI) UpdateApplicationRequest(arg0 *iotfleethub.UpdateApplicationInput) (*request.Request, *iotfleethub.UpdateApplicationOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateApplicationRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*iotfleethub.UpdateApplicationOutput)
	return ret0, ret1
}

// UpdateApplicationRequest indicates an expected call of UpdateApplicationRequest
func (mr *MockIoTFleetHubAPIMockRecorder) UpdateApplicationRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateApplicationRequest", reflect.TypeOf((*MockIoTFleetHubAPI)(nil).UpdateApplicationRequest), arg0)
}

// UpdateApplicationWithContext mocks base method
func (m *MockIoTFleetHubAPI) UpdateApplicationWithContext(arg0 context.Context, arg1 *iotfleethub.UpdateApplicationInput, arg2 ...request.Option) (*iotfleethub.UpdateApplicationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateApplicationWithContext", varargs...)
	ret0, _ := ret[0].(*iotfleethub.UpdateApplicationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateApplicationWithContext indicates an expected call of UpdateApplicationWithContext
func (mr *MockIoTFleetHubAPIMockRecorder) UpdateApplicationWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateApplicationWithContext", reflect.TypeOf((*MockIoTFleetHubAPI)(nil).UpdateApplicationWithContext), varargs...)
}