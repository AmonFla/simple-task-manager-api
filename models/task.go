package models

import "time"

type Task struct {
	ID          uint      `json:"id,omitempty"`
	Title       string    `json:"title,omitempty"`
	Description string    `json:"description,omitempty"`
	ProjectId   string    `json:"project_id,omitempty"`
	UserId      string    `json:"user_id,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}
