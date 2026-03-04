package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Book struct {
	ID        uuid.UUID      `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Title     string         `json:"title"`
	Author    string         `json:"author"`
	Cover     string         `json:"cover"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index,column:deleted_at"`
}
