package utils

import (
	"fmt"
	"os"
	"path"
	"strings"
)

func GetApplicationName() (string, error) {
	current, err := os.Getwd()

	if err != nil {
		return "", err
	}

	extension := "mod"
	filename := GetFilename("go", extension)
	filepath := path.Join(current, filename)

	isExist := PathIsExist(filepath)

	if !isExist {
		return "", fmt.Errorf("%s not found in %s", filename, current)
	}

	read, err := ReadFileData(filepath)

	if err != nil {
		return "", err
	}

	separator := GetSeparator()

	temporary := strings.Split(string(read), separator)[0]
	module := strings.Split(temporary, " ")[1]

	return module, nil
}
