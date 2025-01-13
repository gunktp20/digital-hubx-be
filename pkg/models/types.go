package models

import "time"

type Type struct {
	ID          string `gorm:"type:uuid;primaryKey;not null" json:"id"`
	Name        string `gorm:"unique;not null" json:"name"`
	Description string `gorm:"not null" json:"description"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
