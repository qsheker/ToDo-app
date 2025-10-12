package inMemDB

import (
	"time"

	"github.com/qsheker/ToDo-app/internal/models"
)

type TodoRepository interface {
	GetAll() []models.Todo
	GetById(id int64) (*models.Todo, bool)
	Save(todo models.TodoRequest)
	UpdateFully(todo models.Todo) bool
	DeleteById(id int64)
	ValidInList(todo models.TodoRequest) bool
}

type InMemoryRepo struct {
	todoList []models.Todo
	lastID   int64
}

func NewInMemoryRepo() *InMemoryRepo {
	return &InMemoryRepo{
		todoList: make([]models.Todo, 0),
		lastID:   0,
	}
}

func (r *InMemoryRepo) GetAll() []models.Todo {
	return r.todoList
}

func (r *InMemoryRepo) GetById(id int64) (*models.Todo, bool) {
	for _, item := range r.todoList {
		if item.ID == id {
			return &item, true
		}
	}
	return nil, false
}

func (r *InMemoryRepo) ValidInList(req models.TodoRequest) bool {
	for _, item := range r.todoList {
		if req.Title == item.Title {
			return true
		}
	}
	return false
}

func (r *InMemoryRepo) Save(req models.TodoRequest) {
	r.lastID++
	newTodo := models.Todo{
		ID:          r.lastID,
		Title:       req.Title,
		Description: req.Description,
		Completed:   false,
		CreatedAt:   time.Now(),
	}
	r.todoList = append(r.todoList, newTodo)
}

func (r *InMemoryRepo) GetByTitle(title string) (*models.Todo, bool) {
	for _, item := range r.todoList {
		if item.Title == title {
			return &item, true
		}
	}
	return nil, false
}

func (r *InMemoryRepo) UpdateFully(todo models.Todo) bool {
	for i := range r.todoList {
		if r.todoList[i].ID == todo.ID {
			originalCreatedAt := r.todoList[i].CreatedAt
			r.todoList[i] = todo
			r.todoList[i].CreatedAt = originalCreatedAt

			return true
		}
	}
	return false
}

func (r *InMemoryRepo) DeleteById(id int64) {
	result := make([]models.Todo, 0)
	for _, item := range r.todoList {
		if item.ID != id {
			result = append(result, item)
		}
	}
	r.todoList = result
}
