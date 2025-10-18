package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Todo struct {
	ID          int64          `json:"id" gorm:"primaryKey;autoIncrement"`
	Title       string         `json:"title" gorm:"type:varchar(255);not null"`
	Description string         `json:"description,omitempty" gorm:"type:text"`
	Completed   bool           `json:"completed" gorm:"default:false"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`

	UserID uuid.UUID `json:"user_id" gorm:"type:uuid;not null;index"`
	User   User      `json:"user,omitempty" gorm:"foreignKey:UserID"`
}

type TodoRequest struct {
	Title       string    `json:"title" validate:"required"`
	Description string    `json:"description,omitempty"`
	Completed   bool      `json:"completed"`
	UserID      uuid.UUID `json:"user_id,omitempty"`
}
