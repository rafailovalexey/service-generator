package structure

import (
	"github.com/rafailovalexey/service-generator/internal/template"
	"github.com/rafailovalexey/service-generator/internal/utils"
	"path/filepath"
	"sort"
)

type Node struct {
	IsDirectory bool
	IsFile      bool
	Name        string
	Template    []byte
	Parent      *[]Node
}

func Recursion(path string, nodes *[]Node) error {
	if nodes == nil {
		return nil
	}

	if len(*nodes) == 0 {
		return nil
	}

	for _, node := range *nodes {
		current := filepath.Join(path, node.Name)

		if node.IsDirectory {
			if !utils.PathIsExist(current) {
				err := utils.CreateDirectory(current)

				if err != nil {
					return err
				}
			}
		}

		if node.IsFile {
			err := utils.CreateFile(current)

			if err != nil {
				return err
			}

			err = utils.SetFileData(current, node.Template)

			if err != nil {
				return err
			}
		}

		err := Recursion(current, node.Parent)

		if err != nil {
			return err
		}
	}

	return nil
}

func Generate(wd string, application string, version string, module string, name string) {
	structure := &[]Node{
		{
			IsDirectory: true,
			Name:        "bin",
			Parent: &[]Node{
				{
					IsFile:   true,
					Name:     utils.GetFilename("mock-generate", "sh"),
					Template: template.GetMockGenerateShellScriptTemplate(),
				},
			},
		},
		{
			IsDirectory: true,
			Name:        "cmd",
			Parent: &[]Node{
				{
					IsDirectory: true,
					Name:        "application",
					Parent: &[]Node{
						{
							IsFile:   true,
							Name:     utils.GetFilename("application", "go"),
							Template: template.GetApplicationTemplate(module, application, name),
						},
					},
				},
			},
		},
		{
			IsDirectory: true,
			Name:        "internal",
			Parent: &[]Node{
				*GetBaseDefinitionAndImplementationStructure(module, "controller", name),
				*GetBaseDefinitionAndImplementationStructure(module, "validation", name),
				*GetBaseDefinitionAndImplementationStructure(module, "converter", name),
				*GetBaseDefinitionAndImplementationStructure(module, "service", name),
				*GetBaseDefinitionAndImplementationStructure(module, "repository", name),
				*GetDataTransferObjectStructure("model", name),
				*GetDataTransferObjectStructure("dto", name),
			},
		},
		{
			IsFile:   true,
			Name:     utils.GetFilename(".gitignore", ""),
			Template: template.GetGitIgnoreTemplate(),
		},
		{
			IsFile:   true,
			Name:     utils.GetFilename(".dockerignore", ""),
			Template: template.GetDockerIgnoreTemplate(),
		},
		{
			IsFile:   true,
			Name:     utils.GetFilename(".env", ""),
			Template: template.GetEnvironmentTemplate(),
		},
		{
			IsFile:   true,
			Name:     utils.GetFilename(".example.env", ""),
			Template: template.GetExampleEnvironmentTemplate(),
		},
		{
			IsFile:   true,
			Name:     utils.GetFilename("main", "go"),
			Template: template.GetMainTemplate(module),
		},
		{
			IsFile:   true,
			Name:     utils.GetFilename("go", "mod"),
			Template: template.GetGoTemplate(module, version),
		},
		{
			IsFile:   true,
			Name:     utils.GetFilename("Makefile", ""),
			Template: template.GetMakefileTemplate(application, name),
		},
		{
			IsFile:   true,
			Name:     utils.GetFilename("README", "md"),
			Template: template.GetReadmeTemplate(module),
		},
	}

	switch application {
	case "grpc":
		temporary := &[]Node{
			{
				IsDirectory: true,
				Name:        "api",
				Parent: &[]Node{
					{
						IsDirectory: true,
						Name:        name + "_" + "v1",
						Parent: &[]Node{
							{
								IsFile:   true,
								Name:     utils.GetFilename(name, "proto"),
								Template: template.GetProtoTemplate(module, name),
							},
						},
					},
				},
			},
			{
				IsDirectory: true,
				Name:        "bin",
				Parent: &[]Node{
					{
						IsFile:   true,
						Name:     utils.GetFilename("grpc-generate", "sh"),
						Template: template.GetGrpcGenerateShellScriptTemplate(),
					},
				},
			},
			{
				IsDirectory: true,
				Name:        "cmd",
				Parent: &[]Node{
					{
						IsDirectory: true,
						Name:        "grpc_server",
						Parent: &[]Node{
							{
								IsFile:   true,
								Name:     utils.GetFilename("grpc_server", "go"),
								Template: template.GetGrpcServerTemplate(module, name),
							},
							{
								IsDirectory: true,
								Name:        "interceptor",
								Parent: &[]Node{
									{
										IsFile:   true,
										Name:     utils.GetFilename("logging", "go"),
										Template: template.GetGrpcLoggingInterceptorTemplate(),
									},
									{
										IsFile:   true,
										Name:     utils.GetFilename("tracing", "go"),
										Template: template.GetGrpcTracingInterceptorTemplate(),
									},
								},
							},
							{
								IsDirectory: true,
								Name:        "middleware",
								Parent: &[]Node{
									{
										IsFile:   true,
										Name:     utils.GetFilename("authentication", "go"),
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
				Name:        "internal",
				Parent: &[]Node{
					{
						IsDirectory: true,
						Name:        "implementation",
						Parent: &[]Node{
							{
								IsDirectory: true,
								Name:        name,
								Parent: &[]Node{
									{
										IsFile:   true,
										Name:     utils.GetFilename("implementation", "go"),
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
				Template: template.GetDockerTemplate(true),
			},
		}

		*structure = append(*structure, *temporary...)
	case "http":
		temporary := &[]Node{
			{
				IsDirectory: true,
				Name:        "cmd",
				Parent: &[]Node{
					{
						IsDirectory: true,
						Name:        "http_server",
						Parent: &[]Node{
							{
								IsFile:   true,
								Name:     utils.GetFilename("http_server", "go"),
								Template: template.GetHttpServerTemplate(module, name),
							},
							{
								IsDirectory: true,
								Name:        "middleware",
								Parent: &[]Node{
									{
										IsFile:   true,
										Name:     utils.GetFilename("authentication", "go"),
										Template: template.GetHttpAuthenticationMiddlewareTemplate(module),
									},
									{
										IsFile:   true,
										Name:     utils.GetFilename("cors", "go"),
										Template: template.GetHttpCorsMiddlewareTemplate(),
									},
									{
										IsFile:   true,
										Name:     utils.GetFilename("chain", "go"),
										Template: template.GetHttpChainMiddlewareTemplate(),
									},
								},
							},
							{
								IsDirectory: true,
								Name:        "interceptor",
								Parent: &[]Node{
									{
										IsFile:   true,
										Name:     utils.GetFilename("logging", "go"),
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
				Parent: &[]Node{
					{
						IsDirectory: true,
						Name:        "handler",
						Parent: &[]Node{
							{
								IsFile:   true,
								Name:     utils.GetFilename("handler", "go"),
								Template: template.GetHttpHandlerDefinitionTemplate(name),
							},
							{
								IsDirectory: true,
								Name:        name,
								Parent: &[]Node{
									{
										IsFile:   true,
										Name:     utils.GetFilename("handler", "go"),
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
				Name:        "utils",
				Parent: &[]Node{
					{
						IsFile:   true,
						Name:     utils.GetFilename("converter_error", "go"),
						Template: template.GetUtilsConverterErrorTemplate(),
					},
					{
						IsFile:   true,
						Name:     utils.GetFilename("response", "go"),
						Template: template.GetUtilsResponseTemplate(),
					},
				},
			},
			{
				IsFile:   true,
				Name:     "application.dockerfile",
				Template: template.GetDockerTemplate(true),
			},
		}

		*structure = append(*structure, *temporary...)
	}

	err := Recursion(wd, structure)

	if err != nil {
		panic(err)
	}
}

func GetBaseDefinitionAndImplementationStructure(module string, layer string, name string) *Node {
	structure := &Node{
		IsDirectory: true,
		Name:        layer,
		Parent: &[]Node{
			{
				IsFile:   true,
				Name:     utils.GetFilename(layer, "go"),
				Template: template.GetBaseDefinitionTemplate(layer, name),
			},
			{
				IsDirectory: true,
				Name:        name,
				Parent: &[]Node{
					{
						IsFile:   true,
						Name:     utils.GetFilename(layer, "go"),
						Template: template.GetBaseImplementationTemplate(module, layer, name),
					},
				},
			},
		},
	}

	return structure
}

func GetDataTransferObjectStructure(layer string, name string) *Node {
	structure := &Node{
		IsDirectory: true,
		Name:        layer,
		Parent: &[]Node{
			{
				IsDirectory: true,
				Name:        name,
				Parent: &[]Node{
					{
						IsFile:   true,
						Name:     utils.GetFilename(layer, "go"),
						Template: template.GetDataTransferObjectTemplate(layer, name),
					},
				},
			},
		},
	}

	return structure
}

func GenerateProvider(wd string, module string, name string) {
	available := map[string]struct{}{
		"handler":        {},
		"implementation": {},
		"api":            {},
		"controller":     {},
		"validation":     {},
		"converter":      {},
		"service":        {},
		"repository":     {},
		"client":         {},
	}

	directories, err := utils.GetDirectories(filepath.Join(wd, "internal"))

	if err != nil {
		panic(err)
	}

	layers := make([]string, 0, 10)

	for _, directory := range directories {
		if _, isExist := available[directory]; isExist {
			layers = append(layers, directory)
		}
	}

	sort.Strings(layers)

	structure := &[]Node{
		{
			IsDirectory: true,
			Name:        "internal",
			Parent: &[]Node{
				{
					IsDirectory: true,
					Name:        "provider",
					Parent: &[]Node{
						{
							IsFile:   true,
							Name:     utils.GetFilename("provider", "go"),
							Template: template.GetProviderDefinitionTemplate(module, name, layers),
						},
						{
							IsDirectory: true,
							Name:        name,
							Parent: &[]Node{
								{
									IsFile:   true,
									Name:     utils.GetFilename("provider", "go"),
									Template: template.GetProviderImplementationTemplate(module, name, layers),
								},
							},
						},
					},
				},
			},
		},
	}

	err = Recursion(wd, structure)

	if err != nil {
		panic(err)
	}
}
