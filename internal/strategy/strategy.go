package strategy

import (
	"github.com/rafailovalexey/service-generator/internal/facade"
	"github.com/rafailovalexey/service-generator/internal/utils"
	"log"
)

type GenerationStrategyInterface interface {
	Generate(string, string, string) error
}

type DataTransferObjectGenerationStrategy struct{}

var _ GenerationStrategyInterface = (*DataTransferObjectGenerationStrategy)(nil)

func (d *DataTransferObjectGenerationStrategy) Generate(_ string, layer string, name string) error {
	wd, err := utils.GetWorkDirectory()

	if err != nil {
		return err
	}

	err = facade.CreateDataTransferObject(wd, layer, name)

	if err != nil {
		return err
	}

	return nil
}

type RealisationGenerationStrategy struct{}

var _ GenerationStrategyInterface = (*RealisationGenerationStrategy)(nil)

func (r *RealisationGenerationStrategy) Generate(module string, layer string, name string) error {
	wd, err := utils.GetWorkDirectory()

	if err != nil {
		return err
	}

	err = facade.CreateInterface(wd, layer, name)

	if err != nil {
		return err
	}

	err = facade.CreateRealisationInterface(wd, module, layer, name)

	if err != nil {
		return err
	}

	return nil
}

type IncomingGenerationStrategy struct{}

var _ GenerationStrategyInterface = (*IncomingGenerationStrategy)(nil)

func (c *IncomingGenerationStrategy) Generate(module string, layer string, name string) error {
	wd, err := utils.GetWorkDirectory()

	if err != nil {
		return err
	}

	err = facade.CreateInterface(wd, layer, name)

	if err != nil {
		log.Panicf("%v\n", err)
	}

	err = facade.CreateRealisationInterface(wd, module, layer, name)

	if err != nil {
		log.Panicf("%v\n", err)
	}

	err = facade.CreateRequestObject(wd, layer, name)

	if err != nil {
		log.Panicf("%v\n", err)
	}

	err = facade.CreateResponseObject(wd, layer, name)

	if err != nil {
		log.Panicf("%v\n", err)
	}

	return nil
}

type ProviderGenerationStrategy struct{}

var _ GenerationStrategyInterface = (*ProviderGenerationStrategy)(nil)

func (c *ProviderGenerationStrategy) Generate(module string, _ string, name string) error {
	wd, err := utils.GetWorkDirectory()

	if err != nil {
		return err
	}

	err = facade.CreateProviderInterface(wd, module, name)

	if err != nil {
		log.Panicf("%v\n", err)
	}

	err = facade.CreateProvider(wd, module, name)

	if err != nil {
		log.Panicf("%v\n", err)
	}

	return nil
}

type ImplementationGenerationStrategy struct{}

var _ GenerationStrategyInterface = (*ImplementationGenerationStrategy)(nil)

func (c *ImplementationGenerationStrategy) Generate(_ string, _ string, name string) error {
	wd, err := utils.GetWorkDirectory()

	if err != nil {
		return err
	}

	err = facade.CreateImplementation(wd, name)

	if err != nil {
		log.Panicf("%v\n", err)
	}

	return nil
}
