package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ClassCategory struct {
	ID           string `gorm:"type:uuid;primaryKey;not null" json:"id"`
	CategoryName string `gorm:"not null" json:"category_name"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (classCategory *ClassCategory) BeforeCreate(tx *gorm.DB) (err error) {
	if classCategory.ID == "" {
		classCategory.ID = uuid.New().String()
	}
	return
}
