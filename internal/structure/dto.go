package structure

import (
	"github.com/rafailovalexey/service-generator/internal/dto"
	"github.com/rafailovalexey/service-generator/internal/template"
	"github.com/rafailovalexey/service-generator/internal/util"
)

func GetDataTransferObjectStructure(layer string, name *dto.NameDto) *dto.NodeDto {
	structure := &dto.NodeDto{
		IsDirectory: true,
		Name:        layer,
		Parent: &[]dto.NodeDto{
			{
				IsDirectory: true,
				Name:        name.SnakeCasePlural,
				Parent: &[]dto.NodeDto{
					{
						IsFile:   true,
						Name:     util.GetFilename(layer, "go"),
						Template: template.GetDataTransferObjectTemplate(layer, name),
					},
				},
			},
		},
	}

	return structure
}
