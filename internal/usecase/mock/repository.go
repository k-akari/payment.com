// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go
//
// Generated by this command:
//
//	mockgen -source=repository.go -destination=mock/repository.go -package=mock
//
// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	domain "github.com/k-akari/payment.com/internal/domain"
	gomock "go.uber.org/mock/gomock"
)

// MockcompanyRepository is a mock of companyRepository interface.
type MockcompanyRepository struct {
	ctrl     *gomock.Controller
	recorder *MockcompanyRepositoryMockRecorder
}

// MockcompanyRepositoryMockRecorder is the mock recorder for MockcompanyRepository.
type MockcompanyRepositoryMockRecorder struct {
	mock *MockcompanyRepository
}

// NewMockcompanyRepository creates a new mock instance.
func NewMockcompanyRepository(ctrl *gomock.Controller) *MockcompanyRepository {
	mock := &MockcompanyRepository{ctrl: ctrl}
	mock.recorder = &MockcompanyRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockcompanyRepository) EXPECT() *MockcompanyRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockcompanyRepository) Create(ctx context.Context, company *domain.Company) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, company)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockcompanyRepositoryMockRecorder) Create(ctx, company any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockcompanyRepository)(nil).Create), ctx, company)
}

// GetByID mocks base method.
func (m *MockcompanyRepository) GetByID(ctx context.Context, companyID domain.CompanyID) (*domain.Company, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, companyID)
	ret0, _ := ret[0].(*domain.Company)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockcompanyRepositoryMockRecorder) GetByID(ctx, companyID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockcompanyRepository)(nil).GetByID), ctx, companyID)
}
