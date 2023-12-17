package main

import (
	"github.com/rafailovalexey/service-generator/internal/modules/flags"
	"github.com/rafailovalexey/service-generator/internal/modules/layer"
	"log"
)

func main() {
	f, err := flags.NewFlags()
	l := layer.NewLayer()

	if err != nil {
		log.Fatalf("error %v", err)
	}

	isExist := l.IsExist(f.Layer)

	if !isExist {
		log.Fatalf("layer not found")
	}

	value, err := l.GetLayer(f.Layer)

	if err != nil {
		log.Fatalf("error %v", err)
	}

	err = value.Generate(f)

	if err != nil {
		log.Fatalf("error %v", err)
	}
}
