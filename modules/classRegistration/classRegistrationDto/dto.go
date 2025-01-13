package dto

import (
	"time"

	"github.com/gunktp20/digital-hubx-be/pkg/models"
)

type (
	CreateClassRegistrationReq struct {
		ClassID        string `json:"class_id" validate:"required,uuid"`
		ClassSessionID string `json:"class_session_id" validate:"required,uuid"`
	}

	CreateClassRegistrationRes struct {
		ID                   string           `json:"id"`
		UserID               string           `json:"user_id"`
		ClassID              string           `json:"class_id"`
		ClassSessionID       string           `json:"class_session_id"`
		UnattendedQuota      int              `json:"unattended_quota"`
		IsBanned             bool             `json:"is_banned"`
		CancellationDeadline time.Time        `json:"cancellation_deadline"`
		RegisteredAt         time.Time        `json:"registered_at"`
		RegStatus            models.RegStatus `json:"reg_status"`
		CreatedAt            time.Time        `json:"created_at"`
		UpdatedAt            time.Time        `json:"updated_at"`
	}

	GetUserRegistrationsRes struct {
		ID                   string           `json:"id"`
		UserID               string           `json:"user_id"`
		ClassID              string           `json:"class_id"`
		ClassSessionID       string           `json:"class_session_id"`
		UnattendedQuota      int              `json:"unattended_quota"`
		IsBanned             bool             `json:"is_banned"`
		CancellationDeadline time.Time        `json:"cancellation_deadline"`
		RegisteredAt         time.Time        `json:"registered_at"`
		RegStatus            models.RegStatus `json:"reg_status"`
		CreatedAt            time.Time        `json:"created_at"`
		UpdatedAt            time.Time        `json:"updated_at"`
	}
)
