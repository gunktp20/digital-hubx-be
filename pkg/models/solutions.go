package models

import "time"

type Solution struct {
	ID          string   `gorm:"type:uuid;primaryKey;not null" json:"id"`
	Name        string   `gorm:"unique;not null" json:"name"` // เพิ่ม index
	Description string   `gorm:"not null" json:"description"`
	AppGroupID  string   `gorm:"type:uuid;not null" json:"app_group_id"`
	AppGroup    AppGroup `gorm:"foreignKey:AppGroupID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"app_group"`
	IsActive    bool     `gorm:"default:true;not null" json:"is_active"`
	IsRemove    bool     `gorm:"default:false;not null" json:"is_remove"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
