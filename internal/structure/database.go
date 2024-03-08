package structure

import (
	"github.com/rafailovalexey/service-generator/internal/dto"
	"github.com/rafailovalexey/service-generator/internal/template"
	"github.com/rafailovalexey/service-generator/internal/util"
)

func GetPostgresStructure() *[]dto.NodeDto {
	structure := &[]dto.NodeDto{
		{
			IsDirectory: true,
			Name:        "database",
			Parent: &[]dto.NodeDto{
				{
					IsDirectory: true,
					Name:        "postgres",
					Parent: &[]dto.NodeDto{
						{
							IsFile:   true,
							Name:     util.GetFilename("postgres", "go"),
							Template: template.GetDatabasePostgresTemplate(),
						},
					},
				},
			},
		},
	}

	return structure
}

func GetMySqlStructure() *[]dto.NodeDto {
	structure := &[]dto.NodeDto{
		{
			IsDirectory: true,
			Name:        "database",
			Parent: &[]dto.NodeDto{
				{
					IsDirectory: true,
					Name:        "mysql",
					Parent: &[]dto.NodeDto{
						{
							IsFile:   true,
							Name:     util.GetFilename("mysql", "go"),
							Template: template.GetDatabaseMySqlTemplate(),
						},
					},
				},
			},
		},
	}

	return structure
}
