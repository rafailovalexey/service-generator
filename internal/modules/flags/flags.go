package flags

import (
	"flag"
	"fmt"
)

type FlagsInterface interface {
	InitializeFlags() error
}

type Flags struct {
	Layer string
	Name  string
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

	f.Layer = *layer
	f.Name = *name

	return nil
}
