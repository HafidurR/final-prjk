package controller

import (
	"net/http"

	"github.com/HafidurR/final-prjk/src/config"
	"github.com/HafidurR/final-prjk/src/helper"
	"github.com/HafidurR/final-prjk/src/models"
	"github.com/gin-gonic/gin"
)


// var db = config.Connect()

func GetAllAuthor(c *gin.Context)  {
	var author []models.Author
	if err := config.Connect().Preload("Books").Find(&author).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": author})
}

func GetDetailAuthor(c *gin.Context)  {
	var author models.Author
	id := c.Param("id")

	if err := config.Connect().Preload("Books").First(&author, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Data not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": author})
}

func CreateAuthor(c *gin.Context) {
	var author models.Author

	if err := c.ShouldBindJSON(&author); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashPassword, _ := helper.HashPassword(author.Password)
	author.Password = string(hashPassword)

	if err := config.Connect().Create(&author).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "success", "success": true})
}

func UpdateAuthor(c *gin.Context)  {
	var author models.Author
	id := c.Param("id")

	if err := config.Connect().First(&author, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Data not found"})
		return
	}

	if err := c.ShouldBindJSON(&author); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.Connect().Save(&author).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"author": author})	
}

func DeleteAuthor(c *gin.Context)  {
	var author models.Author
	id := c.Param("id")

	if err := config.Connect().First(&author, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Data not found"})
		return
	}

	if err := config.Connect().Delete(&author).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Author deleted"})
}