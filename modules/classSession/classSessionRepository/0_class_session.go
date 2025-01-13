package repository

import (
	"errors"
	"fmt"
	"time"

	classSessionDto "github.com/gunktp20/digital-hubx-be/modules/classSession/classSessionDto"
	"github.com/gunktp20/digital-hubx-be/pkg/models"
)

type (
	ClassSessionRepositoryService interface {
		CreateClassSession(createClassSessionReq *classSessionDto.CreateClassSessionReq, cancellationDeadline time.Time) (*classSessionDto.CreateClassSessionRes, error)
		GetAllClassSessions(class_id, class_tier string, page int, limit int) (*[]classSessionDto.ClassSessionsRes, int64, error)
		CheckSessionDateConflict(classID string, date time.Time) (bool, error)
		GetClassSessionById(classSessionID string) (*models.ClassSession, error)
		GetMaxCapacityOfClassSessionById(classSessionID string) (int, error)
	}
)

func (r *classSessionGormRepository) CheckSessionDateConflict(classID string, date time.Time) (bool, error) {
	var count int64
	err := r.db.Model(&models.ClassSession{}).
		Where("class_id = ? AND date = ?", classID, date).
		Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil // ถ้า count > 0 แสดงว่ามี session ซ้ำ
}

func (r *classSessionGormRepository) CreateClassSession(createClassSessionReq *classSessionDto.CreateClassSessionReq, cancellationDeadline time.Time) (*classSessionDto.CreateClassSessionRes, error) {

	classSession := models.ClassSession{
		ClassID:              createClassSessionReq.ClassID,
		Date:                 createClassSessionReq.Date,
		MaxCapacity:          createClassSessionReq.MaxCapacity,
		CancellationDeadline: cancellationDeadline,
		StartTime:            createClassSessionReq.StartTime,
		EndTime:              createClassSessionReq.EndTime,
		Location:             createClassSessionReq.Location,
	}

	if err := r.db.Create(&classSession).Error; err != nil {
		return &classSessionDto.CreateClassSessionRes{}, err
	}

	return &classSessionDto.CreateClassSessionRes{
		ID:                   classSession.ID,
		ClassID:              classSession.ClassID,
		Date:                 classSession.Date,
		MaxCapacity:          classSession.MaxCapacity,
		CancellationDeadline: classSession.CancellationDeadline,
		ClassSessionStatus:   classSession.ClassSessionStatus,
		StartTime:            classSession.StartTime,
		EndTime:              classSession.EndTime,
		Location:             classSession.Location,
		CreatedAt:            classSession.CreatedAt,
		UpdatedAt:            classSession.UpdatedAt,
	}, nil
}

func (r *classSessionGormRepository) GetAllClassSessions(class_id, class_tier string, page int, limit int) (*[]classSessionDto.ClassSessionsRes, int64, error) {
	var classSessions []models.ClassSession
	var total int64

	query := r.db.Model(&models.ClassSession{})

	// Filter by class_id
	if class_id != "" {
		query = query.Where("class_id = ?", class_id)
	}

	fmt.Print("class_tier", class_tier)

	// Filter by class_tier
	if class_tier != "" {
		query = query.Joins("JOIN classes ON classes.id = class_sessions.class_id").
			Where("classes.class_tier = ?", class_tier)
	}

	query.Count(&total)

	offset := (page - 1) * limit
	result := query.
		Preload("Class"). // Load related Class data
		Limit(limit).
		Offset(offset).
		Find(&classSessions)

	if result.Error != nil {
		return &[]classSessionDto.ClassSessionsRes{}, 0, result.Error
	}

	// Add remaining slots field
	var classSessionsRes []classSessionDto.ClassSessionsRes

	for i := range classSessions {
		var totalRegistrations int64

		r.db.Model(&models.UserClassRegistration{}).
			Where("class_session_id = ?", classSessions[i].ID).
			Count(&totalRegistrations)

		// Use append to add elements to the slice
		classSessionsRes = append(classSessionsRes, classSessionDto.ClassSessionsRes{
			ID:                 classSessions[i].ID,
			ClassID:            classSessions[i].ClassID,
			Date:               classSessions[i].Date,
			MaxCapacity:        classSessions[i].MaxCapacity,
			ClassSessionStatus: classSessions[i].ClassSessionStatus,
			Class:              classSessions[i].Class,
			StartTime:          classSessions[i].StartTime,
			EndTime:            classSessions[i].EndTime,
			Location:           classSessions[i].Location,
			RemainingSeats:     int(classSessions[i].MaxCapacity) - int(totalRegistrations),
			CreatedAt:          classSessions[i].CreatedAt,
			UpdatedAt:          classSessions[i].UpdatedAt,
		})
	}

	return &classSessionsRes, total, nil
}

func (r *classSessionGormRepository) GetClassSessionById(classSessionID string) (*models.ClassSession, error) {
	var classSession = new(models.ClassSession)
	result := r.db.First(&classSession, "id = ?", classSessionID)

	if result.Error != nil {
		return &models.ClassSession{}, result.Error
	}

	if result.RowsAffected == 0 {
		return &models.ClassSession{}, errors.New("class session record not found")
	}

	return classSession, nil
}

func (r *classSessionGormRepository) GetMaxCapacityOfClassSessionById(classSessionID string) (int, error) {
	var classSession = new(models.ClassSession)
	result := r.db.First(&classSession, "id = ?", classSessionID)

	if result.Error != nil {
		return 0, result.Error
	}

	if result.RowsAffected == 0 {
		return 0, errors.New("class session record not found")
	}

	return classSession.MaxCapacity, nil
}
