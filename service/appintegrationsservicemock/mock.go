// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/aws-sdk-go/service/appintegrationsservice/appintegrationsserviceiface (interfaces: AppIntegrationsServiceAPI)

// Package appintegrationsservicemock is a generated GoMock package.
package appintegrationsservicemock

import (
	context "context"
	reflect "reflect"

	request "github.com/aws/aws-sdk-go/aws/request"
	appintegrationsservice "github.com/aws/aws-sdk-go/service/appintegrationsservice"
	gomock "github.com/golang/mock/gomock"
)

// MockAppIntegrationsServiceAPI is a mock of AppIntegrationsServiceAPI interface.
type MockAppIntegrationsServiceAPI struct {
	ctrl     *gomock.Controller
	recorder *MockAppIntegrationsServiceAPIMockRecorder
}

// MockAppIntegrationsServiceAPIMockRecorder is the mock recorder for MockAppIntegrationsServiceAPI.
type MockAppIntegrationsServiceAPIMockRecorder struct {
	mock *MockAppIntegrationsServiceAPI
}

// NewMockAppIntegrationsServiceAPI creates a new mock instance.
func NewMockAppIntegrationsServiceAPI(ctrl *gomock.Controller) *MockAppIntegrationsServiceAPI {
	mock := &MockAppIntegrationsServiceAPI{ctrl: ctrl}
	mock.recorder = &MockAppIntegrationsServiceAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAppIntegrationsServiceAPI) EXPECT() *MockAppIntegrationsServiceAPIMockRecorder {
	return m.recorder
}

// CreateEventIntegration mocks base method.
func (m *MockAppIntegrationsServiceAPI) CreateEventIntegration(arg0 *appintegrationsservice.CreateEventIntegrationInput) (*appintegrationsservice.CreateEventIntegrationOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEventIntegration", arg0)
	ret0, _ := ret[0].(*appintegrationsservice.CreateEventIntegrationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateEventIntegration indicates an expected call of CreateEventIntegration.
func (mr *MockAppIntegrationsServiceAPIMockRecorder) CreateEventIntegration(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEventIntegration", reflect.TypeOf((*MockAppIntegrationsServiceAPI)(nil).CreateEventIntegration), arg0)
}

// CreateEventIntegrationRequest mocks base method.
func (m *MockAppIntegrationsServiceAPI) CreateEventIntegrationRequest(arg0 *appintegrationsservice.CreateEventIntegrationInput) (*request.Request, *appintegrationsservice.CreateEventIntegrationOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEventIntegrationRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*appintegrationsservice.CreateEventIntegrationOutput)
	return ret0, ret1
}

// CreateEventIntegrationRequest indicates an expected call of CreateEventIntegrationRequest.
func (mr *MockAppIntegrationsServiceAPIMockRecorder) CreateEventIntegrationRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEventIntegrationRequest", reflect.TypeOf((*MockAppIntegrationsServiceAPI)(nil).CreateEventIntegrationRequest), arg0)
}

// CreateEventIntegrationWithContext mocks base method.
func (m *MockAppIntegrationsServiceAPI) CreateEventIntegrationWithContext(arg0 context.Context, arg1 *appintegrationsservice.CreateEventIntegrationInput, arg2 ...request.Option) (*appintegrationsservice.CreateEventIntegrationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateEventIntegrationWithContext", varargs...)
	ret0, _ := ret[0].(*appintegrationsservice.CreateEventIntegrationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateEventIntegrationWithContext indicates an expected call of CreateEventIntegrationWithContext.
func (mr *MockAppIntegrationsServiceAPIMockRecorder) CreateEventIntegrationWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEventIntegrationWithContext", reflect.TypeOf((*MockAppIntegrationsServiceAPI)(nil).CreateEventIntegrationWithContext), varargs...)
}

// DeleteEventIntegration mocks base method.
func (m *MockAppIntegrationsServiceAPI) DeleteEventIntegration(arg0 *appintegrationsservice.DeleteEventIntegrationInput) (*appintegrationsservice.DeleteEventIntegrationOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteEventIntegration", arg0)
	ret0, _ := ret[0].(*appintegrationsservice.DeleteEventIntegrationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteEventIntegration indicates an expected call of DeleteEventIntegration.
func (mr *MockAppIntegrationsServiceAPIMockRecorder) DeleteEventIntegration(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteEventIntegration", reflect.TypeOf((*MockAppIntegrationsServiceAPI)(nil).DeleteEventIntegration), arg0)
}

// DeleteEventIntegrationRequest mocks base method.
func (m *MockAppIntegrationsServiceAPI) DeleteEventIntegrationRequest(arg0 *appintegrationsservice.DeleteEventIntegrationInput) (*request.Request, *appintegrationsservice.DeleteEventIntegrationOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteEventIntegrationRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*appintegrationsservice.DeleteEventIntegrationOutput)
	return ret0, ret1
}

// DeleteEventIntegrationRequest indicates an expected call of DeleteEventIntegrationRequest.
func (mr *MockAppIntegrationsServiceAPIMockRecorder) DeleteEventIntegrationRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteEventIntegrationRequest", reflect.TypeOf((*MockAppIntegrationsServiceAPI)(nil).DeleteEventIntegrationRequest), arg0)
}

// DeleteEventIntegrationWithContext mocks base method.
func (m *MockAppIntegrationsServiceAPI) DeleteEventIntegrationWithContext(arg0 context.Context, arg1 *appintegrationsservice.DeleteEventIntegrationInput, arg2 ...request.Option) (*appintegrationsservice.DeleteEventIntegrationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteEventIntegrationWithContext", varargs...)
	ret0, _ := ret[0].(*appintegrationsservice.DeleteEventIntegrationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteEventIntegrationWithContext indicates an expected call of DeleteEventIntegrationWithContext.
func (mr *MockAppIntegrationsServiceAPIMockRecorder) DeleteEventIntegrationWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteEventIntegrationWithContext", reflect.TypeOf((*MockAppIntegrationsServiceAPI)(nil).DeleteEventIntegrationWithContext), varargs...)
}

// GetEventIntegration mocks base method.
func (m *MockAppIntegrationsServiceAPI) GetEventIntegration(arg0 *appintegrationsservice.GetEventIntegrationInput) (*appintegrationsservice.GetEventIntegrationOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEventIntegration", arg0)
	ret0, _ := ret[0].(*appintegrationsservice.GetEventIntegrationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEventIntegration indicates an expected call of GetEventIntegration.
func (mr *MockAppIntegrationsServiceAPIMockRecorder) GetEventIntegration(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEventIntegration", reflect.TypeOf((*MockAppIntegrationsServiceAPI)(nil).GetEventIntegration), arg0)
}

// GetEventIntegrationRequest mocks base method.
func (m *MockAppIntegrationsServiceAPI) GetEventIntegrationRequest(arg0 *appintegrationsservice.GetEventIntegrationInput) (*request.Request, *appintegrationsservice.GetEventIntegrationOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEventIntegrationRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*appintegrationsservice.GetEventIntegrationOutput)
	return ret0, ret1
}

// GetEventIntegrationRequest indicates an expected call of GetEventIntegrationRequest.
func (mr *MockAppIntegrationsServiceAPIMockRecorder) GetEventIntegrationRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEventIntegrationRequest", reflect.TypeOf((*MockAppIntegrationsServiceAPI)(nil).GetEventIntegrationRequest), arg0)
}

// GetEventIntegrationWithContext mocks base method.
func (m *MockAppIntegrationsServiceAPI) GetEventIntegrationWithContext(arg0 context.Context, arg1 *appintegrationsservice.GetEventIntegrationInput, arg2 ...request.Option) (*appintegrationsservice.GetEventIntegrationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetEventIntegrationWithContext", varargs...)
	ret0, _ := ret[0].(*appintegrationsservice.GetEventIntegrationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEventIntegrationWithContext indicates an expected call of GetEventIntegrationWithContext.
func (mr *MockAppIntegrationsServiceAPIMockRecorder) GetEventIntegrationWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEventIntegrationWithContext", reflect.TypeOf((*MockAppIntegrationsServiceAPI)(nil).GetEventIntegrationWithContext), varargs...)
}

// ListEventIntegrationAssociations mocks base method.
func (m *MockAppIntegrationsServiceAPI) ListEventIntegrationAssociations(arg0 *appintegrationsservice.ListEventIntegrationAssociationsInput) (*appintegrationsservice.ListEventIntegrationAssociationsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListEventIntegrationAssociations", arg0)
	ret0, _ := ret[0].(*appintegrationsservice.ListEventIntegrationAssociationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListEventIntegrationAssociations indicates an expected call of ListEventIntegrationAssociations.
func (mr *MockAppIntegrationsServiceAPIMockRecorder) ListEventIntegrationAssociations(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEventIntegrationAssociations", reflect.TypeOf((*MockAppIntegrationsServiceAPI)(nil).ListEventIntegrationAssociations), arg0)
}

// ListEventIntegrationAssociationsRequest mocks base method.
func (m *MockAppIntegrationsServiceAPI) ListEventIntegrationAssociationsRequest(arg0 *appintegrationsservice.ListEventIntegrationAssociationsInput) (*request.Request, *appintegrationsservice.ListEventIntegrationAssociationsOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListEventIntegrationAssociationsRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*appintegrationsservice.ListEventIntegrationAssociationsOutput)
	return ret0, ret1
}

// ListEventIntegrationAssociationsRequest indicates an expected call of ListEventIntegrationAssociationsRequest.
func (mr *MockAppIntegrationsServiceAPIMockRecorder) ListEventIntegrationAssociationsRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEventIntegrationAssociationsRequest", reflect.TypeOf((*MockAppIntegrationsServiceAPI)(nil).ListEventIntegrationAssociationsRequest), arg0)
}

// ListEventIntegrationAssociationsWithContext mocks base method.
func (m *MockAppIntegrationsServiceAPI) ListEventIntegrationAssociationsWithContext(arg0 context.Context, arg1 *appintegrationsservice.ListEventIntegrationAssociationsInput, arg2 ...request.Option) (*appintegrationsservice.ListEventIntegrationAssociationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListEventIntegrationAssociationsWithContext", varargs...)
	ret0, _ := ret[0].(*appintegrationsservice.ListEventIntegrationAssociationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListEventIntegrationAssociationsWithContext indicates an expected call of ListEventIntegrationAssociationsWithContext.
func (mr *MockAppIntegrationsServiceAPIMockRecorder) ListEventIntegrationAssociationsWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEventIntegrationAssociationsWithContext", reflect.TypeOf((*MockAppIntegrationsServiceAPI)(nil).ListEventIntegrationAssociationsWithContext), varargs...)
}

// ListEventIntegrations mocks base method.
func (m *MockAppIntegrationsServiceAPI) ListEventIntegrations(arg0 *appintegrationsservice.ListEventIntegrationsInput) (*appintegrationsservice.ListEventIntegrationsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListEventIntegrations", arg0)
	ret0, _ := ret[0].(*appintegrationsservice.ListEventIntegrationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListEventIntegrations indicates an expected call of ListEventIntegrations.
func (mr *MockAppIntegrationsServiceAPIMockRecorder) ListEventIntegrations(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEventIntegrations", reflect.TypeOf((*MockAppIntegrationsServiceAPI)(nil).ListEventIntegrations), arg0)
}

// ListEventIntegrationsRequest mocks base method.
func (m *MockAppIntegrationsServiceAPI) ListEventIntegrationsRequest(arg0 *appintegrationsservice.ListEventIntegrationsInput) (*request.Request, *appintegrationsservice.ListEventIntegrationsOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListEventIntegrationsRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*appintegrationsservice.ListEventIntegrationsOutput)
	return ret0, ret1
}

// ListEventIntegrationsRequest indicates an expected call of ListEventIntegrationsRequest.
func (mr *MockAppIntegrationsServiceAPIMockRecorder) ListEventIntegrationsRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEventIntegrationsRequest", reflect.TypeOf((*MockAppIntegrationsServiceAPI)(nil).ListEventIntegrationsRequest), arg0)
}

// ListEventIntegrationsWithContext mocks base method.
func (m *MockAppIntegrationsServiceAPI) ListEventIntegrationsWithContext(arg0 context.Context, arg1 *appintegrationsservice.ListEventIntegrationsInput, arg2 ...request.Option) (*appintegrationsservice.ListEventIntegrationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListEventIntegrationsWithContext", varargs...)
	ret0, _ := ret[0].(*appintegrationsservice.ListEventIntegrationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListEventIntegrationsWithContext indicates an expected call of ListEventIntegrationsWithContext.
func (mr *MockAppIntegrationsServiceAPIMockRecorder) ListEventIntegrationsWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEventIntegrationsWithContext", reflect.TypeOf((*MockAppIntegrationsServiceAPI)(nil).ListEventIntegrationsWithContext), varargs...)
}

// ListTagsForResource mocks base method.
func (m *MockAppIntegrationsServiceAPI) ListTagsForResource(arg0 *appintegrationsservice.ListTagsForResourceInput) (*appintegrationsservice.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTagsForResource", arg0)
	ret0, _ := ret[0].(*appintegrationsservice.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTagsForResource indicates an expected call of ListTagsForResource.
func (mr *MockAppIntegrationsServiceAPIMockRecorder) ListTagsForResource(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResource", reflect.TypeOf((*MockAppIntegrationsServiceAPI)(nil).ListTagsForResource), arg0)
}

// ListTagsForResourceRequest mocks base method.
func (m *MockAppIntegrationsServiceAPI) ListTagsForResourceRequest(arg0 *appintegrationsservice.ListTagsForResourceInput) (*request.Request, *appintegrationsservice.ListTagsForResourceOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTagsForResourceRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*appintegrationsservice.ListTagsForResourceOutput)
	return ret0, ret1
}

// ListTagsForResourceRequest indicates an expected call of ListTagsForResourceRequest.
func (mr *MockAppIntegrationsServiceAPIMockRecorder) ListTagsForResourceRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResourceRequest", reflect.TypeOf((*MockAppIntegrationsServiceAPI)(nil).ListTagsForResourceRequest), arg0)
}

// ListTagsForResourceWithContext mocks base method.
func (m *MockAppIntegrationsServiceAPI) ListTagsForResourceWithContext(arg0 context.Context, arg1 *appintegrationsservice.ListTagsForResourceInput, arg2 ...request.Option) (*appintegrationsservice.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTagsForResourceWithContext", varargs...)
	ret0, _ := ret[0].(*appintegrationsservice.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTagsForResourceWithContext indicates an expected call of ListTagsForResourceWithContext.
func (mr *MockAppIntegrationsServiceAPIMockRecorder) ListTagsForResourceWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResourceWithContext", reflect.TypeOf((*MockAppIntegrationsServiceAPI)(nil).ListTagsForResourceWithContext), varargs...)
}

// TagResource mocks base method.
func (m *MockAppIntegrationsServiceAPI) TagResource(arg0 *appintegrationsservice.TagResourceInput) (*appintegrationsservice.TagResourceOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TagResource", arg0)
	ret0, _ := ret[0].(*appintegrationsservice.TagResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TagResource indicates an expected call of TagResource.
func (mr *MockAppIntegrationsServiceAPIMockRecorder) TagResource(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TagResource", reflect.TypeOf((*MockAppIntegrationsServiceAPI)(nil).TagResource), arg0)
}

// TagResourceRequest mocks base method.
func (m *MockAppIntegrationsServiceAPI) TagResourceRequest(arg0 *appintegrationsservice.TagResourceInput) (*request.Request, *appintegrationsservice.TagResourceOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TagResourceRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*appintegrationsservice.TagResourceOutput)
	return ret0, ret1
}

// TagResourceRequest indicates an expected call of TagResourceRequest.
func (mr *MockAppIntegrationsServiceAPIMockRecorder) TagResourceRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TagResourceRequest", reflect.TypeOf((*MockAppIntegrationsServiceAPI)(nil).TagResourceRequest), arg0)
}

// TagResourceWithContext mocks base method.
func (m *MockAppIntegrationsServiceAPI) TagResourceWithContext(arg0 context.Context, arg1 *appintegrationsservice.TagResourceInput, arg2 ...request.Option) (*appintegrationsservice.TagResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "TagResourceWithContext", varargs...)
	ret0, _ := ret[0].(*appintegrationsservice.TagResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TagResourceWithContext indicates an expected call of TagResourceWithContext.
func (mr *MockAppIntegrationsServiceAPIMockRecorder) TagResourceWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TagResourceWithContext", reflect.TypeOf((*MockAppIntegrationsServiceAPI)(nil).TagResourceWithContext), varargs...)
}

// UntagResource mocks base method.
func (m *MockAppIntegrationsServiceAPI) UntagResource(arg0 *appintegrationsservice.UntagResourceInput) (*appintegrationsservice.UntagResourceOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UntagResource", arg0)
	ret0, _ := ret[0].(*appintegrationsservice.UntagResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UntagResource indicates an expected call of UntagResource.
func (mr *MockAppIntegrationsServiceAPIMockRecorder) UntagResource(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UntagResource", reflect.TypeOf((*MockAppIntegrationsServiceAPI)(nil).UntagResource), arg0)
}

// UntagResourceRequest mocks base method.
func (m *MockAppIntegrationsServiceAPI) UntagResourceRequest(arg0 *appintegrationsservice.UntagResourceInput) (*request.Request, *appintegrationsservice.UntagResourceOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UntagResourceRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*appintegrationsservice.UntagResourceOutput)
	return ret0, ret1
}

// UntagResourceRequest indicates an expected call of UntagResourceRequest.
func (mr *MockAppIntegrationsServiceAPIMockRecorder) UntagResourceRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UntagResourceRequest", reflect.TypeOf((*MockAppIntegrationsServiceAPI)(nil).UntagResourceRequest), arg0)
}

// UntagResourceWithContext mocks base method.
func (m *MockAppIntegrationsServiceAPI) UntagResourceWithContext(arg0 context.Context, arg1 *appintegrationsservice.UntagResourceInput, arg2 ...request.Option) (*appintegrationsservice.UntagResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UntagResourceWithContext", varargs...)
	ret0, _ := ret[0].(*appintegrationsservice.UntagResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UntagResourceWithContext indicates an expected call of UntagResourceWithContext.
func (mr *MockAppIntegrationsServiceAPIMockRecorder) UntagResourceWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UntagResourceWithContext", reflect.TypeOf((*MockAppIntegrationsServiceAPI)(nil).UntagResourceWithContext), varargs...)
}

// UpdateEventIntegration mocks base method.
func (m *MockAppIntegrationsServiceAPI) UpdateEventIntegration(arg0 *appintegrationsservice.UpdateEventIntegrationInput) (*appintegrationsservice.UpdateEventIntegrationOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateEventIntegration", arg0)
	ret0, _ := ret[0].(*appintegrationsservice.UpdateEventIntegrationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateEventIntegration indicates an expected call of UpdateEventIntegration.
func (mr *MockAppIntegrationsServiceAPIMockRecorder) UpdateEventIntegration(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateEventIntegration", reflect.TypeOf((*MockAppIntegrationsServiceAPI)(nil).UpdateEventIntegration), arg0)
}

// UpdateEventIntegrationRequest mocks base method.
func (m *MockAppIntegrationsServiceAPI) UpdateEventIntegrationRequest(arg0 *appintegrationsservice.UpdateEventIntegrationInput) (*request.Request, *appintegrationsservice.UpdateEventIntegrationOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateEventIntegrationRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*appintegrationsservice.UpdateEventIntegrationOutput)
	return ret0, ret1
}

// UpdateEventIntegrationRequest indicates an expected call of UpdateEventIntegrationRequest.
func (mr *MockAppIntegrationsServiceAPIMockRecorder) UpdateEventIntegrationRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateEventIntegrationRequest", reflect.TypeOf((*MockAppIntegrationsServiceAPI)(nil).UpdateEventIntegrationRequest), arg0)
}

// UpdateEventIntegrationWithContext mocks base method.
func (m *MockAppIntegrationsServiceAPI) UpdateEventIntegrationWithContext(arg0 context.Context, arg1 *appintegrationsservice.UpdateEventIntegrationInput, arg2 ...request.Option) (*appintegrationsservice.UpdateEventIntegrationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateEventIntegrationWithContext", varargs...)
	ret0, _ := ret[0].(*appintegrationsservice.UpdateEventIntegrationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateEventIntegrationWithContext indicates an expected call of UpdateEventIntegrationWithContext.
func (mr *MockAppIntegrationsServiceAPIMockRecorder) UpdateEventIntegrationWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateEventIntegrationWithContext", reflect.TypeOf((*MockAppIntegrationsServiceAPI)(nil).UpdateEventIntegrationWithContext), varargs...)
}
