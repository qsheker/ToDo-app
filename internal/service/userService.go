package service

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/qsheker/ToDo-app/internal/models"
	"github.com/qsheker/ToDo-app/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Create(user *models.CreateUserRequest) error
	GetByID(id uuid.UUID) (*models.UserResponse, error)
	GetByUsername(username string) (*models.UserResponse, error)
	Update(user *models.User) error
	Delete(id uuid.UUID) error
	ValidatePassword(user *models.User, password string) error
	hashPassword(password string) (string, error)
}

type UserServiceImpl struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &UserServiceImpl{repo: repo}
}
func (s *UserServiceImpl) Create(user *models.CreateUserRequest) error {
	hashedPass, err := s.hashPassword(user.Password)
	if err != nil {
		return err
	}

	entity := &models.User{
		ID:       uuid.New(),
		Name:     user.Name,
		Username: user.Username,
		Password: hashedPass,
	}
	if err := s.repo.Create(entity); err != nil {
		return err
	}
	return nil

}
func (s *UserServiceImpl) GetByID(id uuid.UUID) (*models.UserResponse, error) {
	user, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	userResponse := &models.UserResponse{
		ID:       user.ID,
		Name:     user.Name,
		Username: user.Username,
	}
	return userResponse, nil
}
func (s *UserServiceImpl) GetByUsername(username string) (*models.UserResponse, error) {
	user, err := s.repo.GetByUsername(username)
	if err != nil {
		return nil, err
	}
	userResponse := &models.UserResponse{
		ID:       user.ID,
		Name:     user.Name,
		Username: user.Username,
	}
	return userResponse, nil
}
func (s *UserServiceImpl) Update(user *models.User) error {
	user.UpdatedAt = time.Now()
	return s.repo.Update(user)
}
func (s *UserServiceImpl) Delete(id uuid.UUID) error {
	return s.repo.Delete(id)
}
func (s *UserServiceImpl) ValidatePassword(user *models.User, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return errors.New("invalid password")
	}
	return nil
}
func (s *UserServiceImpl) hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}
