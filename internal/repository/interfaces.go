package repository

import (
	"github.com/google/uuid"
	"github.com/qsheker/ToDo-app/internal/models"
)

type TodoRepository interface {
	Create(todo *models.Todo) error
	GetByID(id int64) (*models.TodoResponse, error)
	GetAll() ([]models.TodoResponse, error)
	GetByUserID(userID uuid.UUID) ([]models.TodoResponse, error)
	Update(todo *models.TodoResponse) error
	Delete(id int64) error
	ToggleComplete(id int64) error
}

type UserRepository interface {
	Create(user *models.User) error
	GetByID(id uuid.UUID) (*models.User, error)
	GetByUsername(username string) (*models.User, error)
	Update(user *models.User) error
	Delete(id uuid.UUID) error
}
