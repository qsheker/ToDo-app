package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qsheker/ToDo-app/internal/models"
	"github.com/qsheker/ToDo-app/internal/service"
)

type AuthHandler struct {
	userService service.UserService
	jwtService  service.JwtService
}

func NewAuthHandler(userService service.UserService, jwtService service.JwtService) *AuthHandler {
	return &AuthHandler{userService: userService, jwtService: jwtService}
}

// @Summary      Sign in
// @Description  Authenticate user and return a JWT token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        input  body      models.LoginRequest  true  "User credentials"
// @Success      200    {object}  map[string]string     "token"
// @Failure      400    {object}  map[string]string     "Invalid request"
// @Failure      502    {object}  map[string]string     "Failed to generate token"
// @Router       /auth/sign-in [post]
func (h *AuthHandler) SignIn(c *gin.Context) {
	var user models.LoginRequest
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	token, err := h.jwtService.GenerateToken(user.Username, user.Password)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

// SignUp godoc
// @Summary      Sign up
// @Description  Register a new user with name, username, and password
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        input  body      models.CreateUserRequest  true  "User registration info"
// @Success      200  {object}  map[string]string  "Successfully registered"
// @Failure      400  {object}  map[string]string  "Invalid input"
// @Failure      502  {object}  map[string]string  "Server error"
// @Router       /auth/sign-up [post]
func (h *AuthHandler) SignUp(c *gin.Context) {
	var input models.CreateUserRequest
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := h.userService.Create(&input)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"username": input.Username,
	})
}
