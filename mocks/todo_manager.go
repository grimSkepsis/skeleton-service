package mocks

import (
	dbmodel "skeleton-service/database/model"

	"github.com/stretchr/testify/mock"
)

type TodoManager struct {
	mock.Mock
}

func (m *TodoManager) Create(todo *dbmodel.Todo) error {
	args := m.Called(todo)
	return args.Error(0)
}

func (m *TodoManager) Update(todo *dbmodel.Todo) error {
	args := m.Called(todo)
	return args.Error(0)
}

func (m *TodoManager) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *TodoManager) GetByID(id string) (*dbmodel.Todo, error) {
	args := m.Called(id)
	return args.Get(0).(*dbmodel.Todo), args.Error(1)
}

func (m *TodoManager) GetAll() ([]*dbmodel.Todo, error) {
	args := m.Called()
	return args.Get(0).([]*dbmodel.Todo), args.Error(1)
}

func (m *TodoManager) GetPage(page int, limit int) ([]*dbmodel.Todo, error) {
	args := m.Called(page, limit)
	return args.Get(0).([]*dbmodel.Todo), args.Error(1)
}

func (m *TodoManager) GetCount() (int64, error) {
	args := m.Called()
	return args.Get(0).(int64), args.Error(1)
}

func (m *TodoManager) GetStats() (*dbmodel.TodoStats, error) {
	args := m.Called()
	return args.Get(0).(*dbmodel.TodoStats), args.Error(1)
}
