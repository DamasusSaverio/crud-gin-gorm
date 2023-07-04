package main

import (
	"github.com/damasussaverio/crud-gin-gorm/config"
	"github.com/damasussaverio/crud-gin-gorm/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	config.ConnectDatabase()
	r.GET("/api/products", controllers.GetAll)
	r.GET("/api/products/:id", controllers.GetById)
	r.POST("/api/products", controllers.Create)
	// r.PUT("/api/products/:id", controllers.Update)
	// r.DELETE("/api/products/:id", controllers.Delete)

	r.Run()
}
