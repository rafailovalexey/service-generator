package strategy

import (
	"github.com/rafailovalexey/service-generator/internal/modules/builder"
	"github.com/rafailovalexey/service-generator/internal/modules/flags"
)

type GenerationStrategyInterface interface {
	Generate(*flags.Flags) error
}

type DataTransferObjectGeneration struct{}

var _ GenerationStrategyInterface = (*DataTransferObjectGeneration)(nil)

func (c *DataTransferObjectGeneration) Generate(f *flags.Flags) error {
	err := builder.CreateDataTransferObject(f.Layer, f.Name)

	if err != nil {
		return err
	}

	return nil
}

type RealisationGeneration struct{}

var _ GenerationStrategyInterface = (*RealisationGeneration)(nil)

func (p *RealisationGeneration) Generate(f *flags.Flags) error {
	err := builder.CreateLayer(f.Layer, f.Name)

	if err != nil {
		return err
	}

	err = builder.CreateImplementation(f.Layer, f.Name)

	if err != nil {
		return err
	}

	return nil
}
