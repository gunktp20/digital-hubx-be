package utils

import (
	"bytes"
	"io"
	"mime/multipart"
)

func ConvertMultipartFileToBytes(fileHeader *multipart.FileHeader) ([]byte, error) {
	file, err := fileHeader.Open()
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var buf bytes.Buffer
	if _, err := io.Copy(&buf, file); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
