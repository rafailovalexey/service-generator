package structure

import (
	"github.com/rafailovalexey/service-generator/internal/dto"
	"github.com/rafailovalexey/service-generator/internal/template"
	"github.com/rafailovalexey/service-generator/internal/util"
)

func GetCronStructure(application *dto.ApplicationDto) *[]dto.NodeDto {
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
							Template: template.GetCronSchedulerTemplate(application),
						},
					},
				},
			},
		},
		{
			IsFile:   true,
			Name:     util.GetFilename("application", "dockerfile"),
			Template: template.GetDockerTemplate(application, false),
		},
		{
			IsFile:   true,
			Name:     util.GetFilename("docker-compose-application", "yml"),
			Template: template.GetDockerComposeTemplate(application, false),
		},
	}

	return structure
}
