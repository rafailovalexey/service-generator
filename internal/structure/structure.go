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

func Generate(wd string, application *dto.ApplicationDto) error {
	structure := GetCoreStructure(application)

	switch application.Type {
	case "grpc":
		*structure = append(*structure, *GetGrpcStructure(application)...)
	case "http":
		*structure = append(*structure, *GetHttpStructure(application)...)
	case "cron":
		*structure = append(*structure, *GetCronStructure(application)...)
	}

	switch application.Database {
	case "postgres":
		*structure = append(*structure, *GetPostgresStructure()...)
	case "mysql":
		*structure = append(*structure, *GetMySqlStructure()...)
	}

	err := Recursion(wd, structure)

	if err != nil {
		return err
	}

	return nil
}

func GenerateProvider(wd string, application *dto.ApplicationDto) error {
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
		return err
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
							Template: template.GetProviderDefinitionTemplate(application, layers),
						},
						{
							IsDirectory: true,
							Name:        application.Names.SnakeCasePlural,
							Parent: &[]dto.NodeDto{
								{
									IsFile:   true,
									Name:     util.GetFilename("provider", "go"),
									Template: template.GetProviderImplementationTemplate(application, layers),
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
		return err
	}

	return nil
}
