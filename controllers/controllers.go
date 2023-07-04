package controllers

import (
	"net/http"

	"github.com/damasussaverio/crud-gin-gorm/config"
	"github.com/damasussaverio/crud-gin-gorm/models"
	"github.com/gin-gonic/gin"
)

func GetAll(c *gin.Context) {
	var products []models.Products
	config.DB.Find(&products)
	c.JSON(http.StatusOK, gin.H{"products": products})
}

func GetById(c *gin.Context) {

}

func Create(c *gin.Context) {
	var product models.Products

	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	config.DB.Create(&product)
	c.JSON(http.StatusOK, gin.H{"product": product})
}

func Update(c *gin.Context) {

}

func Delete(c *gin.Context) {

}
