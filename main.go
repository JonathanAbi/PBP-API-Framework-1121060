package main

import (
	"github.com/LATIHAN_GIN/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/users", controller.GetUsers)
	router.POST("/users", controller.InsertUser)
	router.PUT("/users/:id", controller.UpdateUser)
	router.DELETE("/users/:id", controller.DeleteUser)

	router.Run("localhost:8080")
}
