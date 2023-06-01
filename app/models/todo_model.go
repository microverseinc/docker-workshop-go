package models

import "gorm.io/gorm"

// Todo represents a todo item
type Todo struct {
	gorm.Model
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}