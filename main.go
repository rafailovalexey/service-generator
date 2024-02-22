package main

import (
	"github.com/rafailovalexey/service-generator/internal/dto"
	"github.com/rafailovalexey/service-generator/internal/structure"
	"github.com/rafailovalexey/service-generator/internal/util"
	"path/filepath"
)

func main() {
	application := "http"

	module := "git.amocrm.ru/whatsapp-lite/whatsapp-lite-messages"

	name := &dto.NameDto{
		LowerCaseFirstLetter:   "w",
		CamelCaseSingular:      "WhatsappLiteMessage",
		CamelCasePlural:        "WhatsappLiteMessages",
		LowerCamelCaseSingular: "whatsappLiteMessage",
		LowerCamelCasePlural:   "whatsappLiteMessages",
		SnakeCaseSingular:      "whatsapp_lite_message",
		SnakeCasePlural:        "whatsapp_lite_messages",
	}

	database := "mysql"

	version := "1.19"

	wd, err := util.GetWorkDirectory()

	if err != nil {
		panic(err)
	}

	wd = filepath.Join(wd, "whatsapp-lite-messages")

	structure.Generate(wd, application, version, database, module, name)
	structure.GenerateProvider(wd, module, name)
}
