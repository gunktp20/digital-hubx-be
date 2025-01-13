package repository

import (
	userDto "github.com/gunktp20/digital-hubx-be/modules/user/userDto"
	"github.com/gunktp20/digital-hubx-be/pkg/models"
)

type (
	UserRepositoryService interface {
		IsUniqueUser(email string) bool
		CreateOneUser(createUserReq *userDto.CreateUserReq) (*userDto.CreateUserRes, error)
		GetOneUserByEmail(email string) (*models.User, error)
	}
)
