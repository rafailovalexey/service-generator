package main

import (
	"github.com/rafailovalexey/service-generator/internal/structure"
	"github.com/rafailovalexey/service-generator/internal/utils"
	"path/filepath"
)

func main() {
	application := "http"
	module := "github.com/emptyhopes/template"
	name := "employees"
	version := "1.19"

	wd, err := utils.GetWorkDirectory()

	if err != nil {
		panic(err)
	}

	wd = filepath.Join(wd, "template")

	structure.Generate(wd, application, version, module, name)
	structure.GenerateProvider(wd, module, name)
}
