package template

import (
	"bytes"
	"fmt"
	"github.com/rafailovalexey/service-generator/internal/dto"
	"github.com/rafailovalexey/service-generator/internal/util"
	"sort"
)

func GetProviderDefinitionTemplate(module string, name *dto.NameDto, layers []string) []byte {
	data := bytes.Buffer{}
	separator := util.GetSeparator()

	imports := make([]string, 0, 10)

	for _, l := range layers {
		switch l {
		case "implementation":
			imports = append(imports, fmt.Sprintf("%s%s \"%s/%s/%s/%s\"", name.LowerCamelCaseSingular, util.GetWithUpperCaseFirstLetter(l), module, "internal", l, name.SnakeCasePlural))
		default:
			imports = append(imports, fmt.Sprintf("\"%s/%s/%s\"", module, "internal", l))
		}
	}

	sort.Strings(imports)

	data.WriteString(fmt.Sprintf("package %s", "provider"))
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
		data.WriteString(fmt.Sprintf("type %s%sInterface any", name.CamelCaseSingular, util.GetWithUpperCaseFirstLetter("provider")))
		data.WriteString(separator)
	default:
		data.WriteString(fmt.Sprintf("type %s%sInterface interface {", name.CamelCaseSingular, util.GetWithUpperCaseFirstLetter("provider")))
		data.WriteString(separator)

		for _, layer := range layers {
			switch layer {
			case "implementation":
				data.WriteString(fmt.Sprintf("\tGet%s%s() *%s%s.%sImplementation", name.CamelCaseSingular, util.GetWithUpperCaseFirstLetter(layer), name.LowerCamelCaseSingular, util.GetWithUpperCaseFirstLetter(layer), name.CamelCaseSingular))
				data.WriteString(separator)
			default:
				data.WriteString(fmt.Sprintf("\tGet%s%s() %s.%s%sInterface", name.CamelCaseSingular, util.GetWithUpperCaseFirstLetter(layer), layer, name.CamelCaseSingular, util.GetWithUpperCaseFirstLetter(layer)))
				data.WriteString(separator)
			}
		}

		data.WriteString(fmt.Sprintf("}"))
		data.WriteString(separator)
	}

	return data.Bytes()
}

func GetProviderImplementationTemplate(module string, name *dto.NameDto, layers []string) []byte {
	data := bytes.Buffer{}
	separator := util.GetSeparator()

	imports := []string{
		fmt.Sprintf("definition \"%s/%s/%s\"", module, "internal", "provider"),
	}

	for _, layer := range layers {
		switch layer {
		case "implementation":
			imports = append(imports, fmt.Sprintf("%s%s \"%s/%s/implementation/%s\"", name.LowerCamelCaseSingular, util.GetWithUpperCaseFirstLetter(layer), module, "internal", name.SnakeCasePlural))
		default:
			imports = append(imports, fmt.Sprintf("\"%s/%s/%s\"", module, "internal", layer))
			imports = append(imports, fmt.Sprintf("%s%s \"%s/%s/%s/%s\"", name.LowerCamelCaseSingular, util.GetWithUpperCaseFirstLetter(layer), module, "internal", layer, name.SnakeCasePlural))
		}
	}

	sort.Strings(imports)

	data.WriteString(fmt.Sprintf("package %s", name.SnakeCasePlural))
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
		data.WriteString(fmt.Sprintf("type %s%s struct {}", name.CamelCaseSingular, util.GetWithUpperCaseFirstLetter("provider")))
		data.WriteString(separator)
		data.WriteString(separator)
	default:
		data.WriteString(fmt.Sprintf("type %s%s struct {", name.CamelCaseSingular, util.GetWithUpperCaseFirstLetter("provider")))
		data.WriteString(separator)

		for _, layer := range layers {
			switch layer {
			case "implementation":
				data.WriteString(fmt.Sprintf("\t%s%s *%s%s.%s%s", name.LowerCamelCaseSingular, util.GetWithUpperCaseFirstLetter(layer), name.LowerCamelCaseSingular, util.GetWithUpperCaseFirstLetter(layer), name.CamelCaseSingular, util.GetWithUpperCaseFirstLetter(layer)))
				data.WriteString(separator)
			default:
				data.WriteString(fmt.Sprintf("\t%s%s %s.%s%sInterface", name.LowerCamelCaseSingular, util.GetWithUpperCaseFirstLetter(layer), layer, name.CamelCaseSingular, util.GetWithUpperCaseFirstLetter(layer)))
				data.WriteString(separator)
			}
		}

		data.WriteString(fmt.Sprintf("}"))
		data.WriteString(separator)
		data.WriteString(separator)
	}

	data.WriteString(fmt.Sprintf("var _ definition.%s%sInterface = (*%s%s)(nil)", name.CamelCaseSingular, util.GetWithUpperCaseFirstLetter("provider"), name.CamelCaseSingular, util.GetWithUpperCaseFirstLetter("provider")))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("func New%s%s() definition.%s%sInterface {", name.CamelCaseSingular, util.GetWithUpperCaseFirstLetter("provider"), name.CamelCaseSingular, util.GetWithUpperCaseFirstLetter("provider")))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\treturn &%s%s{}", name.CamelCaseSingular, util.GetWithUpperCaseFirstLetter("provider")))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)

	for _, l := range layers {
		switch l {
		case "implementation":
			data.WriteString(separator)
			data.WriteString(fmt.Sprintf("func (%s *%s%s) Get%s%s() *%s%s.%sImplementation {", util.GetFirstLetterLowerCase("provider"), name.CamelCaseSingular, util.GetWithUpperCaseFirstLetter("provider"), name.CamelCaseSingular, util.GetWithUpperCaseFirstLetter(l), name.LowerCamelCaseSingular, util.GetWithUpperCaseFirstLetter(l), name.CamelCaseSingular))
			data.WriteString(separator)
			data.WriteString(fmt.Sprintf("\tif %s.%s%s == nil {", util.GetFirstLetterLowerCase("provider"), name.LowerCamelCaseSingular, util.GetWithUpperCaseFirstLetter(l)))
			data.WriteString(separator)
			data.WriteString(fmt.Sprintf("\t\t%s.%s%s = %s%s.New%s%s()", util.GetFirstLetterLowerCase("provider"), name.LowerCamelCaseSingular, util.GetWithUpperCaseFirstLetter(l), name.LowerCamelCaseSingular, util.GetWithUpperCaseFirstLetter(l), name.CamelCaseSingular, util.GetWithUpperCaseFirstLetter(l)))
			data.WriteString(separator)
			data.WriteString(fmt.Sprintf("\t}"))
			data.WriteString(separator)
			data.WriteString(separator)

			data.WriteString(fmt.Sprintf("\treturn %s.%s%s", util.GetFirstLetterLowerCase("provider"), name.LowerCamelCaseSingular, util.GetWithUpperCaseFirstLetter(l)))
			data.WriteString(separator)
			data.WriteString(fmt.Sprintf("}"))
			data.WriteString(separator)
		default:
			data.WriteString(separator)
			data.WriteString(fmt.Sprintf("func (%s *%s%s) Get%s%s() %s.%s%sInterface {", util.GetFirstLetterLowerCase("provider"), name.CamelCaseSingular, util.GetWithUpperCaseFirstLetter("provider"), name.CamelCaseSingular, util.GetWithUpperCaseFirstLetter(l), l, name.CamelCaseSingular, util.GetWithUpperCaseFirstLetter(l)))
			data.WriteString(separator)
			data.WriteString(fmt.Sprintf("\tif %s.%s%s == nil {", util.GetFirstLetterLowerCase("provider"), name.LowerCamelCaseSingular, util.GetWithUpperCaseFirstLetter(l)))
			data.WriteString(separator)
			data.WriteString(fmt.Sprintf("\t\t%s.%s%s = %s%s.New%s%s()", util.GetFirstLetterLowerCase("provider"), name.LowerCamelCaseSingular, util.GetWithUpperCaseFirstLetter(l), name.LowerCamelCaseSingular, util.GetWithUpperCaseFirstLetter(l), name.CamelCaseSingular, util.GetWithUpperCaseFirstLetter(l)))
			data.WriteString(separator)
			data.WriteString(fmt.Sprintf("\t}"))
			data.WriteString(separator)
			data.WriteString(separator)

			data.WriteString(fmt.Sprintf("\treturn %s.%s%s", util.GetFirstLetterLowerCase("provider"), name.LowerCamelCaseSingular, util.GetWithUpperCaseFirstLetter(l)))
			data.WriteString(separator)
			data.WriteString(fmt.Sprintf("}"))
			data.WriteString(separator)
		}
	}

	return data.Bytes()
}
