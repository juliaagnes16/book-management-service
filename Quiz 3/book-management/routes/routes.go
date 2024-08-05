package routes

import (
	"main.go/controllers"
	"main.go/middlewares"

	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		api.Use(middlewares.Auth())
		api.GET("/books", controllers.GetBooks)
		api.POST("/books", controllers.AddBook)
		api.GET("/books/:id", controllers.GetBookByID)
		api.PUT("/books/:id", controllers.UpdateBook)
		api.DELETE("/books/:id", controllers.DeleteBook)

		api.GET("/categories", controllers.GetCategories)
		api.POST("/categories", controllers.AddCategory)
		api.GET("/categories/:id", controllers.GetCategoryByID)
		api.PUT("/categories/:id", controllers.UpdateCategory)
		api.DELETE("/categories/:id", controllers.DeleteCategory)
	}
	return router
}
