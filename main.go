package main

import (
	"github.com/rafailovalexey/service-generator/internal/strategy"
	"github.com/rafailovalexey/service-generator/internal/utils"
	"log"
	"path"
)

func main() {
	wd, _ := utils.GetWorkDirectory()
	wd = path.Join(wd, "test")

	module := "github.com/rafailovalexey/service-test"
	name := "employees"

	err := strategy.GenerateSubscribeApplication(wd, module, name)

	if err != nil {
		log.Panicf("%s\v", err)
	}

	err = strategy.GenerateRealisation(wd, module, "repository", name)

	if err != nil {
		log.Panicf("%s\v", err)
	}

	err = strategy.GenerateRealisation(wd, module, "service", name)

	if err != nil {
		log.Panicf("%s\v", err)
	}

	err = strategy.GenerateRealisation(wd, module, "client", name)

	if err != nil {
		log.Panicf("%s\v", err)
	}

	err = strategy.GenerateRealisation(wd, module, "converter", name)

	if err != nil {
		log.Panicf("%s\v", err)
	}

	err = strategy.GenerateRealisation(wd, module, "validation", name)

	if err != nil {
		log.Panicf("%s\v", err)
	}

	err = strategy.GenerateProvider(wd, module, name)

	if err != nil {
		log.Panicf("%s\v", err)
	}
}
