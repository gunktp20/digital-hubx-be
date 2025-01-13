package usecase

import (
	authDto "github.com/gunktp20/digital-hubx-be/modules/auth/authDto"
	userRepository "github.com/gunktp20/digital-hubx-be/modules/user/userRepository"
	"github.com/gunktp20/digital-hubx-be/pkg/config"
)

type (
	AuthUsecaseService interface {
		Authenticate(authenticationReq *authDto.AuthenticationReq) (*authDto.AuthenticationRes, error)
	}

	authUsecase struct {
		userRepo userRepository.UserRepositoryService
		cfg      *config.Config
	}
)

func NewAuthUsecase(userRepo userRepository.UserRepositoryService, cfg *config.Config) AuthUsecaseService {
	return &authUsecase{
		userRepo: userRepo,
		cfg:      cfg,
	}
}
