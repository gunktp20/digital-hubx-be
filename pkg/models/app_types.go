package models

import "time"

type AppType struct {
	ID        string `gorm:"type:uuid;primaryKey;not null" json:"id"`
	AppID     string `gorm:"type:uuid;not null" json:"app_id"`
	TypeID    string `gorm:"type:uuid;not null" json:"type_id"`
	CreatedAt time.Time
	UpdatedAt time.Time

	App  App  `gorm:"foreignKey:AppID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"app"`
	Type Type `gorm:"foreignKey:TypeID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"type"`
}
