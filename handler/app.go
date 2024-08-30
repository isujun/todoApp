package handler

import (
	"portfolio/todoApp/infrastructure/config"
	"portfolio/todoApp/infrastructure/database"
	todo_repository "portfolio/todoApp/repository/todo_repository/todo_repo"
	"portfolio/todoApp/service"

	"github.com/gin-gonic/gin"
)

func StartApp() {

	config.LoadAppConfig()
	db := database.GetDatabaseInstance()

	// Dependency injection
	todoRepo := todo_repository.NewTodoRepo(db)
	todoService := service.NewTodoService(todoRepo)
	todoHandler := NewTodoHandler(todoService)

	route := gin.Default()

	todoRoute := route.Group("/todos")
	{
		todoRoute.GET("/", todoHandler.GetTodos)
		todoRoute.POST("/", todoHandler.CreateTodo)
		todoRoute.GET("/:todoId", todoHandler.GetTodo)
		todoRoute.PUT("/:todoId", todoHandler.UpdateTodo)
		todoRoute.DELETE("/:todoId", todoHandler.DeleteTodo)
	}

	route.Run(":" + config.GetAppConfig().Port)
}
