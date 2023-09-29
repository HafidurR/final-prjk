package routes

import (
	"github.com/HafidurR/final-prjk/src/controller"
	"github.com/gin-gonic/gin"
)

func HandleRequest() {
	router := gin.Default()

	//Route for login
	router.POST("/login", controller.Login)

	// Route for user
	router.GET("/author", controller.GetAllAuthor)
	router.GET("/author/:id", controller.GetDetailAuthor)
	router.POST("/author", controller.CreateAuthor)
	router.PUT("/author/:id", controller.UpdateAuthor)
	router.DELETE("/author/:id", controller.DeleteAuthor)

	// Route for book
	router.GET("/book", controller.GetAllBook)
	router.GET("/book/:id", controller.GetDetailBook)
	router.POST("/book", controller.CreateBook)
	router.PUT("/book/:id", controller.UpdateBook)
	router.DELETE("/book/:id", controller.DeleteBook)

	router.Run()
}