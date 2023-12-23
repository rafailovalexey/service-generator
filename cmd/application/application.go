package application

import (
	"context"
	"github.com/rafailovalexey/service-generator/cmd/flags"
	"github.com/rafailovalexey/service-generator/internal/layer"
	"log"
)

type ApplicationInterface interface {
	Run()
}

type Application struct{}

var _ ApplicationInterface = (*Application)(nil)

func NewApplication(ctx context.Context) (*Application, error) {
	return &Application{}, nil
}

func (a *Application) Run() {
	f, err := flags.NewFlags()
	l := layer.NewLayer()

	if err != nil {
		log.Panicf("error %v\n", err)
	}

	last := make(map[string]struct{}, 10)

	for key, isNeedLast := range f.Layers {
		if isNeedLast {
			last[key] = struct{}{}

			continue
		}

		isExist := l.IsExist(key)

		if !isExist {
			log.Panicf("layer not found\n")
		}

		value, err := l.GetLayer(key)

		if err != nil {
			log.Panicf("error %v\n", err)
		}

		err = value.Generate(key, f.Name)

		if err != nil {
			log.Panicf("error %v\n", err)
		}
	}

	for key, _ := range last {
		isExist := l.IsExist(key)

		if !isExist {
			log.Panicf("layer not found\n")
		}

		value, err := l.GetLayer(key)

		if err != nil {
			log.Panicf("error %v\n", err)
		}

		err = value.Generate(key, f.Name)

		if err != nil {
			log.Panicf("error %v\n", err)
		}
	}
}
