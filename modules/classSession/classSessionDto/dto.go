package dto

import (
	"time"

	"github.com/gunktp20/digital-hubx-be/pkg/models"
)

type (
	CreateClassSessionReq struct {
		ClassID     string    `json:"class_id" validate:"required,uuid"`
		Date        time.Time `json:"date" validate:"required"`
		MaxCapacity int       `json:"max_capacity" validate:"required"`
		// ClassSessionStatus models.ClassSessionStatus `json:"class_session_status" validate:"required"`
		StartTime time.Time `json:"start_time" validate:"required"`
		EndTime   time.Time `json:"end_time" validate:"required"`
		Location  string    `json:"location" validate:"required"`
	}

	CreateClassSessionRes struct {
		ID                   string                    `json:"id"`
		ClassID              string                    `json:"class_id"`
		Date                 time.Time                 `json:"date" `
		MaxCapacity          int                       `json:"max_capacity" `
		CancellationDeadline time.Time                 `json:"cancellation_deadline"`
		ClassSessionStatus   models.ClassSessionStatus `json:"class_session_status" `
		StartTime            time.Time                 `json:"start_time" `
		EndTime              time.Time                 `json:"end_time" `
		Location             string                    `json:"location" `
		CreatedAt            time.Time                 `json:"created_at"`
		UpdatedAt            time.Time                 `json:"updated_at"`
	}

	ClassSessionsRes struct {
		ID                 string                    `json:"id"`
		ClassID            string                    `json:"class_id"`
		Class              models.Class              `json:"class"`
		Date               time.Time                 `json:"date" `
		MaxCapacity        int                       `json:"max_capacity" `
		ClassSessionStatus models.ClassSessionStatus `json:"class_session_status" `
		StartTime          time.Time                 `json:"start_time" `
		EndTime            time.Time                 `json:"end_time" `
		Location           string                    `json:"location" `
		RemainingSeats     int                       `json:"remaining_seats" `
		CreatedAt          time.Time                 `json:"created_at"`
		UpdatedAt          time.Time                 `json:"updated_at"`
	}
)
