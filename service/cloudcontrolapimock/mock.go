// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/aws-sdk-go/service/cloudcontrolapi/cloudcontrolapiiface (interfaces: CloudControlApiAPI)

// Package cloudcontrolapimock is a generated GoMock package.
package cloudcontrolapimock

import (
	context "context"
	reflect "reflect"

	request "github.com/aws/aws-sdk-go/aws/request"
	cloudcontrolapi "github.com/aws/aws-sdk-go/service/cloudcontrolapi"
	gomock "github.com/golang/mock/gomock"
)

// MockCloudControlApiAPI is a mock of CloudControlApiAPI interface.
type MockCloudControlApiAPI struct {
	ctrl     *gomock.Controller
	recorder *MockCloudControlApiAPIMockRecorder
}

// MockCloudControlApiAPIMockRecorder is the mock recorder for MockCloudControlApiAPI.
type MockCloudControlApiAPIMockRecorder struct {
	mock *MockCloudControlApiAPI
}

// NewMockCloudControlApiAPI creates a new mock instance.
func NewMockCloudControlApiAPI(ctrl *gomock.Controller) *MockCloudControlApiAPI {
	mock := &MockCloudControlApiAPI{ctrl: ctrl}
	mock.recorder = &MockCloudControlApiAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCloudControlApiAPI) EXPECT() *MockCloudControlApiAPIMockRecorder {
	return m.recorder
}

// CancelResourceRequest mocks base method.
func (m *MockCloudControlApiAPI) CancelResourceRequest(arg0 *cloudcontrolapi.CancelResourceRequestInput) (*cloudcontrolapi.CancelResourceRequestOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CancelResourceRequest", arg0)
	ret0, _ := ret[0].(*cloudcontrolapi.CancelResourceRequestOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CancelResourceRequest indicates an expected call of CancelResourceRequest.
func (mr *MockCloudControlApiAPIMockRecorder) CancelResourceRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelResourceRequest", reflect.TypeOf((*MockCloudControlApiAPI)(nil).CancelResourceRequest), arg0)
}

// CancelResourceRequestRequest mocks base method.
func (m *MockCloudControlApiAPI) CancelResourceRequestRequest(arg0 *cloudcontrolapi.CancelResourceRequestInput) (*request.Request, *cloudcontrolapi.CancelResourceRequestOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CancelResourceRequestRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*cloudcontrolapi.CancelResourceRequestOutput)
	return ret0, ret1
}

// CancelResourceRequestRequest indicates an expected call of CancelResourceRequestRequest.
func (mr *MockCloudControlApiAPIMockRecorder) CancelResourceRequestRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelResourceRequestRequest", reflect.TypeOf((*MockCloudControlApiAPI)(nil).CancelResourceRequestRequest), arg0)
}

// CancelResourceRequestWithContext mocks base method.
func (m *MockCloudControlApiAPI) CancelResourceRequestWithContext(arg0 context.Context, arg1 *cloudcontrolapi.CancelResourceRequestInput, arg2 ...request.Option) (*cloudcontrolapi.CancelResourceRequestOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CancelResourceRequestWithContext", varargs...)
	ret0, _ := ret[0].(*cloudcontrolapi.CancelResourceRequestOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CancelResourceRequestWithContext indicates an expected call of CancelResourceRequestWithContext.
func (mr *MockCloudControlApiAPIMockRecorder) CancelResourceRequestWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelResourceRequestWithContext", reflect.TypeOf((*MockCloudControlApiAPI)(nil).CancelResourceRequestWithContext), varargs...)
}

// CreateResource mocks base method.
func (m *MockCloudControlApiAPI) CreateResource(arg0 *cloudcontrolapi.CreateResourceInput) (*cloudcontrolapi.CreateResourceOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateResource", arg0)
	ret0, _ := ret[0].(*cloudcontrolapi.CreateResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateResource indicates an expected call of CreateResource.
func (mr *MockCloudControlApiAPIMockRecorder) CreateResource(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateResource", reflect.TypeOf((*MockCloudControlApiAPI)(nil).CreateResource), arg0)
}

// CreateResourceRequest mocks base method.
func (m *MockCloudControlApiAPI) CreateResourceRequest(arg0 *cloudcontrolapi.CreateResourceInput) (*request.Request, *cloudcontrolapi.CreateResourceOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateResourceRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*cloudcontrolapi.CreateResourceOutput)
	return ret0, ret1
}

// CreateResourceRequest indicates an expected call of CreateResourceRequest.
func (mr *MockCloudControlApiAPIMockRecorder) CreateResourceRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateResourceRequest", reflect.TypeOf((*MockCloudControlApiAPI)(nil).CreateResourceRequest), arg0)
}

// CreateResourceWithContext mocks base method.
func (m *MockCloudControlApiAPI) CreateResourceWithContext(arg0 context.Context, arg1 *cloudcontrolapi.CreateResourceInput, arg2 ...request.Option) (*cloudcontrolapi.CreateResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateResourceWithContext", varargs...)
	ret0, _ := ret[0].(*cloudcontrolapi.CreateResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateResourceWithContext indicates an expected call of CreateResourceWithContext.
func (mr *MockCloudControlApiAPIMockRecorder) CreateResourceWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateResourceWithContext", reflect.TypeOf((*MockCloudControlApiAPI)(nil).CreateResourceWithContext), varargs...)
}

// DeleteResource mocks base method.
func (m *MockCloudControlApiAPI) DeleteResource(arg0 *cloudcontrolapi.DeleteResourceInput) (*cloudcontrolapi.DeleteResourceOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteResource", arg0)
	ret0, _ := ret[0].(*cloudcontrolapi.DeleteResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteResource indicates an expected call of DeleteResource.
func (mr *MockCloudControlApiAPIMockRecorder) DeleteResource(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteResource", reflect.TypeOf((*MockCloudControlApiAPI)(nil).DeleteResource), arg0)
}

// DeleteResourceRequest mocks base method.
func (m *MockCloudControlApiAPI) DeleteResourceRequest(arg0 *cloudcontrolapi.DeleteResourceInput) (*request.Request, *cloudcontrolapi.DeleteResourceOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteResourceRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*cloudcontrolapi.DeleteResourceOutput)
	return ret0, ret1
}

// DeleteResourceRequest indicates an expected call of DeleteResourceRequest.
func (mr *MockCloudControlApiAPIMockRecorder) DeleteResourceRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteResourceRequest", reflect.TypeOf((*MockCloudControlApiAPI)(nil).DeleteResourceRequest), arg0)
}

// DeleteResourceWithContext mocks base method.
func (m *MockCloudControlApiAPI) DeleteResourceWithContext(arg0 context.Context, arg1 *cloudcontrolapi.DeleteResourceInput, arg2 ...request.Option) (*cloudcontrolapi.DeleteResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteResourceWithContext", varargs...)
	ret0, _ := ret[0].(*cloudcontrolapi.DeleteResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteResourceWithContext indicates an expected call of DeleteResourceWithContext.
func (mr *MockCloudControlApiAPIMockRecorder) DeleteResourceWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteResourceWithContext", reflect.TypeOf((*MockCloudControlApiAPI)(nil).DeleteResourceWithContext), varargs...)
}

// GetResource mocks base method.
func (m *MockCloudControlApiAPI) GetResource(arg0 *cloudcontrolapi.GetResourceInput) (*cloudcontrolapi.GetResourceOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetResource", arg0)
	ret0, _ := ret[0].(*cloudcontrolapi.GetResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetResource indicates an expected call of GetResource.
func (mr *MockCloudControlApiAPIMockRecorder) GetResource(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResource", reflect.TypeOf((*MockCloudControlApiAPI)(nil).GetResource), arg0)
}

// GetResourceRequest mocks base method.
func (m *MockCloudControlApiAPI) GetResourceRequest(arg0 *cloudcontrolapi.GetResourceInput) (*request.Request, *cloudcontrolapi.GetResourceOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetResourceRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*cloudcontrolapi.GetResourceOutput)
	return ret0, ret1
}

// GetResourceRequest indicates an expected call of GetResourceRequest.
func (mr *MockCloudControlApiAPIMockRecorder) GetResourceRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResourceRequest", reflect.TypeOf((*MockCloudControlApiAPI)(nil).GetResourceRequest), arg0)
}

// GetResourceRequestStatus mocks base method.
func (m *MockCloudControlApiAPI) GetResourceRequestStatus(arg0 *cloudcontrolapi.GetResourceRequestStatusInput) (*cloudcontrolapi.GetResourceRequestStatusOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetResourceRequestStatus", arg0)
	ret0, _ := ret[0].(*cloudcontrolapi.GetResourceRequestStatusOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetResourceRequestStatus indicates an expected call of GetResourceRequestStatus.
func (mr *MockCloudControlApiAPIMockRecorder) GetResourceRequestStatus(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResourceRequestStatus", reflect.TypeOf((*MockCloudControlApiAPI)(nil).GetResourceRequestStatus), arg0)
}

// GetResourceRequestStatusRequest mocks base method.
func (m *MockCloudControlApiAPI) GetResourceRequestStatusRequest(arg0 *cloudcontrolapi.GetResourceRequestStatusInput) (*request.Request, *cloudcontrolapi.GetResourceRequestStatusOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetResourceRequestStatusRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*cloudcontrolapi.GetResourceRequestStatusOutput)
	return ret0, ret1
}

// GetResourceRequestStatusRequest indicates an expected call of GetResourceRequestStatusRequest.
func (mr *MockCloudControlApiAPIMockRecorder) GetResourceRequestStatusRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResourceRequestStatusRequest", reflect.TypeOf((*MockCloudControlApiAPI)(nil).GetResourceRequestStatusRequest), arg0)
}

// GetResourceRequestStatusWithContext mocks base method.
func (m *MockCloudControlApiAPI) GetResourceRequestStatusWithContext(arg0 context.Context, arg1 *cloudcontrolapi.GetResourceRequestStatusInput, arg2 ...request.Option) (*cloudcontrolapi.GetResourceRequestStatusOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetResourceRequestStatusWithContext", varargs...)
	ret0, _ := ret[0].(*cloudcontrolapi.GetResourceRequestStatusOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetResourceRequestStatusWithContext indicates an expected call of GetResourceRequestStatusWithContext.
func (mr *MockCloudControlApiAPIMockRecorder) GetResourceRequestStatusWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResourceRequestStatusWithContext", reflect.TypeOf((*MockCloudControlApiAPI)(nil).GetResourceRequestStatusWithContext), varargs...)
}

// GetResourceWithContext mocks base method.
func (m *MockCloudControlApiAPI) GetResourceWithContext(arg0 context.Context, arg1 *cloudcontrolapi.GetResourceInput, arg2 ...request.Option) (*cloudcontrolapi.GetResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetResourceWithContext", varargs...)
	ret0, _ := ret[0].(*cloudcontrolapi.GetResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetResourceWithContext indicates an expected call of GetResourceWithContext.
func (mr *MockCloudControlApiAPIMockRecorder) GetResourceWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResourceWithContext", reflect.TypeOf((*MockCloudControlApiAPI)(nil).GetResourceWithContext), varargs...)
}

// ListResourceRequests mocks base method.
func (m *MockCloudControlApiAPI) ListResourceRequests(arg0 *cloudcontrolapi.ListResourceRequestsInput) (*cloudcontrolapi.ListResourceRequestsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListResourceRequests", arg0)
	ret0, _ := ret[0].(*cloudcontrolapi.ListResourceRequestsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListResourceRequests indicates an expected call of ListResourceRequests.
func (mr *MockCloudControlApiAPIMockRecorder) ListResourceRequests(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListResourceRequests", reflect.TypeOf((*MockCloudControlApiAPI)(nil).ListResourceRequests), arg0)
}

// ListResourceRequestsPages mocks base method.
func (m *MockCloudControlApiAPI) ListResourceRequestsPages(arg0 *cloudcontrolapi.ListResourceRequestsInput, arg1 func(*cloudcontrolapi.ListResourceRequestsOutput, bool) bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListResourceRequestsPages", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListResourceRequestsPages indicates an expected call of ListResourceRequestsPages.
func (mr *MockCloudControlApiAPIMockRecorder) ListResourceRequestsPages(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListResourceRequestsPages", reflect.TypeOf((*MockCloudControlApiAPI)(nil).ListResourceRequestsPages), arg0, arg1)
}

// ListResourceRequestsPagesWithContext mocks base method.
func (m *MockCloudControlApiAPI) ListResourceRequestsPagesWithContext(arg0 context.Context, arg1 *cloudcontrolapi.ListResourceRequestsInput, arg2 func(*cloudcontrolapi.ListResourceRequestsOutput, bool) bool, arg3 ...request.Option) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListResourceRequestsPagesWithContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListResourceRequestsPagesWithContext indicates an expected call of ListResourceRequestsPagesWithContext.
func (mr *MockCloudControlApiAPIMockRecorder) ListResourceRequestsPagesWithContext(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListResourceRequestsPagesWithContext", reflect.TypeOf((*MockCloudControlApiAPI)(nil).ListResourceRequestsPagesWithContext), varargs...)
}

// ListResourceRequestsRequest mocks base method.
func (m *MockCloudControlApiAPI) ListResourceRequestsRequest(arg0 *cloudcontrolapi.ListResourceRequestsInput) (*request.Request, *cloudcontrolapi.ListResourceRequestsOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListResourceRequestsRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*cloudcontrolapi.ListResourceRequestsOutput)
	return ret0, ret1
}

// ListResourceRequestsRequest indicates an expected call of ListResourceRequestsRequest.
func (mr *MockCloudControlApiAPIMockRecorder) ListResourceRequestsRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListResourceRequestsRequest", reflect.TypeOf((*MockCloudControlApiAPI)(nil).ListResourceRequestsRequest), arg0)
}

// ListResourceRequestsWithContext mocks base method.
func (m *MockCloudControlApiAPI) ListResourceRequestsWithContext(arg0 context.Context, arg1 *cloudcontrolapi.ListResourceRequestsInput, arg2 ...request.Option) (*cloudcontrolapi.ListResourceRequestsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListResourceRequestsWithContext", varargs...)
	ret0, _ := ret[0].(*cloudcontrolapi.ListResourceRequestsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListResourceRequestsWithContext indicates an expected call of ListResourceRequestsWithContext.
func (mr *MockCloudControlApiAPIMockRecorder) ListResourceRequestsWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListResourceRequestsWithContext", reflect.TypeOf((*MockCloudControlApiAPI)(nil).ListResourceRequestsWithContext), varargs...)
}

// ListResources mocks base method.
func (m *MockCloudControlApiAPI) ListResources(arg0 *cloudcontrolapi.ListResourcesInput) (*cloudcontrolapi.ListResourcesOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListResources", arg0)
	ret0, _ := ret[0].(*cloudcontrolapi.ListResourcesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListResources indicates an expected call of ListResources.
func (mr *MockCloudControlApiAPIMockRecorder) ListResources(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListResources", reflect.TypeOf((*MockCloudControlApiAPI)(nil).ListResources), arg0)
}

// ListResourcesPages mocks base method.
func (m *MockCloudControlApiAPI) ListResourcesPages(arg0 *cloudcontrolapi.ListResourcesInput, arg1 func(*cloudcontrolapi.ListResourcesOutput, bool) bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListResourcesPages", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListResourcesPages indicates an expected call of ListResourcesPages.
func (mr *MockCloudControlApiAPIMockRecorder) ListResourcesPages(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListResourcesPages", reflect.TypeOf((*MockCloudControlApiAPI)(nil).ListResourcesPages), arg0, arg1)
}

// ListResourcesPagesWithContext mocks base method.
func (m *MockCloudControlApiAPI) ListResourcesPagesWithContext(arg0 context.Context, arg1 *cloudcontrolapi.ListResourcesInput, arg2 func(*cloudcontrolapi.ListResourcesOutput, bool) bool, arg3 ...request.Option) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListResourcesPagesWithContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListResourcesPagesWithContext indicates an expected call of ListResourcesPagesWithContext.
func (mr *MockCloudControlApiAPIMockRecorder) ListResourcesPagesWithContext(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListResourcesPagesWithContext", reflect.TypeOf((*MockCloudControlApiAPI)(nil).ListResourcesPagesWithContext), varargs...)
}

// ListResourcesRequest mocks base method.
func (m *MockCloudControlApiAPI) ListResourcesRequest(arg0 *cloudcontrolapi.ListResourcesInput) (*request.Request, *cloudcontrolapi.ListResourcesOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListResourcesRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*cloudcontrolapi.ListResourcesOutput)
	return ret0, ret1
}

// ListResourcesRequest indicates an expected call of ListResourcesRequest.
func (mr *MockCloudControlApiAPIMockRecorder) ListResourcesRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListResourcesRequest", reflect.TypeOf((*MockCloudControlApiAPI)(nil).ListResourcesRequest), arg0)
}

// ListResourcesWithContext mocks base method.
func (m *MockCloudControlApiAPI) ListResourcesWithContext(arg0 context.Context, arg1 *cloudcontrolapi.ListResourcesInput, arg2 ...request.Option) (*cloudcontrolapi.ListResourcesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListResourcesWithContext", varargs...)
	ret0, _ := ret[0].(*cloudcontrolapi.ListResourcesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListResourcesWithContext indicates an expected call of ListResourcesWithContext.
func (mr *MockCloudControlApiAPIMockRecorder) ListResourcesWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListResourcesWithContext", reflect.TypeOf((*MockCloudControlApiAPI)(nil).ListResourcesWithContext), varargs...)
}

// UpdateResource mocks base method.
func (m *MockCloudControlApiAPI) UpdateResource(arg0 *cloudcontrolapi.UpdateResourceInput) (*cloudcontrolapi.UpdateResourceOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateResource", arg0)
	ret0, _ := ret[0].(*cloudcontrolapi.UpdateResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateResource indicates an expected call of UpdateResource.
func (mr *MockCloudControlApiAPIMockRecorder) UpdateResource(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateResource", reflect.TypeOf((*MockCloudControlApiAPI)(nil).UpdateResource), arg0)
}

// UpdateResourceRequest mocks base method.
func (m *MockCloudControlApiAPI) UpdateResourceRequest(arg0 *cloudcontrolapi.UpdateResourceInput) (*request.Request, *cloudcontrolapi.UpdateResourceOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateResourceRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*cloudcontrolapi.UpdateResourceOutput)
	return ret0, ret1
}

// UpdateResourceRequest indicates an expected call of UpdateResourceRequest.
func (mr *MockCloudControlApiAPIMockRecorder) UpdateResourceRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateResourceRequest", reflect.TypeOf((*MockCloudControlApiAPI)(nil).UpdateResourceRequest), arg0)
}

// UpdateResourceWithContext mocks base method.
func (m *MockCloudControlApiAPI) UpdateResourceWithContext(arg0 context.Context, arg1 *cloudcontrolapi.UpdateResourceInput, arg2 ...request.Option) (*cloudcontrolapi.UpdateResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateResourceWithContext", varargs...)
	ret0, _ := ret[0].(*cloudcontrolapi.UpdateResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateResourceWithContext indicates an expected call of UpdateResourceWithContext.
func (mr *MockCloudControlApiAPIMockRecorder) UpdateResourceWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateResourceWithContext", reflect.TypeOf((*MockCloudControlApiAPI)(nil).UpdateResourceWithContext), varargs...)
}

// WaitUntilResourceRequestSuccess mocks base method.
func (m *MockCloudControlApiAPI) WaitUntilResourceRequestSuccess(arg0 *cloudcontrolapi.GetResourceRequestStatusInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitUntilResourceRequestSuccess", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// WaitUntilResourceRequestSuccess indicates an expected call of WaitUntilResourceRequestSuccess.
func (mr *MockCloudControlApiAPIMockRecorder) WaitUntilResourceRequestSuccess(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitUntilResourceRequestSuccess", reflect.TypeOf((*MockCloudControlApiAPI)(nil).WaitUntilResourceRequestSuccess), arg0)
}

// WaitUntilResourceRequestSuccessWithContext mocks base method.
func (m *MockCloudControlApiAPI) WaitUntilResourceRequestSuccessWithContext(arg0 context.Context, arg1 *cloudcontrolapi.GetResourceRequestStatusInput, arg2 ...request.WaiterOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "WaitUntilResourceRequestSuccessWithContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// WaitUntilResourceRequestSuccessWithContext indicates an expected call of WaitUntilResourceRequestSuccessWithContext.
func (mr *MockCloudControlApiAPIMockRecorder) WaitUntilResourceRequestSuccessWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitUntilResourceRequestSuccessWithContext", reflect.TypeOf((*MockCloudControlApiAPI)(nil).WaitUntilResourceRequestSuccessWithContext), varargs...)
}
