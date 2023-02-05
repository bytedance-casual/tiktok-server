package utils

import (
	"io"
	"mime/multipart"
	"os"
)

func ExistFile(filePath string) bool {
	_, err := os.Stat(filePath)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func UploadedFile2Bytes(file *multipart.FileHeader) ([]byte, error) {
	src, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	bytes, err := io.ReadAll(src)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}
