package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ClassSessionStatus string

const (
	FullyBooked ClassSessionStatus = "fullybook"
	Available   ClassSessionStatus = "available"
)

type ClassSession struct {
	ID                   string             `gorm:"type:uuid;primaryKey;not null" json:"id"`
	ClassID              string             `gorm:"type:uuid;not null" json:"class_id"`
	Class                Class              `gorm:"foreignKey:ClassID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"class"`
	Date                 time.Time          `gorm:"not null" json:"date"`
	MaxCapacity          int                `gorm:"not null" json:"max_capacity"`
	ClassSessionStatus   ClassSessionStatus `gorm:"type:class_session_status;not null;default:'available'" json:"class_session_status"`
	StartTime            time.Time          `gorm:"not null" json:"start_time"`
	EndTime              time.Time          `gorm:"not null" json:"end_time"`
	Location             string             `gorm:"not null" json:"location"`
	CancellationDeadline time.Time          `gorm:"not null" json:"cancellation_deadline"`
	CreatedAt            time.Time
	UpdatedAt            time.Time
}

func (classSession *ClassSession) BeforeCreate(tx *gorm.DB) (err error) {
	if classSession.ID == "" {
		classSession.ID = uuid.New().String()
	}
	return
}
