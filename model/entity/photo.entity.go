package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Photo struct {
	ID         uuid.UUID      `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Image      string         `json:"image"`
	CategoryID uuid.UUID      `json:"category_id" gorm:"type:uuid"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"-" gorm:"index,column:deleted_at"`
}
