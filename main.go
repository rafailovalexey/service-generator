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
}
