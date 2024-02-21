package structure

import (
	"github.com/rafailovalexey/service-generator/internal/dto"
	"github.com/rafailovalexey/service-generator/internal/template"
	"github.com/rafailovalexey/service-generator/internal/util"
)

func GetCronStructure(module string, name *dto.NameDto) *[]dto.NodeDto {
	structure := &[]dto.NodeDto{
		{
			IsDirectory: true,
			Name:        "cmd",
			Parent: &[]dto.NodeDto{
				{
					IsDirectory: true,
					Name:        "cron_scheduler",
					Parent: &[]dto.NodeDto{
						{
							IsFile:   true,
							Name:     util.GetFilename("cron_scheduler", "go"),
							Template: template.GetCronSchedulerTemplate(module, name),
						},
					},
				},
			},
		},
	}

	return structure
}
