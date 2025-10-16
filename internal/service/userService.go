package service

import "github.com/qsheker/ToDo-app/internal/models"

type UserService interface {
	Create(user *models.User) error
	GetByID(id int64) (*models.User, error)
	GetByUsername(username string) (*models.User, error)
	Update(user *models.User) error
	Delete(id int64) error

	CreateUserFromRequest(req *models.CreateUserRequest) (*models.User, error)
	GetUserResponse(id int64) (*models.UserResponse, error)
	ValidatePassword(user *models.User, password string) error
	HashPassword(password string) (string, error)
}
