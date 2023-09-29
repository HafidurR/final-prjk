package controller

import (
	"net/http"

	"github.com/HafidurR/final-prjk/src/config"
	"github.com/HafidurR/final-prjk/src/models"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var author models.Author
 
	if err := config.Connect().Where("email = ?", author.Email).First(&author).Error; 
	err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": author})
}