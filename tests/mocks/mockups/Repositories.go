package mockups

import (
	"github.com/golang/mock/gomock"
	"hexagonal/src/core/domain"
	"reflect"
)

// MockGamesRepository is a mock of GamesRepository interface
type MockGamesRepository struct {
	ctrl     *gomock.Controller
	recorder *MockGamesRepositoryMockRecorder
}

// MockGamesRepositoryMockRecorder is the mock recorder for MockGamesRepository
type MockGamesRepositoryMockRecorder struct {
	mock *MockGamesRepository
}

// NewMockGamesRepository creates a new mock instance
func NewMockGamesRepository(ctrl *gomock.Controller) *MockGamesRepository {
	mock := &MockGamesRepository{ctrl: ctrl}
	mock.recorder = &MockGamesRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGamesRepository) EXPECT() *MockGamesRepositoryMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockGamesRepository) Get(id string) (domain.Game, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", id)
	ret0, _ := ret[0].(domain.Game)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockGamesRepositoryMockRecorder) Get(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockGamesRepository)(nil).Get), id)
}

// Save mocks base method
func (m *MockGamesRepository) Save(arg0 domain.Game) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save
func (mr *MockGamesRepositoryMockRecorder) Save(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockGamesRepository)(nil).Save), arg0)
}
