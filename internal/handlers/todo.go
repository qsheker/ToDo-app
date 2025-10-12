package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/qsheker/ToDo-app/internal/models"
	"github.com/qsheker/ToDo-app/internal/service"
)

type TodoHandler struct {
	todoService service.TodoService
}

func NewTodoHandler(todoService service.TodoService) *TodoHandler {
	return &TodoHandler{todoService: todoService}
}

// CreateTodo creates new todo
// @Summary Create todo
// @Description Create a new todo item
// @Tags todos
// @Accept json
// @Produce json
// @Param todo body models.TodoRequest true "Todo data"
// @Success 200 {object} models.Todo
// @Failure 400 {object} map[string]string
// @Router /todos [post]
func (h *TodoHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.TodoRequest
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	model, err := h.todoService.CreateTodo(todo)
	if err != nil {
		http.Error(w, "Todo is not valid", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model)
}

// GetTodo returns task by id
// @Summary Get todo by ID
// @Description Get todo item by ID
// @Tags todos
// @Produce json
// @Param id path int true "Todo ID"
// @Success 200 {object} models.Todo
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /todos/{id} [get]
func (h *TodoHandler) GetTodo(w http.ResponseWriter, r *http.Request) {
	res := r.PathValue("id")
	id, err := strconv.ParseInt(res, 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	model, err := h.todoService.GetTodo(id)
	if err != nil {
		http.Error(w, "Todo not found!", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model)
}

// GetAll return all todos
// @Summary Get all todos
// @Description Get list of all todos
// @Tags todos
// @Produce json
// @Success 200 {array} models.Todo
// @Router /todos [get]
func (h *TodoHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	todoList := h.todoService.GetAllTodos()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(todoList)
}

// UpdateTodo updates todo
// @Summary Update todo
// @Description Update todo item
// @Tags todos
// @Accept json
// @Produce json
// @Param id path int true "Todo ID"
// @Param todo body models.TodoRequest true "Todo data"
// @Success 200 {object} models.Todo
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /todos/{id} [put]
func (h *TodoHandler) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.TodoRequest

	res := r.PathValue("id")
	id, err := strconv.ParseInt(res, 10, 64)

	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	model, err := h.todoService.UpdateTodo(id, todo)
	if err != nil {
		http.Error(w, "Failed to update: ", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model)
}

// DeleteById delete's the todo task by id
// @Summary Delete todo
// @Description Delete todo item by ID
// @Tags todos
// @Produce json
// @Param id path int true "Todo ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /todos/{id} [delete]
func (h *TodoHandler) DeleteById(w http.ResponseWriter, r *http.Request) {
	res := r.PathValue("id")
	id, err := strconv.ParseInt(res, 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	result := h.todoService.DeleteTodo(id)
	if result != nil {
		http.Error(w, "Failed to delete!", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Todo deleted successfully"})
}
