package graph

import (
	dbmanager "skeleton-service/database/manager"

	"go.uber.org/zap"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	logger      *zap.Logger
	todoManager dbmanager.TodoManager
}

func NewResolver(logger *zap.Logger, todoManager dbmanager.TodoManager) *Resolver {
	return &Resolver{logger: logger, todoManager: todoManager}
}
