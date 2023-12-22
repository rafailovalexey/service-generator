package template

import (
	"bytes"
	"fmt"
	"strings"
)

type Imports = []string

type Function struct {
	Name          string
	ArgumentsType []string
	ArgumentsName []string
	Outputs       []string
	Code          []string
}

type Functions = []Function

type Method struct {
	Name          string
	ArgumentsType []string
	ArgumentsName []string
	Outputs       []string
	Code          []string
}

type Methods = []Method

func GetInterfaceTemplate(separator string, layer string, name string, imports *Imports, methods *Methods) []byte {
	data := bytes.Buffer{}

	if len(*methods) == 0 {
		data.WriteString(fmt.Sprintf("package %s", lowercase(layer)))
		data.WriteString(separator)
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("type %s%sInterface interface {}", capitalize(name), capitalize(layer)))
		data.WriteString(separator)
	}

	if len(*methods) != 0 {
		data.WriteString(fmt.Sprintf("package %s", lowercase(layer)))
		data.WriteString(separator)
		data.WriteString(separator)

		if len(*imports) != 0 {
			data.WriteString(fmt.Sprintf("import ("))
			data.WriteString(separator)

			for _, library := range *imports {
				data.WriteString(fmt.Sprintf("\t\"%s\"", library))
				data.WriteString(separator)
			}

			data.WriteString(fmt.Sprintf(")"))
			data.WriteString(separator)
		}

		data.WriteString(separator)

		data.WriteString(fmt.Sprintf("type %s%sInterface interface {", capitalize(name), capitalize(layer)))
		data.WriteString(separator)

		for _, method := range *methods {
			arguments := strings.Join(method.ArgumentsType, ", ")
			outputs := strings.Join(method.Outputs, ", ")

			if len(method.Outputs) > 1 {
				data.WriteString(fmt.Sprintf("\t%s(%s) (%s)", method.Name, arguments, outputs))
				data.WriteString(separator)
			}

			if len(method.Outputs) == 1 {
				data.WriteString(fmt.Sprintf("\t%s(%s) %s", method.Name, arguments, outputs))
				data.WriteString(separator)
			}

			if len(method.Outputs) == 0 {
				data.WriteString(fmt.Sprintf("\t%s(%s)", method.Name, arguments))
				data.WriteString(separator)
			}
		}

		data.WriteString(fmt.Sprintf("}"))
		data.WriteString(separator)
	}

	return data.Bytes()
}

func GetRealisationInterfaceTemplate(application string, separator string, kind string, layer string, name string, imports *Imports, methods *Methods, functions *Functions) []byte {
	data := bytes.Buffer{}

	data.WriteString(fmt.Sprintf("package %s", lowercase(name)))
	data.WriteString(separator)
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("import ("))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tdefinition \"%s/%s/%s\"", application, kind, lowercase(layer)))
	data.WriteString(separator)

	if len(*imports) != 0 {
		for _, library := range *imports {
			data.WriteString(fmt.Sprintf("\t\"%s\"", library))
			data.WriteString(separator)
		}
	}

	data.WriteString(fmt.Sprintf(")"))
	data.WriteString(separator)
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("type %s%s struct {}", capitalize(name), capitalize(layer)))
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

	if len(*methods) != 0 {
		for _, method := range *methods {
			args := make([]string, len(method.ArgumentsType))

			for index, _ := range method.ArgumentsType {
				args[index] = fmt.Sprintf("%s %s", method.ArgumentsName[index], method.ArgumentsType[index])
			}

			arguments := strings.Join(args, ", ")
			outputs := strings.Join(method.Outputs, ", ")

			data.WriteString(separator)

			if len(method.Outputs) > 1 {
				data.WriteString(fmt.Sprintf("func (%s *%s%s) %s(%s) (%s) {", lowercase(first(name)), capitalize(name), capitalize(layer), method.Name, arguments, outputs))

				data.WriteString(separator)

				for _, code := range method.Code {
					data.WriteString(fmt.Sprintf("\t%s", code))
					data.WriteString(separator)
				}

				data.WriteString(fmt.Sprintf("}"))
			}

			if len(method.Outputs) == 1 {
				data.WriteString(fmt.Sprintf("func (%s *%s%s) %s(%s) %s {", lowercase(first(name)), capitalize(name), capitalize(layer), method.Name, arguments, outputs))

				data.WriteString(separator)

				for _, code := range method.Code {
					data.WriteString(fmt.Sprintf("\t%s", code))
					data.WriteString(separator)
				}

				data.WriteString(fmt.Sprintf("}"))
			}

			if len(method.Outputs) == 0 {
				data.WriteString(fmt.Sprintf("func (%s *%s%s) %s(%s) {", lowercase(first(name)), capitalize(name), capitalize(layer), method.Name, arguments))

				data.WriteString(separator)

				for _, code := range method.Code {
					data.WriteString(fmt.Sprintf("\t%s", code))
					data.WriteString(separator)
				}

				data.WriteString(fmt.Sprintf("}"))
			}

			data.WriteString(separator)
		}
	}

	if len(*functions) != 0 {
		for _, function := range *functions {
			args := make([]string, len(function.ArgumentsType))

			for index, _ := range function.ArgumentsType {
				args[index] = fmt.Sprintf("%s %s", function.ArgumentsName[index], function.ArgumentsType[index])
			}

			arguments := strings.Join(args, ", ")
			outputs := strings.Join(function.Outputs, ", ")

			data.WriteString(separator)

			if len(function.Outputs) > 1 {
				data.WriteString(fmt.Sprintf("func %s(%s) (%s) {", function.Name, arguments, outputs))

				data.WriteString(separator)

				for _, code := range function.Code {
					data.WriteString(fmt.Sprintf("\t%s", code))
					data.WriteString(separator)
				}

				data.WriteString(fmt.Sprintf("}"))
			}

			if len(function.Outputs) == 1 {
				data.WriteString(fmt.Sprintf("func %s(%s) %s {", function.Name, arguments, outputs))

				data.WriteString(separator)

				for _, code := range function.Code {
					data.WriteString(fmt.Sprintf("\t%s", code))
					data.WriteString(separator)
				}

				data.WriteString(fmt.Sprintf("}"))
			}

			if len(function.Outputs) == 0 {
				data.WriteString(fmt.Sprintf("func %s(%s) {", function.Name, arguments))

				data.WriteString(separator)

				for _, code := range function.Code {
					data.WriteString(fmt.Sprintf("\t%s", code))
					data.WriteString(separator)
				}

				data.WriteString(fmt.Sprintf("}"))
			}

			data.WriteString(separator)
		}

	}

	return data.Bytes()
}

func GetDataTransferObjectTemplate(separator string, layer string, name string) []byte {
	data := bytes.Buffer{}

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
	data.WriteString(fmt.Sprintf("\t%s := make([]%s%s, len(%s))", lowercase(layer), capitalize(singular(name)), capitalize(layer), lowercase(name)))
	data.WriteString(separator)
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tfor _, %s := range %s {", lowercase(singular(name)), lowercase(name)))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\t%s = append(%s, %s)", lowercase(layer), lowercase(layer), lowercase(singular(name))))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t}"))
	data.WriteString(separator)
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\treturn &%s", lowercase(layer)))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)

	return data.Bytes()
}

func GetRequestTemplate(separator string, name string) []byte {
	data := bytes.Buffer{}

	data.WriteString(fmt.Sprintf("package %s", lowercase(name)))
	data.WriteString(separator)

	return data.Bytes()
}

func GetResponseTemplate(separator string, name string) []byte {
	data := bytes.Buffer{}

	data.WriteString(fmt.Sprintf("package %s", lowercase(name)))
	data.WriteString(separator)

	return data.Bytes()
}

func GetProviderInterfaceTemplate(application string, separator string, kind string, layers []string, layer string, name string) []byte {
	data := bytes.Buffer{}

	data.WriteString(fmt.Sprintf("package %s", lowercase(layer)))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("import ("))
	data.WriteString(separator)

	for _, l := range layers {
		data.WriteString(fmt.Sprintf("\t\"%s/%s/%s\"", application, kind, l))
		data.WriteString(separator)
	}

	data.WriteString(fmt.Sprintf(")"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("type %s%sInterface interface {", capitalize(name), capitalize(layer)))
	data.WriteString(separator)

	for _, l := range layers {
		data.WriteString(fmt.Sprintf("\tGet%s%s() %s.%s%sInterface", capitalize(name), capitalize(l), l, capitalize(name), capitalize(l)))
		data.WriteString(separator)
	}

	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)

	return data.Bytes()
}

func GetProviderRealisationTemplate(application string, separator string, kind string, layers []string, layer string, name string) []byte {
	data := bytes.Buffer{}

	data.WriteString(fmt.Sprintf("package %s", lowercase(name)))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("import ("))
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\tdefinition \"%s/%s/%s\"", application, kind, lowercase(layer)))
	data.WriteString(separator)

	for _, l := range layers {
		data.WriteString(fmt.Sprintf("\t\"%s/%s/%s\"", application, kind, l))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t%s%s \"%s/%s/%s/%s\"", name, capitalize(l), application, kind, l, name))
		data.WriteString(separator)
	}

	data.WriteString(fmt.Sprintf(")"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("type %s%s struct {", capitalize(name), capitalize(layer)))
	data.WriteString(separator)

	for _, l := range layers {
		data.WriteString(fmt.Sprintf("\t%s%s %s.%s%sInterface", name, capitalize(l), l, capitalize(name), capitalize(l)))
		data.WriteString(separator)
	}

	data.WriteString(fmt.Sprintf("}"))
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
	data.WriteString(separator)

	for _, l := range layers {
		data.WriteString(fmt.Sprintf("func (%s *%s%s) Get%s%s() %s.%s%sInterface {", first(layer), capitalize(name), capitalize(layer), capitalize(name), capitalize(l), l, capitalize(name), capitalize(l)))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\tif %s.%s%s == nil {", first(layer), name, capitalize(l)))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t\t%s.%s%s = %s%s.New%s%s()", first(layer), name, capitalize(l), name, capitalize(l), capitalize(name), capitalize(l)))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t}"))
		data.WriteString(separator)
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\treturn %s.%s%s", first(layer), name, capitalize(l)))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("}"))
		data.WriteString(separator)
	}

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

func first(value string) string {
	return value[:1]
}
