package template

import (
	"bytes"
	"fmt"
	"github.com/rafailovalexey/service-generator/internal/modules/system"
	"strings"
)

func GetLayerData(layer string, name string) []byte {
	data := bytes.Buffer{}
	separator := system.GetSeparator()

	data.WriteString(fmt.Sprintf("package %s", lowercase(layer)))
	data.WriteString(separator)
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("type %s%sInterface interface {}", capitalize(name), capitalize(layer)))
	data.WriteString(separator)

	return data.Bytes()
}

func GetImplementationData(layer string, name string, module string, kind string) []byte {
	data := bytes.Buffer{}
	separator := system.GetSeparator()

	data.WriteString(fmt.Sprintf("package %s", lowercase(name)))
	data.WriteString(separator)
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("import ("))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tdefinition \"%s/%s/%s\"", module, kind, lowercase(layer)))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf(")"))
	data.WriteString(separator)
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("type %s%s struct{}", capitalize(name), capitalize(layer)))
	data.WriteString(separator)
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("var _ definition.%s%sInterface = (*%s%s)(nil)", capitalize(name), capitalize(layer), capitalize(name), capitalize(layer)))
	data.WriteString(separator)
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("func New%s%s() definition.%s%sInterface {", capitalize(name), capitalize(layer), capitalize(name), capitalize(layer)))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\treturn &%s%s{}", capitalize(name), capitalize(layer)))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)

	return data.Bytes()
}

func GetDtoData(layer string, name string) []byte {
	data := bytes.Buffer{}
	separator := system.GetSeparator()

	data.WriteString(fmt.Sprintf("package %s", lowercase(name)))
	data.WriteString(separator)
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("type %s%s struct {}", capitalize(singular(name)), capitalize(layer)))
	data.WriteString(separator)
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("type %s%s = []%s%s", capitalize(name), capitalize(layer), capitalize(singular(name)), capitalize(layer)))
	data.WriteString(separator)
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("func New%s%s() *%s%s {", capitalize(singular(name)), capitalize(layer), capitalize(singular(name)), capitalize(layer)))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\treturn &%s%s{}", capitalize(singular(name)), capitalize(layer)))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("func New%s%s(%s ...%s%s) *%s%s {", capitalize(name), capitalize(layer), lowercase(name), capitalize(singular(name)), capitalize(layer), capitalize(name), capitalize(layer)))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t%s := make([]%s%s, len(%s))", lowercase(plural(layer)), capitalize(singular(name)), capitalize(layer), lowercase(name)))
	data.WriteString(separator)
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tfor _, %s := range %s {", lowercase(singular(name)), lowercase(name)))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\t%s = append(%s, %s)", lowercase(plural(layer)), lowercase(plural(layer)), lowercase(singular(name))))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t}"))
	data.WriteString(separator)
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\treturn &%s", lowercase(plural(layer))))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)

	return data.Bytes()
}

func capitalize(value string) string {
	if value == "" {
		return value
	}

	return strings.ToUpper(value[:1]) + value[1:]
}

func lowercase(value string) string {
	return strings.ToLower(value)
}

func singular(value string) string {
	if value == "" {
		return value
	}

	return value[:1] + value[1:len(value)-1]
}

func plural(value string) string {
	if value == "" {
		return value
	}

	return value + "s"
}
