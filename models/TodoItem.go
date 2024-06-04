package models

import "time"

type TodoItem struct {
	Id          int        `json:"id" gorm:"column:id;"`
	Title       string     `json:"title" gorm:"column:title;"`
	Description string     `json:"description" gorm:"column:description;"`
	Status      string     `json:"status" gorm:"column:status;"`
	CreateAt    *time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdateAt    *time.Time `json:"updated_at" gorm:"column:updated_at;"`
}

func (TodoItem) ItemsTableName() string {
	return "todo_items"
}
