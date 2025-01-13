package dto

import (
	"time"

	"github.com/gunktp20/digital-hubx-be/pkg/models"
)

type (
	CreateClassReq struct {
		Title           string           `json:"title" validate:"required,min=5"`
		Description     string           `json:"description" validate:"required,min=5"`
		CoverImage      string           `json:"cover_image"`
		ClassCategoryID string           `json:"class_category_id" validate:"required"`
		ClassLevel      int              `json:"class_level"`
		ClassTier       models.ClassTier `json:"class_tier" validate:"required"`
	}

	CreateClassRes struct {
		ID              string           `json:"id"`
		Title           string           `json:"title"`
		Description     string           `json:"description"`
		CoverImage      string           `json:"cover_image"`
		ClassCategoryID string           `json:"class_category_id"`
		ClassLevel      int              `json:"class_level"`
		ClassTier       models.ClassTier `json:"class_tier"`
		IsActive        bool             `json:"is_active"`
		IsRemove        bool             `json:"is_remove"`
		CreatedAt       time.Time        `json:"created_at"`
		UpdatedAt       time.Time        `json:"updated_at"`
	}
)
