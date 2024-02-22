package structure

import (
	"github.com/rafailovalexey/service-generator/internal/dto"
	"github.com/rafailovalexey/service-generator/internal/template"
	"github.com/rafailovalexey/service-generator/internal/util"
)

func GetMySQLStructure() *[]dto.NodeDto {
	structure := &[]dto.NodeDto{
		{
			IsDirectory: true,
			Name:        "database",
			Parent: &[]dto.NodeDto{
				{
					IsFile:   true,
					Name:     util.GetFilename("mysql", "go"),
					Template: template.GetDatabaseMySQLTemplate(),
				},
			},
		},
	}

	return structure
}
