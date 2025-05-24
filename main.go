package main

import (
	"rest_test/internal"

	"github.com/gin-gonic/gin"
)

func main() {
	db := internal.ConnectDB()

	repos := internal.NewRepository(db)
	handler := internal.NewHandler(repos)

	router := gin.Default()

	// router.POST("/tasks/", Create)
	router.GET("/tasks/", handler.GetAll)
	router.GET("/tasks/:id", handler.GetById)
	// router.PUT("/tasks/:id", Update)
	router.DELETE("/tasks/:id", handler.DeleteById)

	router.Run("localhost:8080")
}
