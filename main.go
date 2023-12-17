package main

import (
	"github.com/rafailovalexey/service-generator/internal/modules/builder"
	"github.com/rafailovalexey/service-generator/internal/modules/flags"
	"log"
)

func main() {
	f, err := flags.NewFlags()

	if err != nil {
		log.Fatalf("error %v", err)
	}

	//err = builder.CreateDataTransferObject(f.Layer, f.Name)
	//
	//if err != nil {
	//	log.Fatalf("error %v", err)
	//}

	err = builder.CreateLayer(f.Layer, f.Name)

	if err != nil {
		log.Fatalf("error %v", err)
	}

	err = builder.CreateImplementation(f.Layer, f.Name)

	if err != nil {
		log.Fatalf("error %v", err)
	}
}
