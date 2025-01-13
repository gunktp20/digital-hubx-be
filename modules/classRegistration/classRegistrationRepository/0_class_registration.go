package repository

import (
	classRegistrationDto "github.com/gunktp20/digital-hubx-be/modules/classRegistration/classRegistrationDto"
	"github.com/gunktp20/digital-hubx-be/pkg/models"
)

type (
	ClassRegistrationRepositoryService interface {
		CreateClassRegistration(createClassRegistrationReq *classRegistrationDto.CreateClassRegistrationReq, userID string) (*classRegistrationDto.CreateClassRegistrationRes, error)
		HasUserRegistered(userID string, classSessionID string) (bool, error)
		GetUserRegistrations(userID string, page int, limit int) (*[]classRegistrationDto.GetUserRegistrationsRes, int64, error)
		CountRegistrationWithClassSessionID(classSessionID string) (int, error)
	}
)

func (r *classRegistrationGormRepository) CreateClassRegistration(createClassRegistrationReq *classRegistrationDto.CreateClassRegistrationReq, userID string) (*classRegistrationDto.CreateClassRegistrationRes, error) {

	classRegistration := models.UserClassRegistration{
		UserID:         userID,
		ClassID:        createClassRegistrationReq.ClassID,
		ClassSessionID: createClassRegistrationReq.ClassSessionID,
	}

	if err := r.db.Create(&classRegistration).Error; err != nil {
		return &classRegistrationDto.CreateClassRegistrationRes{}, err
	}

	return &classRegistrationDto.CreateClassRegistrationRes{
		ID:              classRegistration.ID,
		UserID:          classRegistration.UserID,
		ClassID:         classRegistration.ClassID,
		ClassSessionID:  classRegistration.ClassSessionID,
		UnattendedQuota: classRegistration.UnattendedQuota,
		IsBanned:        classRegistration.IsBanned,
		RegisteredAt:    classRegistration.RegisteredAt,
		RegStatus:       classRegistration.RegStatus,
		CreatedAt:       classRegistration.CreatedAt,
		UpdatedAt:       classRegistration.UpdatedAt,
	}, nil

}

func (r *classRegistrationGormRepository) HasUserRegistered(userID string, classSessionID string) (bool, error) {
	var count int64
	err := r.db.Model(&models.UserClassRegistration{}).
		Where("user_id = ? AND class_session_id = ?", userID, classSessionID).
		Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *classRegistrationGormRepository) GetUserRegistrations(userID string, page int, limit int) (*[]classRegistrationDto.GetUserRegistrationsRes, int64, error) {
	var userClassRegistrations []models.UserClassRegistration
	var total int64

	query := r.db.Model(&models.UserClassRegistration{})

	// Filter by class_level
	if userID != "" {
		query = query.Where("user_id = ?", userID)
	}

	query.Count(&total)

	offset := (page - 1) * limit
	result := query.
		// Preload("ClassSessions", func(db *gorm.DB) *gorm.DB {
		// 	return db.Select("id,class_id,max_capacity,date,start_time,end_time")
		// }).
		Limit(limit).
		Offset(offset).
		Find(&userClassRegistrations)

	var userClassRegistrationsRes []classRegistrationDto.GetUserRegistrationsRes

	for i := range userClassRegistrations {

		userClassRegistrationsRes = append(userClassRegistrationsRes, classRegistrationDto.GetUserRegistrationsRes{
			ID:              userClassRegistrations[i].ID,
			UserID:          userClassRegistrations[i].UserID,
			ClassID:         userClassRegistrations[i].ClassID,
			ClassSessionID:  userClassRegistrations[i].ClassSessionID,
			UnattendedQuota: userClassRegistrations[i].UnattendedQuota,
			IsBanned:        userClassRegistrations[i].IsBanned,
			RegisteredAt:    userClassRegistrations[i].RegisteredAt,
			RegStatus:       userClassRegistrations[i].RegStatus,
			CreatedAt:       userClassRegistrations[i].CreatedAt,
			UpdatedAt:       userClassRegistrations[i].UpdatedAt,
		})
	}

	if result.Error != nil {
		return &[]classRegistrationDto.GetUserRegistrationsRes{}, 0, result.Error
	}

	return &userClassRegistrationsRes, total, nil
}

func (r *classRegistrationGormRepository) CountRegistrationWithClassSessionID(classSessionID string) (int, error) {

	var totalRegistrations int64
	result := r.db.Model(&models.UserClassRegistration{}).
		Where("class_session_id = ?", classSessionID).
		Count(&totalRegistrations)

	if result.Error != nil {
		return 0, result.Error
	}

	return int(totalRegistrations), nil
}
