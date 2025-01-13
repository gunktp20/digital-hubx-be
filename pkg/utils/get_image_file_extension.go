package utils

import (
	"errors"
	"net/http"
)

func GetImageFileExtension(fileBytes []byte) (string, error) {
	contentType := http.DetectContentType(fileBytes)

	allowedExtensions := map[string]string{
		"image/jpeg":                ".jpg",
		"image/png":                 ".png",
		"image/gif":                 ".gif",
		"image/svg+xml":             ".svg",
		"text/xml;":                 ".svg",
		"text/plain;":               ".svg",
		"text/xml; charset=utf-8":   ".svg",
		"text/plain; charset=utf-8": ".svg",
	}

	if ext, ok := allowedExtensions[contentType]; ok {
		return ext, nil
	}

	return "", errors.New("unsupported file type")
}
