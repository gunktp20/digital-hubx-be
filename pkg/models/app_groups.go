package models

import "time"

type AppGroup struct {
	ID          string `gorm:"type:uuid;primaryKey;not null" json:"id"`
	Name        string `gorm:"unique;not null" json:"name"`
	IconURL     string `gorm:"unique;not null" json:"icon_url"`
	Description string `gorm:"not null" json:"description"`
	IsActive    bool   `gorm:"default:true;not null" json:"is_active"`
	IsRemove    bool   `gorm:"default:false;not null" json:"is_remove"`
	CreatedAt   time.Time
	UpdatedAt   time.Time

	Solutions []Solution `json:"solutions"`
}
