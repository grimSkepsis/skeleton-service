package graph

import (
	dbmodel "skeleton-service/database/model"
	"skeleton-service/graph/model"
)

func convertDBTodoToModel(dbTodo *dbmodel.Todo) *model.Todo {
	return &model.Todo{
		ID:   dbTodo.ID,
		Text: dbTodo.Text,
		Done: dbTodo.Done,
		User: &model.User{ID: dbTodo.UserID, Name: "Test User"},
	}
}

func convertDBTodosToModels(dbTodos []*dbmodel.Todo) []*model.Todo {
	var todos []*model.Todo
	for _, dbTodo := range dbTodos {
		todos = append(todos, convertDBTodoToModel(dbTodo))
	}
	return todos
}
