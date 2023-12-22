package facade

import (
	"github.com/rafailovalexey/service-generator/internal/template"
	"github.com/rafailovalexey/service-generator/internal/utils"
	"os"
	"path"
)

func CreateInterface(layer string, name string, imports *template.Imports, methods *template.Methods) error {
	current, err := os.Getwd()

	if err != nil {
		return err
	}

	separator := utils.GetSeparator()

	kind := "internal"
	extension := "go"

	filename := utils.GetFilename(layer, extension)
	directory := path.Join(current, kind, layer)
	filepath := path.Join(current, kind, layer, filename)

	isExist := utils.PathIsExist(directory)

	if !isExist {
		err = utils.CreateDirectory(directory)

		if err != nil {
			return err
		}
	}

	err = utils.CreateFile(filepath)

	if err != nil {
		return err
	}

	data := template.GetInterfaceTemplate(layer, name, separator, imports, methods)

	err = utils.SetFileData(filepath, data)

	if err != nil {
		return err
	}

	return nil
}

func CreateRealisationInterface(layer string, name string, imports *template.Imports, methods *template.Methods, functions *template.Functions) error {
	current, err := os.Getwd()

	if err != nil {
		return err
	}

	application, err := utils.GetApplicationName()

	if err != nil {
		return err
	}

	separator := utils.GetSeparator()

	kind := "internal"
	extension := "go"

	if err != nil {
		return err
	}

	filename := utils.GetFilename(layer, extension)
	directory := path.Join(current, kind, layer, name)
	filepath := path.Join(current, kind, layer, name, filename)

	isExist := utils.PathIsExist(directory)

	if !isExist {
		err = utils.CreateDirectory(directory)

		if err != nil {
			return err
		}
	}

	err = utils.CreateFile(filepath)

	if err != nil {
		return err
	}

	data := template.GetRealisationInterfaceTemplate(layer, name, application, kind, separator, imports, methods, functions)

	err = utils.SetFileData(filepath, data)

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

	separator := utils.GetSeparator()

	kind := "internal"
	extension := "go"

	if err != nil {
		return err
	}

	filename := utils.GetFilename(layer, extension)
	directory := path.Join(current, kind, layer, name)
	filepath := path.Join(current, kind, layer, name, filename)

	isExist := utils.PathIsExist(directory)

	if !isExist {
		err = utils.CreateDirectory(directory)

		if err != nil {
			return err
		}
	}

	err = utils.CreateFile(filepath)

	if err != nil {
		return err
	}

	data := template.GetDataTransferObjectTemplate(layer, name, separator)

	err = utils.SetFileData(filepath, data)

	if err != nil {
		return err
	}

	return nil
}

func CreateRequestObject(layer string, name string) error {
	current, err := os.Getwd()

	if err != nil {
		return err
	}

	separator := utils.GetSeparator()

	kind := "internal"
	request := "request"
	extension := "go"

	if err != nil {
		return err
	}

	filename := utils.GetFilename(request, extension)
	directory := path.Join(current, kind, layer, name, request)
	filepath := path.Join(current, kind, layer, name, request, filename)

	isExist := utils.PathIsExist(directory)

	if !isExist {
		err = utils.CreateDirectory(directory)

		if err != nil {
			return err
		}
	}

	err = utils.CreateFile(filepath)

	if err != nil {
		return err
	}

	data := template.GetRequestTemplate(name, separator)

	err = utils.SetFileData(filepath, data)

	if err != nil {
		return err
	}

	return nil
}

func CreateResponseObject(layer string, name string) error {
	current, err := os.Getwd()

	if err != nil {
		return err
	}

	separator := utils.GetSeparator()

	kind := "internal"
	response := "response"
	extension := "go"

	if err != nil {
		return err
	}

	filename := utils.GetFilename(response, extension)
	directory := path.Join(current, kind, layer, name, response)
	filepath := path.Join(current, kind, layer, name, response, filename)

	isExist := utils.PathIsExist(directory)

	if !isExist {
		err = utils.CreateDirectory(directory)

		if err != nil {
			return err
		}
	}

	err = utils.CreateFile(filepath)

	if err != nil {
		return err
	}

	data := template.GetRequestTemplate(name, separator)

	err = utils.SetFileData(filepath, data)

	if err != nil {
		return err
	}

	return nil
}
