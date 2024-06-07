package models

import "time"

type Task struct {
	TaskId      uint      `json:"task_id" gorm:"primary_key"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Priority    string    `json:"priority"`
	Status      bool      `json:"Status"`
	DueDate     time.Time `json:"due_date"`
	Tag         string    `json:"tag"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
