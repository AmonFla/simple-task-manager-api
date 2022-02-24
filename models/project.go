package models

import "time"

type Project struct {
	ID          uint         `json:"id,omitempty"`
	Name        string       `json:"name,omitempty"`
	Description string       `json:"description,omitempty"`
	States      ProjectState `json:"states,omitempty"`
	Users       []int        `json:"users,omitempty"`
	CreatedAt   time.Time    `json:"created_at,omitempty"`
	UpdatedAt   time.Time    `json:"updated_at,omitempty"`
	ClosedAt    time.Time    `json:"closed_at,omitempty"`
}
