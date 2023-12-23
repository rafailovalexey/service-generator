package strategy

import (
	"github.com/rafailovalexey/service-generator/internal/facade"
	"log"
)

type GenerationStrategyInterface interface {
	Generate(string, string) error
}

type DataTransferObjectGenerationStrategy struct{}

var _ GenerationStrategyInterface = (*DataTransferObjectGenerationStrategy)(nil)

func (d *DataTransferObjectGenerationStrategy) Generate(layer string, name string) error {
	err := facade.CreateDataTransferObject(layer, name)

	if err != nil {
		return err
	}

	return nil
}

type RealisationGenerationStrategy struct{}

var _ GenerationStrategyInterface = (*RealisationGenerationStrategy)(nil)

func (r *RealisationGenerationStrategy) Generate(layer string, name string) error {
	err := facade.CreateInterface(layer, name)

	if err != nil {
		return err
	}

	err = facade.CreateRealisationInterface(layer, name)

	if err != nil {
		return err
	}

	return nil
}

type IncomingGenerationStrategy struct{}

var _ GenerationStrategyInterface = (*IncomingGenerationStrategy)(nil)

func (c *IncomingGenerationStrategy) Generate(layer string, name string) error {
	err := facade.CreateInterface(layer, name)

	if err != nil {
		log.Panicf("%v\n", err)
	}

	err = facade.CreateRealisationInterface(layer, name)

	if err != nil {
		log.Panicf("%v\n", err)
	}

	err = facade.CreateRequestObject(layer, name)

	if err != nil {
		log.Panicf("%v\n", err)
	}

	err = facade.CreateResponseObject(layer, name)

	if err != nil {
		log.Panicf("%v\n", err)
	}

	return nil
}

type ProviderGenerationStrategy struct{}

var _ GenerationStrategyInterface = (*ProviderGenerationStrategy)(nil)

func (c *ProviderGenerationStrategy) Generate(layer string, name string) error {
	err := facade.CreateProviderInterface(layer, name)

	if err != nil {
		log.Panicf("%v\n", err)
	}

	err = facade.CreateProvider(layer, name)

	if err != nil {
		log.Panicf("%v\n", err)
	}

	return nil
}

type ImplementationGenerationStrategy struct{}

var _ GenerationStrategyInterface = (*ImplementationGenerationStrategy)(nil)

func (c *ImplementationGenerationStrategy) Generate(layer string, name string) error {
	err := facade.CreateImplementation(layer, name)

	if err != nil {
		log.Panicf("%v\n", err)
	}

	return nil
}
