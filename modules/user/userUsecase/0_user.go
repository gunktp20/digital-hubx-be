package usecase

import (
	userDto "github.com/gunktp20/digital-hubx-be/modules/user/userDto"
	repository "github.com/gunktp20/digital-hubx-be/modules/user/userRepository"
)

type (
	UserUsecaseService interface {
		CreateUser(createUserReq *userDto.CreateUserReq) (*userDto.CreateUserRes, error)
	}

	userUsecase struct {
		userRepo repository.UserRepositoryService
	}
)

func NewUserUsecase(userRepo repository.UserRepositoryService) UserUsecaseService {
	return &userUsecase{userRepo: userRepo}
}
