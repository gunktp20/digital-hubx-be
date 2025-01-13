package dto

import "time"

type (
	CreateClassCategoryReq struct {
		ClassCategoryName string `json:"class_category_name" validate:"required"`
	}

	CreateClassCategoryRes struct {
		ID                string    `json:"id"`
		ClassCategoryName string    `json:"class_category_name"`
		CreatedAt         time.Time `json:"created_at"`
		UpdatedAt         time.Time `json:"updated_at"`
	}
)
