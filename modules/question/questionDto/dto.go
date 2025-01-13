package dto

import (
	"time"

	"github.com/gunktp20/digital-hubx-be/pkg/models"
)

type (
	CreateQuestionReq struct {
		Description  string               `json:"description" validate:"required"`
		ClassID      string               `json:"class_id" validate:"required,uuid"`
		QuestionType models.QuestionTypes `json:"question_type" validate:"required"`
	}

	QuestionClass struct {
		Title       string           `json:"title"`
		Description string           `json:"description"`
		ClassTier   models.ClassTier `json:"class_tier"`
		ClassLevel  int              `json:"class_level"`
	}

	CreateQuestionRes struct {
		ID           string               `json:"id"`
		Description  string               `json:"description"`
		ClassID      string               `json:"class_id"`
		Class        QuestionClass        `json:"class"`
		QuestionType models.QuestionTypes `json:"question_type"`
		CreatedAt    time.Time            `json:"created_at"`
		UpdatedAt    time.Time            `json:"updated_at"`

		Choices []models.Choice `json:"choices" validate:"required"`
	}
)
