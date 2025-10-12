package service

import (
	"errors"
	"strings"

	"github.com/qsheker/ToDo-app/internal/inMemDB"
	"github.com/qsheker/ToDo-app/internal/models"
)

var (
	ErrTodoNotFound = errors.New("todo not found")
	ErrInvalidTitle = errors.New("title cannot be empty")
)

type TodoService interface {
	CreateTodo(req models.TodoRequest) (*models.Todo, error)
	GetTodo(id int64) (*models.Todo, error)
	GetAllTodos() []models.Todo
	UpdateTodo(id int64, req models.TodoRequest) (*models.Todo, error)
	DeleteTodo(id int64) error

	ToggleComplete(id int64) (*models.Todo, error)
	MarkCompleted(id int64) (*models.Todo, error)
	MarkActive(id int64) (*models.Todo, error)

	GetActiveTodos() []models.Todo
	GetCompletedTodos() []models.Todo
	SearchByTitle(title string) []models.Todo

	ValidateTodo(req models.TodoRequest) error
	Exists(id int64) bool
}

type TodoServiceImpl struct {
	repo inMemDB.TodoRepository
}

func NewTodoService(repo inMemDB.TodoRepository) *TodoServiceImpl {
	return &TodoServiceImpl{repo: repo}
}

func (ts *TodoServiceImpl) ValidateTodo(req models.TodoRequest) error {
	if strings.TrimSpace(req.Title) == "" {
		return ErrInvalidTitle
	}
	return nil
}

func (ts *TodoServiceImpl) Exists(id int64) bool {
	_, found := ts.repo.GetById(id)
	return found
}

func (ts *TodoServiceImpl) CreateTodo(req models.TodoRequest) (*models.Todo, error) {
	if err := ts.ValidateTodo(req); err != nil {
		return nil, err
	}

	ts.repo.Save(req)

	allTodos := ts.repo.GetAll()
	if len(allTodos) == 0 {
		return nil, errors.New("failed to create todo")
	}

	return &allTodos[len(allTodos)-1], nil
}

func (ts *TodoServiceImpl) GetTodo(id int64) (*models.Todo, error) {
	todo, found := ts.repo.GetById(id)
	if !found {
		return nil, ErrTodoNotFound
	}
	return todo, nil
}

func (ts *TodoServiceImpl) GetAllTodos() []models.Todo {
	return ts.repo.GetAll()
}

func (ts *TodoServiceImpl) UpdateTodo(id int64, req models.TodoRequest) (*models.Todo, error) {
	if err := ts.ValidateTodo(req); err != nil {
		return nil, err
	}

	existing, found := ts.repo.GetById(id)
	if !found {
		return nil, ErrTodoNotFound
	}

	updatedTodo := models.Todo{
		ID:          existing.ID,
		Title:       req.Title,
		Description: req.Description,
		Completed:   req.Completed,
		CreatedAt:   existing.CreatedAt,
	}

	success := ts.repo.UpdateFully(updatedTodo)
	if !success {
		return nil, errors.New("failed to update todo")
	}

	return &updatedTodo, nil
}

func (ts *TodoServiceImpl) DeleteTodo(id int64) error {
	if !ts.Exists(id) {
		return ErrTodoNotFound
	}
	ts.repo.DeleteById(id)
	return nil
}

func (ts *TodoServiceImpl) ToggleComplete(id int64) (*models.Todo, error) {
	todo, found := ts.repo.GetById(id)
	if !found {
		return nil, ErrTodoNotFound
	}
	todo.Completed = !todo.Completed
	success := ts.repo.UpdateFully(*todo)
	if !success {
		return nil, errors.New("failed to toggle todo status")
	}
	return todo, nil
}

func (ts *TodoServiceImpl) MarkCompleted(id int64) (*models.Todo, error) {
	todo, found := ts.repo.GetById(id)
	if !found {
		return nil, ErrTodoNotFound
	}
	todo.Completed = true
	success := ts.repo.UpdateFully(*todo)
	if !success {
		return nil, errors.New("failed to mark todo as completed")
	}
	return todo, nil
}

func (ts *TodoServiceImpl) MarkActive(id int64) (*models.Todo, error) {
	todo, found := ts.repo.GetById(id)
	if !found {
		return nil, ErrTodoNotFound
	}
	todo.Completed = false
	success := ts.repo.UpdateFully(*todo)
	if !success {
		return nil, errors.New("failed to mark todo as active")
	}
	return todo, nil
}

func (ts *TodoServiceImpl) GetActiveTodos() []models.Todo {
	allTodos := ts.repo.GetAll()
	var active []models.Todo
	for _, todo := range allTodos {
		if !todo.Completed {
			active = append(active, todo)
		}
	}
	return active
}

func (ts *TodoServiceImpl) GetCompletedTodos() []models.Todo {
	allTodos := ts.repo.GetAll()
	var completed []models.Todo
	for _, todo := range allTodos {
		if todo.Completed {
			completed = append(completed, todo)
		}
	}
	return completed
}

func (ts *TodoServiceImpl) SearchByTitle(title string) []models.Todo {
	allTodos := ts.repo.GetAll()
	var result []models.Todo
	for _, todo := range allTodos {
		if strings.Contains(strings.ToLower(todo.Title), strings.ToLower(title)) {
			result = append(result, todo)
		}
	}
	return result
}
