package utils

import (
	"errors"
	"fmt"
	"net/http"
)

func ValidateFile(fileBytes []byte, allowedTypes []string, maxFileSize int64) error {
	if int64(len(fileBytes)) > maxFileSize {
		return errors.New("file size exceeds the allowed limit")
	}

	contentType := http.DetectContentType(fileBytes)

	fmt.Println(contentType)

	isValidType := false
	for _, allowedType := range allowedTypes {
		if contentType == allowedType {
			isValidType = true
			break
		}
	}

	if !isValidType {
		formattedAllowedTypes := formatAllowedTypes(allowedTypes)
		return fmt.Errorf("invalid file type. Only %s are allowed \n", formattedAllowedTypes)
	}

	return nil
}
