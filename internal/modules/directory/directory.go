package directory

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
	err := os.MkdirAll(path, os.ModePerm)

	if err != nil {
		return err
	}

	return nil
}
