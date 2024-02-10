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
	Type        string
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

	structure := &[]Node{
		{
			IsDirectory: true,
			Name:        "bin",
			Type:        "core",
			Parent: &[]Node{
				{
					IsFile:   true,
					Name:     utils.GetFilename("grpc-generate", "sh"),
					Type:     "core",
					Template: test.GetGrpcGenerateShellScriptTemplate(),
				},
				{
					IsFile:   true,
					Name:     utils.GetFilename("mock-generate", "sh"),
					Type:     "core",
					Template: test.GetMockGenerateShellScriptTemplate(),
				},
			},
		},
		{
			IsDirectory: true,
			Name:        "internal",
			Type:        "core",
			Parent: &[]Node{
				{
					IsDirectory: true,
					Name:        "handler",
					Type:        "handler",
					Parent: &[]Node{
						{
							IsFile:   true,
							Name:     utils.GetFilename("handler", "go"),
							Type:     "handler",
							Template: test.GetHttpHandlerDefinitionTemplate(name),
						},
						{
							IsDirectory: true,
							Name:        name,
							Type:        "handler",
							Parent: &[]Node{
								{
									IsFile:   true,
									Name:     utils.GetFilename("handler", "go"),
									Type:     "handler",
									Template: test.GetHttpHandlerImplementationTemplate(module, name),
								},
							},
						},
					},
				},
				{
					IsDirectory: true,
					Name:        "implementation",
					Type:        "implementation",
					Parent: &[]Node{
						{
							IsDirectory: true,
							Name:        name,
							Type:        "implementation",
							Parent: &[]Node{
								{
									IsFile:   true,
									Name:     utils.GetFilename("implementation", "go"),
									Type:     "implementation",
									Template: test.GetGrpcServerImplementationTemplate(module, name),
								},
							},
						},
					},
				},
				{
					IsDirectory: true,
					Name:        "controller",
					Type:        "controller",
					Parent: &[]Node{
						{
							IsFile:   true,
							Name:     utils.GetFilename("controller", "go"),
							Type:     "controller",
							Template: test.GetBaseDefinitionTemplate("controller", name),
						},
						{
							IsDirectory: true,
							Name:        name,
							Type:        "controller",
							Parent: &[]Node{
								{
									IsFile:   true,
									Name:     utils.GetFilename("controller", "go"),
									Type:     "controller",
									Template: test.GetBaseImplementationTemplate(module, "controller", name),
								},
							},
						},
					},
				},
				{
					IsDirectory: true,
					Name:        "converter",
					Type:        "converter",
					Parent: &[]Node{
						{
							IsFile:   true,
							Name:     utils.GetFilename("converter", "go"),
							Type:     "converter",
							Template: test.GetBaseDefinitionTemplate("converter", name),
						},
						{
							IsDirectory: true,
							Name:        name,
							Type:        "converter",
							Parent: &[]Node{
								{
									IsFile:   true,
									Name:     utils.GetFilename("converter", "go"),
									Type:     "converter",
									Template: test.GetBaseImplementationTemplate(module, "converter", name),
								},
							},
						},
					},
				},
				{
					IsDirectory: true,
					Name:        "validation",
					Type:        "validation",
					Parent: &[]Node{
						{
							IsFile:   true,
							Name:     utils.GetFilename("validation", "go"),
							Type:     "validation",
							Template: test.GetBaseDefinitionTemplate("validation", name),
						},
						{
							IsDirectory: true,
							Name:        name,
							Type:        "validation",
							Parent: &[]Node{
								{
									IsFile:   true,
									Name:     utils.GetFilename("validation", "go"),
									Type:     "validation",
									Template: test.GetBaseImplementationTemplate(module, "validation", name),
								},
							},
						},
					},
				},
				{
					IsDirectory: true,
					Name:        "service",
					Type:        "service",
					Parent: &[]Node{
						{
							IsFile:   true,
							Name:     utils.GetFilename("service", "go"),
							Type:     "service",
							Template: test.GetBaseDefinitionTemplate("service", name),
						},
						{
							IsDirectory: true,
							Name:        name,
							Type:        "service",
							Parent: &[]Node{
								{
									IsFile:   true,
									Name:     utils.GetFilename("service", "go"),
									Type:     "service",
									Template: test.GetBaseImplementationTemplate(module, "service", name),
								},
							},
						},
					},
				},
				{
					IsDirectory: true,
					Name:        "repository",
					Type:        "repository",
					Parent: &[]Node{
						{
							IsFile:   true,
							Name:     utils.GetFilename("repository", "go"),
							Type:     "repository",
							Template: test.GetBaseDefinitionTemplate("repository", name),
						},
						{
							IsDirectory: true,
							Name:        name,
							Type:        "repository",
							Parent: &[]Node{
								{
									IsFile:   true,
									Name:     utils.GetFilename("repository", "go"),
									Type:     "repository",
									Template: test.GetBaseImplementationTemplate(module, "repository", name),
								},
							},
						},
					},
				},
				{
					IsDirectory: true,
					Name:        "dto",
					Type:        "dto",
					Parent: &[]Node{
						{
							IsDirectory: true,
							Name:        name,
							Type:        "dto",
							Parent: &[]Node{
								{
									IsFile:   true,
									Name:     utils.GetFilename("dto", "go"),
									Type:     "dto",
									Template: test.GetDataTransferObjectTemplate("dto", name),
								},
							},
						},
					},
				},
				{
					IsDirectory: true,
					Name:        "model",
					Type:        "model",
					Parent: &[]Node{
						{
							IsDirectory: true,
							Name:        name,
							Type:        "model",
							Parent: &[]Node{
								{
									IsFile:   true,
									Name:     utils.GetFilename("model", "go"),
									Type:     "model",
									Template: test.GetDataTransferObjectTemplate("model", name),
								},
							},
						},
					},
				},
			},
		},
		{
			IsFile:   true,
			Name:     "Makefile",
			Type:     "core",
			Template: test.GetMakefileTemplate(application, name),
		},
		{
			IsFile:   true,
			Name:     "go.mod",
			Type:     "core",
			Template: test.GetGoTemplate(module, version),
		},
	}

	err = Recursion(wd, structure)

	if err != nil {
		panic(err)
	}

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

	structure = &[]Node{
		{
			IsDirectory: true,
			Name:        "internal",
			Type:        "core",
			Parent: &[]Node{
				{
					IsDirectory: true,
					Name:        "provider",
					Type:        "provider",
					Parent: &[]Node{
						{
							IsFile:   true,
							Name:     utils.GetFilename("provider", "go"),
							Type:     "provider",
							Template: test.GetProviderDefinitionTemplate(module, name, layers),
						},
						{
							IsDirectory: true,
							Name:        name,
							Type:        "provider",
							Parent: &[]Node{
								{
									IsFile:   true,
									Name:     utils.GetFilename("provider", "go"),
									Type:     "provider",
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
