package structure

import (
	"github.com/rafailovalexey/service-generator/internal/dto"
	"github.com/rafailovalexey/service-generator/internal/template"
	"github.com/rafailovalexey/service-generator/internal/util"
)

func GetGrpcStructure(application *dto.ApplicationDto) *[]dto.NodeDto {
	structure := &[]dto.NodeDto{
		{
			IsDirectory: true,
			Name:        "api",
			Parent: &[]dto.NodeDto{
				{
					IsDirectory: true,
					Name:        application.Names.SnakeCasePlural + "_" + "v1",
					Parent: &[]dto.NodeDto{
						{
							IsFile:   true,
							Name:     util.GetFilename(application.Names.SnakeCasePlural, "proto"),
							Template: template.GetProtoTemplate(application),
						},
					},
				},
			},
		},
		{
			IsDirectory: true,
			Name:        "bin",
			Parent: &[]dto.NodeDto{
				{
					IsFile:   true,
					Name:     util.GetFilename("grpc-generate", "sh"),
					Template: template.GetGrpcGenerateShellScriptTemplate(),
				},
			},
		},
		{
			IsDirectory: true,
			Name:        "cmd",
			Parent: &[]dto.NodeDto{
				{
					IsDirectory: true,
					Name:        "grpc_server",
					Parent: &[]dto.NodeDto{
						{
							IsFile:   true,
							Name:     util.GetFilename("grpc_server", "go"),
							Template: template.GetGrpcServerTemplate(application),
						},
						{
							IsDirectory: true,
							Name:        "interceptor",
							Parent: &[]dto.NodeDto{
								{
									IsFile:   true,
									Name:     util.GetFilename("logging", "go"),
									Template: template.GetGrpcLoggingInterceptorTemplate(),
								},
								{
									IsFile:   true,
									Name:     util.GetFilename("tracing", "go"),
									Template: template.GetGrpcTracingInterceptorTemplate(),
								},
							},
						},
						{
							IsDirectory: true,
							Name:        "middleware",
							Parent: &[]dto.NodeDto{
								{
									IsFile:   true,
									Name:     util.GetFilename("authentication", "go"),
									Template: template.GetGrpcAuthenticationMiddlewareTemplate(),
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
					Name:        "implementation",
					Parent: &[]dto.NodeDto{
						{
							IsDirectory: true,
							Name:        application.Names.SnakeCasePlural,
							Parent: &[]dto.NodeDto{
								{
									IsFile:   true,
									Name:     util.GetFilename("implementation", "go"),
									Template: template.GetGrpcServerImplementationTemplate(application),
								},
							},
						},
					},
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
