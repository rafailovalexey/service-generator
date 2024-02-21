package structure

import (
	"github.com/rafailovalexey/service-generator/internal/dto"
	"github.com/rafailovalexey/service-generator/internal/template"
	"github.com/rafailovalexey/service-generator/internal/util"
)

func GetHttpStructure(module string, name *dto.NameDto) *[]dto.NodeDto {
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
							Template: template.GetHttpServerTemplate(module, name),
						},
						{
							IsDirectory: true,
							Name:        "middleware",
							Parent: &[]dto.NodeDto{
								{
									IsFile:   true,
									Name:     util.GetFilename("authentication", "go"),
									Template: template.GetHttpAuthenticationMiddlewareTemplate(module),
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
									Name:     util.GetFilename("logging", "go"),
									Template: template.GetHttpLoggingInterceptorTemplate(),
								},
							},
						},
					},
				},
			},
		},
		{
			IsDirectory: true,
			Name:        "internal",
			Parent: &[]dto.NodeDto{
				*GetBaseDefinitionAndImplementationStructure(module, "controller", name),
				*GetBaseDefinitionAndImplementationStructure(module, "validation", name),
				{
					IsDirectory: true,
					Name:        "handler",
					Parent: &[]dto.NodeDto{
						{
							IsFile:   true,
							Name:     util.GetFilename("handler", "go"),
							Template: template.GetHttpHandlerDefinitionTemplate(name),
						},
						{
							IsDirectory: true,
							Name:        name.SnakeCasePlural,
							Parent: &[]dto.NodeDto{
								{
									IsFile:   true,
									Name:     util.GetFilename("handler", "go"),
									Template: template.GetHttpHandlerImplementationTemplate(module, name),
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
			Name:     "application.dockerfile",
			Template: template.GetDockerTemplate(true),
		},
	}

	return structure
}
