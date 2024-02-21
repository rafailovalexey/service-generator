package structure

import (
	"github.com/rafailovalexey/service-generator/internal/dto"
	"github.com/rafailovalexey/service-generator/internal/template"
	"github.com/rafailovalexey/service-generator/internal/util"
)

func GetBaseDefinitionAndImplementationStructure(module string, layer string, name *dto.NameDto) *dto.NodeDto {
	structure := &dto.NodeDto{
		IsDirectory: true,
		Name:        layer,
		Parent: &[]dto.NodeDto{
			{
				IsFile:   true,
				Name:     util.GetFilename(layer, "go"),
				Template: template.GetBaseDefinitionTemplate(layer, name),
			},
			{
				IsDirectory: true,
				Name:        name.SnakeCasePlural,
				Parent: &[]dto.NodeDto{
					{
						IsFile:   true,
						Name:     util.GetFilename(layer, "go"),
						Template: template.GetBaseImplementationTemplate(module, layer, name),
					},
				},
			},
		},
	}

	return structure
}
