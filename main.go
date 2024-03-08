package main

import (
	"fmt"
	"github.com/rafailovalexey/service-generator/internal/dto"
	"github.com/rafailovalexey/service-generator/internal/structure"
	"github.com/rafailovalexey/service-generator/internal/util"
	"log"
	"path/filepath"
)

func main() {
	version := "1.20"
	application := "grpc"
	database := "mysql"

	organization := "github.com"
	project := "rafailovalexey"
	directory := "whatsapp-clients"

	name := &dto.NameDto{
		LowerCaseFirstLetter:   "w",
		CamelCaseSingular:      "WhatsappClient",
		CamelCasePlural:        "WhatsappClients",
		LowerCamelCaseSingular: "whatsappClient",
		LowerCamelCasePlural:   "whatsappClients",
		SnakeCaseSingular:      "whatsapp_client",
		SnakeCasePlural:        "whatsapp_clients",
	}

	module := fmt.Sprintf("%s/%s/%s", organization, project, directory)

	wd, err := util.GetWorkDirectory()

	if err != nil {
		log.Panicf(err.Error())
	}

	wd = filepath.Join(wd, directory)

	err = structure.Generate(wd, application, organization, version, database, module, name)

	if err != nil {
		log.Panicf(err.Error())
	}

	err = structure.GenerateProvider(wd, module, name)

	if err != nil {
		log.Panicf(err.Error())
	}
}
