package usecase

import (
	"errors"

	"github.com/golang-jwt/jwt/v4"
)

func (u *middlewareUsecase) JwtAuthorization(tokenString string) (map[string]interface{}, error) {

	secretKey := u.cfg.Jwt.AccessSecretKey

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("token is invalid")
		}
		return []byte(secretKey), nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return claims, nil
	}

	return nil, errors.New("unable to parse token claims")
}
