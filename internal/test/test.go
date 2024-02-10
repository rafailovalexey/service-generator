package test

import (
	"bytes"
	"fmt"
	"github.com/rafailovalexey/service-generator/internal/utils"
	"sort"
)

func GetBaseDefinitionTemplate(layer string, name string) []byte {
	data := bytes.Buffer{}
	separator := utils.GetSeparator()

	data.WriteString(fmt.Sprintf("package %s", utils.Lowercase(layer)))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("type %s%sInterface interface {}", utils.Capitalize(utils.SingularForm(name)), utils.Capitalize(layer)))
	data.WriteString(separator)

	return data.Bytes()
}

func GetBaseImplementationTemplate(module string, layer string, name string) []byte {
	data := bytes.Buffer{}
	separator := utils.GetSeparator()

	imports := []string{
		fmt.Sprintf("definition \"%s/%s/%s\"", module, "internal", utils.Lowercase(layer)),
	}

	sort.Strings(imports)

	data.WriteString(fmt.Sprintf("package %s", utils.Lowercase(name)))
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

	data.WriteString(fmt.Sprintf("type %s%s struct {}", utils.Capitalize(utils.SingularForm(name)), utils.Capitalize(layer)))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("var _ definition.%s%sInterface = (*%s%s)(nil)", utils.Capitalize(utils.SingularForm(name)), utils.Capitalize(layer), utils.Capitalize(utils.SingularForm(name)), utils.Capitalize(layer)))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("func New%s%s() definition.%s%sInterface {", utils.Capitalize(utils.SingularForm(name)), utils.Capitalize(layer), utils.Capitalize(utils.SingularForm(name)), utils.Capitalize(layer)))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\treturn &%s%s{}", utils.Capitalize(utils.SingularForm(name)), utils.Capitalize(layer)))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)

	return data.Bytes()
}

func GetDataTransferObjectTemplate(layer string, name string) []byte {
	data := bytes.Buffer{}
	separator := utils.GetSeparator()

	data.WriteString(fmt.Sprintf("package %s", utils.Lowercase(name)))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("type %s%s struct {}", utils.Capitalize(utils.SingularForm(name)), utils.Capitalize(layer)))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("type %s%s = []%s%s", utils.Capitalize(name), utils.Capitalize(layer), utils.Capitalize(utils.SingularForm(name)), utils.Capitalize(layer)))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("func New%s%s() *%s%s {", utils.Capitalize(utils.SingularForm(name)), utils.Capitalize(layer), utils.Capitalize(utils.SingularForm(name)), utils.Capitalize(layer)))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\treturn &%s%s{}", utils.Capitalize(utils.SingularForm(name)), utils.Capitalize(layer)))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("func New%s%s(%s ...%s%s) *%s%s {", utils.Capitalize(name), utils.Capitalize(layer), utils.Lowercase(name), utils.Capitalize(utils.SingularForm(name)), utils.Capitalize(layer), utils.Capitalize(name), utils.Capitalize(layer)))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t%s := make(%s%s, len(%s))", utils.Lowercase(layer), utils.Capitalize(name), utils.Capitalize(layer), utils.Lowercase(name)))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\tfor _, %s := range %s {", utils.Lowercase(utils.SingularForm(name)), utils.Lowercase(name)))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\t%s = append(%s, %s)", utils.Lowercase(layer), utils.Lowercase(layer), utils.Lowercase(utils.SingularForm(name))))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\treturn &%s", utils.Lowercase(layer)))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)

	return data.Bytes()
}

func GetGrpcServerImplementationTemplate(module string, name string) []byte {
	data := bytes.Buffer{}
	separator := utils.GetSeparator()

	imports := []string{
		fmt.Sprintf("\"%s/pkg/%s_v1\"", module, name),
	}

	sort.Strings(imports)

	data.WriteString(fmt.Sprintf("package %s", utils.Lowercase(name)))
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

	data.WriteString(fmt.Sprintf("type %s%s struct {", utils.Capitalize(utils.SingularForm(name)), utils.Capitalize("implementation")))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t%s_v1.Unimplemented%sV1Server", name, utils.Capitalize(name)))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("func New%s%s() *%s%s {", utils.Capitalize(utils.SingularForm(name)), utils.Capitalize("implementation"), utils.Capitalize(utils.SingularForm(name)), utils.Capitalize("implementation")))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\treturn &%s%s{}", utils.Capitalize(utils.SingularForm(name)), utils.Capitalize("implementation")))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("func (%s *%s%s) mustEmbedUnimplemented%sV1Server() {", utils.FirstLetter(name), utils.Capitalize(utils.SingularForm(name)), utils.Capitalize("implementation"), utils.Capitalize(utils.SingularForm(name))))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\treturn"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)

	return data.Bytes()
}

func GetHttpHandlerDefinitionTemplate(name string) []byte {
	data := bytes.Buffer{}
	separator := utils.GetSeparator()

	imports := []string{
		"\"net/http\"",
	}

	sort.Strings(imports)

	data.WriteString(fmt.Sprintf("package %s", utils.Lowercase("handler")))
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

	data.WriteString(fmt.Sprintf("type %s%sInterface interface {", utils.Capitalize(utils.SingularForm(name)), utils.Capitalize("handler")))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t%sHandle(response http.ResponseWriter, request *http.Request)", utils.Capitalize(utils.SingularForm(name))))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)

	return data.Bytes()
}

func GetHttpHandlerImplementationTemplate(module string, name string) []byte {
	data := bytes.Buffer{}
	separator := utils.GetSeparator()

	imports := []string{
		fmt.Sprintf("\"%s/utils\"", module),
		fmt.Sprintf("\"net/http\""),
		fmt.Sprintf("definition \"%s/%s/%s\"", module, "internal", utils.Lowercase("handler")),
	}

	sort.Strings(imports)

	data.WriteString(fmt.Sprintf("package %s", utils.Lowercase(name)))
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

	data.WriteString(fmt.Sprintf("type %s%s struct {}", utils.Capitalize(utils.SingularForm(name)), utils.Capitalize("handler")))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("var _ definition.%s%sInterface = (*%s%s)(nil)", utils.Capitalize(utils.SingularForm(name)), utils.Capitalize("handler"), utils.Capitalize(utils.SingularForm(name)), utils.Capitalize("handler")))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("func New%s%s() *%s%s {", utils.Capitalize(utils.SingularForm(name)), utils.Capitalize("handler"), utils.Capitalize(utils.SingularForm(name)), utils.Capitalize("handler")))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\treturn &%s%s{}", utils.Capitalize(utils.SingularForm(name)), utils.Capitalize("handler")))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("func (%s *%s%s) %sHandle(response http.ResponseWriter, request *http.Request) {", utils.FirstLetter(name), utils.Capitalize(utils.SingularForm(name)), utils.Capitalize("handler"), utils.Capitalize(utils.SingularForm(name))))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tswitch request.Method {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tdefault:"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\tutils.ResponseMethodNotAllowed(response, request)"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\t\treturn"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t}"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)

	return data.Bytes()
}

func GetProviderDefinitionTemplate(module string, name string, layers []string) []byte {
	data := bytes.Buffer{}
	separator := utils.GetSeparator()

	imports := make([]string, 0, 10)

	for _, l := range layers {
		switch l {
		case "implementation":
			imports = append(imports, fmt.Sprintf("%s%s \"%s/%s/%s/%s\"", utils.SingularForm(name), utils.Capitalize(l), module, "internal", l, name))
		default:
			imports = append(imports, fmt.Sprintf("\"%s/%s/%s\"", module, "internal", l))
		}
	}

	sort.Strings(imports)

	data.WriteString(fmt.Sprintf("package %s", utils.Lowercase("provider")))
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
		data.WriteString(fmt.Sprintf("type %s%sInterface any", utils.Capitalize(utils.SingularForm(name)), utils.Capitalize("provider")))
		data.WriteString(separator)
	default:
		data.WriteString(fmt.Sprintf("type %s%sInterface interface {", utils.Capitalize(utils.SingularForm(name)), utils.Capitalize("provider")))
		data.WriteString(separator)

		for _, layer := range layers {
			switch layer {
			case "implementation":
				data.WriteString(fmt.Sprintf("\tGet%s%s() *%s%s.%sImplementation", utils.Capitalize(utils.SingularForm(name)), utils.Capitalize(layer), utils.SingularForm(name), utils.Capitalize(layer), utils.Capitalize(utils.SingularForm(name))))
				data.WriteString(separator)
			default:
				data.WriteString(fmt.Sprintf("\tGet%s%s() %s.%s%sInterface", utils.Capitalize(utils.SingularForm(name)), utils.Capitalize(layer), layer, utils.Capitalize(utils.SingularForm(name)), utils.Capitalize(layer)))
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
	separator := utils.GetSeparator()

	imports := []string{
		fmt.Sprintf("definition \"%s/%s/%s\"", module, "internal", utils.Lowercase("provider")),
	}

	for _, layer := range layers {
		switch layer {
		case "implementation":
			imports = append(imports, fmt.Sprintf("%s%s \"%s/%s/implementation/%s\"", utils.SingularForm(name), utils.Capitalize(layer), module, "internal", name))
		default:
			imports = append(imports, fmt.Sprintf("\"%s/%s/%s\"", module, "internal", layer))
			imports = append(imports, fmt.Sprintf("%s%s \"%s/%s/%s/%s\"", utils.SingularForm(name), utils.Capitalize(layer), module, "internal", layer, name))
		}
	}

	sort.Strings(imports)

	data.WriteString(fmt.Sprintf("package %s", utils.Lowercase(name)))
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
		data.WriteString(fmt.Sprintf("type %s%s struct {}", utils.Capitalize(utils.SingularForm(name)), utils.Capitalize("provider")))
		data.WriteString(separator)
		data.WriteString(separator)
	default:
		data.WriteString(fmt.Sprintf("type %s%s struct {", utils.Capitalize(utils.SingularForm(name)), utils.Capitalize("provider")))
		data.WriteString(separator)

		for _, layer := range layers {
			switch layer {
			case "implementation":
				data.WriteString(fmt.Sprintf("\t%s%s *%s%s.%s%s", utils.SingularForm(name), utils.Capitalize(layer), utils.SingularForm(name), utils.Capitalize(layer), utils.Capitalize(utils.SingularForm(name)), utils.Capitalize(layer)))
				data.WriteString(separator)
			default:
				data.WriteString(fmt.Sprintf("\t%s%s %s.%s%sInterface", utils.SingularForm(name), utils.Capitalize(layer), layer, utils.Capitalize(utils.SingularForm(name)), utils.Capitalize(layer)))
				data.WriteString(separator)
			}
		}

		data.WriteString(fmt.Sprintf("}"))
		data.WriteString(separator)
		data.WriteString(separator)
	}

	data.WriteString(fmt.Sprintf("var _ definition.%s%sInterface = (*%s%s)(nil)", utils.Capitalize(utils.SingularForm(name)), utils.Capitalize("provider"), utils.Capitalize(utils.SingularForm(name)), utils.Capitalize("provider")))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("func New%s%s() definition.%s%sInterface {", utils.Capitalize(utils.SingularForm(name)), utils.Capitalize("provider"), utils.Capitalize(utils.SingularForm(name)), utils.Capitalize("provider")))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\treturn &%s%s{}", utils.Capitalize(utils.SingularForm(name)), utils.Capitalize("provider")))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)

	for _, l := range layers {
		switch l {
		case "implementation":
			data.WriteString(separator)
			data.WriteString(fmt.Sprintf("func (%s *%s%s) Get%s%s() *%s%s.%sImplementation {", utils.FirstLetter("provider"), utils.Capitalize(utils.SingularForm(name)), utils.Capitalize("provider"), utils.Capitalize(utils.SingularForm(name)), utils.Capitalize(l), utils.SingularForm(name), utils.Capitalize(l), utils.Capitalize(utils.SingularForm(name))))
			data.WriteString(separator)
			data.WriteString(fmt.Sprintf("\tif %s.%s%s == nil {", utils.FirstLetter("provider"), utils.SingularForm(name), utils.Capitalize(l)))
			data.WriteString(separator)
			data.WriteString(fmt.Sprintf("\t\t%s.%s%s = %s%s.New%s%s()", utils.FirstLetter("provider"), utils.SingularForm(name), utils.Capitalize(l), utils.SingularForm(name), utils.Capitalize(l), utils.Capitalize(utils.SingularForm(name)), utils.Capitalize(l)))
			data.WriteString(separator)
			data.WriteString(fmt.Sprintf("\t}"))
			data.WriteString(separator)
			data.WriteString(separator)

			data.WriteString(fmt.Sprintf("\treturn %s.%s%s", utils.FirstLetter("provider"), utils.SingularForm(name), utils.Capitalize(l)))
			data.WriteString(separator)
			data.WriteString(fmt.Sprintf("}"))
			data.WriteString(separator)
		default:
			data.WriteString(separator)
			data.WriteString(fmt.Sprintf("func (%s *%s%s) Get%s%s() %s.%s%sInterface {", utils.FirstLetter("provider"), utils.Capitalize(utils.SingularForm(name)), utils.Capitalize("provider"), utils.Capitalize(utils.SingularForm(name)), utils.Capitalize(l), l, utils.Capitalize(utils.SingularForm(name)), utils.Capitalize(l)))
			data.WriteString(separator)
			data.WriteString(fmt.Sprintf("\tif %s.%s%s == nil {", utils.FirstLetter("provider"), utils.SingularForm(name), utils.Capitalize(l)))
			data.WriteString(separator)
			data.WriteString(fmt.Sprintf("\t\t%s.%s%s = %s%s.New%s%s()", utils.FirstLetter("provider"), utils.SingularForm(name), utils.Capitalize(l), utils.SingularForm(name), utils.Capitalize(l), utils.Capitalize(utils.SingularForm(name)), utils.Capitalize(l)))
			data.WriteString(separator)
			data.WriteString(fmt.Sprintf("\t}"))
			data.WriteString(separator)
			data.WriteString(separator)

			data.WriteString(fmt.Sprintf("\treturn %s.%s%s", utils.FirstLetter("provider"), utils.SingularForm(name), utils.Capitalize(l)))
			data.WriteString(separator)
			data.WriteString(fmt.Sprintf("}"))
			data.WriteString(separator)
		}
	}

	return data.Bytes()
}
