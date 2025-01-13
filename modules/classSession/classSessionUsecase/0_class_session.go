package usecase

import (
	"errors"
	"time"

	classSessionDto "github.com/gunktp20/digital-hubx-be/modules/classSession/classSessionDto"
	classSessionRepository "github.com/gunktp20/digital-hubx-be/modules/classSession/classSessionRepository"
	"github.com/gunktp20/digital-hubx-be/pkg/bucket"
	"github.com/gunktp20/digital-hubx-be/pkg/utils"
)

type (
	ClassSessionUsecaseService interface {
		CreateClassSession(createClassSessionReq *classSessionDto.CreateClassSessionReq) (*classSessionDto.CreateClassSessionRes, error)
		GetAllClassSessions(class_id, class_tier string, page int, limit int) (*[]classSessionDto.ClassSessionsRes, int64, error)
	}

	classSessionUsecase struct {
		classSessionRepo classSessionRepository.ClassSessionRepositoryService
		bucketClient     bucket.BucketClientService
	}
)

func NewClassSessionUsecase(classSessionRepo classSessionRepository.ClassSessionRepositoryService, bucketClient bucket.BucketClientService) ClassSessionUsecaseService {
	return &classSessionUsecase{classSessionRepo: classSessionRepo, bucketClient: bucketClient}
}

func (u *classSessionUsecase) CreateClassSession(createClassSessionReq *classSessionDto.CreateClassSessionReq) (*classSessionDto.CreateClassSessionRes, error) {

	dateFields := []time.Time{
		createClassSessionReq.Date, createClassSessionReq.StartTime, createClassSessionReq.EndTime,
	}

	// ? Check all fields about date and time is future
	_, err := utils.AreAllFutureDates(dateFields)
	if err != nil {
		return &classSessionDto.CreateClassSessionRes{}, err
	}

	// ? Check is date conflict ?
	isDateConflict, err := u.classSessionRepo.CheckSessionDateConflict(createClassSessionReq.ClassID, createClassSessionReq.Date)

	// ? Calculate cancellation deadline
	var cancellationDeadline = createClassSessionReq.Date.AddDate(0, 0, -7)

	if err != nil {
		return &classSessionDto.CreateClassSessionRes{}, err
	}

	if isDateConflict {
		return &classSessionDto.CreateClassSessionRes{}, errors.New("class session date conflicts with an existing session")
	}

	return u.classSessionRepo.CreateClassSession(createClassSessionReq, cancellationDeadline)
}

func (u *classSessionUsecase) GetAllClassSessions(class_id, class_tier string, page int, limit int) (*[]classSessionDto.ClassSessionsRes, int64, error) {
	classSessiones, total, err := u.classSessionRepo.GetAllClassSessions(class_id, class_tier, page, limit)
	if err != nil {
		return nil, 0, err
	}

	return classSessiones, total, nil
}
