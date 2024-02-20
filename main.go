package main

import (
	"github.com/rafailovalexey/service-generator/internal/dto"
	"github.com/rafailovalexey/service-generator/internal/structure"
	"github.com/rafailovalexey/service-generator/internal/util"
	"path/filepath"
)

func main() {
	application := "grpc"

	module := "github.com/emptyhopes/whatsapp-messages-cron"

	name := &dto.NameDto{
		LowerCaseFirstLetter:   "w",
		CamelCaseSingular:      "WhatsappMessage",
		CamelCasePlural:        "WhatsappMessages",
		LowerCamelCaseSingular: "whatsappMessage",
		LowerCamelCasePlural:   "whatsappMessages",
		SnakeCaseSingular:      "whatsapp_message",
		SnakeCasePlural:        "whatsapp_messages",
	}

	version := "1.19"

	wd, err := util.GetWorkDirectory()

	if err != nil {
		panic(err)
	}

	wd = filepath.Join(wd, "whatsapp-messages-cron")

	structure.Generate(wd, application, version, module, name)
	structure.GenerateProvider(wd, module, name)
}
