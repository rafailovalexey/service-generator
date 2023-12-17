package application

import (
	"context"
	"github.com/rafailovalexey/service-generator/internal/modules/flags"
	"github.com/rafailovalexey/service-generator/internal/modules/layer"
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
