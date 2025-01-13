package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        string `gorm:"type:uuid;primaryKey;not null" json:"id"`
	Email     string `gorm:"not null;unique" form:"email" json:"email" validate:"required,min=7"`
	Password  string `gorm:"not null" json:"password"`
	Role      string `gorm:"not null;default:'user'" json:"role"`
	CreatedAt time.Time
	UpdatedAt time.Time

	UserClassRegistrations []UserClassRegistration `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user_class_registrations"`
	UserQuestionAnswers    []UserQuestionAnswer    `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user_question_answers"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	if user.ID == "" {
		user.ID = uuid.New().String()
	}
	return
}
