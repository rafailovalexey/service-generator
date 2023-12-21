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
		log.Fatalf("error %v", err)
	}

	for _, key := range f.Layers {
		isExist := l.IsExist(key)

		if !isExist {
			log.Fatalf("layer not found")
		}

		value, err := l.GetLayer(key)

		if err != nil {
			log.Fatalf("error %v", err)
		}

		err = value.Generate(key, f.Name)

		if err != nil {
			log.Fatalf("error %v", err)
		}
	}
}
