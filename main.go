package main

import (
	"github.com/rafailovalexey/service-generator/internal/dto"
	"github.com/rafailovalexey/service-generator/internal/facade"
	"log"
)

func main() {
	names := dto.NewNameDto(
		"w",
		"WhatsappClientHttp",
		"WhatsappClientsHttp",
		"whatsappClientHttp",
		"whatsappClientsHttp",
		"whatsapp_client_http",
		"whatsapp_clients_http",
	)

	application := dto.NewApplicationDto(
		"1.20",
		"http",
		"postgres",
		"github.com",
		"rafailovalexey",
		"whatsapp-clients-http",
		names,
	)

	err := facade.Create(application)

	if err != nil {
		log.Panicf("error %v", err)
	}
}
