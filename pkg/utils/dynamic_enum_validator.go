package utils

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

func DynamicEnumValidator(enumValues []string) validator.Func {
	return func(fl validator.FieldLevel) bool {
		fieldValue := fl.Field().String()
		for _, validValue := range enumValues {
			if strings.ToLower(fieldValue) == strings.ToLower(validValue) {
				return true
			}
		}
		return false
	}
}
