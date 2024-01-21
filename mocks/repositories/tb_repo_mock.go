// Code generated by MockGen. DO NOT EDIT.
// Source: adapter/repositories/tb/tb_repo.go

// Package mocks is a generated GoMock package.
package mocks

import (
	models "be-cadidate-votes/models"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockTbRepo is a mock of TbRepo interface.
type MockTbRepo struct {
	ctrl     *gomock.Controller
	recorder *MockTbRepoMockRecorder
}

// MockTbRepoMockRecorder is the mock recorder for MockTbRepo.
type MockTbRepoMockRecorder struct {
	mock *MockTbRepo
}

// NewMockTbRepo creates a new mock instance.
func NewMockTbRepo(ctrl *gomock.Controller) *MockTbRepo {
	mock := &MockTbRepo{ctrl: ctrl}
	mock.recorder = &MockTbRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTbRepo) EXPECT() *MockTbRepoMockRecorder {
	return m.recorder
}

// CountCandidate mocks base method.
func (m *MockTbRepo) CountCandidate(ctx context.Context, filter models.ListCandidateRequest) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountCandidate", ctx, filter)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountCandidate indicates an expected call of CountCandidate.
func (mr *MockTbRepoMockRecorder) CountCandidate(ctx, filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountCandidate", reflect.TypeOf((*MockTbRepo)(nil).CountCandidate), ctx, filter)
}

// CreateCandidate mocks base method.
func (m *MockTbRepo) CreateCandidate(ctx context.Context, user *models.Candidate) (*models.Candidate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCandidate", ctx, user)
	ret0, _ := ret[0].(*models.Candidate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCandidate indicates an expected call of CreateCandidate.
func (mr *MockTbRepoMockRecorder) CreateCandidate(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCandidate", reflect.TypeOf((*MockTbRepo)(nil).CreateCandidate), ctx, user)
}

// CreateUser mocks base method.
func (m *MockTbRepo) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, user)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockTbRepoMockRecorder) CreateUser(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockTbRepo)(nil).CreateUser), ctx, user)
}

// DeleteVotes mocks base method.
func (m *MockTbRepo) DeleteVotes(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteVotes", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteVotes indicates an expected call of DeleteVotes.
func (mr *MockTbRepoMockRecorder) DeleteVotes(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteVotes", reflect.TypeOf((*MockTbRepo)(nil).DeleteVotes), ctx)
}

// GetUser mocks base method.
func (m *MockTbRepo) GetUser(ctx context.Context, filter models.User) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", ctx, filter)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockTbRepoMockRecorder) GetUser(ctx, filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockTbRepo)(nil).GetUser), ctx, filter)
}

// ListCandidateWithVote mocks base method.
func (m *MockTbRepo) ListCandidateWithVote(ctx context.Context, filter models.ListCandidateRequest) ([]models.CandidateResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListCandidateWithVote", ctx, filter)
	ret0, _ := ret[0].([]models.CandidateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCandidateWithVote indicates an expected call of ListCandidateWithVote.
func (mr *MockTbRepoMockRecorder) ListCandidateWithVote(ctx, filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCandidateWithVote", reflect.TypeOf((*MockTbRepo)(nil).ListCandidateWithVote), ctx, filter)
}

// UpdateCandidate mocks base method.
func (m *MockTbRepo) UpdateCandidate(ctx context.Context, user models.Candidate) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCandidate", ctx, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCandidate indicates an expected call of UpdateCandidate.
func (mr *MockTbRepoMockRecorder) UpdateCandidate(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCandidate", reflect.TypeOf((*MockTbRepo)(nil).UpdateCandidate), ctx, user)
}
