package structure

import (
	"github.com/rafailovalexey/service-generator/internal/dto"
	"github.com/rafailovalexey/service-generator/internal/template"
	"github.com/rafailovalexey/service-generator/internal/util"
)

func GetCoreStructure(application string, version string, database string, module string, name *dto.NameDto) *[]dto.NodeDto {
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
							Template: template.GetApplicationTemplate(module, application, name),
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
							Template: template.GetMigrationTemplate(module),
						},
					},
				},
			},
		},
		{
			IsDirectory: true,
			Name:        "internal",
			Parent: &[]dto.NodeDto{
				*GetBaseDefinitionAndImplementationStructure(module, "converter", name),
				*GetBaseDefinitionAndImplementationStructure(module, "service", name),
				*GetBaseDefinitionAndImplementationStructure(module, "repository", name),
				*GetDataTransferObjectStructure("model", name),
				*GetDataTransferObjectStructure("dto", name),
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
			Template: template.GetEnvironmentTemplate(application, database),
		},
		{
			IsFile:   true,
			Name:     util.GetFilename(".example.env", ""),
			Template: template.GetEnvironmentTemplate(application, database),
		},
		{
			IsFile:   true,
			Name:     util.GetFilename("main", "go"),
			Template: template.GetMainTemplate(module),
		},
		{
			IsFile:   true,
			Name:     util.GetFilename("go", "mod"),
			Template: template.GetGoTemplate(module, version),
		},
		{
			IsFile:   true,
			Name:     util.GetFilename("Makefile", ""),
			Template: template.GetMakefileTemplate(application, name),
		},
		{
			IsFile:   true,
			Name:     util.GetFilename("README", "md"),
			Template: template.GetReadmeTemplate(module),
		},
	}

	return structure
}
