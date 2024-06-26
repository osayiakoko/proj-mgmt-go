package data

import (
	"time"

	"github.com/osayiakoko/project-mgmt-sys/internal/validator"
)

type Task struct {
	ID          int64        `json:"id"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Priority    TaskPriority `json:"priority"`
	Status      TaskStatus   `json:"status"`
	CreatedAt   time.Time    `json:"createdAt"`
	UpdatedAt   time.Time    `json:"updatedAt"`
}

func ValidateTask(v *validator.Validator, task *Task) {
	v.Check(task.Title != "", "title", "must be provided")
	v.Check(task.Description != "", "description", "must be provided")
	v.Check(task.Priority != "", "priority", "must be provided")
	v.Check(task.Status != "", "status", "must be provided")
}
