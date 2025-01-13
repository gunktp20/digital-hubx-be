package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RegStatus string

const (
	Completed  RegStatus = "completed"
	Registered RegStatus = "registered"
)

type UserClassRegistration struct {
	ID              string       `gorm:"type:uuid;primaryKey;not null" json:"id"`
	UserID          string       `gorm:"type:uuid;not null" json:"user_id"`
	User            User         `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user"`
	ClassID         string       `gorm:"type:uuid;not null" json:"class_id"`
	Class           Class        `gorm:"foreignKey:ClassID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"class"`
	ClassSessionID  string       `gorm:"type:uuid;not null" json:"class_session_id"`
	ClassSession    ClassSession `gorm:"foreignKey:ClassSessionID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"class_session"`
	UnattendedQuota int          `gorm:"default:0;not null" json:"unattended_quota"`
	IsBanned        bool         `gorm:"default:false;not null" json:"is_banned"`
	RegisteredAt    time.Time    `gorm:"autoCreateTime" json:"registered_at"`
	RegStatus       RegStatus    `gorm:"type:reg_status;not null;default:registered" json:"reg_status"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (userClassRegistration *UserClassRegistration) BeforeCreate(tx *gorm.DB) (err error) {
	if userClassRegistration.ID == "" {
		userClassRegistration.ID = uuid.New().String()
	}
	return
}
