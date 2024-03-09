package structure

import (
	"github.com/rafailovalexey/service-generator/internal/dto"
	"github.com/rafailovalexey/service-generator/internal/template"
	"github.com/rafailovalexey/service-generator/internal/util"
)

func GetCoreStructure(application *dto.ApplicationDto) *[]dto.NodeDto {
	structure := &[]dto.NodeDto{
		{
			IsDirectory: true,
			Name:        "bin",
			Parent: &[]dto.NodeDto{
				{
					IsFile:   true,
					Name:     util.GetFilename("mock-generate", "sh"),
					Template: template.GetMockGenerateShellScriptTemplate(),
				},
			},
		},
		{
			IsDirectory: true,
			Name:        "cmd",
			Parent: &[]dto.NodeDto{
				{
					IsDirectory: true,
					Name:        "application",
					Parent: &[]dto.NodeDto{
						{
							IsFile:   true,
							Name:     util.GetFilename("application", "go"),
							Template: template.GetApplicationTemplate(application),
						},
					},
				},
			},
		},
		{
			IsDirectory: true,
			Name:        "config",
			Parent: &[]dto.NodeDto{
				{
					IsFile:   true,
					Name:     util.GetFilename("config", "go"),
					Template: template.GetConfigTemplate(application),
				},
			},
		},
		{
			IsDirectory: true,
			Name:        "internal",
			Parent: &[]dto.NodeDto{
				*GetBaseDefinitionAndImplementationStructure(application, "converter"),
				*GetBaseDefinitionAndImplementationStructure(application, "service"),
				*GetBaseDefinitionAndImplementationStructure(application, "repository"),
				*GetDataTransferObjectStructure(application, "model"),
				*GetDataTransferObjectStructure(application, "dto"),
			},
		},
		{
			IsFile:   true,
			Name:     util.GetFilename(".gitignore", ""),
			Template: template.GetGitIgnoreTemplate(),
		},
		{
			IsFile:   true,
			Name:     util.GetFilename(".dockerignore", ""),
			Template: template.GetDockerIgnoreTemplate(),
		},
		{
			IsFile:   true,
			Name:     util.GetFilename(".env", ""),
			Template: template.GetEnvironmentTemplate(application),
		},
		{
			IsFile:   true,
			Name:     util.GetFilename(".example.env", ""),
			Template: template.GetEnvironmentTemplate(application),
		},
		{
			IsFile:   true,
			Name:     util.GetFilename("main", "go"),
			Template: template.GetMainTemplate(application),
		},
		{
			IsFile:   true,
			Name:     util.GetFilename("go", "mod"),
			Template: template.GetGoTemplate(application),
		},
		{
			IsFile:   true,
			Name:     util.GetFilename("Makefile", ""),
			Template: template.GetMakefileTemplate(application),
		},
		{
			IsFile:   true,
			Name:     util.GetFilename("README", "md"),
			Template: template.GetReadmeTemplate(application),
		},
	}

	return structure
}
