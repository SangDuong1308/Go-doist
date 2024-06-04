package dtos

type TodoItemCreate struct {
	Id          int    `json:"-" gorm:"column:id;"`
	Title       string `json:"title" gorm:"column:title;" binding:"required"`
	Description string `json:"description" gorm:"column:description" binding:"required"`
	Status      string `json:"status" gorm:"column:status" binding:"required,oneof=doing done deleted"`
}

func (TodoItemCreate) TableName() string {
	return "todo_items"
}
