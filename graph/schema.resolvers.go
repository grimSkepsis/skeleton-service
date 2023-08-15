package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"fmt"
	"math"
	dbmodel "skeleton-service/database/model"
	"skeleton-service/graph/model"

	"go.uber.org/zap"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	r.logger.Info("createTodo", zap.Any("input", input))
	newTodo := dbmodel.Todo{Text: input.Text, UserID: input.UserID, Done: false}
	result := r.db.Create(&newTodo)

	if result.Error != nil {
		return nil, result.Error
	}

	return convertDBTodoToModel(&newTodo), nil
}

// DeleteTodoByID is the resolver for the deleteTodoById field.
func (r *mutationResolver) DeleteTodoByID(ctx context.Context, id string) (bool, error) {
	r.logger.Info("deleteTodoById", zap.String("id", id))

	result := r.db.Delete(&dbmodel.Todo{ID: id})
	if result.Error != nil {
		return false, result.Error
	}
	if result.RowsAffected == 0 {
		return false, fmt.Errorf("no todo with id %s", id)
	}
	return true, nil
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	r.logger.Info("todos")
	var dbTodos []dbmodel.Todo
	result := r.db.Find(&dbTodos)
	if result.Error != nil {
		return nil, result.Error
	}

	var todos []*model.Todo
	for _, dbTodo := range dbTodos {
		todos = append(todos, convertDBTodoToModel(&dbTodo))
	}

	return todos, nil
}

// TodosPaginated is the resolver for the todosPaginated field.
func (r *queryResolver) TodosPaginated(ctx context.Context, page int, limit int) (*model.TodoConnection, error) {
	r.logger.Info("todosPaginated", zap.Int("page", page), zap.Int("limit", limit))
	var dbTodos []dbmodel.Todo
	result := r.db.Limit(limit).Offset((page - 1) * limit).Find(&dbTodos)
	if result.Error != nil {
		return nil, result.Error
	}
	var todoCount int64
	countResult := r.db.Model(&dbmodel.Todo{}).Count(&todoCount)
	if countResult.Error != nil {
		return nil, countResult.Error
	}
	return &model.TodoConnection{
		Edges: convertDBTodosToModels(dbTodos),
		PageInfo: &model.PageInfo{
			Total:           int(todoCount),
			TotalPages:      int(math.Ceil(float64(todoCount) / float64(limit))),
			CurrentPage:     int(page),
			HasNextPage:     todoCount > int64((page-1)*limit),
			HasPreviousPage: page > 1,
		},
	}, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
