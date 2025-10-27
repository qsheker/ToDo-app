package repository

import (
	"github.com/google/uuid"
	"github.com/qsheker/ToDo-app/internal/models"
)

type TodoRepository interface {
	Create(todo *models.Todo) error
	GetByID(id int64) (*models.Todo, error)
	GetAll() ([]models.Todo, error)
	GetByUserID(userID uuid.UUID) ([]models.Todo, error)
	Update(todo *models.Todo) error
	Delete(id int64) error
	ToggleComplete(id int64) error
}

type UserRepository interface {
	Create(user *models.User) error
	GetByID(id uuid.UUID) (*models.User, error)
	GetByUsername(username string) (*models.User, error)
	Update(user *models.User) error
	Delete(id uuid.UUID) error
	GetUser(username, password string) (models.User, error)
}
