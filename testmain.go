package main

import (
	"github.com/rafailovalexey/service-generator/internal/test"
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

func main() {
	application := "grpc"
	module := "github.com/emptyhopes/test"
	name := "employees"
	version := "1.20"

	wd, err := utils.GetWorkDirectory()

	if err != nil {
		panic(err)
	}

	wd = filepath.Join(wd, "test")

	Generate(wd, application, version, module, name)
	Provider(wd, module, name)
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
					Template: test.GetMockGenerateShellScriptTemplate(),
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
			Template: test.GetGitIgnoreTemplate(),
		},
		{
			IsFile:   true,
			Name:     utils.GetFilename(".dockerignore", ""),
			Template: test.GetDockerIgnoreTemplate(),
		},
		{
			IsFile:   true,
			Name:     utils.GetFilename(".env", ""),
			Template: test.GetEnvironmentTemplate(),
		},
		{
			IsFile:   true,
			Name:     utils.GetFilename(".example.env", ""),
			Template: test.GetExampleEnvironmentTemplate(),
		},
		{
			IsFile:   true,
			Name:     utils.GetFilename("main", "go"),
			Template: test.GetMainTemplate(module, application),
		},
		{
			IsFile:   true,
			Name:     utils.GetFilename("go", "mod"),
			Template: test.GetGoTemplate(module, version),
		},
		{
			IsFile:   true,
			Name:     utils.GetFilename("Makefile", ""),
			Template: test.GetMakefileTemplate(application, name),
		},
		{
			IsFile:   true,
			Name:     utils.GetFilename("README", "md"),
			Template: test.GetReadmeTemplate(module),
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
						IsFile:   true,
						Name:     utils.GetFilename(name, "proto"),
						Template: test.GetProtoTemplate(module, name),
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
						Template: test.GetGrpcGenerateShellScriptTemplate(),
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
								Name:     "grpc_server",
								Template: test.GetGrpcServerTemplate(module, name),
							},
							{
								IsDirectory: true,
								Name:        "interceptor",
								Parent: &[]Node{
									{
										IsFile:   true,
										Name:     utils.GetFilename("logging", "sh"),
										Template: test.GetGrpcLoggingInterceptorTemplate(),
									},
									{
										IsFile:   true,
										Name:     utils.GetFilename("tracing", "sh"),
										Template: test.GetGrpcTracingInterceptorTemplate(),
									},
								},
							},
							{
								IsDirectory: true,
								Name:        "middleware",
								Parent: &[]Node{
									{
										IsFile:   true,
										Name:     utils.GetFilename("authentication", "sh"),
										Template: test.GetGrpcAuthenticationMiddlewareTemplate(),
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
										Template: test.GetGrpcServerImplementationTemplate(module, name),
									},
								},
							},
						},
					},
				},
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
								Template: test.GetHttpServerTemplate(module, name),
							},
							{
								IsDirectory: true,
								Name:        "middleware",
								Parent: &[]Node{
									{
										IsFile:   true,
										Name:     utils.GetFilename("logging", "go"),
										Template: test.GetHttpLoggingInterceptorTemplate(),
									},
									{
										IsFile:   true,
										Name:     utils.GetFilename("authentication", "go"),
										Template: test.GetHttpAuthenticationMiddlewareTemplate(module),
									},
									{
										IsFile:   true,
										Name:     utils.GetFilename("cors", "go"),
										Template: test.GetHttpCorsMiddlewareTemplate(),
									},
									{
										IsFile:   true,
										Name:     utils.GetFilename("chain", "go"),
										Template: test.GetHttpChainMiddlewareTemplate(),
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
								Template: test.GetHttpHandlerDefinitionTemplate(name),
							},
							{
								IsDirectory: true,
								Name:        name,
								Parent: &[]Node{
									{
										IsFile:   true,
										Name:     utils.GetFilename("handler", "go"),
										Template: test.GetHttpHandlerImplementationTemplate(module, name),
									},
								},
							},
						},
					},
				},
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
				Template: test.GetBaseDefinitionTemplate(layer, name),
			},
			{
				IsDirectory: true,
				Name:        name,
				Parent: &[]Node{
					{
						IsFile:   true,
						Name:     utils.GetFilename(layer, "go"),
						Template: test.GetBaseImplementationTemplate(module, layer, name),
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
						Template: test.GetDataTransferObjectTemplate(layer, name),
					},
				},
			},
		},
	}

	return structure
}

func Provider(wd string, module string, name string) {
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
							Template: test.GetProviderDefinitionTemplate(module, name, layers),
						},
						{
							IsDirectory: true,
							Name:        name,
							Parent: &[]Node{
								{
									IsFile:   true,
									Name:     utils.GetFilename("provider", "go"),
									Template: test.GetProviderImplementationTemplate(module, name, layers),
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
