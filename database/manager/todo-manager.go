package dbmanager

import (
	dbmodel "skeleton-service/database/model"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type todoManager struct {
	db *gorm.DB
}

func NewTodoManager(db *gorm.DB) TodoManager {
	return &todoManager{db: db}
}

func (m *todoManager) Create(todo *dbmodel.Todo) error {
	result := m.db.Create(todo)
	if result.Error != nil {
		return errors.Wrap(result.Error, "failed to create todo")
	}
	return nil
}

func (m *todoManager) Update(todo *dbmodel.Todo) error {
	result := m.db.Save(todo)
	if result.Error != nil {
		return errors.Wrap(result.Error, "failed to update todo")
	}
	return nil
}

func (m *todoManager) Delete(id string) error {
	result := m.db.Delete(&dbmodel.Todo{ID: id})
	if result.Error != nil {
		return errors.Wrap(result.Error, "failed to delete todo")
	}
	if result.RowsAffected == 0 {
		return errors.New("todo not found")
	}
	return nil
}

func (m *todoManager) GetByID(id string) (*dbmodel.Todo, error) {
	var dbTodo dbmodel.Todo
	result := m.db.First(&dbTodo, "id = ?", id)
	if result.Error != nil {
		return nil, errors.Wrap(result.Error, "failed to get todo by ID")
	}
	return &dbTodo, nil
}

func (m *todoManager) GetAll() ([]*dbmodel.Todo, error) {
	var dbTodos []*dbmodel.Todo
	result := m.db.Find(&dbTodos)
	if result.Error != nil {
		return nil, errors.Wrap(result.Error, "failed to get all todos")
	}
	return dbTodos, nil
}

func (m *todoManager) GetPage(page int, limit int) ([]*dbmodel.Todo, error) {
	var dbTodos []*dbmodel.Todo
	result := m.db.Limit(limit).Offset((page - 1) * limit).Order("created_at asc").Find(&dbTodos)
	if result.Error != nil {
		return nil, errors.Wrap(result.Error, "failed to get todos by page")
	}
	return dbTodos, nil
}

func (m *todoManager) GetCount() (int64, error) {
	var count int64
	result := m.db.Model(&dbmodel.Todo{}).Count(&count)
	if result.Error != nil {
		return 0, errors.Wrap(result.Error, "failed to get todo count")
	}
	return count, nil
}

func (m *todoManager) GetStats() (*dbmodel.TodoStats, error) {
	stats := dbmodel.TodoStats{}
	result := m.db.Model(&dbmodel.Todo{}).Select(`count(*) as total, sum(case when done = TRUE then 1 else 0 end) as total_completed, string_agg(text, ', ' order by created_at asc) AS aggregate_text`).Scan(&stats)
	if result.Error != nil {
		return nil, result.Error
	}
	return &stats, nil
}
