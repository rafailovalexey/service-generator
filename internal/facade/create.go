package facade

import (
	"github.com/rafailovalexey/service-generator/internal/dto"
	"github.com/rafailovalexey/service-generator/internal/structure"
	"github.com/rafailovalexey/service-generator/internal/util"
	"path/filepath"
)

func Create(application *dto.ApplicationDto) error {
	wd, err := util.GetWorkDirectory()

	if err != nil {
		return err
	}

	wd = filepath.Join(wd, application.Directory)

	err = structure.Generate(wd, application)

	if err != nil {
		return err
	}

	err = structure.GenerateProvider(wd, application)

	if err != nil {
		return err
	}

	return nil
}
