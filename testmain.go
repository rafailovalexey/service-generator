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
	module := "github.com/emptyhopes/test"
	name := "employees"

	wd, _ := utils.GetWorkDirectory()
	wd = filepath.Join(wd, "test")

	_ = Recursion(GetStructure(module, name), wd)
	_ = Recursion(GetProviderStructure(wd, module, name), wd)
}

func Recursion(nodes *[]Node, path string) error {
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

		err := Recursion(node.Parent, current)

		if err != nil {
			return err
		}
	}

	return nil
}

func GetStructure(module string, name string) *[]Node {
	structure := &[]Node{
		{
			IsDirectory: true,
			Name:        "internal",
			Type:        "core",
			Parent: &[]Node{
				*GetGrcpServerLayerStructure(module, name),
				*GetHttpHandlerLayerStructure(module, name),
				*GetBaseLayerStructure(module, "controller", name),
				*GetBaseLayerStructure(module, "converter", name),
				*GetBaseLayerStructure(module, "validation", name),
				*GetBaseLayerStructure(module, "service", name),
				*GetBaseLayerStructure(module, "repository", name),
				*GetBaseLayerStructure(module, "client", name),
				*GetDataTransferObjectLayerStructure("dto", name),
				*GetDataTransferObjectLayerStructure("model", name),
			},
		},
	}

	return structure
}

func GetBaseLayerStructure(module string, layer string, name string) *Node {
	structure := &Node{
		IsDirectory: true,
		Name:        layer,
		Type:        layer,
		Parent: &[]Node{
			{
				IsFile:   true,
				Name:     utils.GetFilename(layer, "go"),
				Type:     layer,
				Template: test.GetBaseDefinitionTemplate(layer, name),
			},
			{
				IsDirectory: true,
				Name:        name,
				Type:        layer,
				Parent: &[]Node{
					{
						IsFile:   true,
						Name:     utils.GetFilename(layer, "go"),
						Type:     layer,
						Template: test.GetBaseImplementationTemplate(module, layer, name),
					},
				},
			},
		},
	}

	return structure
}

func GetDataTransferObjectLayerStructure(layer string, name string) *Node {
	structure := &Node{
		IsDirectory: true,
		Name:        layer,
		Type:        layer,
		Parent: &[]Node{
			{
				IsDirectory: true,
				Name:        name,
				Type:        layer,
				Parent: &[]Node{
					{
						IsFile:   true,
						Name:     utils.GetFilename(layer, "go"),
						Type:     layer,
						Template: test.GetDataTransferObjectTemplate(layer, name),
					},
				},
			},
		},
	}

	return structure
}

func GetGrcpServerLayerStructure(module string, name string) *Node {
	layer := "implementation"

	structure := &Node{
		IsDirectory: true,
		Name:        layer,
		Type:        layer,
		Parent: &[]Node{
			{
				IsDirectory: true,
				Name:        name,
				Type:        layer,
				Parent: &[]Node{
					{
						IsFile:   true,
						Name:     utils.GetFilename(layer, "go"),
						Type:     layer,
						Template: test.GetGrpcServerImplementationTemplate(module, name),
					},
				},
			},
		},
	}

	return structure
}

func GetHttpHandlerLayerStructure(module string, name string) *Node {
	layer := "handler"

	structure := &Node{
		IsDirectory: true,
		Name:        layer,
		Type:        layer,
		Parent: &[]Node{
			{
				IsFile:   true,
				Name:     utils.GetFilename(layer, "go"),
				Type:     layer,
				Template: test.GetHttpHandlerDefinitionTemplate(name),
			},
			{
				IsDirectory: true,
				Name:        name,
				Type:        layer,
				Parent: &[]Node{
					{
						IsFile:   true,
						Name:     utils.GetFilename(layer, "go"),
						Type:     layer,
						Template: test.GetHttpHandlerImplementationTemplate(module, name),
					},
				},
			},
		},
	}

	return structure
}

func GetProviderStructure(wd string, module string, name string) *[]Node {
	structure := &[]Node{
		{
			IsDirectory: true,
			Name:        "internal",
			Type:        "core",
			Parent: &[]Node{
				*GetProviderLayerStructure(wd, module, name),
			},
		},
	}

	return structure
}

func GetProviderLayerStructure(wd string, module string, name string) *Node {
	layer := "provider"

	available := make(map[string]struct{}, 10)

	available["handler"] = struct{}{}
	available["api"] = struct{}{}
	available["controller"] = struct{}{}
	available["implementation"] = struct{}{}
	available["validation"] = struct{}{}
	available["converter"] = struct{}{}
	available["service"] = struct{}{}
	available["repository"] = struct{}{}
	available["client"] = struct{}{}

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

	structure := &Node{
		IsDirectory: true,
		Name:        layer,
		Type:        layer,
		Parent: &[]Node{
			{
				IsFile:   true,
				Name:     utils.GetFilename(layer, "go"),
				Type:     layer,
				Template: test.GetProviderDefinitionTemplate(module, name, layers),
			},
			{
				IsDirectory: true,
				Name:        name,
				Type:        layer,
				Parent: &[]Node{
					{
						IsFile:   true,
						Name:     utils.GetFilename(layer, "go"),
						Type:     layer,
						Template: test.GetProviderImplementationTemplate(module, name, layers),
					},
				},
			},
		},
	}

	return structure
}
