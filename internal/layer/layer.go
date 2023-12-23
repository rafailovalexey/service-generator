package layer

import (
	"fmt"
	"github.com/rafailovalexey/service-generator/internal/strategy"
	"sync"
)

type LayerInterface interface {
	IsExist(string) bool
	GetLayer(string) (strategy.GenerationStrategyInterface, error)
}

type Layer struct {
	dictionary map[string]strategy.GenerationStrategyInterface
	rwmutex    sync.RWMutex
}

var _ LayerInterface = (*Layer)(nil)

func NewLayer() *Layer {
	dictionary := GetDictionary()

	return &Layer{
		dictionary: dictionary,
		rwmutex:    sync.RWMutex{},
	}
}

func (l *Layer) IsExist(value string) bool {
	l.rwmutex.RLock()
	defer l.rwmutex.RUnlock()

	_, isExist := l.dictionary[value]

	return isExist
}

func (l *Layer) GetLayer(value string) (strategy.GenerationStrategyInterface, error) {
	l.rwmutex.RLock()
	defer l.rwmutex.RUnlock()

	element, isExist := l.dictionary[value]

	if !isExist {
		return nil, fmt.Errorf("element not found")
	}

	return element, nil
}

func GetDictionary() map[string]strategy.GenerationStrategyInterface {
	dictionary := make(map[string]strategy.GenerationStrategyInterface, 100)

	dictionary["api"] = &strategy.IncomingGenerationStrategy{}
	dictionary["controller"] = &strategy.IncomingGenerationStrategy{}
	dictionary["client"] = &strategy.IncomingGenerationStrategy{}

	dictionary["provider"] = &strategy.ProviderGenerationStrategy{}

	dictionary["implementation"] = &strategy.ImplementationGenerationStrategy{}

	// handler

	dictionary["service"] = &strategy.RealisationGenerationStrategy{}
	dictionary["repository"] = &strategy.RealisationGenerationStrategy{}
	dictionary["converter"] = &strategy.RealisationGenerationStrategy{}
	dictionary["validation"] = &strategy.RealisationGenerationStrategy{}

	dictionary["dto"] = &strategy.DataTransferObjectGenerationStrategy{}
	dictionary["model"] = &strategy.DataTransferObjectGenerationStrategy{}

	return dictionary
}
