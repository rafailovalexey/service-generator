package builder

import (
	"github.com/rafailovalexey/service-generator/internal/modules/application"
	"github.com/rafailovalexey/service-generator/internal/modules/directory"
	"github.com/rafailovalexey/service-generator/internal/modules/file"
	"github.com/rafailovalexey/service-generator/internal/modules/template"
	"os"
	"path"
)

func CreateLayer(layer string, name string) error {
	current, err := os.Getwd()

	if err != nil {
		return err
	}

	kind := "internal"
	extension := "go"

	if err != nil {
		return err
	}

	filename := file.GetFilename(layer, extension)
	absolute := path.Join(current, kind, layer)
	filepath := path.Join(current, kind, layer, filename)

	isExist := directory.IsExist(absolute)

	if !isExist {
		err = directory.Create(absolute)

		if err != nil {
			return err
		}
	}

	err = file.Create(filepath)

	if err != nil {
		return err
	}

	data := template.GetLayer(layer, name)

	err = file.Set(filepath, data)

	if err != nil {
		return err
	}

	return nil
}

func CreateImplementation(layer string, name string) error {
	current, err := os.Getwd()

	if err != nil {
		return err
	}

	module, err := application.GetApplicationModuleName()

	if err != nil {
		return err
	}

	kind := "internal"
	extension := "go"

	if err != nil {
		return err
	}

	filename := file.GetFilename(layer, extension)
	absolute := path.Join(current, kind, layer, name)
	filepath := path.Join(current, kind, layer, name, filename)

	isExist := directory.IsExist(absolute)

	if !isExist {
		err = directory.Create(absolute)

		if err != nil {
			return err
		}
	}

	err = file.Create(filepath)

	if err != nil {
		return err
	}

	data := template.GetImplementation(layer, name, module, kind)

	err = file.Set(filepath, data)

	if err != nil {
		return err
	}

	return nil
}

func CreateDataTransferObject(layer string, name string) error {
	current, err := os.Getwd()

	if err != nil {
		return err
	}

	kind := "internal"
	extension := "go"

	if err != nil {
		return err
	}

	filename := file.GetFilename(layer, extension)
	absolute := path.Join(current, kind, layer, name)
	filepath := path.Join(current, kind, layer, name, filename)

	isExist := directory.IsExist(absolute)

	if !isExist {
		err = directory.Create(absolute)

		if err != nil {
			return err
		}
	}

	err = file.Create(filepath)

	if err != nil {
		return err
	}

	data := template.GetDataTransferObject(layer, name)

	err = file.Set(filepath, data)

	if err != nil {
		return err
	}

	return nil
}
