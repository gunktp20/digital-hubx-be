package models

import "time"

type UserAppFavorite struct {
	ID        string `gorm:"type:uuid;primaryKey;not null" json:"id"`
	UserID    string `gorm:"type:uuid;not null" json:"user_id"`                                           // เพิ่ม index
	AppID     string `gorm:"type:uuid;not null" json:"app_id"`                                            // เพิ่ม index
	App       App    `gorm:"foreignKey:AppID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"app"`   // เพิ่ม index
	User      User   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user"` // เพิ่ม index
	CreatedAt time.Time
	UpdatedAt time.Time
}
