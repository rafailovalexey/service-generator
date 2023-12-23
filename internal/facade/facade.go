package facade

import (
	"github.com/rafailovalexey/service-generator/internal/template"
	"github.com/rafailovalexey/service-generator/internal/utils"
	"os"
	"path"
	"sort"
)

func CreateInterface(layer string, name string) error {
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

	data := template.GetInterfaceTemplate(separator, layer, name)

	err = utils.SetFileData(filepath, data)

	if err != nil {
		return err
	}

	return nil
}

func CreateRealisationInterface(layer string, name string) error {
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

	data := template.GetRealisationInterfaceTemplate(application, separator, kind, layer, name)

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

	data := template.GetDataTransferObjectTemplate(separator, layer, name)

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

	data := template.GetRequestTemplate(separator, name)

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

	data := template.GetResponseTemplate(separator, name)

	err = utils.SetFileData(filepath, data)

	if err != nil {
		return err
	}

	return nil
}

func CreateProvider(layer string, name string) error {
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

	available := make(map[string]struct{}, 10)

	available["api"] = struct{}{}
	available["controller"] = struct{}{}
	available["implementation"] = struct{}{}
	available["client"] = struct{}{}
	available["validation"] = struct{}{}
	available["converter"] = struct{}{}
	available["repository"] = struct{}{}
	available["service"] = struct{}{}

	directory = path.Join(current, kind)
	directories, err := utils.GetDirectories(directory)

	if err != nil {
		return err
	}

	layers := make([]string, 0, 10)

	for _, d := range directories {
		if _, isExist = available[d]; isExist {
			layers = append(layers, d)
		}
	}

	sort.Strings(layers)

	data := template.GetProviderRealisationTemplate(application, separator, kind, layers, layer, name)

	err = utils.SetFileData(filepath, data)

	if err != nil {
		return err
	}

	return nil
}

func CreateProviderInterface(layer string, name string) error {
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

	available := make(map[string]struct{}, 10)

	available["api"] = struct{}{}
	available["controller"] = struct{}{}
	available["implementation"] = struct{}{}
	available["client"] = struct{}{}
	available["validation"] = struct{}{}
	available["converter"] = struct{}{}
	available["repository"] = struct{}{}
	available["service"] = struct{}{}

	directory = path.Join(current, kind)
	directories, err := utils.GetDirectories(directory)

	if err != nil {
		return err
	}

	layers := make([]string, 0, 10)

	for _, d := range directories {
		if _, isExist = available[d]; isExist {
			layers = append(layers, d)
		}
	}

	sort.Strings(layers)

	data := template.GetProviderInterfaceTemplate(application, separator, kind, layers, layer, name)

	err = utils.SetFileData(filepath, data)

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

	data := template.GetImplementationRealisationTemplate(separator, layer, name)

	err = utils.SetFileData(filepath, data)

	if err != nil {
		return err
	}

	return nil
}
