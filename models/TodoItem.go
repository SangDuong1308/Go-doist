package models

import "time"

type TodoItem struct {
	Id          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      string     `json:"status"`
	CreateAt    *time.Time `json:"create_at"`
	UpdateAt    *time.Time `json:"update_at,omitempty"`
}

func (TodoItem) ItemsTableName() string {
	return "todo_items"
}
