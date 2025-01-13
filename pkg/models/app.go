package models

import "time"

type App struct {
	ID                 string   `gorm:"type:uuid;primaryKey;not null" json:"id"`
	AppName            string   `gorm:"unique;not null" json:"app_name"` // เพิ่ม index
	AppLink            string   `gorm:"not null" json:"app_link"`
	ImageURL           string   `gorm:"not null" json:"image_url"`
	ShortDescription   string   `gorm:"not null" json:"short_description"`
	KeyFeatures        []string `gorm:"type:text[]" json:"key_features"` // ใช้ json ในการจัดการ array
	Overview           string   `gorm:"null" json:"overview"`
	AppGroupID         string   `gorm:"type:uuid;not null" json:"app_group_id"` // เพิ่ม index
	AppGroup           AppGroup `gorm:"foreignKey:AppGroupID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"app_group"`
	CostReduction      bool     `gorm:"default:false;not null" json:"cost_reduction"`
	EfficiencyIncrease bool     `gorm:"default:false;not null" json:"efficiency_increase"`
	ProductionIncrease bool     `gorm:"default:false;not null" json:"production_increase"`
	IsActive           bool     `gorm:"default:true;not null" json:"is_active"`
	IsRemove           bool     `gorm:"default:false;not null" json:"is_remove"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
}
