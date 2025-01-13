package models

import "time"

type AppProcess struct {
	ID        string `gorm:"type:uuid;primaryKey;not null" json:"id"`
	AppID     string `gorm:"type:uuid;not null" json:"app_id"`     // เพิ่ม index
	ProcessID string `gorm:"type:uuid;not null" json:"process_id"` // เพิ่ม index
	CreatedAt time.Time
	UpdatedAt time.Time

	Process Process `gorm:"foreignKey:ProcessID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"process"`
	App     App     `gorm:"foreignKey:AppID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"app"`
}
