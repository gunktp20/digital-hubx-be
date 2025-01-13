package dto

import (
	"github.com/gunktp20/digital-hubx-be/pkg/models"
)

type (
	CreateChoiceReq struct {
		Description string `json:"description" validate:"required"`
		QuestionID  string `json:"question_id" validate:"required,uuid"`
	}

	ChoiceClass struct {
		Title       string           `json:"title"`
		Description string           `json:"description"`
		ClassTier   models.ClassTier `json:"class_tier"`
		ClassLevel  int              `json:"class_level"`
	}

	CreateChoiceRes struct {
		ID          string `json:"id"`
		Description string `json:"description"`
		QuestionID  string `json:"question_id"`
	}
)
