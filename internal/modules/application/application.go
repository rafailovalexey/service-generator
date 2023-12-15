package application

import (
	"fmt"
	"github.com/rafailovalexey/service-generator/internal/modules/file"
	"github.com/rafailovalexey/service-generator/internal/modules/system"
	"os"
	"path"
	"strings"
)

func GetApplicationModule() (string, error) {
	current, err := os.Getwd()

	if err != nil {
		return "", err
	}

	extension := "mod"
	filename := file.GetFilename("go", extension)
	filepath := path.Join(current, filename)

	isExist := file.IsExist(filepath)

	if !isExist {
		return "", fmt.Errorf("%s not found in %s", filename, current)
	}

	read, err := file.Read(filepath)

	if err != nil {
		return "", err
	}

	separator := system.GetSeparator()

	temporary := strings.Split(string(read), separator)[0]
	module := strings.Split(temporary, " ")[1]

	return module, nil
}
