package dbmodel

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	ID     string `gorm:"unique;default:gen_random_uuid()"`
	Text   string
	UserID string
	Done   bool
}
