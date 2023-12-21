package file

import (
	"os"
)

func IsExist(path string) bool {
	_, err := os.Stat(path)

	if err != nil {
		return false
	}

	return true
}

func Create(path string) error {
	file, err := os.Create(path)

	if err != nil {
		return err
	}

	defer file.Close()

	return nil
}

func Read(path string) (string, error) {
	data, err := os.ReadFile(path)

	if err != nil {
		return "", err
	}

	return string(data), nil
}

func Set(filepath string, data []byte) error {
	err := os.WriteFile(filepath, data, os.ModePerm)

	if err != nil {
		return err
	}

	return nil
}

func GetFilename(name string, extension string) string {
	filename := name + "." + extension

	return filename
}
