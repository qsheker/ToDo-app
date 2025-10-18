package service

import (
	"errors"
	"log"
	"time"

	"github.com/qsheker/ToDo-app/internal/models"
	"github.com/qsheker/ToDo-app/internal/repository"
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

type TodoServiceImpl struct {
	repo repository.TodoRepository
}

func NewTodoService(repo repository.TodoRepository) TodoService {
	return &TodoServiceImpl{repo: repo}
}
func (s *TodoServiceImpl) CreateTodo(req *models.TodoRequest) (*models.Todo, error) {
	if req.Title == "" {
		return nil, errors.New("title is required")
	}
	todo := &models.Todo{
		Title:       req.Title,
		Description: req.Description,
		Completed:   req.Completed,
	}
	if err := s.repo.Create(todo); err != nil {
		log.Println("Error creating a todo: ", err)
		return nil, err
	}
	return todo, nil
}
func (s *TodoServiceImpl) GetTodoByID(id int64) (*models.Todo, error) {
	todo, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return todo, nil
}
func (s *TodoServiceImpl) GetAllTodos() ([]models.Todo, error) {
	return s.repo.GetAll()
}
func (s *TodoServiceImpl) GetTodosByUserID(userID int64) ([]models.Todo, error) {
	todo, err := s.repo.GetByUserID(userID)
	if err != nil {
		return nil, err
	}
	return todo, nil
}
func (s *TodoServiceImpl) UpdateTodo(id int64, req *models.TodoRequest) (*models.Todo, error) {
	todo, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	todo.Title = req.Title
	todo.Description = req.Description
	todo.UpdatedAt = time.Now()

	if err := s.repo.Update(todo); err != nil {
		return nil, err
	}
	return todo, nil
}
func (s *TodoServiceImpl) DeleteTodo(id int64) error {
	return s.repo.Delete(id)
}
func (s *TodoServiceImpl) ToggleComplete(id int64) (*models.Todo, error) {
	if err := s.repo.ToggleComplete(id); err != nil {
		return nil, err
	}
	todo, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return todo, nil
}
