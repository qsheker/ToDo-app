package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/qsheker/ToDo-app/docs"
	"github.com/qsheker/ToDo-app/internal/handlers"
	"github.com/qsheker/ToDo-app/internal/repository"
	"github.com/qsheker/ToDo-app/internal/routes"
	"github.com/qsheker/ToDo-app/internal/service"
)

// @title ToDo App API
// @version 1.0
// @description REST API for working with todo's

// @host localhost:8081
// @BasePath /
func main() {
	r := gin.Default()
	injector := repository.Injector()

	todoRepo := repository.NewTodoRepository(injector)
	userRepo := repository.NewUserRepository(injector)

	todoService := service.NewTodoService(todoRepo)
	userService := service.NewUserService(userRepo)

	todoHandler := handlers.NewTodoHandler(todoService)
	userHandler := handlers.NewUserHandler(userService)

	routes.RegisterRoutes(r, todoHandler, userHandler)

	r.Run("localhost:8081")
}
