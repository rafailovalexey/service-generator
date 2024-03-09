package structure

import (
	"github.com/rafailovalexey/service-generator/internal/dto"
	"github.com/rafailovalexey/service-generator/internal/template"
	"github.com/rafailovalexey/service-generator/internal/util"
)

func GetHttpStructure(application *dto.ApplicationDto) *[]dto.NodeDto {
	structure := &[]dto.NodeDto{
		{
			IsDirectory: true,
			Name:        "cmd",
			Parent: &[]dto.NodeDto{
				{
					IsDirectory: true,
					Name:        "http_server",
					Parent: &[]dto.NodeDto{
						{
							IsFile:   true,
							Name:     util.GetFilename("http_server", "go"),
							Template: template.GetHttpServerTemplate(application),
						},
						{
							IsDirectory: true,
							Name:        "middleware",
							Parent: &[]dto.NodeDto{
								{
									IsFile:   true,
									Name:     util.GetFilename("middleware", "go"),
									Template: template.GetHttpMiddlewareTemplate(),
								},
								{
									IsFile:   true,
									Name:     util.GetFilename("authentication", "go"),
									Template: template.GetHttpAuthenticationMiddlewareTemplate(application),
								},
								{
									IsFile:   true,
									Name:     util.GetFilename("cors", "go"),
									Template: template.GetHttpCorsMiddlewareTemplate(),
								},
								{
									IsFile:   true,
									Name:     util.GetFilename("chain", "go"),
									Template: template.GetHttpChainMiddlewareTemplate(),
								},
							},
						},
						{
							IsDirectory: true,
							Name:        "interceptor",
							Parent: &[]dto.NodeDto{
								{
									IsFile:   true,
									Name:     util.GetFilename("interceptor", "go"),
									Template: template.GetHttpInterceptorTemplate(),
								},
								{
									IsFile:   true,
									Name:     util.GetFilename("logging", "go"),
									Template: template.GetHttpLoggingInterceptorTemplate(),
								},
							},
						},
					},
				},
				{
					IsDirectory: true,
					Name:        "migration",
					Parent: &[]dto.NodeDto{
						{
							IsFile:   true,
							Name:     util.GetFilename("migration", "go"),
							Template: template.GetMigrationTemplate(application),
						},
					},
				},
			},
		},
		{
			IsDirectory: true,
			Name:        "database",
			Parent: &[]dto.NodeDto{
				{
					IsDirectory: true,
					Name:        "migration",
					Parent:      &[]dto.NodeDto{},
				},
			},
		},
		{
			IsDirectory: true,
			Name:        "internal",
			Parent: &[]dto.NodeDto{
				*GetBaseDefinitionAndImplementationStructure(application, "controller"),
				*GetBaseDefinitionAndImplementationStructure(application, "validation"),
				{
					IsDirectory: true,
					Name:        "handler",
					Parent: &[]dto.NodeDto{
						{
							IsFile:   true,
							Name:     util.GetFilename("handler", "go"),
							Template: template.GetHttpHandlerDefinitionTemplate(application),
						},
						{
							IsDirectory: true,
							Name:        application.Names.SnakeCasePlural,
							Parent: &[]dto.NodeDto{
								{
									IsFile:   true,
									Name:     util.GetFilename("handler", "go"),
									Template: template.GetHttpHandlerImplementationTemplate(application),
								},
							},
						},
					},
				},
			},
		},
		{
			IsDirectory: true,
			Name:        "util",
			Parent: &[]dto.NodeDto{
				{
					IsFile:   true,
					Name:     util.GetFilename("converter_error", "go"),
					Template: template.GetUtilConverterErrorTemplate(),
				},
				{
					IsFile:   true,
					Name:     util.GetFilename("response", "go"),
					Template: template.GetUtilResponseTemplate(),
				},
			},
		},
		{
			IsFile:   true,
			Name:     util.GetFilename("application", "dockerfile"),
			Template: template.GetDockerTemplate(application, true),
		},
		{
			IsFile:   true,
			Name:     util.GetFilename("docker-compose-application", "yml"),
			Template: template.GetDockerComposeTemplate(application, true),
		},
	}

	return structure
}
