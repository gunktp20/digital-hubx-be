package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Choice struct {
	ID          string   `gorm:"type:uuid;primaryKey;not null" json:"id"`
	Description string   `gorm:"unique;not null" json:"description"`    // เพิ่ม index
	QuestionID  string   `gorm:"type:uuid;not null" json:"question_id"` // เพิ่ม index
	Question    Question `gorm:"foreignKey:QuestionID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"question"`
	CreatedAt   time.Time
	UpdatedAt   time.Time

	UserQuestionAnswers []UserQuestionAnswer `json:"user_question_answers"`
}

func (choice *Choice) BeforeCreate(tx *gorm.DB) (err error) {
	if choice.ID == "" {
		choice.ID = uuid.New().String()
	}
	return
}
