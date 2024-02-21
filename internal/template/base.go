package template

import (
	"bytes"
	"fmt"
	"github.com/rafailovalexey/service-generator/internal/dto"
	"github.com/rafailovalexey/service-generator/internal/util"
	"sort"
)

func GetBaseDefinitionTemplate(layer string, name *dto.NameDto) []byte {
	data := bytes.Buffer{}
	separator := util.GetSeparator()

	data.WriteString(fmt.Sprintf("package %s", layer))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("type %s%sInterface interface {}", name.CamelCaseSingular, util.GetWithUpperCaseFirstLetter(layer)))
	data.WriteString(separator)

	return data.Bytes()
}

func GetBaseImplementationTemplate(module string, layer string, name *dto.NameDto) []byte {
	data := bytes.Buffer{}
	separator := util.GetSeparator()

	imports := []string{
		fmt.Sprintf("definition \"%s/%s/%s\"", module, "internal", layer),
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

	data.WriteString(fmt.Sprintf("type %s%s struct {}", name.CamelCaseSingular, util.GetWithUpperCaseFirstLetter(layer)))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("var _ definition.%s%sInterface = (*%s%s)(nil)", name.CamelCaseSingular, util.GetWithUpperCaseFirstLetter(layer), name.CamelCaseSingular, util.GetWithUpperCaseFirstLetter(layer)))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("func New%s%s() definition.%s%sInterface {", name.CamelCaseSingular, util.GetWithUpperCaseFirstLetter(layer), name.CamelCaseSingular, util.GetWithUpperCaseFirstLetter(layer)))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\treturn &%s%s{}", name.CamelCaseSingular, util.GetWithUpperCaseFirstLetter(layer)))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)

	return data.Bytes()
}
