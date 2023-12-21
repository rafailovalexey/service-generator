package strategy

import (
	"github.com/rafailovalexey/service-generator/internal/facade"
)

type GenerationStrategyInterface interface {
	Generate(string, string) error
}

type DataTransferObjectGeneration struct{}

var _ GenerationStrategyInterface = (*DataTransferObjectGeneration)(nil)

func (c *DataTransferObjectGeneration) Generate(layer string, name string) error {
	err := facade.CreateDataTransferObject(layer, name)

	if err != nil {
		return err
	}

	return nil
}

type RealisationGeneration struct{}

var _ GenerationStrategyInterface = (*RealisationGeneration)(nil)

func (p *RealisationGeneration) Generate(layer string, name string) error {
	err := facade.CreateLayer(layer, name)

	if err != nil {
		return err
	}

	err = facade.CreateImplementation(layer, name)

	if err != nil {
		return err
	}

	return nil
}
