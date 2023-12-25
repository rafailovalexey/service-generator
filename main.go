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
	version := "1.19"

	name := "employees"

	application := "application"
	//application := "subscribe"
	//application := "cron"

	//implementing := "grpc"
	implementing := "http"
	//implementing := "subscribe"
	//implementing := "cron"

	switch implementing {
	case "grpc":
		implementing = "grpc_server"
	case "http":
		implementing = "http_server"
	case "cron":
		implementing = "cron_scheduler"
	case "subscribe":
		implementing = "nats_subscribe"
	}

	err := strategy.GenerateHttpApplication(wd, module, version, application, name, implementing)

	if err != nil {
		log.Panicf("%s\v", err)
	}
}
