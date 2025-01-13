package utils

import (
	"strings"

	"github.com/go-playground/validator"
	"github.com/gunktp20/digital-hubx-be/pkg/config"
)

func TranslateValidationError(err validator.ValidationErrors) map[string]string {
	errors := make(map[string]string)
	for _, e := range err {
		field := e.Field()
		tag := e.Tag()
		param := e.Param()

		// ดึงข้อความจาก Config
		messageTemplate, exists := config.ValidationMessages[tag]
		if !exists {
			messageTemplate = "Invalid value for {field}" // แก้ไข: ใช้ string แทน map[string]string
		}

		// แทนที่ {field} และ {param}
		message := strings.ReplaceAll(messageTemplate, "{field}", field) // แก้ไข: messageTemplate เป็น string
		message = strings.ReplaceAll(message, "{param}", param)

		errors[field] = message
	}
	return errors
}
