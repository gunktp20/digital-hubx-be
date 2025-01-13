package usecase

import (
	"errors"

	classRepository "github.com/gunktp20/digital-hubx-be/modules/class/classRepository"
	classRegistrationDto "github.com/gunktp20/digital-hubx-be/modules/classRegistration/classRegistrationDto"
	classRegistrationRepository "github.com/gunktp20/digital-hubx-be/modules/classRegistration/classRegistrationRepository"
	classSessionRepository "github.com/gunktp20/digital-hubx-be/modules/classSession/classSessionRepository"
	"github.com/gunktp20/digital-hubx-be/pkg/utils"
)

type (
	ClassRegistrationUsecaseService interface {
		CreateClassRegistration(createClassRegistrationReq *classRegistrationDto.CreateClassRegistrationReq, userID string) (*classRegistrationDto.CreateClassRegistrationRes, error)
		GetUserRegistrations(userID string, page int, limit int) (*[]classRegistrationDto.GetUserRegistrationsRes, int64, error)
	}

	classRegistrationUsecase struct {
		classRegistrationRepo classRegistrationRepository.ClassRegistrationRepositoryService
		classSessionRepo      classSessionRepository.ClassSessionRepositoryService
		classRepository       classRepository.ClassRepositoryService
	}
)

func NewClassRegistrationUsecase(
	classRegistrationRepo classRegistrationRepository.ClassRegistrationRepositoryService,
	classSessionRepo classSessionRepository.ClassSessionRepositoryService,
	classRepository classRepository.ClassRepositoryService,
) ClassRegistrationUsecaseService {
	return &classRegistrationUsecase{
		classRegistrationRepo: classRegistrationRepo,
		classSessionRepo:      classSessionRepo,
		classRepository:       classRepository,
	}
}

func (u *classRegistrationUsecase) CreateClassRegistration(createClassRegistrationReq *classRegistrationDto.CreateClassRegistrationReq, userID string) (*classRegistrationDto.CreateClassRegistrationRes, error) {

	classSession, err := u.classSessionRepo.GetClassSessionById(createClassRegistrationReq.ClassSessionID)
	if err != nil {
		return &classRegistrationDto.CreateClassRegistrationRes{}, err
	}

	if classSession.ClassID != createClassRegistrationReq.ClassID {
		return &classRegistrationDto.CreateClassRegistrationRes{}, errors.New("class id request doesn't match with class id of class session that you provided")
	}

	//  ? Check if the user is already registered for the class session.
	isRegistered, err := u.classRegistrationRepo.HasUserRegistered(userID, createClassRegistrationReq.ClassSessionID)
	if err != nil {
		return &classRegistrationDto.CreateClassRegistrationRes{}, err
	}
	if isRegistered {
		return &classRegistrationDto.CreateClassRegistrationRes{}, errors.New("user has already registered for this class session")
	}

	// ? Check if the registration has reached the maximum capacity
	maxCapacity, err := u.classSessionRepo.GetMaxCapacityOfClassSessionById(createClassRegistrationReq.ClassSessionID)
	if err != nil {
		return &classRegistrationDto.CreateClassRegistrationRes{}, err
	}
	totalRegistrations, err := u.classRegistrationRepo.CountRegistrationWithClassSessionID(createClassRegistrationReq.ClassSessionID)
	if err != nil {
		return &classRegistrationDto.CreateClassRegistrationRes{}, err
	}
	if totalRegistrations >= maxCapacity {
		return &classRegistrationDto.CreateClassRegistrationRes{}, errors.New("registration has reached the maximum capacity")
	}

	// ? Check is event date valid for register
	eventDateValidForReg, err := utils.IsEventDateValidForReg(classSession.Date)
	if err != nil {
		return &classRegistrationDto.CreateClassRegistrationRes{}, err
	}
	if !eventDateValidForReg {
		return &classRegistrationDto.CreateClassRegistrationRes{}, errors.New("registration not allowed for this date")
	}

	return u.classRegistrationRepo.CreateClassRegistration(createClassRegistrationReq, userID)
}

func (u *classRegistrationUsecase) GetUserRegistrations(userID string, page int, limit int) (*[]classRegistrationDto.GetUserRegistrationsRes, int64, error) {

	userClassRegistration, total, err := u.classRegistrationRepo.GetUserRegistrations(userID, page, limit)
	if err != nil {
		return nil, 0, err
	}

	return userClassRegistration, total, nil
}
