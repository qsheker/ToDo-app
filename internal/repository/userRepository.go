package repository

import (
	"github.com/google/uuid"
	"github.com/qsheker/ToDo-app/internal/models"
	"gorm.io/gorm"
)

type gormUserRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) TodoRepository {
	return &gormTodoRepo{db: db}
}

func (repo *gormUserRepo) Create(user *models.User) error {
	return repo.db.Create(&user).Error
}
func (repo *gormUserRepo) GetByID(id uuid.UUID) (*models.User, error) {
	var user models.User
	err := repo.db.Preload("Todos").First(&user, id).Error
	return &user, err
}
func (repo *gormUserRepo) GetByUsername(username string) (*models.User, error) {
	var user models.User
	err := repo.db.Where("username = ?", username).First(&user).Error
	return &user, err
}
func (repo *gormUserRepo) Update(user *models.User) error {
	return repo.db.Save(user).Error
}
func (repo *gormUserRepo) Delete(id uuid.UUID) error {
	return repo.db.Delete(&models.User{}, id).Error
}
