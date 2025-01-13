package models

import "time"

type AppPreviewImage struct {
	ID        string `gorm:"type:uuid;primaryKey;not null" json:"id"`
	AppID     string `gorm:"type:uuid;not null" json:"app_id"` // เพิ่ม index
	Image     string `gorm:"not null" json:"image"`
	CreatedAt time.Time
	UpdatedAt time.Time

	App App `gorm:"foreignKey:AppID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"app"`
}
