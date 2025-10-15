package repository

import (
	"github.com/qsheker/ToDo-app/internal/models"
	"gorm.io/gorm"
)

type gormTodoRepo struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) TodoRepository {
	return &gormTodoRepo{db: db}
}

func (repo *gormTodoRepo) Create(todo *models.Todo) error {
	return repo.db.Create(&todo).Error
}
func (repo *gormTodoRepo) GetByID(id int64) (*models.Todo, error) {
	var todo models.Todo
	err := repo.db.Preload("User").First(&todo, id).Error
	return &todo, err
}
func (repo *gormTodoRepo) GetAll() ([]models.Todo, error) {
	var todos []models.Todo
	err := repo.db.Preload("User").Order("created_at DESC").Find(&todos).Error
	return todos, err
}
func (repo *gormTodoRepo) GetByUserID(userID int64) ([]models.Todo, error) {
	var todos []models.Todo
	err := repo.db.Where("user_id = ?", userID).Order("created_at DESC").Find(&todos).Error
	return todos, err
}
func (repo *gormTodoRepo) Update(todo *models.Todo) error {
	return repo.db.Save(todo).Error
}
func (repo *gormTodoRepo) Delete(id int64) error {
	return repo.db.Delete(&models.Todo{}, id).Error
}
func (repo *gormTodoRepo) ToggleComplete(id int64) error {
	return repo.db.Exec("UPDATE todos SET completed = NOT completed WHERE id = ?", id).Error
}
