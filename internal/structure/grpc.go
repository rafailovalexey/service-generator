package structure

import (
	"github.com/rafailovalexey/service-generator/internal/dto"
	"github.com/rafailovalexey/service-generator/internal/template"
	"github.com/rafailovalexey/service-generator/internal/util"
)

func GetGrpcStructure(module string, organization string, version string, name *dto.NameDto) *[]dto.NodeDto {
	structure := &[]dto.NodeDto{
		{
			IsDirectory: true,
			Name:        "api",
			Parent: &[]dto.NodeDto{
				{
					IsDirectory: true,
					Name:        name.SnakeCasePlural + "_" + "v1",
					Parent: &[]dto.NodeDto{
						{
							IsFile:   true,
							Name:     util.GetFilename(name.SnakeCasePlural, "proto"),
							Template: template.GetProtoTemplate(module, name),
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
							Template: template.GetGrpcServerTemplate(module, name),
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
			},
		},
		{
			IsDirectory: true,
			Name:        "database",
			Parent: &[]dto.NodeDto{
				{
					IsDirectory: true,
					Name:        "migrations",
					Parent:      &[]dto.NodeDto{},
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
					Name:        "implementation",
					Parent: &[]dto.NodeDto{
						{
							IsDirectory: true,
							Name:        name.SnakeCasePlural,
							Parent: &[]dto.NodeDto{
								{
									IsFile:   true,
									Name:     util.GetFilename("implementation", "go"),
									Template: template.GetGrpcServerImplementationTemplate(module, name),
								},
							},
						},
					},
				},
			},
		},
		{
			IsFile:   true,
			Name:     "application.dockerfile",
			Template: template.GetDockerTemplate(organization, version, true),
		},
	}

	return structure
}
