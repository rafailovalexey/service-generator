package facade

import (
	"github.com/rafailovalexey/service-generator/internal/dto"
	"github.com/rafailovalexey/service-generator/internal/structure"
	"github.com/rafailovalexey/service-generator/internal/util"
	"path/filepath"
)

func GenerateLayers(application *dto.ApplicationDto, layers []string) error {
	wd, err := util.GetWorkDirectory()

	if err != nil {
		return err
	}

	wd = filepath.Join(wd, application.Directory)

	err = structure.GenerateLayers(wd, application, layers)

	if err != nil {
		return err
	}

	err = structure.GenerateProvider(wd, application)

	if err != nil {
		return err
	}

	return nil
}
