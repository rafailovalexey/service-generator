package main

import (
	"fmt"
	"github.com/rafailovalexey/service-generator/internal/template"
	"github.com/rafailovalexey/service-generator/internal/utils"
	"path/filepath"
)

type Node struct {
	IsDirectory bool
	IsFile      bool
	Name        string
	Type        string
	Template    []byte
	Parent      []Node
}

func main() {
	structure := []Node{
		{
			IsDirectory: true,
			Name:        "internal",
			Type:        "core",
			Parent: []Node{
				*GetBaseLayerStructure("service", "employees"),
			},
		},
	}

	wd, _ := utils.GetWorkDirectory()

	wd = filepath.Join(wd, "test2")

	err := Recursion(structure, wd)

	if err != nil {
		fmt.Println(err)
	}
}

func GetBaseLayerStructure(layer string, name string) *Node {
	structure := &Node{
		IsDirectory: true,
		Name:        layer,
		Type:        layer,
		Parent: []Node{
			{
				IsFile:   true,
				Name:     utils.GetFilename(layer, "go"),
				Type:     layer,
				Template: template.GetInterfaceTemplate(layer, name),
			},
			{
				IsDirectory: true,
				Name:        name,
				Type:        layer,
				Parent: []Node{
					{
						IsFile:   true,
						Name:     utils.GetFilename(layer, "go"),
						Type:     layer,
						Template: template.GetRealisationInterfaceTemplate(module, kind, layer, name),
					},
				},
			},
		},
	}

	return structure
}

func Recursion(nodes []Node, path string) error {
	if nodes == nil {
		return nil
	}

	if len(nodes) == 0 {
		return nil
	}

	for _, node := range nodes {
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
