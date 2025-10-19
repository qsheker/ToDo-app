package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/qsheker/ToDo-app/internal/models"
	"github.com/qsheker/ToDo-app/internal/service"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(s service.UserService) *UserHandler {
	return &UserHandler{service: s}
}

// @Summary      Create a new user
// @Description  Register a new user with name, username, and password
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        input  body      models.CreateUserRequest  true  "User info"
// @Success      201    {object}  map[string]string
// @Failure      400    {object}  map[string]string
// @Failure      500    {object}  map[string]string
// @Router       /users [post]
func (h *UserHandler) CreateUser(c *gin.Context) {
	var req models.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.Create(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "user created"})
}

// @Summary      Get user by ID
// @Description  Retrieve a user by their unique UUID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "User ID"
// @Success      200  {object}  models.UserResponse
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /users/{id} [get]
func (h *UserHandler) GetUserById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid uuid"})
		return
	}
	user, err := h.service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

// @Summary      Get user by Username
// @Description  Retrieve a user by Username
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        username   path      string  true  "Username"
// @Success      200  {object}  models.UserResponse
// @Failure      404  {object}  map[string]string
// @Router       /users/{username} [get]
func (h *UserHandler) GetByUsername(c *gin.Context) {
	usernameParam := c.Param("username")
	user, err := h.service.GetByUsername(usernameParam)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "invalid username",
		})
		return
	}
	c.JSON(http.StatusOK, user)
}

// @Summary      Update user info
// @Description  Update name, username or password for a user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        input  body      models.UpdateUserRequest  true  "Updated user info"
// @Success      200    {object}  map[string]string
// @Failure      400    {object}  map[string]string
// @Router       /users [put]
func (h *UserHandler) Update(c *gin.Context) {
	var user models.UpdateUserRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := h.service.Update(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "User updated",
	})
}

// @Summary      Delete user by ID
// @Description  Delete user from database by their UUID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "User ID"
// @Success      200  {object}  map[string]string
// @Failure      400  {object}  map[string]string
// @Router       /users/{id} [delete]
func (h *UserHandler) Delete(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid uuid"})
		return
	}
	if err := h.service.Delete(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "deleted successfully",
	})
}
