// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/aws-sdk-go/service/honeycode/honeycodeiface (interfaces: HoneycodeAPI)

// Package honeycodemock is a generated GoMock package.
package honeycodemock

import (
	context "context"
	reflect "reflect"

	request "github.com/aws/aws-sdk-go/aws/request"
	honeycode "github.com/aws/aws-sdk-go/service/honeycode"
	gomock "github.com/golang/mock/gomock"
)

// MockHoneycodeAPI is a mock of HoneycodeAPI interface.
type MockHoneycodeAPI struct {
	ctrl     *gomock.Controller
	recorder *MockHoneycodeAPIMockRecorder
}

// MockHoneycodeAPIMockRecorder is the mock recorder for MockHoneycodeAPI.
type MockHoneycodeAPIMockRecorder struct {
	mock *MockHoneycodeAPI
}

// NewMockHoneycodeAPI creates a new mock instance.
func NewMockHoneycodeAPI(ctrl *gomock.Controller) *MockHoneycodeAPI {
	mock := &MockHoneycodeAPI{ctrl: ctrl}
	mock.recorder = &MockHoneycodeAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHoneycodeAPI) EXPECT() *MockHoneycodeAPIMockRecorder {
	return m.recorder
}

// BatchCreateTableRows mocks base method.
func (m *MockHoneycodeAPI) BatchCreateTableRows(arg0 *honeycode.BatchCreateTableRowsInput) (*honeycode.BatchCreateTableRowsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BatchCreateTableRows", arg0)
	ret0, _ := ret[0].(*honeycode.BatchCreateTableRowsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BatchCreateTableRows indicates an expected call of BatchCreateTableRows.
func (mr *MockHoneycodeAPIMockRecorder) BatchCreateTableRows(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchCreateTableRows", reflect.TypeOf((*MockHoneycodeAPI)(nil).BatchCreateTableRows), arg0)
}

// BatchCreateTableRowsRequest mocks base method.
func (m *MockHoneycodeAPI) BatchCreateTableRowsRequest(arg0 *honeycode.BatchCreateTableRowsInput) (*request.Request, *honeycode.BatchCreateTableRowsOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BatchCreateTableRowsRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*honeycode.BatchCreateTableRowsOutput)
	return ret0, ret1
}

// BatchCreateTableRowsRequest indicates an expected call of BatchCreateTableRowsRequest.
func (mr *MockHoneycodeAPIMockRecorder) BatchCreateTableRowsRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchCreateTableRowsRequest", reflect.TypeOf((*MockHoneycodeAPI)(nil).BatchCreateTableRowsRequest), arg0)
}

// BatchCreateTableRowsWithContext mocks base method.
func (m *MockHoneycodeAPI) BatchCreateTableRowsWithContext(arg0 context.Context, arg1 *honeycode.BatchCreateTableRowsInput, arg2 ...request.Option) (*honeycode.BatchCreateTableRowsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BatchCreateTableRowsWithContext", varargs...)
	ret0, _ := ret[0].(*honeycode.BatchCreateTableRowsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BatchCreateTableRowsWithContext indicates an expected call of BatchCreateTableRowsWithContext.
func (mr *MockHoneycodeAPIMockRecorder) BatchCreateTableRowsWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchCreateTableRowsWithContext", reflect.TypeOf((*MockHoneycodeAPI)(nil).BatchCreateTableRowsWithContext), varargs...)
}

// BatchDeleteTableRows mocks base method.
func (m *MockHoneycodeAPI) BatchDeleteTableRows(arg0 *honeycode.BatchDeleteTableRowsInput) (*honeycode.BatchDeleteTableRowsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BatchDeleteTableRows", arg0)
	ret0, _ := ret[0].(*honeycode.BatchDeleteTableRowsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BatchDeleteTableRows indicates an expected call of BatchDeleteTableRows.
func (mr *MockHoneycodeAPIMockRecorder) BatchDeleteTableRows(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchDeleteTableRows", reflect.TypeOf((*MockHoneycodeAPI)(nil).BatchDeleteTableRows), arg0)
}

// BatchDeleteTableRowsRequest mocks base method.
func (m *MockHoneycodeAPI) BatchDeleteTableRowsRequest(arg0 *honeycode.BatchDeleteTableRowsInput) (*request.Request, *honeycode.BatchDeleteTableRowsOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BatchDeleteTableRowsRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*honeycode.BatchDeleteTableRowsOutput)
	return ret0, ret1
}

// BatchDeleteTableRowsRequest indicates an expected call of BatchDeleteTableRowsRequest.
func (mr *MockHoneycodeAPIMockRecorder) BatchDeleteTableRowsRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchDeleteTableRowsRequest", reflect.TypeOf((*MockHoneycodeAPI)(nil).BatchDeleteTableRowsRequest), arg0)
}

// BatchDeleteTableRowsWithContext mocks base method.
func (m *MockHoneycodeAPI) BatchDeleteTableRowsWithContext(arg0 context.Context, arg1 *honeycode.BatchDeleteTableRowsInput, arg2 ...request.Option) (*honeycode.BatchDeleteTableRowsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BatchDeleteTableRowsWithContext", varargs...)
	ret0, _ := ret[0].(*honeycode.BatchDeleteTableRowsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BatchDeleteTableRowsWithContext indicates an expected call of BatchDeleteTableRowsWithContext.
func (mr *MockHoneycodeAPIMockRecorder) BatchDeleteTableRowsWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchDeleteTableRowsWithContext", reflect.TypeOf((*MockHoneycodeAPI)(nil).BatchDeleteTableRowsWithContext), varargs...)
}

// BatchUpdateTableRows mocks base method.
func (m *MockHoneycodeAPI) BatchUpdateTableRows(arg0 *honeycode.BatchUpdateTableRowsInput) (*honeycode.BatchUpdateTableRowsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BatchUpdateTableRows", arg0)
	ret0, _ := ret[0].(*honeycode.BatchUpdateTableRowsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BatchUpdateTableRows indicates an expected call of BatchUpdateTableRows.
func (mr *MockHoneycodeAPIMockRecorder) BatchUpdateTableRows(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchUpdateTableRows", reflect.TypeOf((*MockHoneycodeAPI)(nil).BatchUpdateTableRows), arg0)
}

// BatchUpdateTableRowsRequest mocks base method.
func (m *MockHoneycodeAPI) BatchUpdateTableRowsRequest(arg0 *honeycode.BatchUpdateTableRowsInput) (*request.Request, *honeycode.BatchUpdateTableRowsOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BatchUpdateTableRowsRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*honeycode.BatchUpdateTableRowsOutput)
	return ret0, ret1
}

// BatchUpdateTableRowsRequest indicates an expected call of BatchUpdateTableRowsRequest.
func (mr *MockHoneycodeAPIMockRecorder) BatchUpdateTableRowsRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchUpdateTableRowsRequest", reflect.TypeOf((*MockHoneycodeAPI)(nil).BatchUpdateTableRowsRequest), arg0)
}

// BatchUpdateTableRowsWithContext mocks base method.
func (m *MockHoneycodeAPI) BatchUpdateTableRowsWithContext(arg0 context.Context, arg1 *honeycode.BatchUpdateTableRowsInput, arg2 ...request.Option) (*honeycode.BatchUpdateTableRowsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BatchUpdateTableRowsWithContext", varargs...)
	ret0, _ := ret[0].(*honeycode.BatchUpdateTableRowsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BatchUpdateTableRowsWithContext indicates an expected call of BatchUpdateTableRowsWithContext.
func (mr *MockHoneycodeAPIMockRecorder) BatchUpdateTableRowsWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchUpdateTableRowsWithContext", reflect.TypeOf((*MockHoneycodeAPI)(nil).BatchUpdateTableRowsWithContext), varargs...)
}

// BatchUpsertTableRows mocks base method.
func (m *MockHoneycodeAPI) BatchUpsertTableRows(arg0 *honeycode.BatchUpsertTableRowsInput) (*honeycode.BatchUpsertTableRowsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BatchUpsertTableRows", arg0)
	ret0, _ := ret[0].(*honeycode.BatchUpsertTableRowsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BatchUpsertTableRows indicates an expected call of BatchUpsertTableRows.
func (mr *MockHoneycodeAPIMockRecorder) BatchUpsertTableRows(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchUpsertTableRows", reflect.TypeOf((*MockHoneycodeAPI)(nil).BatchUpsertTableRows), arg0)
}

// BatchUpsertTableRowsRequest mocks base method.
func (m *MockHoneycodeAPI) BatchUpsertTableRowsRequest(arg0 *honeycode.BatchUpsertTableRowsInput) (*request.Request, *honeycode.BatchUpsertTableRowsOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BatchUpsertTableRowsRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*honeycode.BatchUpsertTableRowsOutput)
	return ret0, ret1
}

// BatchUpsertTableRowsRequest indicates an expected call of BatchUpsertTableRowsRequest.
func (mr *MockHoneycodeAPIMockRecorder) BatchUpsertTableRowsRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchUpsertTableRowsRequest", reflect.TypeOf((*MockHoneycodeAPI)(nil).BatchUpsertTableRowsRequest), arg0)
}

// BatchUpsertTableRowsWithContext mocks base method.
func (m *MockHoneycodeAPI) BatchUpsertTableRowsWithContext(arg0 context.Context, arg1 *honeycode.BatchUpsertTableRowsInput, arg2 ...request.Option) (*honeycode.BatchUpsertTableRowsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BatchUpsertTableRowsWithContext", varargs...)
	ret0, _ := ret[0].(*honeycode.BatchUpsertTableRowsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BatchUpsertTableRowsWithContext indicates an expected call of BatchUpsertTableRowsWithContext.
func (mr *MockHoneycodeAPIMockRecorder) BatchUpsertTableRowsWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchUpsertTableRowsWithContext", reflect.TypeOf((*MockHoneycodeAPI)(nil).BatchUpsertTableRowsWithContext), varargs...)
}

// DescribeTableDataImportJob mocks base method.
func (m *MockHoneycodeAPI) DescribeTableDataImportJob(arg0 *honeycode.DescribeTableDataImportJobInput) (*honeycode.DescribeTableDataImportJobOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeTableDataImportJob", arg0)
	ret0, _ := ret[0].(*honeycode.DescribeTableDataImportJobOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeTableDataImportJob indicates an expected call of DescribeTableDataImportJob.
func (mr *MockHoneycodeAPIMockRecorder) DescribeTableDataImportJob(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeTableDataImportJob", reflect.TypeOf((*MockHoneycodeAPI)(nil).DescribeTableDataImportJob), arg0)
}

// DescribeTableDataImportJobRequest mocks base method.
func (m *MockHoneycodeAPI) DescribeTableDataImportJobRequest(arg0 *honeycode.DescribeTableDataImportJobInput) (*request.Request, *honeycode.DescribeTableDataImportJobOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeTableDataImportJobRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*honeycode.DescribeTableDataImportJobOutput)
	return ret0, ret1
}

// DescribeTableDataImportJobRequest indicates an expected call of DescribeTableDataImportJobRequest.
func (mr *MockHoneycodeAPIMockRecorder) DescribeTableDataImportJobRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeTableDataImportJobRequest", reflect.TypeOf((*MockHoneycodeAPI)(nil).DescribeTableDataImportJobRequest), arg0)
}

// DescribeTableDataImportJobWithContext mocks base method.
func (m *MockHoneycodeAPI) DescribeTableDataImportJobWithContext(arg0 context.Context, arg1 *honeycode.DescribeTableDataImportJobInput, arg2 ...request.Option) (*honeycode.DescribeTableDataImportJobOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeTableDataImportJobWithContext", varargs...)
	ret0, _ := ret[0].(*honeycode.DescribeTableDataImportJobOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeTableDataImportJobWithContext indicates an expected call of DescribeTableDataImportJobWithContext.
func (mr *MockHoneycodeAPIMockRecorder) DescribeTableDataImportJobWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeTableDataImportJobWithContext", reflect.TypeOf((*MockHoneycodeAPI)(nil).DescribeTableDataImportJobWithContext), varargs...)
}

// GetScreenData mocks base method.
func (m *MockHoneycodeAPI) GetScreenData(arg0 *honeycode.GetScreenDataInput) (*honeycode.GetScreenDataOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetScreenData", arg0)
	ret0, _ := ret[0].(*honeycode.GetScreenDataOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetScreenData indicates an expected call of GetScreenData.
func (mr *MockHoneycodeAPIMockRecorder) GetScreenData(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetScreenData", reflect.TypeOf((*MockHoneycodeAPI)(nil).GetScreenData), arg0)
}

// GetScreenDataRequest mocks base method.
func (m *MockHoneycodeAPI) GetScreenDataRequest(arg0 *honeycode.GetScreenDataInput) (*request.Request, *honeycode.GetScreenDataOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetScreenDataRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*honeycode.GetScreenDataOutput)
	return ret0, ret1
}

// GetScreenDataRequest indicates an expected call of GetScreenDataRequest.
func (mr *MockHoneycodeAPIMockRecorder) GetScreenDataRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetScreenDataRequest", reflect.TypeOf((*MockHoneycodeAPI)(nil).GetScreenDataRequest), arg0)
}

// GetScreenDataWithContext mocks base method.
func (m *MockHoneycodeAPI) GetScreenDataWithContext(arg0 context.Context, arg1 *honeycode.GetScreenDataInput, arg2 ...request.Option) (*honeycode.GetScreenDataOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetScreenDataWithContext", varargs...)
	ret0, _ := ret[0].(*honeycode.GetScreenDataOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetScreenDataWithContext indicates an expected call of GetScreenDataWithContext.
func (mr *MockHoneycodeAPIMockRecorder) GetScreenDataWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetScreenDataWithContext", reflect.TypeOf((*MockHoneycodeAPI)(nil).GetScreenDataWithContext), varargs...)
}

// InvokeScreenAutomation mocks base method.
func (m *MockHoneycodeAPI) InvokeScreenAutomation(arg0 *honeycode.InvokeScreenAutomationInput) (*honeycode.InvokeScreenAutomationOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InvokeScreenAutomation", arg0)
	ret0, _ := ret[0].(*honeycode.InvokeScreenAutomationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InvokeScreenAutomation indicates an expected call of InvokeScreenAutomation.
func (mr *MockHoneycodeAPIMockRecorder) InvokeScreenAutomation(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InvokeScreenAutomation", reflect.TypeOf((*MockHoneycodeAPI)(nil).InvokeScreenAutomation), arg0)
}

// InvokeScreenAutomationRequest mocks base method.
func (m *MockHoneycodeAPI) InvokeScreenAutomationRequest(arg0 *honeycode.InvokeScreenAutomationInput) (*request.Request, *honeycode.InvokeScreenAutomationOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InvokeScreenAutomationRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*honeycode.InvokeScreenAutomationOutput)
	return ret0, ret1
}

// InvokeScreenAutomationRequest indicates an expected call of InvokeScreenAutomationRequest.
func (mr *MockHoneycodeAPIMockRecorder) InvokeScreenAutomationRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InvokeScreenAutomationRequest", reflect.TypeOf((*MockHoneycodeAPI)(nil).InvokeScreenAutomationRequest), arg0)
}

// InvokeScreenAutomationWithContext mocks base method.
func (m *MockHoneycodeAPI) InvokeScreenAutomationWithContext(arg0 context.Context, arg1 *honeycode.InvokeScreenAutomationInput, arg2 ...request.Option) (*honeycode.InvokeScreenAutomationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "InvokeScreenAutomationWithContext", varargs...)
	ret0, _ := ret[0].(*honeycode.InvokeScreenAutomationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InvokeScreenAutomationWithContext indicates an expected call of InvokeScreenAutomationWithContext.
func (mr *MockHoneycodeAPIMockRecorder) InvokeScreenAutomationWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InvokeScreenAutomationWithContext", reflect.TypeOf((*MockHoneycodeAPI)(nil).InvokeScreenAutomationWithContext), varargs...)
}

// ListTableColumns mocks base method.
func (m *MockHoneycodeAPI) ListTableColumns(arg0 *honeycode.ListTableColumnsInput) (*honeycode.ListTableColumnsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTableColumns", arg0)
	ret0, _ := ret[0].(*honeycode.ListTableColumnsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTableColumns indicates an expected call of ListTableColumns.
func (mr *MockHoneycodeAPIMockRecorder) ListTableColumns(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTableColumns", reflect.TypeOf((*MockHoneycodeAPI)(nil).ListTableColumns), arg0)
}

// ListTableColumnsPages mocks base method.
func (m *MockHoneycodeAPI) ListTableColumnsPages(arg0 *honeycode.ListTableColumnsInput, arg1 func(*honeycode.ListTableColumnsOutput, bool) bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTableColumnsPages", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListTableColumnsPages indicates an expected call of ListTableColumnsPages.
func (mr *MockHoneycodeAPIMockRecorder) ListTableColumnsPages(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTableColumnsPages", reflect.TypeOf((*MockHoneycodeAPI)(nil).ListTableColumnsPages), arg0, arg1)
}

// ListTableColumnsPagesWithContext mocks base method.
func (m *MockHoneycodeAPI) ListTableColumnsPagesWithContext(arg0 context.Context, arg1 *honeycode.ListTableColumnsInput, arg2 func(*honeycode.ListTableColumnsOutput, bool) bool, arg3 ...request.Option) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTableColumnsPagesWithContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListTableColumnsPagesWithContext indicates an expected call of ListTableColumnsPagesWithContext.
func (mr *MockHoneycodeAPIMockRecorder) ListTableColumnsPagesWithContext(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTableColumnsPagesWithContext", reflect.TypeOf((*MockHoneycodeAPI)(nil).ListTableColumnsPagesWithContext), varargs...)
}

// ListTableColumnsRequest mocks base method.
func (m *MockHoneycodeAPI) ListTableColumnsRequest(arg0 *honeycode.ListTableColumnsInput) (*request.Request, *honeycode.ListTableColumnsOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTableColumnsRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*honeycode.ListTableColumnsOutput)
	return ret0, ret1
}

// ListTableColumnsRequest indicates an expected call of ListTableColumnsRequest.
func (mr *MockHoneycodeAPIMockRecorder) ListTableColumnsRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTableColumnsRequest", reflect.TypeOf((*MockHoneycodeAPI)(nil).ListTableColumnsRequest), arg0)
}

// ListTableColumnsWithContext mocks base method.
func (m *MockHoneycodeAPI) ListTableColumnsWithContext(arg0 context.Context, arg1 *honeycode.ListTableColumnsInput, arg2 ...request.Option) (*honeycode.ListTableColumnsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTableColumnsWithContext", varargs...)
	ret0, _ := ret[0].(*honeycode.ListTableColumnsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTableColumnsWithContext indicates an expected call of ListTableColumnsWithContext.
func (mr *MockHoneycodeAPIMockRecorder) ListTableColumnsWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTableColumnsWithContext", reflect.TypeOf((*MockHoneycodeAPI)(nil).ListTableColumnsWithContext), varargs...)
}

// ListTableRows mocks base method.
func (m *MockHoneycodeAPI) ListTableRows(arg0 *honeycode.ListTableRowsInput) (*honeycode.ListTableRowsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTableRows", arg0)
	ret0, _ := ret[0].(*honeycode.ListTableRowsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTableRows indicates an expected call of ListTableRows.
func (mr *MockHoneycodeAPIMockRecorder) ListTableRows(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTableRows", reflect.TypeOf((*MockHoneycodeAPI)(nil).ListTableRows), arg0)
}

// ListTableRowsPages mocks base method.
func (m *MockHoneycodeAPI) ListTableRowsPages(arg0 *honeycode.ListTableRowsInput, arg1 func(*honeycode.ListTableRowsOutput, bool) bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTableRowsPages", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListTableRowsPages indicates an expected call of ListTableRowsPages.
func (mr *MockHoneycodeAPIMockRecorder) ListTableRowsPages(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTableRowsPages", reflect.TypeOf((*MockHoneycodeAPI)(nil).ListTableRowsPages), arg0, arg1)
}

// ListTableRowsPagesWithContext mocks base method.
func (m *MockHoneycodeAPI) ListTableRowsPagesWithContext(arg0 context.Context, arg1 *honeycode.ListTableRowsInput, arg2 func(*honeycode.ListTableRowsOutput, bool) bool, arg3 ...request.Option) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTableRowsPagesWithContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListTableRowsPagesWithContext indicates an expected call of ListTableRowsPagesWithContext.
func (mr *MockHoneycodeAPIMockRecorder) ListTableRowsPagesWithContext(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTableRowsPagesWithContext", reflect.TypeOf((*MockHoneycodeAPI)(nil).ListTableRowsPagesWithContext), varargs...)
}

// ListTableRowsRequest mocks base method.
func (m *MockHoneycodeAPI) ListTableRowsRequest(arg0 *honeycode.ListTableRowsInput) (*request.Request, *honeycode.ListTableRowsOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTableRowsRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*honeycode.ListTableRowsOutput)
	return ret0, ret1
}

// ListTableRowsRequest indicates an expected call of ListTableRowsRequest.
func (mr *MockHoneycodeAPIMockRecorder) ListTableRowsRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTableRowsRequest", reflect.TypeOf((*MockHoneycodeAPI)(nil).ListTableRowsRequest), arg0)
}

// ListTableRowsWithContext mocks base method.
func (m *MockHoneycodeAPI) ListTableRowsWithContext(arg0 context.Context, arg1 *honeycode.ListTableRowsInput, arg2 ...request.Option) (*honeycode.ListTableRowsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTableRowsWithContext", varargs...)
	ret0, _ := ret[0].(*honeycode.ListTableRowsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTableRowsWithContext indicates an expected call of ListTableRowsWithContext.
func (mr *MockHoneycodeAPIMockRecorder) ListTableRowsWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTableRowsWithContext", reflect.TypeOf((*MockHoneycodeAPI)(nil).ListTableRowsWithContext), varargs...)
}

// ListTables mocks base method.
func (m *MockHoneycodeAPI) ListTables(arg0 *honeycode.ListTablesInput) (*honeycode.ListTablesOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTables", arg0)
	ret0, _ := ret[0].(*honeycode.ListTablesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTables indicates an expected call of ListTables.
func (mr *MockHoneycodeAPIMockRecorder) ListTables(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTables", reflect.TypeOf((*MockHoneycodeAPI)(nil).ListTables), arg0)
}

// ListTablesPages mocks base method.
func (m *MockHoneycodeAPI) ListTablesPages(arg0 *honeycode.ListTablesInput, arg1 func(*honeycode.ListTablesOutput, bool) bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTablesPages", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListTablesPages indicates an expected call of ListTablesPages.
func (mr *MockHoneycodeAPIMockRecorder) ListTablesPages(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTablesPages", reflect.TypeOf((*MockHoneycodeAPI)(nil).ListTablesPages), arg0, arg1)
}

// ListTablesPagesWithContext mocks base method.
func (m *MockHoneycodeAPI) ListTablesPagesWithContext(arg0 context.Context, arg1 *honeycode.ListTablesInput, arg2 func(*honeycode.ListTablesOutput, bool) bool, arg3 ...request.Option) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTablesPagesWithContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListTablesPagesWithContext indicates an expected call of ListTablesPagesWithContext.
func (mr *MockHoneycodeAPIMockRecorder) ListTablesPagesWithContext(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTablesPagesWithContext", reflect.TypeOf((*MockHoneycodeAPI)(nil).ListTablesPagesWithContext), varargs...)
}

// ListTablesRequest mocks base method.
func (m *MockHoneycodeAPI) ListTablesRequest(arg0 *honeycode.ListTablesInput) (*request.Request, *honeycode.ListTablesOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTablesRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*honeycode.ListTablesOutput)
	return ret0, ret1
}

// ListTablesRequest indicates an expected call of ListTablesRequest.
func (mr *MockHoneycodeAPIMockRecorder) ListTablesRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTablesRequest", reflect.TypeOf((*MockHoneycodeAPI)(nil).ListTablesRequest), arg0)
}

// ListTablesWithContext mocks base method.
func (m *MockHoneycodeAPI) ListTablesWithContext(arg0 context.Context, arg1 *honeycode.ListTablesInput, arg2 ...request.Option) (*honeycode.ListTablesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTablesWithContext", varargs...)
	ret0, _ := ret[0].(*honeycode.ListTablesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTablesWithContext indicates an expected call of ListTablesWithContext.
func (mr *MockHoneycodeAPIMockRecorder) ListTablesWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTablesWithContext", reflect.TypeOf((*MockHoneycodeAPI)(nil).ListTablesWithContext), varargs...)
}

// QueryTableRows mocks base method.
func (m *MockHoneycodeAPI) QueryTableRows(arg0 *honeycode.QueryTableRowsInput) (*honeycode.QueryTableRowsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryTableRows", arg0)
	ret0, _ := ret[0].(*honeycode.QueryTableRowsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryTableRows indicates an expected call of QueryTableRows.
func (mr *MockHoneycodeAPIMockRecorder) QueryTableRows(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryTableRows", reflect.TypeOf((*MockHoneycodeAPI)(nil).QueryTableRows), arg0)
}

// QueryTableRowsPages mocks base method.
func (m *MockHoneycodeAPI) QueryTableRowsPages(arg0 *honeycode.QueryTableRowsInput, arg1 func(*honeycode.QueryTableRowsOutput, bool) bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryTableRowsPages", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// QueryTableRowsPages indicates an expected call of QueryTableRowsPages.
func (mr *MockHoneycodeAPIMockRecorder) QueryTableRowsPages(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryTableRowsPages", reflect.TypeOf((*MockHoneycodeAPI)(nil).QueryTableRowsPages), arg0, arg1)
}

// QueryTableRowsPagesWithContext mocks base method.
func (m *MockHoneycodeAPI) QueryTableRowsPagesWithContext(arg0 context.Context, arg1 *honeycode.QueryTableRowsInput, arg2 func(*honeycode.QueryTableRowsOutput, bool) bool, arg3 ...request.Option) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "QueryTableRowsPagesWithContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// QueryTableRowsPagesWithContext indicates an expected call of QueryTableRowsPagesWithContext.
func (mr *MockHoneycodeAPIMockRecorder) QueryTableRowsPagesWithContext(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryTableRowsPagesWithContext", reflect.TypeOf((*MockHoneycodeAPI)(nil).QueryTableRowsPagesWithContext), varargs...)
}

// QueryTableRowsRequest mocks base method.
func (m *MockHoneycodeAPI) QueryTableRowsRequest(arg0 *honeycode.QueryTableRowsInput) (*request.Request, *honeycode.QueryTableRowsOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryTableRowsRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*honeycode.QueryTableRowsOutput)
	return ret0, ret1
}

// QueryTableRowsRequest indicates an expected call of QueryTableRowsRequest.
func (mr *MockHoneycodeAPIMockRecorder) QueryTableRowsRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryTableRowsRequest", reflect.TypeOf((*MockHoneycodeAPI)(nil).QueryTableRowsRequest), arg0)
}

// QueryTableRowsWithContext mocks base method.
func (m *MockHoneycodeAPI) QueryTableRowsWithContext(arg0 context.Context, arg1 *honeycode.QueryTableRowsInput, arg2 ...request.Option) (*honeycode.QueryTableRowsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "QueryTableRowsWithContext", varargs...)
	ret0, _ := ret[0].(*honeycode.QueryTableRowsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryTableRowsWithContext indicates an expected call of QueryTableRowsWithContext.
func (mr *MockHoneycodeAPIMockRecorder) QueryTableRowsWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryTableRowsWithContext", reflect.TypeOf((*MockHoneycodeAPI)(nil).QueryTableRowsWithContext), varargs...)
}

// StartTableDataImportJob mocks base method.
func (m *MockHoneycodeAPI) StartTableDataImportJob(arg0 *honeycode.StartTableDataImportJobInput) (*honeycode.StartTableDataImportJobOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartTableDataImportJob", arg0)
	ret0, _ := ret[0].(*honeycode.StartTableDataImportJobOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StartTableDataImportJob indicates an expected call of StartTableDataImportJob.
func (mr *MockHoneycodeAPIMockRecorder) StartTableDataImportJob(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartTableDataImportJob", reflect.TypeOf((*MockHoneycodeAPI)(nil).StartTableDataImportJob), arg0)
}

// StartTableDataImportJobRequest mocks base method.
func (m *MockHoneycodeAPI) StartTableDataImportJobRequest(arg0 *honeycode.StartTableDataImportJobInput) (*request.Request, *honeycode.StartTableDataImportJobOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartTableDataImportJobRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*honeycode.StartTableDataImportJobOutput)
	return ret0, ret1
}

// StartTableDataImportJobRequest indicates an expected call of StartTableDataImportJobRequest.
func (mr *MockHoneycodeAPIMockRecorder) StartTableDataImportJobRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartTableDataImportJobRequest", reflect.TypeOf((*MockHoneycodeAPI)(nil).StartTableDataImportJobRequest), arg0)
}

// StartTableDataImportJobWithContext mocks base method.
func (m *MockHoneycodeAPI) StartTableDataImportJobWithContext(arg0 context.Context, arg1 *honeycode.StartTableDataImportJobInput, arg2 ...request.Option) (*honeycode.StartTableDataImportJobOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "StartTableDataImportJobWithContext", varargs...)
	ret0, _ := ret[0].(*honeycode.StartTableDataImportJobOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StartTableDataImportJobWithContext indicates an expected call of StartTableDataImportJobWithContext.
func (mr *MockHoneycodeAPIMockRecorder) StartTableDataImportJobWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartTableDataImportJobWithContext", reflect.TypeOf((*MockHoneycodeAPI)(nil).StartTableDataImportJobWithContext), varargs...)
}
