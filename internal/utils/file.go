package utils

import (
	"os"
)

func CreateFile(path string) error {
	file, err := os.Create(path)

	if err != nil {
		return err
	}

	defer file.Close()

	return nil
}

func SetFileData(filepath string, data []byte) error {
	err := os.WriteFile(filepath, data, os.ModePerm)

	if err != nil {
		return err
	}

	return nil
}

func GetFilename(name string, extension string) string {
	if extension == "" {
		return name
	}

	filename := name + "." + extension

	return filename
}
