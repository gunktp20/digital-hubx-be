package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type QuestionTypes string

const (
	QChoice QuestionTypes = "choice"
	QText   QuestionTypes = "text"
)

type Question struct {
	ID           string        `gorm:"type:uuid;primaryKey;not null" json:"id"`
	Description  string        `gorm:"unique;not null" json:"description"` // เพิ่ม index
	ClassID      string        `gorm:"type:uuid;not null" json:"class_id"` // เพิ่ม index
	Class        Class         `gorm:"foreignKey:ClassID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"class"`
	QuestionType QuestionTypes `gorm:"type:question_types;not null" json:"question_type"`
	CreatedAt    time.Time
	UpdatedAt    time.Time

	Choices             []Choice             `json:"choices"`
	UserQuestionAnswers []UserQuestionAnswer `json:"user_question_answers"`
}

func (question *Question) BeforeCreate(tx *gorm.DB) (err error) {
	if question.ID == "" {
		question.ID = uuid.New().String()
	}
	return
}
