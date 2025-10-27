package service

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/qsheker/ToDo-app/internal/repository"
)

const signingKey = "fbce1ceef702296950744b17f161021bc9bcee13bb9063a2b524eef6f3c285dc"

type tokenClaims struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type JwtService interface {
	GenerateToken(username, password string) (string, error)
}

type JwtServiceImpl struct {
	repo repository.UserRepository
}

func NewJwtService(repo repository.UserRepository) JwtService {
	return &JwtServiceImpl{repo: repo}
}

func (s *JwtServiceImpl) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username, password)
	if err != nil {
		return "", err
	}
	claims := tokenClaims{
		UserID:   user.ID.String(),
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(signingKey))
}
