package models

import "time"

type UserQuestionAnswer struct {
	ID           string `gorm:"type:uuid;primaryKey;not null" json:"id"`
	UserID       string `gorm:"type:uuid;not null" json:"user_id"`     // เพิ่ม index
	QuestionID   string `gorm:"type:uuid;not null" json:"question_id"` // เพิ่ม index
	ChoiceID     string `gorm:"type:uuid;null" json:"choice_id"`       // เพิ่ม index
	AnswerText   string `gorm:"type:uuid;null" json:"answer_text"`     // เพิ่ม index
	QuestionType Class  `gorm:"type:question_types;not null" json:"question_type"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
