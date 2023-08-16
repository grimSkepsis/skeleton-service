package dbmodel

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	ID     string `gorm:"unique;default:gen_random_uuid()"`
	Text   string
	UserID string
	Done   bool
}

type TodoStats struct {
	Total          int64  `json:"total"`
	TotalCompleted int64  `json:"total_completed"`
	AggregateText  string `json:"aggregate_text"`
}
