package controller

import (
	"net/http"

	"github.com/HafidurR/final-prjk/src/config"
	"github.com/HafidurR/final-prjk/src/models"
	"github.com/gin-gonic/gin"
)

func GetAllBook(c *gin.Context)  {
	var book []models.Book

	if err := config.Connect().Find(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

func GetDetailBook(c *gin.Context)  {
	var book models.Book
	id := c.Param("id")

	if err := config.Connect().First(&book, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Data not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

func CreateBook(c *gin.Context) {
	var book models.Book
	var author models.Author

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.Connect().First(&author, book.AuthorID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Author not found"})
		return
	}

	if err := config.Connect().Create(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "book created successfully"})
}

func UpdateBook(c *gin.Context)  {
	var book models.Book
	var author models.Author
	id := c.Param("id")

	
	if err := config.Connect().First(&author, book.AuthorID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Author not found"})
		return
	}

	if err := config.Connect().First(&book, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Data not found"})
		return
	}

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.Connect().Save(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"book": book})	
}

func DeleteBook(c *gin.Context)  {
	var book models.Book
	id := c.Param("id")

	if err := config.Connect().First(&book, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Data not found"})
		return
	}

	if err := config.Connect().Delete(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
}