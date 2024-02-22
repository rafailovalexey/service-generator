package structure

import (
	"github.com/rafailovalexey/service-generator/internal/dto"
	"github.com/rafailovalexey/service-generator/internal/template"
	"github.com/rafailovalexey/service-generator/internal/util"
	"path/filepath"
	"sort"
)

func Recursion(path string, nodes *[]dto.NodeDto) error {
	if nodes == nil {
		return nil
	}

	if len(*nodes) == 0 {
		return nil
	}

	for _, node := range *nodes {
		current := filepath.Join(path, node.Name)

		if node.IsDirectory {
			if !util.PathIsExist(current) {
				err := util.CreateDirectory(current)

				if err != nil {
					return err
				}
			}
		}

		if node.IsFile {
			err := util.CreateFile(current)

			if err != nil {
				return err
			}

			err = util.SetFileData(current, node.Template)

			if err != nil {
				return err
			}
		}

		err := Recursion(current, node.Parent)

		if err != nil {
			return err
		}
	}

	return nil
}

func Generate(wd string, application string, version string, database string, module string, name *dto.NameDto) {
	structure := GetCoreStructure(application, version, database, module, name)

	switch application {
	case "grpc":
		*structure = append(*structure, *GetGrpcStructure(module, name)...)
	case "http":
		*structure = append(*structure, *GetHttpStructure(module, name)...)
	case "cron":
		*structure = append(*structure, *GetCronStructure(module, name)...)
	}

	switch database {
	case "mysql":
		*structure = append(*structure, *GetMySQLStructure()...)
	}

	err := Recursion(wd, structure)

	if err != nil {
		panic(err)
	}
}

func GenerateProvider(wd string, module string, name *dto.NameDto) {
	available := map[string]struct{}{
		"handler":        {},
		"implementation": {},
		"api":            {},
		"controller":     {},
		"validation":     {},
		"converter":      {},
		"service":        {},
		"repository":     {},
		"client":         {},
	}

	directories, err := util.GetDirectories(filepath.Join(wd, "internal"))

	if err != nil {
		panic(err)
	}

	layers := make([]string, 0, 10)

	for _, directory := range directories {
		if _, isExist := available[directory]; isExist {
			layers = append(layers, directory)
		}
	}

	sort.Strings(layers)

	structure := &[]dto.NodeDto{
		{
			IsDirectory: true,
			Name:        "internal",
			Parent: &[]dto.NodeDto{
				{
					IsDirectory: true,
					Name:        "provider",
					Parent: &[]dto.NodeDto{
						{
							IsFile:   true,
							Name:     util.GetFilename("provider", "go"),
							Template: template.GetProviderDefinitionTemplate(module, name, layers),
						},
						{
							IsDirectory: true,
							Name:        name.SnakeCasePlural,
							Parent: &[]dto.NodeDto{
								{
									IsFile:   true,
									Name:     util.GetFilename("provider", "go"),
									Template: template.GetProviderImplementationTemplate(module, name, layers),
								},
							},
						},
					},
				},
			},
		},
	}

	err = Recursion(wd, structure)

	if err != nil {
		panic(err)
	}
}
