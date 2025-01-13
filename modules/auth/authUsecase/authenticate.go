package usecase

import (
	"errors"
	"fmt"
	"time"

	authDto "github.com/gunktp20/digital-hubx-be/modules/auth/authDto"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func (u *authUsecase) Authenticate(authenticationReq *authDto.AuthenticationReq) (*authDto.AuthenticationRes, error) {

	user, err := u.userRepo.GetOneUserByEmail(authenticationReq.Email)

	if err != nil {
		if err.Error() == "record not found" {
			return &authDto.AuthenticationRes{}, errors.New("email or password is invalid")
		}
		return &authDto.AuthenticationRes{}, err
	}

	if authenticationReq.Password == "" {
		return &authDto.AuthenticationRes{}, errors.New("provided password is empty")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(authenticationReq.Password)); err != nil {
		fmt.Printf("failed compare hash and password %s", err.Error())
		return &authDto.AuthenticationRes{}, errors.New("email or password is invalid")
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["email"] = user.Email
	claims["user_id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	accessToken, err := token.SignedString([]byte(u.cfg.Jwt.AccessSecretKey))

	if err != nil {
		return &authDto.AuthenticationRes{}, err
	}

	return &authDto.AuthenticationRes{
		AccessToken: accessToken,
		Email:       user.Email,
		UserID:      user.ID,
	}, nil
}
