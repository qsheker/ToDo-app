package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/qsheker/ToDo-app/internal/models"
	"github.com/qsheker/ToDo-app/internal/service"
)

type TodoHandler struct {
	service service.TodoService
}

func NewTodoHandler(s service.TodoService) *TodoHandler {
	return &TodoHandler{service: s}
}

// @Summary      Create a new todo
// @Description  Create a new todo task with title, description and completion status
// @Tags         todos
// @Accept       json
// @Produce      json
// @Param        input  body      models.TodoRequest  true  "Todo data"
// @Success      201  {object}  models.Todo
// @Failure      400    {object}  map[string]string
// @Failure      500    {object}  map[string]string
// @Router       /todos [post]
func (h *TodoHandler) CreateTodo(c *gin.Context) {
	var req models.TodoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo, err := h.service.CreateTodo(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, todo)
}

// @Summary      Get todo by ID
// @Description  Retrieve a todo item by its ID
// @Tags         todos
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Todo ID"
// @Success      200  {object}  models.Todo
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /todos/{id} [get]
func (h *TodoHandler) GetTodoByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}
	todo, err := h.service.GetTodoByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "invalid id",
		})
		return
	}
	c.JSON(http.StatusFound, todo)
}

// @Summary      Get all todos
// @Description  Retrieve all todo items from database
// @Tags         todos
// @Accept       json
// @Produce      json
// @Success      200  {array}   models.Todo
// @Failure      400  {object}  map[string]string
// @Router       /todos [get]
func (h *TodoHandler) GetAllTodo(c *gin.Context) {
	todos, err := h.service.GetAllTodos()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, todos)
}

// @Summary      Get todos by user ID
// @Description  Retrieve all todos that belong to a specific user
// @Tags         todos
// @Accept       json
// @Produce      json
// @Param        userID   path      string  true  "User UUID"
// @Success      200      {array}   models.Todo
// @Failure      400      {object}  map[string]string
// @Failure      404      {object}  map[string]string
// @Router       /todos/user/{userID} [get]
func (h *TodoHandler) GetTodosByUserID(c *gin.Context) {
	userIDParam := c.Param("userID")
	userID, err := uuid.Parse(userIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	todos, err := h.service.GetTodosByUserID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, todos)
}

// @Summary      Update a todo
// @Description  Update an existing todo's title or description
// @Tags         todos
// @Accept       json
// @Produce      json
// @Param        id     path      int                  true  "Todo ID"
// @Param        input  body      models.TodoRequest    true  "Updated todo data"
// @Success      200    {object}  models.Todo
// @Failure      400    {object}  map[string]string
// @Failure      500    {object}  map[string]string
// @Router       /todos/{id} [put]
func (h *TodoHandler) UpdateTodo(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	var req models.TodoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo, err := h.service.UpdateTodo(id, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, todo)
}

// @Summary      Delete a todo
// @Description  Delete a todo by its ID
// @Tags         todos
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Todo ID"
// @Success      200  {object}  map[string]string
// @Failure      400  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /todos/{id} [delete]
func (h *TodoHandler) DeleteTodo(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	if err := h.service.DeleteTodo(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "todo deleted"})
}

// @Summary      Toggle todo completion
// @Description  Mark todo as complete or incomplete
// @Tags         todos
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Todo ID"
// @Success      200  {object}  models.Todo
// @Failure      400  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /todos/{id}/toggle [patch]
func (h *TodoHandler) ToggleComplete(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	todo, err := h.service.ToggleComplete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, todo)
}
