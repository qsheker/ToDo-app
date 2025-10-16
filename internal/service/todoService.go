package service

import (
	"github.com/qsheker/ToDo-app/internal/models"
)

type TodoService interface {
	CreateTodo(req *models.TodoRequest) (*models.Todo, error)
	GetTodoByID(id int64) (*models.Todo, error)
	GetAllTodos() ([]models.Todo, error)
	GetTodosByUserID(userID int64) ([]models.Todo, error)
	UpdateTodo(id int64, req *models.TodoRequest) (*models.Todo, error)
	DeleteTodo(id int64) error
	ToggleComplete(id int64) (*models.Todo, error)
}
