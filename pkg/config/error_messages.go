package config

var ValidationMessages = map[string]string{
	"required": "{field} is required",
	"email":    "{field} must be a valid email",
	"min":      "{field} must be at least {param}",
	"max":      "{field} cannot exceed {param}",
}
