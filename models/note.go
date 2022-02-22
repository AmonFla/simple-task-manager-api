package models

import "time"

type Note struct {
	ID        uint      `json:"id,omitempty"`
	Comment   string    `json:"description,omitempty"`
	UserId    string    `json:"user_id,omitempty"`
	TaskId    string    `json:"task_id,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}
