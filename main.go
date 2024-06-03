package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"time"
)

type TodoItem struct {
	Id          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      string     `json:"status"`
	CreateAt    *time.Time `json:"create_at"`
	UpdateAt    *time.Time `json:"update_at,omitempty"`
}

type TodoItemCreate struct {
	Id          int    `json:"-" gorm:"column:id;"`
	Title       string `json:"title" gorm:"column:title;"`
	Description string `json:"description" gorm:"column:description"`
	//Status      string `json:"status" gorm:"column:status"`
}

func (TodoItemCreate) TableName() string {
	return "todo_items"
}

func main() {
	dsn := os.Getenv("DB_CONN_STRING")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}
	//jsonData, err := json.Marshal(item)
	//
	//if err != nil {
	//	fmt.Println("error:", err)
	//	return
	//}
	//fmt.Println(string(jsonData))
	//
	//jsonStr := "{\"id\":1,\"title\":\"First Item\",\"description\":\"The first one\",\"status\":\"Doing\",\"create_at\":\"2024-06-03T15:56:49.354384Z\",\"update_at\":null}"
	//var item2 TodoItem
	//
	//if err := json.Unmarshal([]byte(jsonStr), &item2); err != nil {
	//	fmt.Println("error:", err)
	//	return
	//}
	//fmt.Println(item2)

	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		items := v1.Group("/items")
		{
			items.POST("", CreateItem(db))
			items.GET("")
			items.GET("/:id", GetItemById(db))
			items.PATCH("/:id")
			items.DELETE("/:id")
		}
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": item,
		})
	})
	r.Run(":3000")
}

func CreateItem(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data TodoItemCreate
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if err := db.Create(&data).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": data.Id,
		})
	}
}

func GetItems(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {

	}
}
