package models

import "time"

type AppSubProcess struct {
	ID           string `gorm:"type:uuid;primaryKey;not null" json:"id"`
	AppID        string `gorm:"type:uuid;not null" json:"app_id"`         // เพิ่ม index
	SubProcessID string `gorm:"type:uuid;not null" json:"sub_process_id"` // เพิ่ม index
	CreatedAt    time.Time
	UpdatedAt    time.Time

	App        App        `gorm:"foreignKey:AppID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"app"`
	SubProcess SubProcess `gorm:"foreignKey:SubProcessID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"sub"`
}
