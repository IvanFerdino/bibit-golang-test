// Code generated by MockGen. DO NOT EDIT.
// Source: ../internal/repository/repository.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	model "IvanFerdino/bibit-golang-test/internal/model"
	context "context"
	reflect "reflect"
	sync "sync"

	gomock "github.com/golang/mock/gomock"
)

// MockMovieRepository is a mock of MovieRepository interface.
type MockMovieRepository struct {
	ctrl     *gomock.Controller
	recorder *MockMovieRepositoryMockRecorder
}

// MockMovieRepositoryMockRecorder is the mock recorder for MockMovieRepository.
type MockMovieRepositoryMockRecorder struct {
	mock *MockMovieRepository
}

// NewMockMovieRepository creates a new mock instance.
func NewMockMovieRepository(ctrl *gomock.Controller) *MockMovieRepository {
	mock := &MockMovieRepository{ctrl: ctrl}
	mock.recorder = &MockMovieRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMovieRepository) EXPECT() *MockMovieRepositoryMockRecorder {
	return m.recorder
}

// Detail mocks base method.
func (m *MockMovieRepository) Detail(ctx context.Context, movieId string) (*model.MovieDetailResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Detail", ctx, movieId)
	ret0, _ := ret[0].(*model.MovieDetailResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Detail indicates an expected call of Detail.
func (mr *MockMovieRepositoryMockRecorder) Detail(ctx, movieId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Detail", reflect.TypeOf((*MockMovieRepository)(nil).Detail), ctx, movieId)
}

// GetDetailRequest mocks base method.
func (m *MockMovieRepository) GetDetailRequest(movieId string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDetailRequest", movieId)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetDetailRequest indicates an expected call of GetDetailRequest.
func (mr *MockMovieRepositoryMockRecorder) GetDetailRequest(movieId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDetailRequest", reflect.TypeOf((*MockMovieRepository)(nil).GetDetailRequest), movieId)
}

// GetSearchRequest mocks base method.
func (m *MockMovieRepository) GetSearchRequest(keyword string, page int) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSearchRequest", keyword, page)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetSearchRequest indicates an expected call of GetSearchRequest.
func (mr *MockMovieRepositoryMockRecorder) GetSearchRequest(keyword, page interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSearchRequest", reflect.TypeOf((*MockMovieRepository)(nil).GetSearchRequest), keyword, page)
}

// Search mocks base method.
func (m *MockMovieRepository) Search(ctx context.Context, keyword string, page int) (*model.MovieSearchResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", ctx, keyword, page)
	ret0, _ := ret[0].(*model.MovieSearchResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search.
func (mr *MockMovieRepositoryMockRecorder) Search(ctx, keyword, page interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockMovieRepository)(nil).Search), ctx, keyword, page)
}

// MockLogRepository is a mock of LogRepository interface.
type MockLogRepository struct {
	ctrl     *gomock.Controller
	recorder *MockLogRepositoryMockRecorder
}

// MockLogRepositoryMockRecorder is the mock recorder for MockLogRepository.
type MockLogRepositoryMockRecorder struct {
	mock *MockLogRepository
}

// NewMockLogRepository creates a new mock instance.
func NewMockLogRepository(ctrl *gomock.Controller) *MockLogRepository {
	mock := &MockLogRepository{ctrl: ctrl}
	mock.recorder = &MockLogRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLogRepository) EXPECT() *MockLogRepositoryMockRecorder {
	return m.recorder
}

// LogApiCall mocks base method.
func (m *MockLogRepository) LogApiCall(ctx context.Context, wg *sync.WaitGroup, data *model.ApiCall) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "LogApiCall", ctx, wg, data)
}

// LogApiCall indicates an expected call of LogApiCall.
func (mr *MockLogRepositoryMockRecorder) LogApiCall(ctx, wg, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LogApiCall", reflect.TypeOf((*MockLogRepository)(nil).LogApiCall), ctx, wg, data)
}