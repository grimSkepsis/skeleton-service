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
	err := r.todoManager.Create(&newTodo)
	if err != nil {
		return nil, err
	}
	return convertDBTodoToModel(&newTodo), nil
}

// DeleteTodoByID is the resolver for the deleteTodoById field.
func (r *mutationResolver) DeleteTodoByID(ctx context.Context, id string) (bool, error) {
	r.logger.Info("deleteTodoById", zap.String("id", id))

	err := r.todoManager.Delete(id)
	if err != nil {
		return false, err
	}
	return true, nil
}

// UpdateTodoByID is the resolver for the updateTodoById field.
func (r *mutationResolver) UpdateTodoByID(ctx context.Context, id string, done bool) (*model.Todo, error) {
	r.logger.Info("updateTodoById", zap.String("id", id), zap.Bool("done", done))

	dbTodo, err := r.todoManager.GetByID(id)
	if err != nil {
		return nil, err
	}

	dbTodo.Done = done
	err = r.todoManager.Update(dbTodo)
	if err != nil {
		return nil, err
	}

	return convertDBTodoToModel(dbTodo), nil
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	r.logger.Info("todos")
	dbTodos, err := r.todoManager.GetAll()
	if err != nil {
		return nil, err
	}

	var todos []*model.Todo
	for _, dbTodo := range dbTodos {
		todos = append(todos, convertDBTodoToModel(dbTodo))
	}

	return todos, nil
}

// TodosPaginated is the resolver for the todosPaginated field.
func (r *queryResolver) TodosPaginated(ctx context.Context, page int, limit int) (*model.TodoConnection, error) {
	r.logger.Info("todosPaginated", zap.Int("page", page), zap.Int("limit", limit))

	dbTodos, err := r.todoManager.GetPage(page, limit)
	if err != nil {
		return nil, err
	}

	todoCount, err := r.todoManager.GetCount()
	if err != nil {
		return nil, err
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

// CompletionRatio is the resolver for the completionRatio field.
func (r *queryResolver) CompletionRatio(ctx context.Context) (float64, error) {
	panic(fmt.Errorf("not implemented: CompletionRatio - completionRatio"))
}

// TodoStats is the resolver for the todoStats field.
func (r *queryResolver) TodoStats(ctx context.Context) (*model.TodoStats, error) {
	stats, err := r.todoManager.GetStats()
	if err != nil {
		return nil, err
	}
	r.logger.Info("todoStats", zap.Any("stats", stats))
	return &model.TodoStats{
		Total:          int(stats.Total),
		TotalCompleted: int(stats.TotalCompleted),
		AggregateText:  stats.AggregateText,
	}, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
