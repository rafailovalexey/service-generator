package main

import (
	"github.com/rafailovalexey/service-generator/internal/dto"
	"github.com/rafailovalexey/service-generator/internal/facade"
	"log"
)

func main() {
	names := dto.NewNameDto(
		"w",
		"WhatsappClientCron",
		"WhatsappClientsCron",
		"whatsappClientCron",
		"whatsappClientsCron",
		"whatsapp_client_cron",
		"whatsapp_clients_cron",
	)

	application := dto.NewApplicationDto(
		"1.20",
		"cron",
		"postgres",
		"github.com",
		"rafailovalexey",
		"whatsapp-clients-cron",
		names,
	)

	err := facade.Create(application)

	if err != nil {
		log.Panicf("error %v", err)
	}
}
