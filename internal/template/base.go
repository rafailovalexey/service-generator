package template

import (
	"bytes"
	"fmt"
	"github.com/rafailovalexey/service-generator/internal/util"
	"sort"
)

func GetBaseDefinitionTemplate(layer string, name string) []byte {
	data := bytes.Buffer{}
	separator := util.GetSeparator()

	data.WriteString(fmt.Sprintf("package %s", util.Lowercase(layer)))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("type %s%sInterface interface {}", util.Capitalize(util.SingularForm(name)), util.Capitalize(layer)))
	data.WriteString(separator)

	return data.Bytes()
}

func GetBaseImplementationTemplate(module string, layer string, name string) []byte {
	data := bytes.Buffer{}
	separator := util.GetSeparator()

	imports := []string{
		fmt.Sprintf("definition \"%s/%s/%s\"", module, "internal", util.Lowercase(layer)),
	}

	sort.Strings(imports)

	data.WriteString(fmt.Sprintf("package %s", util.Lowercase(name)))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("import ("))
	data.WriteString(separator)

	for _, i := range imports {
		data.WriteString(fmt.Sprintf("\t%s", i))
		data.WriteString(separator)
	}

	data.WriteString(fmt.Sprintf(")"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("type %s%s struct {}", util.Capitalize(util.SingularForm(name)), util.Capitalize(layer)))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("var _ definition.%s%sInterface = (*%s%s)(nil)", util.Capitalize(util.SingularForm(name)), util.Capitalize(layer), util.Capitalize(util.SingularForm(name)), util.Capitalize(layer)))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("func New%s%s() definition.%s%sInterface {", util.Capitalize(util.SingularForm(name)), util.Capitalize(layer), util.Capitalize(util.SingularForm(name)), util.Capitalize(layer)))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\treturn &%s%s{}", util.Capitalize(util.SingularForm(name)), util.Capitalize(layer)))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)

	return data.Bytes()
}

func GetDataTransferObjectTemplate(layer string, name string) []byte {
	data := bytes.Buffer{}
	separator := util.GetSeparator()

	data.WriteString(fmt.Sprintf("package %s", util.Lowercase(name)))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("type %s%s struct {}", util.Capitalize(util.SingularForm(name)), util.Capitalize(layer)))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("type %s%s = []%s%s", util.Capitalize(name), util.Capitalize(layer), util.Capitalize(util.SingularForm(name)), util.Capitalize(layer)))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("func New%s%s() *%s%s {", util.Capitalize(util.SingularForm(name)), util.Capitalize(layer), util.Capitalize(util.SingularForm(name)), util.Capitalize(layer)))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\treturn &%s%s{}", util.Capitalize(util.SingularForm(name)), util.Capitalize(layer)))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("func New%s%s(%s ...%s%s) *%s%s {", util.Capitalize(name), util.Capitalize(layer), util.Lowercase(name), util.Capitalize(util.SingularForm(name)), util.Capitalize(layer), util.Capitalize(name), util.Capitalize(layer)))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t%s := make(%s%s, len(%s))", util.Lowercase(layer), util.Capitalize(name), util.Capitalize(layer), util.Lowercase(name)))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\tfor _, %s := range %s {", util.Lowercase(util.SingularForm(name)), util.Lowercase(name)))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\t%s = append(%s, %s)", util.Lowercase(layer), util.Lowercase(layer), util.Lowercase(util.SingularForm(name))))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\treturn &%s", util.Lowercase(layer)))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)

	return data.Bytes()
}

func GetProviderDefinitionTemplate(module string, name string, layers []string) []byte {
	data := bytes.Buffer{}
	separator := util.GetSeparator()

	imports := make([]string, 0, 10)

	for _, l := range layers {
		switch l {
		case "implementation":
			imports = append(imports, fmt.Sprintf("%s%s \"%s/%s/%s/%s\"", util.SingularForm(name), util.Capitalize(l), module, "internal", l, name))
		default:
			imports = append(imports, fmt.Sprintf("\"%s/%s/%s\"", module, "internal", l))
		}
	}

	sort.Strings(imports)

	data.WriteString(fmt.Sprintf("package %s", util.Lowercase("provider")))
	data.WriteString(separator)
	data.WriteString(separator)

	if len(layers) != 0 {
		data.WriteString(fmt.Sprintf("import ("))
		data.WriteString(separator)

		for _, i := range imports {
			data.WriteString(fmt.Sprintf("\t%s", i))
			data.WriteString(separator)
		}

		data.WriteString(fmt.Sprintf(")"))
		data.WriteString(separator)
		data.WriteString(separator)
	}

	switch len(layers) {
	case 0:
		data.WriteString(fmt.Sprintf("type %s%sInterface any", util.Capitalize(util.SingularForm(name)), util.Capitalize("provider")))
		data.WriteString(separator)
	default:
		data.WriteString(fmt.Sprintf("type %s%sInterface interface {", util.Capitalize(util.SingularForm(name)), util.Capitalize("provider")))
		data.WriteString(separator)

		for _, layer := range layers {
			switch layer {
			case "implementation":
				data.WriteString(fmt.Sprintf("\tGet%s%s() *%s%s.%sImplementation", util.Capitalize(util.SingularForm(name)), util.Capitalize(layer), util.SingularForm(name), util.Capitalize(layer), util.Capitalize(util.SingularForm(name))))
				data.WriteString(separator)
			default:
				data.WriteString(fmt.Sprintf("\tGet%s%s() %s.%s%sInterface", util.Capitalize(util.SingularForm(name)), util.Capitalize(layer), layer, util.Capitalize(util.SingularForm(name)), util.Capitalize(layer)))
				data.WriteString(separator)
			}
		}

		data.WriteString(fmt.Sprintf("}"))
		data.WriteString(separator)
	}

	return data.Bytes()
}

func GetProviderImplementationTemplate(module string, name string, layers []string) []byte {
	data := bytes.Buffer{}
	separator := util.GetSeparator()

	imports := []string{
		fmt.Sprintf("definition \"%s/%s/%s\"", module, "internal", util.Lowercase("provider")),
	}

	for _, layer := range layers {
		switch layer {
		case "implementation":
			imports = append(imports, fmt.Sprintf("%s%s \"%s/%s/implementation/%s\"", util.SingularForm(name), util.Capitalize(layer), module, "internal", name))
		default:
			imports = append(imports, fmt.Sprintf("\"%s/%s/%s\"", module, "internal", layer))
			imports = append(imports, fmt.Sprintf("%s%s \"%s/%s/%s/%s\"", util.SingularForm(name), util.Capitalize(layer), module, "internal", layer, name))
		}
	}

	sort.Strings(imports)

	data.WriteString(fmt.Sprintf("package %s", util.Lowercase(name)))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("import ("))
	data.WriteString(separator)

	for _, i := range imports {
		data.WriteString(fmt.Sprintf("\t%s", i))
		data.WriteString(separator)
	}

	data.WriteString(fmt.Sprintf(")"))
	data.WriteString(separator)
	data.WriteString(separator)

	switch len(layers) {
	case 0:
		data.WriteString(fmt.Sprintf("type %s%s struct {}", util.Capitalize(util.SingularForm(name)), util.Capitalize("provider")))
		data.WriteString(separator)
		data.WriteString(separator)
	default:
		data.WriteString(fmt.Sprintf("type %s%s struct {", util.Capitalize(util.SingularForm(name)), util.Capitalize("provider")))
		data.WriteString(separator)

		for _, layer := range layers {
			switch layer {
			case "implementation":
				data.WriteString(fmt.Sprintf("\t%s%s *%s%s.%s%s", util.SingularForm(name), util.Capitalize(layer), util.SingularForm(name), util.Capitalize(layer), util.Capitalize(util.SingularForm(name)), util.Capitalize(layer)))
				data.WriteString(separator)
			default:
				data.WriteString(fmt.Sprintf("\t%s%s %s.%s%sInterface", util.SingularForm(name), util.Capitalize(layer), layer, util.Capitalize(util.SingularForm(name)), util.Capitalize(layer)))
				data.WriteString(separator)
			}
		}

		data.WriteString(fmt.Sprintf("}"))
		data.WriteString(separator)
		data.WriteString(separator)
	}

	data.WriteString(fmt.Sprintf("var _ definition.%s%sInterface = (*%s%s)(nil)", util.Capitalize(util.SingularForm(name)), util.Capitalize("provider"), util.Capitalize(util.SingularForm(name)), util.Capitalize("provider")))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("func New%s%s() definition.%s%sInterface {", util.Capitalize(util.SingularForm(name)), util.Capitalize("provider"), util.Capitalize(util.SingularForm(name)), util.Capitalize("provider")))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\treturn &%s%s{}", util.Capitalize(util.SingularForm(name)), util.Capitalize("provider")))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)

	for _, l := range layers {
		switch l {
		case "implementation":
			data.WriteString(separator)
			data.WriteString(fmt.Sprintf("func (%s *%s%s) Get%s%s() *%s%s.%sImplementation {", util.FirstLetter("provider"), util.Capitalize(util.SingularForm(name)), util.Capitalize("provider"), util.Capitalize(util.SingularForm(name)), util.Capitalize(l), util.SingularForm(name), util.Capitalize(l), util.Capitalize(util.SingularForm(name))))
			data.WriteString(separator)
			data.WriteString(fmt.Sprintf("\tif %s.%s%s == nil {", util.FirstLetter("provider"), util.SingularForm(name), util.Capitalize(l)))
			data.WriteString(separator)
			data.WriteString(fmt.Sprintf("\t\t%s.%s%s = %s%s.New%s%s()", util.FirstLetter("provider"), util.SingularForm(name), util.Capitalize(l), util.SingularForm(name), util.Capitalize(l), util.Capitalize(util.SingularForm(name)), util.Capitalize(l)))
			data.WriteString(separator)
			data.WriteString(fmt.Sprintf("\t}"))
			data.WriteString(separator)
			data.WriteString(separator)

			data.WriteString(fmt.Sprintf("\treturn %s.%s%s", util.FirstLetter("provider"), util.SingularForm(name), util.Capitalize(l)))
			data.WriteString(separator)
			data.WriteString(fmt.Sprintf("}"))
			data.WriteString(separator)
		default:
			data.WriteString(separator)
			data.WriteString(fmt.Sprintf("func (%s *%s%s) Get%s%s() %s.%s%sInterface {", util.FirstLetter("provider"), util.Capitalize(util.SingularForm(name)), util.Capitalize("provider"), util.Capitalize(util.SingularForm(name)), util.Capitalize(l), l, util.Capitalize(util.SingularForm(name)), util.Capitalize(l)))
			data.WriteString(separator)
			data.WriteString(fmt.Sprintf("\tif %s.%s%s == nil {", util.FirstLetter("provider"), util.SingularForm(name), util.Capitalize(l)))
			data.WriteString(separator)
			data.WriteString(fmt.Sprintf("\t\t%s.%s%s = %s%s.New%s%s()", util.FirstLetter("provider"), util.SingularForm(name), util.Capitalize(l), util.SingularForm(name), util.Capitalize(l), util.Capitalize(util.SingularForm(name)), util.Capitalize(l)))
			data.WriteString(separator)
			data.WriteString(fmt.Sprintf("\t}"))
			data.WriteString(separator)
			data.WriteString(separator)

			data.WriteString(fmt.Sprintf("\treturn %s.%s%s", util.FirstLetter("provider"), util.SingularForm(name), util.Capitalize(l)))
			data.WriteString(separator)
			data.WriteString(fmt.Sprintf("}"))
			data.WriteString(separator)
		}
	}

	return data.Bytes()
}
