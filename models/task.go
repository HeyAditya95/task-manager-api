package models

import "gorm.io/gorm"

// Task represent the task table in the database
type Task struct {
	gorm.Model        //adds id , CreatedAt , UpdatedAt , DeletedAt
	Title      string `json:"title"`
	Completed  bool   `json:"completed"`
}
