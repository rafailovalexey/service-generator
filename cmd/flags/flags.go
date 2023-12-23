package flags

import (
	"flag"
	"fmt"
	"strings"
)

type FlagsInterface interface {
	InitializeFlags() error
}

type Flags struct {
	Layers map[string]bool
	Name   string
}

var _ FlagsInterface = (*Flags)(nil)

func NewFlags() (*Flags, error) {
	flags := &Flags{}

	err := flags.InitializeFlags()

	if err != nil {
		return nil, err
	}

	return flags, nil
}

func (f *Flags) InitializeFlags() error {
	layer := flag.String("layer", "", "")
	name := flag.String("name", "", "")

	flag.Parse()

	if *layer == "" {
		return fmt.Errorf("layer is empty")
	}

	if *name == "" {
		return fmt.Errorf("name is empty")
	}

	layers := make(map[string]bool, 100)

	for _, value := range strings.Split(*layer, ",") {
		layers[value] = false

		if value == "provider" {
			layers[value] = true
		}
	}

	f.Layers = layers
	f.Name = *name

	return nil
}
