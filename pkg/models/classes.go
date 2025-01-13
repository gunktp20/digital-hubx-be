package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ClassTier string

const (
	Essential ClassTier = "essential"
	Literacy  ClassTier = "literacy"
	Mastery   ClassTier = "mastery"
)

type Class struct {
	ID              string        `gorm:"type:uuid;primaryKey;not null" json:"id"`
	Title           string        `gorm:"not null;unique" json:"title"`
	Description     string        `gorm:"not null" json:"description"`
	CoverImage      string        `gorm:"not null" json:"cover_image"`
	ClassCategoryID string        `gorm:"type:uuid;null" json:"class_category_id"`
	ClassCategory   ClassCategory `gorm:"foreignKey:ClassCategoryID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"class_category"`
	ClassTier       ClassTier     `gorm:"type:class_tiers;not null" json:"class_tier"`
	ClassLevel      int           `gorm:"not null;default:1" json:"class_level"`
	IsActive        bool          `gorm:"default:true;not null" json:"is_active"`
	IsRemove        bool          `gorm:"default:false;not null" json:"is_remove"`
	EnableQuestion  bool          `gorm:"default:false;not null" json:"enable_question"`
	CreatedAt       time.Time     `json:"created_at"`
	UpdatedAt       time.Time     `json:"updated_at"`

	ClassSessions        []ClassSession        `gorm:"foreignKey:ClassID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"class_sessions"`
	ClassHighLightImages []ClassHighLightImage `gorm:"foreignKey:ClassID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"class_high_light_images"`
	Questions            []Question            `json:"questions"`
}

func (class *Class) BeforeCreate(tx *gorm.DB) (err error) {
	if class.ID == "" {
		class.ID = uuid.New().String()
	}
	return
}
