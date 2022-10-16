// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/repository/internal/ports/company_port.go

// Package mock_ports is a generated GoMock package.
package mock_ports

import (
	context "context"
	reflect "reflect"

	domain "github.com/akhi19/companies/pkg/domain"
	gomock "github.com/golang/mock/gomock"
)

// MockICompany is a mock of ICompany interface.
type MockICompany struct {
	ctrl     *gomock.Controller
	recorder *MockICompanyMockRecorder
}

// MockICompanyMockRecorder is the mock recorder for MockICompany.
type MockICompanyMockRecorder struct {
	mock *MockICompany
}

// NewMockICompany creates a new mock instance.
func NewMockICompany(ctrl *gomock.Controller) *MockICompany {
	mock := &MockICompany{ctrl: ctrl}
	mock.recorder = &MockICompanyMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockICompany) EXPECT() *MockICompanyMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockICompany) Add(ctx context.Context, companyDTO domain.CompanyDTO) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", ctx, companyDTO)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add.
func (mr *MockICompanyMockRecorder) Add(ctx, companyDTO interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockICompany)(nil).Add), ctx, companyDTO)
}

// Delete mocks base method.
func (m *MockICompany) Delete(ctx context.Context, name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, name)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockICompanyMockRecorder) Delete(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockICompany)(nil).Delete), ctx, name)
}

// GetCompanyByName mocks base method.
func (m *MockICompany) GetCompanyByName(ctx context.Context, name string) (*domain.CompanyDTO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCompanyByName", ctx, name)
	ret0, _ := ret[0].(*domain.CompanyDTO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCompanyByName indicates an expected call of GetCompanyByName.
func (mr *MockICompanyMockRecorder) GetCompanyByName(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCompanyByName", reflect.TypeOf((*MockICompany)(nil).GetCompanyByName), ctx, name)
}

// Update mocks base method.
func (m *MockICompany) Update(ctx context.Context, name string, updateCompanyDTO domain.UpdateCompanyDTO) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, name, updateCompanyDTO)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockICompanyMockRecorder) Update(ctx, name, updateCompanyDTO interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockICompany)(nil).Update), ctx, name, updateCompanyDTO)
}