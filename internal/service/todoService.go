package service

import (
	"errors"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/qsheker/ToDo-app/internal/models"
	"github.com/qsheker/ToDo-app/internal/repository"
)

type TodoService interface {
	CreateTodo(req *models.TodoRequest) (*models.TodoResponse, error)
	GetTodoByID(id int64) (*models.TodoResponse, error)
	GetAllTodos() ([]models.TodoResponse, error)
	GetTodosByUserID(userID uuid.UUID) ([]models.TodoResponse, error)
	UpdateTodo(id int64, req *models.TodoRequest) (*models.TodoResponse, error)
	DeleteTodo(id int64) error
	ToggleComplete(id int64) (*models.TodoResponse, error)
}

type TodoServiceImpl struct {
	repo repository.TodoRepository
}

func NewTodoService(repo repository.TodoRepository) TodoService {
	return &TodoServiceImpl{repo: repo}
}

func (s *TodoServiceImpl) CreateTodo(req *models.TodoRequest) (*models.TodoResponse, error) {
	if req.Title == "" {
		return nil, errors.New("title is required")
	}

	todo := &models.Todo{
		Title:       req.Title,
		Description: req.Description,
		Completed:   req.Completed,
		UserID:      req.UserID,
	}

	if err := s.repo.Create(todo); err != nil {
		log.Println("Error creating a todo: ", err)
		return nil, err
	}

	return s.todoToResponse(todo), nil
}

func (s *TodoServiceImpl) GetTodoByID(id int64) (*models.TodoResponse, error) {
	todo, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return s.todoToResponse(todo), nil
}

func (s *TodoServiceImpl) GetAllTodos() ([]models.TodoResponse, error) {
	todos, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	responses := make([]models.TodoResponse, len(todos))
	for i, todo := range todos {
		responses[i] = *s.todoToResponse(&todo)
	}
	return responses, nil
}

func (s *TodoServiceImpl) GetTodosByUserID(userID uuid.UUID) ([]models.TodoResponse, error) {
	todos, err := s.repo.GetByUserID(userID)
	if err != nil {
		return nil, err
	}

	responses := make([]models.TodoResponse, len(todos))
	for i, todo := range todos {
		responses[i] = *s.todoToResponse(&todo)
	}
	return responses, nil
}

func (s *TodoServiceImpl) UpdateTodo(id int64, req *models.TodoRequest) (*models.TodoResponse, error) {
	todo, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	todo.Title = req.Title
	todo.Description = req.Description
	todo.Completed = req.Completed
	todo.UpdatedAt = time.Now()

	if err := s.repo.Update(todo); err != nil {
		return nil, err
	}

	return s.todoToResponse(todo), nil
}

func (s *TodoServiceImpl) DeleteTodo(id int64) error {
	return s.repo.Delete(id)
}

func (s *TodoServiceImpl) ToggleComplete(id int64) (*models.TodoResponse, error) {
	if err := s.repo.ToggleComplete(id); err != nil {
		return nil, err
	}

	todo, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return s.todoToResponse(todo), nil
}

// Helper method to convert Todo to TodoResponse
func (s *TodoServiceImpl) todoToResponse(todo *models.Todo) *models.TodoResponse {
	return &models.TodoResponse{
		ID:          todo.ID,
		Title:       todo.Title,
		Description: todo.Description,
		Completed:   todo.Completed,
		CreatedAt:   todo.CreatedAt,
		UpdatedAt:   todo.UpdatedAt,
		UserID:      todo.UserID,
	}
}
