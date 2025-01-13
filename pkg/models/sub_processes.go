package models

import "time"

type SubProcess struct {
	ID          string  `gorm:"type:uuid;primaryKey;not null" json:"id"`
	Name        string  `gorm:"unique;not null" json:"name"`
	Description string  `gorm:"not null" json:"description"`
	ProcessID   string  `gorm:"type:uuid;not null" json:"process_id"` // เพิ่ม index
	Process     Process `gorm:"foreignKey:ProcessID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"process"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
