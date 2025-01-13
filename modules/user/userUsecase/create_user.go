package usecase

import (
	"errors"

	userDto "github.com/gunktp20/digital-hubx-be/modules/user/userDto"
	"github.com/gunktp20/digital-hubx-be/pkg/models"

	"golang.org/x/crypto/bcrypt"
)

func (r *userUsecase) CreateUser(createUserReq *userDto.CreateUserReq) (*userDto.CreateUserRes, error) {

	user, err := r.userRepo.GetOneUserByEmail(createUserReq.Email)
	if err == nil && user != (&models.User{}) {
		return &userDto.CreateUserRes{}, errors.New("email was taken")

	}

	// Hashing Password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(createUserReq.Password), bcrypt.DefaultCost)
	if err != nil {
		return &userDto.CreateUserRes{}, errors.New("failed to hash password")
	}

	createUserReq.Password = string(hashedPassword)

	return r.userRepo.CreateOneUser(createUserReq)
}
