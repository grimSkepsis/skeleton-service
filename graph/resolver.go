package graph

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	logger *zap.Logger
	db     *gorm.DB
}

func NewResolver(logger *zap.Logger, db *gorm.DB) *Resolver {
	return &Resolver{logger: logger, db: db}
}
