package main

import (
	"log"
	"net/http"

	_ "github.com/qsheker/ToDo-app/docs"
	"github.com/qsheker/ToDo-app/internal/handlers"
	"github.com/qsheker/ToDo-app/internal/inMemDB"
	"github.com/qsheker/ToDo-app/internal/service"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title ToDo App API
// @version 1.0
// @description REST API for working with todo's

// @host localhost:8081
// @BasePath /
func main() {
	repo := inMemDB.NewInMemoryRepo()
	todoService := service.NewTodoService(repo)
	todoHandler := handlers.NewTodoHandler(todoService)
	greetHandler := handlers.NewGreetHandler("Aldik")

	mux := http.NewServeMux()

	mux.HandleFunc("GET /swagger/", httpSwagger.WrapHandler)
	mux.HandleFunc("GET /hello", greetHandler.BasicHandler)
	mux.HandleFunc("POST /todos", todoHandler.CreateTodo)
	mux.HandleFunc("GET /todos", todoHandler.GetAll)
	mux.HandleFunc("GET /todos/{id}", todoHandler.GetTodo)
	mux.HandleFunc("PUT /todos/{id}", todoHandler.UpdateTodo)
	mux.HandleFunc("DELETE /todos/{id}", todoHandler.DeleteById)

	log.Println("Server starting on :8081")
	log.Println("Swagger UI available at: http://localhost:8081/swagger/index.html")
	log.Fatal(http.ListenAndServe(":8081", mux))
}
