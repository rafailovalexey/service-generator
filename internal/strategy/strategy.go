package strategy

import (
	"github.com/rafailovalexey/service-generator/internal/facade"
	"github.com/rafailovalexey/service-generator/internal/template"
	"log"
)

type GenerationStrategyInterface interface {
	Generate(string, string) error
}

type DataTransferObjectGeneration struct{}

var _ GenerationStrategyInterface = (*DataTransferObjectGeneration)(nil)

func (d *DataTransferObjectGeneration) Generate(layer string, name string) error {
	err := facade.CreateDataTransferObject(layer, name)

	if err != nil {
		return err
	}

	return nil
}

type RealisationGeneration struct{}

var _ GenerationStrategyInterface = (*RealisationGeneration)(nil)

func (r *RealisationGeneration) Generate(layer string, name string) error {
	imports := &template.Imports{}
	methods := &template.Methods{}
	functions := &template.Functions{}

	err := facade.CreateInterface(layer, name, imports, methods)

	if err != nil {
		return err
	}

	err = facade.CreateRealisationInterface(layer, name, imports, methods, functions)

	if err != nil {
		return err
	}

	return nil
}

type IncomingGeneration struct{}

var _ GenerationStrategyInterface = (*IncomingGeneration)(nil)

func (c *IncomingGeneration) Generate(layer string, name string) error {
	imports := &template.Imports{}
	methods := &template.Methods{}
	functions := &template.Functions{}

	err := facade.CreateInterface(layer, name, imports, methods)

	if err != nil {
		log.Panicf("%v\n", err)
	}

	err = facade.CreateRealisationInterface(layer, name, imports, methods, functions)

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
