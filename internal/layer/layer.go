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

	dictionary["api"] = &strategy.IncomingGeneration{}
	dictionary["controller"] = &strategy.IncomingGeneration{}
	dictionary["client"] = &strategy.IncomingGeneration{}

	// provider
	// handler
	// implementation

	dictionary["service"] = &strategy.RealisationGeneration{}
	dictionary["repository"] = &strategy.RealisationGeneration{}
	dictionary["converter"] = &strategy.RealisationGeneration{}
	dictionary["validation"] = &strategy.RealisationGeneration{}

	dictionary["dto"] = &strategy.DataTransferObjectGeneration{}
	dictionary["model"] = &strategy.DataTransferObjectGeneration{}

	return dictionary
}
