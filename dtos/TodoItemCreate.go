package dtos

type TodoItemCreate struct {
	Id          int    `json:"-" gorm:"column:id;"`
	Title       string `json:"title" gorm:"column:title;"`
	Description string `json:"description" gorm:"column:description"`
	//Status      string `json:"status" gorm:"column:status"`
}

func (TodoItemCreate) TableName() string {
	return "todo_items"
}
