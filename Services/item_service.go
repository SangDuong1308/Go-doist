package Services

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"social-todo-list/dtos"
	"social-todo-list/models"
	"strconv"
	"strings"
)

func CreateItem(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var dataItem dtos.TodoItemCreate

		if err := c.ShouldBind(&dataItem); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// preprocess title - trim all spaces
		dataItem.Title = strings.TrimSpace(dataItem.Title)

		if dataItem.Title == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "title cannot be blank"})
			return
		}

		if err := db.Create(&dataItem).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
		c.JSON(http.StatusOK, gin.H{
			"item_id": dataItem.Id,
		})
	}
}

func GetItemById(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {

		var itemResponse models.TodoItem

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if err := db.First(&itemResponse, id).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
		c.JSON(http.StatusOK, gin.H{
			"item": itemResponse,
		})
	}
}

func GetAllItems(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var paging dtos.Pagination

		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if paging.Page <= 0 {
			paging.Page = 1
		}

		if paging.Limit <= 0 {
			paging.Limit = 10
		}

		offset := (paging.Page - 1) * paging.Limit

		var result []models.TodoItem

		if err := db.Table(models.TodoItem{}.ItemsTableName()).
			Count(&paging.Total).
			Offset(offset).
			Order("id desc").
			Find(&result).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
		c.JSON(http.StatusOK, gin.H{"data": result})
	}
}

func DeleteById(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var item models.TodoItem

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		result := db.Delete(&item, id)

		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": result.Error.Error(),
			})
			return
		}

		if result.RowsAffected == 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Can not find item",
			})
			return
		}

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprint("Deleted success item", id),
		})
	}
}

func UpdateById(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var dataItem models.TodoItem

		if err := c.ShouldBind(&dataItem); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := db.Where("id = ?", id).Updates(&dataItem).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
		c.JSON(http.StatusOK, gin.H{"New data": dataItem})
	}
}
