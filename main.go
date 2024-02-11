package main

import (
	"github.com/rafailovalexey/service-generator/internal/structure"
	"github.com/rafailovalexey/service-generator/internal/util"
	"path/filepath"
)

func main() {
	application := "grpc"
	module := "github.com/emptyhopes/employees"
	name := "employees"
	version := "1.19"

	wd, err := util.GetWorkDirectory()

	if err != nil {
		panic(err)
	}

	wd = filepath.Join(wd, "employees")

	structure.Generate(wd, application, version, module, name)
	structure.GenerateProvider(wd, module, name)
}
