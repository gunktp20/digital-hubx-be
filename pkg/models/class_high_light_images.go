package models

import "time"

type ClassHighLightImage struct {
	ID        string `gorm:"type:uuid;primaryKey;not null" json:"id"`
	ImageURL  string `gorm:"unique;not null" json:"image_url"`
	ClassID   string `gorm:"type:uuid;not null" json:"class_id"`
	Class     Class  `gorm:"foreignKey:ClassID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"class"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
