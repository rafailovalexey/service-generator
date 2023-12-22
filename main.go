package main

import (
	"context"
	"github.com/rafailovalexey/service-generator/cmd/application"
	"log"
)

func main() {
	ctx := context.Background()

	a, err := application.NewApplication(ctx)

	if err != nil {
		log.Panicf("an error occurred while starting the utils %v\n", err)
	}

	a.Run()
}
