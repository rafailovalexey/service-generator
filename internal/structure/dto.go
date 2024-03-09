package structure

import (
	"github.com/rafailovalexey/service-generator/internal/dto"
	"github.com/rafailovalexey/service-generator/internal/template"
	"github.com/rafailovalexey/service-generator/internal/util"
)

func GetDataTransferObjectStructure(application *dto.ApplicationDto, layer string) *dto.NodeDto {
	structure := &dto.NodeDto{
		IsDirectory: true,
		Name:        layer,
		Parent: &[]dto.NodeDto{
			{
				IsDirectory: true,
				Name:        application.Names.SnakeCasePlural,
				Parent: &[]dto.NodeDto{
					{
						IsFile:   true,
						Name:     util.GetFilename(layer, "go"),
						Template: template.GetDataTransferObjectTemplate(application, layer),
					},
				},
			},
		},
	}

	return structure
}
