package main

import (
	"github.com/rafailovalexey/service-generator/internal/dto"
	"github.com/rafailovalexey/service-generator/internal/structure"
	"github.com/rafailovalexey/service-generator/internal/util"
	"path/filepath"
)

func main() {
	application := "cron"

	module := "git.amocrm.ru/whatsapp-lite/whatsapp-lite-messages-cron"

	name := &dto.NameDto{
		LowerCaseFirstLetter:   "w",
		CamelCaseSingular:      "WhatsappLiteMessage",
		CamelCasePlural:        "WhatsappLiteMessages",
		LowerCamelCaseSingular: "whatsappLiteMessage",
		LowerCamelCasePlural:   "whatsappLiteMessages",
		SnakeCaseSingular:      "whatsapp_lite_message",
		SnakeCasePlural:        "whatsapp_lite_messages",
	}

	version := "1.19"

	wd, err := util.GetWorkDirectory()

	if err != nil {
		panic(err)
	}

	wd = filepath.Join(wd, "whatsapp-lite-messages-cron")

	structure.Generate(wd, application, version, module, name)
	structure.GenerateProvider(wd, module, name)
}
