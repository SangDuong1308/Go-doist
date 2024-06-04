package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"social-todo-list/Services"
	"social-todo-list/database"
	"social-todo-list/models"
)

func indexView(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
	c.JSON(http.StatusOK, gin.H{"message": "TODO APP"})
}

func main() {
	r := gin.Default()
	// CORS configuration
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}

	r.Use(cors.New(config))

	db, err := database.ConnectToDB()

	if err != nil {
		log.Fatalln(err)
	}

	// migration
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		return
	}

	//r := gin.Default()
	//config := cors.DefaultConfig()
	//config.AllowAllOrigins = true
	//r.Use(cors.New(config))

	v1 := r.Group("/api/v1")
	{
		v1.GET("/", indexView)
		items := v1.Group("/items")
		{
			items.POST("", Services.CreateItem(db))
			items.GET("", Services.GetAllItems(db))
			items.GET("/:id", Services.GetItemById(db))
			items.PATCH("/:id", Services.UpdateById(db))
			items.DELETE("/:id", Services.DeleteById(db))
		}
		auths := v1.Group("/auth")
		{
			auths.POST("/register", Services.Register(db))
		}
	}

	r.Run(":5000")
}
