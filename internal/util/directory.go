package util

import (
	"os"
)

func GetDirectories(path string) ([]string, error) {
	var directories []string

	directory, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer directory.Close()

	files, err := directory.Readdir(-1)

	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if file.IsDir() {
			directories = append(directories, file.Name())
		}
	}

	return directories, nil
}

func CreateDirectory(path string) error {
	err := os.MkdirAll(path, os.ModePerm)

	if err != nil {
		return err
	}

	return nil
}

func GetWorkDirectory() (string, error) {
	wd, err := os.Getwd()

	if err != nil {
		return "", err
	}

	return wd, nil
}
