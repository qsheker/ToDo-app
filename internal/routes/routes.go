package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/qsheker/ToDo-app/internal/handlers"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterRoutes(r *gin.Engine, todoHandler *handlers.TodoHandler, userHandler *handlers.UserHandler) {
	todoRoutes := r.Group("/todos")
	{
		todoRoutes.POST("/", todoHandler.CreateTodo)
		todoRoutes.GET("/", todoHandler.GetAllTodo)
		todoRoutes.GET("/:id", todoHandler.GetTodoByID)
		todoRoutes.GET("/user/:userID", todoHandler.GetTodosByUserID)
		todoRoutes.PUT("/:id", todoHandler.UpdateTodo)
		todoRoutes.DELETE("/:id", todoHandler.DeleteTodo)
		todoRoutes.PATCH("/:id/toggle", todoHandler.ToggleComplete)
	}

	userRoutes := r.Group("/users")
	{
		userRoutes.POST("/", userHandler.CreateUser)
		userRoutes.GET("/:id", userHandler.GetUserById)
		userRoutes.GET("/username/:username", userHandler.GetByUsername)
		userRoutes.PUT("/", userHandler.Update)
		userRoutes.DELETE("/:id", userHandler.Delete)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
