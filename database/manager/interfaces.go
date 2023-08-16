package dbmanager

import (
	dbmodel "skeleton-service/database/model"
)

type TodoManager interface {
	Create(todo *dbmodel.Todo) error
	Update(todo *dbmodel.Todo) error
	Delete(id string) error
	GetByID(id string) (*dbmodel.Todo, error)
	GetAll() ([]*dbmodel.Todo, error)
	GetPage(page int, limit int) ([]*dbmodel.Todo, error)
	GetCount() (int64, error)
	GetStats() (*dbmodel.TodoStats, error)
}
