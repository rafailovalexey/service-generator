package template

import (
	"bytes"
	"fmt"
	"github.com/rafailovalexey/service-generator/internal/utils"
	"sort"
)

func GetInterfaceTemplate(layer string, name string) []byte {
	data := bytes.Buffer{}
	separator := utils.GetSeparator()

	data.WriteString(fmt.Sprintf("package %s", utils.Lowercase(layer)))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("type %s%sInterface interface {}", utils.Capitalize(utils.SingularForm(name)), utils.Capitalize(layer)))
	data.WriteString(separator)

	return data.Bytes()
}

func GetRealisationInterfaceTemplate(module string, kind string, layer string, name string) []byte {
	data := bytes.Buffer{}
	separator := utils.GetSeparator()

	imports := []string{
		fmt.Sprintf("definition \"%s/%s/%s\"", module, kind, utils.Lowercase(layer)),
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
	data.WriteString(fmt.Sprintf("\t%s := make([]%s%s, len(%s))", utils.Lowercase(layer), utils.Capitalize(utils.SingularForm(name)), utils.Capitalize(layer), utils.Lowercase(name)))
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

func GetRequestTemplate(name string) []byte {
	data := bytes.Buffer{}
	separator := utils.GetSeparator()

	data.WriteString(fmt.Sprintf("package %s", utils.Lowercase(name)))
	data.WriteString(separator)

	return data.Bytes()
}

func GetResponseTemplate(name string) []byte {
	data := bytes.Buffer{}
	separator := utils.GetSeparator()

	data.WriteString(fmt.Sprintf("package %s", utils.Lowercase(name)))
	data.WriteString(separator)

	return data.Bytes()
}

func GetProviderInterfaceTemplate(module string, kind string, layer string, name string, layers []string) []byte {
	data := bytes.Buffer{}
	separator := utils.GetSeparator()

	imports := make([]string, 0, 10)

	for _, l := range layers {
		switch l {
		case "implementation":
			imports = append(imports, fmt.Sprintf("%s%s \"%s/%s/%s/%s\"", utils.SingularForm(name), utils.Capitalize(l), module, kind, l, name))
		default:
			imports = append(imports, fmt.Sprintf("\"%s/%s/%s\"", module, kind, l))
		}
	}

	sort.Strings(imports)

	data.WriteString(fmt.Sprintf("package %s", utils.Lowercase(layer)))
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
		data.WriteString(fmt.Sprintf("type %s%sInterface any", utils.Capitalize(utils.SingularForm(name)), utils.Capitalize(layer)))
		data.WriteString(separator)
	default:
		data.WriteString(fmt.Sprintf("type %s%sInterface interface {", utils.Capitalize(utils.SingularForm(name)), utils.Capitalize(layer)))
		data.WriteString(separator)

		for _, l := range layers {
			switch l {
			case "implementation":
				data.WriteString(fmt.Sprintf("\tGet%s%s() *%s%s.%sImplementation", utils.Capitalize(utils.SingularForm(name)), utils.Capitalize(l), utils.SingularForm(name), utils.Capitalize(l), utils.Capitalize(utils.SingularForm(name))))
				data.WriteString(separator)
			default:
				data.WriteString(fmt.Sprintf("\tGet%s%s() %s.%s%sInterface", utils.Capitalize(utils.SingularForm(name)), utils.Capitalize(l), l, utils.Capitalize(utils.SingularForm(name)), utils.Capitalize(l)))
				data.WriteString(separator)
			}
		}

		data.WriteString(fmt.Sprintf("}"))
		data.WriteString(separator)
	}

	return data.Bytes()
}

func GetProviderRealisationTemplate(module string, kind string, layer string, name string, layers []string) []byte {
	data := bytes.Buffer{}
	separator := utils.GetSeparator()

	imports := []string{
		fmt.Sprintf("definition \"%s/%s/%s\"", module, kind, utils.Lowercase(layer)),
	}

	for _, l := range layers {
		switch l {
		case "implementation":
			imports = append(imports, fmt.Sprintf("%s%s \"%s/%s/implementation/%s\"", utils.SingularForm(name), utils.Capitalize(l), module, kind, name))
		default:
			imports = append(imports, fmt.Sprintf("\"%s/%s/%s\"", module, kind, l))
			imports = append(imports, fmt.Sprintf("%s%s \"%s/%s/%s/%s\"", utils.SingularForm(name), utils.Capitalize(l), module, kind, l, name))
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
		data.WriteString(fmt.Sprintf("type %s%s struct {}", utils.Capitalize(utils.SingularForm(name)), utils.Capitalize(layer)))
		data.WriteString(separator)
		data.WriteString(separator)
	default:
		data.WriteString(fmt.Sprintf("type %s%s struct {", utils.Capitalize(utils.SingularForm(name)), utils.Capitalize(layer)))
		data.WriteString(separator)

		for _, l := range layers {
			switch l {
			case "implementation":
				data.WriteString(fmt.Sprintf("\t%s%s *%s%s.%s%s", utils.SingularForm(name), utils.Capitalize(l), utils.SingularForm(name), utils.Capitalize(l), utils.Capitalize(utils.SingularForm(name)), utils.Capitalize(l)))
				data.WriteString(separator)
			default:
				data.WriteString(fmt.Sprintf("\t%s%s %s.%s%sInterface", utils.SingularForm(name), utils.Capitalize(l), l, utils.Capitalize(utils.SingularForm(name)), utils.Capitalize(l)))
				data.WriteString(separator)
			}
		}

		data.WriteString(fmt.Sprintf("}"))
		data.WriteString(separator)
		data.WriteString(separator)
	}

	data.WriteString(fmt.Sprintf("var _ definition.%s%sInterface = (*%s%s)(nil)", utils.Capitalize(utils.SingularForm(name)), utils.Capitalize(layer), utils.Capitalize(utils.SingularForm(name)), utils.Capitalize(layer)))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("func New%s%s() definition.%s%sInterface {", utils.Capitalize(utils.SingularForm(name)), utils.Capitalize(layer), utils.Capitalize(utils.SingularForm(name)), utils.Capitalize(layer)))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\treturn &%s%s{}", utils.Capitalize(utils.SingularForm(name)), utils.Capitalize(layer)))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)

	for _, l := range layers {
		switch l {
		case "implementation":
			data.WriteString(separator)
			data.WriteString(fmt.Sprintf("func (%s *%s%s) Get%s%s() *%s%s.%sImplementation {", utils.FirstLetter(layer), utils.Capitalize(utils.SingularForm(name)), utils.Capitalize(layer), utils.Capitalize(utils.SingularForm(name)), utils.Capitalize(l), utils.SingularForm(name), utils.Capitalize(l), utils.Capitalize(utils.SingularForm(name))))
			data.WriteString(separator)
			data.WriteString(fmt.Sprintf("\tif %s.%s%s == nil {", utils.FirstLetter(layer), utils.SingularForm(name), utils.Capitalize(l)))
			data.WriteString(separator)
			data.WriteString(fmt.Sprintf("\t\t%s.%s%s = %s%s.New%s%s()", utils.FirstLetter(layer), utils.SingularForm(name), utils.Capitalize(l), utils.SingularForm(name), utils.Capitalize(l), utils.Capitalize(utils.SingularForm(name)), utils.Capitalize(l)))
			data.WriteString(separator)
			data.WriteString(fmt.Sprintf("\t}"))
			data.WriteString(separator)
			data.WriteString(separator)

			data.WriteString(fmt.Sprintf("\treturn %s.%s%s", utils.FirstLetter(layer), utils.SingularForm(name), utils.Capitalize(l)))
			data.WriteString(separator)
			data.WriteString(fmt.Sprintf("}"))
			data.WriteString(separator)
		default:
			data.WriteString(separator)
			data.WriteString(fmt.Sprintf("func (%s *%s%s) Get%s%s() %s.%s%sInterface {", utils.FirstLetter(layer), utils.Capitalize(utils.SingularForm(name)), utils.Capitalize(layer), utils.Capitalize(utils.SingularForm(name)), utils.Capitalize(l), l, utils.Capitalize(utils.SingularForm(name)), utils.Capitalize(l)))
			data.WriteString(separator)
			data.WriteString(fmt.Sprintf("\tif %s.%s%s == nil {", utils.FirstLetter(layer), utils.SingularForm(name), utils.Capitalize(l)))
			data.WriteString(separator)
			data.WriteString(fmt.Sprintf("\t\t%s.%s%s = %s%s.New%s%s()", utils.FirstLetter(layer), utils.SingularForm(name), utils.Capitalize(l), utils.SingularForm(name), utils.Capitalize(l), utils.Capitalize(utils.SingularForm(name)), utils.Capitalize(l)))
			data.WriteString(separator)
			data.WriteString(fmt.Sprintf("\t}"))
			data.WriteString(separator)
			data.WriteString(separator)

			data.WriteString(fmt.Sprintf("\treturn %s.%s%s", utils.FirstLetter(layer), utils.SingularForm(name), utils.Capitalize(l)))
			data.WriteString(separator)
			data.WriteString(fmt.Sprintf("}"))
			data.WriteString(separator)
		}
	}

	return data.Bytes()
}

func GetImplementationRealisationTemplate(module string, layer string, name string) []byte {
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

	data.WriteString(fmt.Sprintf("type %s%s struct {", utils.Capitalize(utils.SingularForm(name)), utils.Capitalize(layer)))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t%s_v1.Unimplemented%sV1Server", name, utils.Capitalize(name)))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("func New%s%s() *%s%s {", utils.Capitalize(utils.SingularForm(name)), utils.Capitalize(layer), utils.Capitalize(utils.SingularForm(name)), utils.Capitalize(layer)))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\treturn &%s%s{}", utils.Capitalize(utils.SingularForm(name)), utils.Capitalize(layer)))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("func (%s *%s%s) mustEmbedUnimplemented%sV1Server() {", utils.FirstLetter(name), utils.Capitalize(utils.SingularForm(name)), utils.Capitalize(layer), utils.Capitalize(utils.SingularForm(name))))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\treturn"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)

	return data.Bytes()
}

func GetHandlerInterfaceTemplate(layer string, name string) []byte {
	data := bytes.Buffer{}
	separator := utils.GetSeparator()

	imports := []string{
		"\"net/http\"",
	}

	sort.Strings(imports)

	data.WriteString(fmt.Sprintf("package %s", utils.Lowercase(layer)))
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

	data.WriteString(fmt.Sprintf("type %s%sInterface interface {", utils.Capitalize(utils.SingularForm(name)), utils.Capitalize(layer)))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t%sHandle(response http.ResponseWriter, request *http.Request)", utils.Capitalize(utils.SingularForm(name))))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)

	return data.Bytes()
}

func GetHandlerRealisationTemplate(module string, kind string, layer string, name string) []byte {
	data := bytes.Buffer{}
	separator := utils.GetSeparator()

	imports := []string{
		fmt.Sprintf("\"%s/utils\"", module),
		fmt.Sprintf("\"net/http\""),
		fmt.Sprintf("definition \"%s/%s/%s\"", module, kind, utils.Lowercase(layer)),
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

	data.WriteString(fmt.Sprintf("func New%s%s() *%s%s {", utils.Capitalize(utils.SingularForm(name)), utils.Capitalize(layer), utils.Capitalize(utils.SingularForm(name)), utils.Capitalize(layer)))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\treturn &%s%s{}", utils.Capitalize(utils.SingularForm(name)), utils.Capitalize(layer)))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("func (%s *%s%s) %sHandle(response http.ResponseWriter, request *http.Request) {", utils.FirstLetter(name), utils.Capitalize(utils.SingularForm(name)), utils.Capitalize(layer), utils.Capitalize(utils.SingularForm(name))))
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

func GetReadmeTemplate() []byte {
	data := bytes.Buffer{}

	return data.Bytes()
}

func GetGitIgnoreTemplate() []byte {
	data := bytes.Buffer{}
	separator := utils.GetSeparator()

	data.WriteString(fmt.Sprintf("# JetBrains"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf(".idea"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("# Build"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("build"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("# Environment"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf(".env"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("# Mocks"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("*_mock.go"))
	data.WriteString(separator)

	return data.Bytes()
}

func GetEnvironmentTemplate() []byte {
	data := bytes.Buffer{}

	return data.Bytes()
}

func GetExampleEnvironmentTemplate(application string) []byte {
	data := bytes.Buffer{}
	separator := utils.GetSeparator()

	switch application {
	case "application":
		data.WriteString(fmt.Sprintf("HOSTNAME = \"hostname\""))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("PORT = \"port\""))
		data.WriteString(separator)
		data.WriteString(separator)

		data.WriteString(fmt.Sprintf("AUTHENTICATION_TOKEN_HEADER = \"header\""))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("AUTHENTICATION_TOKEN = \"token\""))
		data.WriteString(separator)
	}

	return data.Bytes()
}

func GetGrpcMicroserviceMakefileTemplate(name string) []byte {
	data := bytes.Buffer{}
	separator := utils.GetSeparator()

	data.WriteString(fmt.Sprintf("# Variables"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("PROTO_SOURCE_DIRECTORY = api"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("PROTO_OUTPUT_DIRECTORY = pkg"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("PROTO_FILES = \\"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t%s_v1/%s.proto", name, name))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("MOCKS_OUTPUT_DIRECTORY = mocks"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("MOCKS_FILES = \\"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("# GRPC"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("grpc-generate:"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t@echo \"Generating GRPC...\""))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t@bin/grpc-generate.sh $(PROTO_SOURCE_DIRECTORY) $(PROTO_OUTPUT_DIRECTORY) $(PROTO_FILES)"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("# Mocks"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("mock-generate:"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t@echo \"Generating Mocks...\""))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t@bin/mock-generate.sh $(MOCKS_OUTPUT_DIRECTORY) $(MOCKS_FILES)"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("# Generate"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("generate: grpc-generate mock-generate"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("# Download"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("download:"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t@echo \"Downloading dependencies...\""))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t@go mod download"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("# Build"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("build:"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t@echo \"Building...\""))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t@go build -o build/main main.go"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("# Tidy"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("tidy:"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t@echo \"Tidy...\""))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t@go mod tidy"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("# Tests"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("tests:"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t@echo \"Running tests...\""))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t@go test -v ./..."))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf(".PHONY: grpc-generate, mock-generate, generate, download, build, tests"))
	data.WriteString(separator)

	return data.Bytes()
}

func GetDefaultMicroserviceMakefileTemplate() []byte {
	data := bytes.Buffer{}
	separator := utils.GetSeparator()

	data.WriteString(fmt.Sprintf("# Variables"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("MOCKS_OUTPUT_DIRECTORY = mocks"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("MOCKS_FILES = \\"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("# Mocks"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("mock-generate:"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t@echo \"Generating Mocks...\""))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t@bin/mock-generate.sh $(MOCKS_OUTPUT_DIRECTORY) $(MOCKS_FILES)"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("# Generate"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("generate: mock-generate"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("# Download"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("download:"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t@echo \"Downloading dependencies...\""))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t@go mod download"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("# Build"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("build:"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t@echo \"Building...\""))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t@go build -o build/main main.go"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("# Tidy"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("tidy:"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t@echo \"Tidy...\""))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t@go mod tidy"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("# Tests"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("tests:"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t@echo \"Running tests...\""))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t@go test -v ./..."))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf(".PHONY: mock-generate, generate, download, build, tests"))
	data.WriteString(separator)

	return data.Bytes()
}

func GetDockerIgnoreTemplate() []byte {
	data := bytes.Buffer{}
	separator := utils.GetSeparator()

	data.WriteString(fmt.Sprintf("# JetBrains"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf(".idea"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("# Build"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("build"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("# Mocks"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("*_mock.go"))
	data.WriteString(separator)

	return data.Bytes()
}

func GetDockerWithPortTemplate() []byte {
	data := bytes.Buffer{}
	separator := utils.GetSeparator()

	data.WriteString(fmt.Sprintf("FROM golang:latest"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("WORKDIR /usr/local/application"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("COPY . ."))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("RUN apt-get update --yes"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("RUN apt-get upgrade --yes"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("RUN apt-get install --yes make"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("RUN export PATH=\"$PATH:$(go env GOPATH)/bin\""))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("RUN make download"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("RUN make build"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("EXPOSE 3000"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("CMD [\"./build/main\"]"))
	data.WriteString(separator)

	return data.Bytes()
}

func GetDockerWithoutPortTemplate() []byte {
	data := bytes.Buffer{}
	separator := utils.GetSeparator()

	data.WriteString(fmt.Sprintf("FROM golang:latest"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("WORKDIR /usr/local/application"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("COPY . ."))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("RUN apt-get update --yes"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("RUN apt-get upgrade --yes"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("RUN apt-get install --yes make"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("RUN export PATH=\"$PATH:$(go env GOPATH)/bin\""))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("RUN make download"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("RUN make build"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("CMD [\"./build/main\"]"))
	data.WriteString(separator)

	return data.Bytes()
}

func GetGrpcGenerateShellScriptTemplate() []byte {
	data := bytes.Buffer{}
	separator := utils.GetSeparator()

	data.WriteString(fmt.Sprintf("#!/bin/bash"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("if [ \"$#\" -lt 2 ]; then"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\techo \"Usage: $0 <PROTO_SOURCE_DIRECTORY> <PROTO_OUTPUT_DIRECTORY> <PROTO_FILES>\""))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\texit 1"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("fi"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("PROTO_SOURCE_DIRECTORY=\"$1\""))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("PROTO_OUTPUT_DIRECTORY=\"$2\""))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("PROTO_FILES=\"${*:3}\""))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("for proto_file in $PROTO_FILES; do"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tPROTO_FILE_DIRECTORY=$(dirname \"$proto_file\")"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\tmkdir -p \"$PROTO_OUTPUT_DIRECTORY/$PROTO_FILE_DIRECTORY\""))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\techo \"Generating proto file for $proto_file...\""))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\tprotoc \\"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\t--proto_path=\"$PROTO_SOURCE_DIRECTORY\" \\"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\t--go_out=\"$PROTO_OUTPUT_DIRECTORY\" \\"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\t--go_opt=paths=source_relative \"$proto_file\" \\"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\t--go-grpc_out=\"$PROTO_OUTPUT_DIRECTORY\" \\"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\t--go-grpc_opt=paths=source_relative \"$proto_file\""))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("done"))
	data.WriteString(separator)

	return data.Bytes()
}

func GetMockGenerateShellScriptTemplate() []byte {
	data := bytes.Buffer{}
	separator := utils.GetSeparator()

	data.WriteString(fmt.Sprintf("#!/bin/bash"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("if [ \"$#\" -lt 1 ]; then"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\techo \"Usage: $0 <MOCKS_OUTPUT_DIRECTORY> <MOCKS_FILES>\""))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\texit 1"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("fi"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("MOCKS_OUTPUT_DIRECTORY=\"$1\""))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("MOCKS_FILES=\"${*:2}\""))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("for mock_file in $MOCKS_FILES; do"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tDIRECTORY=$(dirname \"$mock_file\")"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tFILENAME=$(basename \"$mock_file\")"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\tEXTENSION=\"${FILENAME##*.}\""))
	data.WriteString(separator)
	data.WriteString(fmt.Sprint("\tFILENAME_WITHOUT_EXTENSIONS=\"${FILENAME%.*}\""))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tOUTPUT_PATH=\"$DIRECTORY/$MOCKS_OUTPUT_DIRECTORY/${FILENAME_WITHOUT_EXTENSIONS}_mock.$EXTENSION\""))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\tmkdir -p \"$DIRECTORY/$MOCKS_OUTPUT_DIRECTORY\""))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\techo \"Generating mock file for $mock_file...\""))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\tmockgen -source=\"$mock_file\" -destination=\"$OUTPUT_PATH\""))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("done"))
	data.WriteString(separator)

	return data.Bytes()
}

func GetGrpcLoggingInterceptorTemplate() []byte {
	data := bytes.Buffer{}
	separator := utils.GetSeparator()

	imports := []string{
		"\"context\"",
		"\"google.golang.org/grpc\"",
		"\"google.golang.org/grpc/codes\"",
		"\"google.golang.org/grpc/metadata\"",
		"\"google.golang.org/grpc/status\"",
		"\"log\"",
	}

	sort.Strings(imports)

	data.WriteString(fmt.Sprintf("package interceptor"))
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

	data.WriteString(fmt.Sprintf("func LoggingInterceptor() grpc.UnaryServerInterceptor {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\treturn func(ctx context.Context, request any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\tmd, isExist := metadata.FromIncomingContext(ctx)"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\t\tif !isExist {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\t\treturn nil, status.Errorf(codes.Internal, \"failed to read metadata\")"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\t}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\t\ttrace := md[\"trace\"][0]"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprint("\t\tlog.Printf(\"incoming grpc request: %s (%s)\", info.FullMethod, trace)"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\t\tresponse, err := handler(ctx, request)"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\t\tif err != nil {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprint("\t\t\tlog.Printf(\"error in grpc request %s (%s) \\n %v\", info.FullMethod, trace, err)"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\t}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\t\tif err == nil {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprint("\t\t\tlog.Printf(\"outgoing grpc response %s (%s)\", info.FullMethod, trace)"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\t}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\t\treturn response, err"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t}"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)

	return data.Bytes()
}

func GetGrpcTracingInterceptorTemplate() []byte {
	data := bytes.Buffer{}
	separator := utils.GetSeparator()

	imports := []string{
		"\"context\"",
		"\"crypto/rand\"",
		"\"encoding/hex\"",
		"\"google.golang.org/grpc\"",
		"\"google.golang.org/grpc/codes\"",
		"\"google.golang.org/grpc/metadata\"",
		"\"google.golang.org/grpc/status\"",
		"\"log\"",
	}

	sort.Strings(imports)

	data.WriteString(fmt.Sprintf("package interceptor"))
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

	data.WriteString(fmt.Sprintf("func TracingInterceptor() grpc.UnaryServerInterceptor {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\treturn func(ctx context.Context, request interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\tmd, isExist := metadata.FromIncomingContext(ctx)"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\t\tif !isExist {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\t\tlog.Printf(\"metadata not found in the request context\\n\")"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\t\t\treturn nil, status.Errorf(codes.Internal, \"failed to read metadata\")"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\t}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\t\tif len(md[\"trace\"]) != 0 {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\t\treturn handler(ctx, request)"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\t}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\t\ttrace, err := GenerateTrace()"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\t\tif err != nil {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\t\treturn nil, status.Errorf(codes.Internal, \"failed to generate trace\")"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\t}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\t\tmd = metadata.Join(md, metadata.New(map[string]string{\"trace\": trace}))"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\tctx = metadata.NewIncomingContext(ctx, md)"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\t\treturn handler(ctx, request)"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t}"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("func GenerateTrace() (string, error) {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\ttrace := make([]byte, 16)"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\tif _, err := rand.Read(trace); err != nil {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\treturn \"\", err"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\treturn hex.EncodeToString(trace), nil"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)

	return data.Bytes()
}

func GetGrpcAuthenticationMiddlewareTemplate() []byte {
	data := bytes.Buffer{}
	separator := utils.GetSeparator()

	imports := []string{
		"\"context\"",
		"\"google.golang.org/grpc\"",
		"\"google.golang.org/grpc/codes\"",
		"\"google.golang.org/grpc/metadata\"",
		"\"google.golang.org/grpc/status\"",
		"\"log\"",
		"\"os\"",
	}

	sort.Strings(imports)

	data.WriteString(fmt.Sprintf("package middleware"))
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

	data.WriteString(fmt.Sprintf("func AuthenticationMiddleware() grpc.UnaryServerInterceptor {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\treturn func(ctx context.Context, request interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\tmd, isExist := metadata.FromIncomingContext(ctx)"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\t\tif !isExist {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\t\treturn nil, status.Errorf(codes.Unauthenticated, \"authentication token not found\")"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\t}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\t\theader := os.Getenv(\"AUTHENTICATION_TOKEN_HEADER\")"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\t\tif header == \"\" {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\t\tlog.Panicf(\"not found authentication token header in environment\")"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\t}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\t\tlist := md[header]"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\t\tif len(list) == 0 {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\t\treturn nil, status.Errorf(codes.Unauthenticated, \"authentication token not found\")"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\t}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\t\tkey := list[0]"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\t\ttoken := os.Getenv(\"AUTHENTICATION_TOKEN\")"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\t\tif token == \"\" {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\t\tlog.Panicf(\"not found authentication token in environment\")"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\t}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\t\tif token != key {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprint("\t\t\tlog.Printf(\"invalid authentication token: %s\", key)"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\t\t\treturn nil, status.Errorf(codes.PermissionDenied, \"invalid authentication token\")"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\t}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\t\treturn handler(ctx, request)"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t}"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)

	return data.Bytes()
}

func GetGrpcServerTemplate(module string, name string) []byte {
	data := bytes.Buffer{}
	separator := utils.GetSeparator()

	imports := []string{
		"\"fmt\"",
		fmt.Sprintf("\"%s/cmd/grpc_server/interceptor\"", module),
		fmt.Sprintf("\"%s/cmd/grpc_server/middleware\"", module),
		fmt.Sprintf("\"%s/pkg/%s_v1\"", module, name),
		"\"google.golang.org/grpc\"",
		"\"google.golang.org/grpc/reflection\"",
		"\"log\"",
		"\"net\"",
		"\"os\"",
	}

	sort.Strings(imports)

	data.WriteString(fmt.Sprintf("package grpc_server"))
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

	data.WriteString(fmt.Sprintf("func Run(api %s_v1.%sV1Server) {", name, utils.Capitalize(name)))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\thostname := os.Getenv(\"HOSTNAME\")"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tport := os.Getenv(\"PORT\")"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\tif port == \"\" {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\tlog.Panicf(\"specify the port\")"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprint("\taddress := fmt.Sprintf(\"%s:%s\", hostname, port)"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprint("\tlog.Printf(\"%s\\n\", fmt.Sprintf(\"grpc server starts at address %s\", address))"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\tlistener, err := net.Listen(\"tcp\", address)"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\tif err != nil {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprint("\t\tlog.Panicf(\"grpc server startup error %v\", err)"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\tserver := grpc.NewServer("))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\tgrpc.ChainUnaryInterceptor("))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\t\tinterceptor.TracingInterceptor(),"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\t\tinterceptor.LoggingInterceptor(),"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\t\tmiddleware.AuthenticationMiddleware(),"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\t),"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t)"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\treflection.Register(server)"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\t%s_v1.Register%sV1Server(server, api)", name, utils.Capitalize(name)))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprint("\tlog.Printf(\"%s\\n\", fmt.Sprintf(\"grpc server is running at %s\", address))"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\terr = server.Serve(listener)"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\tif err != nil {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprint("\t\tlog.Panicf(\"grpc server startup error %v\", err)"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t}"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)

	return data.Bytes()
}

func GetHttpLoggingInterceptorTemplate() []byte {
	data := bytes.Buffer{}
	separator := utils.GetSeparator()

	data.WriteString(fmt.Sprintf("package interceptor"))
	data.WriteString(separator)
	data.WriteString(separator)

	imports := []string{
		"\"log\"",
		"\"net/http\"",
		"\"time\"",
	}

	sort.Strings(imports)

	data.WriteString(fmt.Sprintf("import ("))
	data.WriteString(separator)

	for _, i := range imports {
		data.WriteString(fmt.Sprintf("\t%s", i))
		data.WriteString(separator)
	}

	data.WriteString(fmt.Sprintf(")"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("func LoggingInterceptor(next http.Handler) http.Handler {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\treturn http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\tstart := time.Now()"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\t\tnext.ServeHTTP(response, request)"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\t\tduration := time.Since(start)"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprint("\t\tlog.Printf(\"%s %s %s - %s %v\\n\", request.Method, request.URL.Path, request.RemoteAddr, request.UserAgent(), duration)"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t})"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)

	return data.Bytes()
}

func GetHttpAuthenticationMiddlewareTemplate(module string) []byte {
	data := bytes.Buffer{}
	separator := utils.GetSeparator()

	imports := []string{
		fmt.Sprintf("\"%s/utils\"", module),
		"\"log\"",
		"\"net/http\"",
		"\"os\"",
	}

	sort.Strings(imports)

	data.WriteString(fmt.Sprintf("package middleware"))
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

	data.WriteString(fmt.Sprintf("func AuthenticationMiddleware(next http.Handler) http.Handler {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\treturn http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\theader := os.Getenv(\"AUTHENTICATION_TOKEN_HEADER\")"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\t\tif header == \"\" {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\t\tlog.Panicf(\"specify the name of the authentication token\")"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\t}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\t\ttoken := os.Getenv(\"AUTHENTICATION_TOKEN\")"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\t\tif token == \"\" {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\t\tlog.Panicf(\"specify the value of the authentication token\")"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\t}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\t\tkey := request.Header.Get(header)"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\t\tif key != token {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\t\tresponse.Header().Set(\"Content-Type\", \"application/json\")"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\t\tresponse.WriteHeader(http.StatusUnauthorized)"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\t\tresponse.Write(utils.ConvertError(\"unauthorized\"))"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\t\t\treturn"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\t}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\t\tnext.ServeHTTP(response, request)"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t})"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)

	return data.Bytes()
}

func GetHttpCorsMiddlewareTemplate() []byte {
	data := bytes.Buffer{}
	separator := utils.GetSeparator()

	imports := []string{
		"\"net/http\"",
	}

	sort.Strings(imports)

	data.WriteString(fmt.Sprintf("package middleware"))
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

	data.WriteString(fmt.Sprintf("func CorsMiddleware(next http.Handler) http.Handler {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\treturn http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\tresponse.Header().Add(\"Access-Control-Allow-Origin\", \"*\")"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\tresponse.Header().Add(\"Access-Control-Allow-Headers\", \"*\")"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\tresponse.Header().Add(\"Access-Control-Allow-Methods\", \"*\")"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\tresponse.Header().Add(\"Access-Control-Allow-Credentials\", \"true\")"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\t\tif request.Method == \"OPTIONS\" {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\t\tresponse.WriteHeader(http.StatusOK)"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\t\t\treturn"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\t}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\t\tnext.ServeHTTP(response, request)"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t})"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)

	return data.Bytes()
}

func GetHttpChainMiddlewareTemplate() []byte {
	data := bytes.Buffer{}
	separator := utils.GetSeparator()

	imports := []string{
		"\"net/http\"",
	}

	sort.Strings(imports)

	data.WriteString(fmt.Sprintf("package middleware"))
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

	data.WriteString(fmt.Sprintf("func ChainMiddleware(middlewares ...func(http.Handler) http.Handler) func(http.Handler) http.Handler {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\treturn func(next http.Handler) http.Handler {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\tfor index := len(middlewares) - 1; index >= 0; index-- {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\t\tnext = middlewares[index](next)"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\t}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\t\treturn next"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t}"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)

	return data.Bytes()
}

func GetHttpServerTemplate(module string, name string) []byte {
	data := bytes.Buffer{}
	separator := utils.GetSeparator()

	imports := []string{
		"\"fmt\"",
		"\"log\"",
		"\"net/http\"",
		"\"os\"",
		fmt.Sprintf("\"%s/internal/handler\"", module),
		fmt.Sprintf("\"%s/cmd/http_server/interceptor\"", module),
		fmt.Sprintf("\"%s/cmd/http_server/middleware\"", module),
		fmt.Sprintf("\"%s/utils\"", module),
	}

	sort.Strings(imports)

	data.WriteString(fmt.Sprintf("package http_server"))
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

	data.WriteString(fmt.Sprintf("func Run(%sHandler handler.%sHandlerInterface) {", utils.SingularForm(name), utils.Capitalize(utils.SingularForm(name))))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\trouter := http.NewServeMux()"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\tmiddlewares := middleware.ChainMiddleware("))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\tinterceptor.LoggingInterceptor,"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\tmiddleware.CorsMiddleware,"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\tmiddleware.AuthenticationMiddleware,"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t)"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\trouter.Handle(\"/v1/%s\", middlewares(http.HandlerFunc(%sHandler.%sHandle)))", name, utils.SingularForm(name), utils.Capitalize(utils.SingularForm(name))))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\trouter.Handle(\"/\", middlewares(http.HandlerFunc(utils.ResponseNotFound)))"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\thostname := os.Getenv(\"HOSTNAME\")"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tport := os.Getenv(\"PORT\")"))
	data.WriteString(separator)
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tif port == \"\" {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\tlog.Panicf(\"specify the port\")"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprint("\taddress := fmt.Sprintf(\"%s:%s\", hostname, port)"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprint("\tlog.Printf(\"http server starts at address %s\\n\", address)"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\terr := http.ListenAndServe(address, router)"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\tif err != nil {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprint("\t\tlog.Panicf(\"error when starting the http server %v\\n\", err)"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t}"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)

	return data.Bytes()
}

func GetApplicationTemplate(module string, application string, name string, implementing string) []byte {
	data := bytes.Buffer{}
	separator := utils.GetSeparator()

	imports := []string{
		"\"context\"",
		"\"github.com/joho/godotenv\"",
		fmt.Sprintf("\"%s/cmd/%s\"", module, implementing),
		fmt.Sprintf("\"%s/internal/provider\"", module),
		fmt.Sprintf("%sProvider \"%s/internal/provider/%s\"", utils.SingularForm(name), module, name),
	}

	sort.Strings(imports)

	data.WriteString(fmt.Sprintf("package %s", application))
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

	data.WriteString(fmt.Sprintf("type %sInterface interface {", utils.Capitalize(application)))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tInitializeDependency(context.Context) error"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tInitializeEnvironment(context.Context) error"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tInitializeProvider(context.Context) error"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tRun()"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("type %s struct {", utils.Capitalize(application)))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t%sProvider provider.%sProviderInterface", utils.SingularForm(name), utils.Capitalize(utils.SingularForm(name))))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("var _ %sInterface = (*%s)(nil)", utils.Capitalize(application), utils.Capitalize(application)))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("func New%s(ctx context.Context) (*%s, error) {", utils.Capitalize(application), utils.Capitalize(application)))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t%s := &%s{}", utils.FirstLetter(application), utils.Capitalize(application)))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\terr := %s.InitializeDependency(ctx)", utils.FirstLetter(application)))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\tif err != nil {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\treturn nil, err"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\treturn %s, nil", utils.FirstLetter(application)))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("func (%s *%s) InitializeDependency(ctx context.Context) error {", utils.FirstLetter(application), utils.Capitalize(application)))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tinits := []func(context.Context) error{"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\t%s.InitializeEnvironment,", utils.FirstLetter(application)))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\t%s.InitializeProvider,", utils.FirstLetter(application)))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\tfor _, function := range inits {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\terr := function(ctx)"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\t\tif err != nil {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\treturn err"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\t}"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\treturn nil"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("func (%s *%s) InitializeEnvironment(_ context.Context) error {", utils.FirstLetter(application), utils.Capitalize(application)))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\terr := godotenv.Load(\".env\")"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\tif err != nil {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\treturn err"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\treturn nil"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("func (%s *%s) InitializeProvider(_ context.Context) error {", utils.FirstLetter(application), utils.Capitalize(application)))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t%s.%sProvider = %sProvider.New%sProvider()", utils.FirstLetter(application), utils.SingularForm(name), utils.SingularForm(name), utils.Capitalize(utils.SingularForm(name))))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\treturn nil"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("func (%s *%s) Run() {", utils.FirstLetter(application), utils.Capitalize(application)))
	data.WriteString(separator)

	switch implementing {
	case "grpc_server":
		data.WriteString(fmt.Sprintf("\timplementation := %s.%sProvider.Get%sImplementation()", utils.FirstLetter(application), utils.SingularForm(name), utils.Capitalize(utils.SingularForm(name))))
		data.WriteString(separator)
		data.WriteString(separator)

		data.WriteString(fmt.Sprintf("\t%s.Run(implementation)", implementing))
		data.WriteString(separator)
	case "http_server":
		data.WriteString(fmt.Sprintf("\thandler := %s.%sProvider.Get%sHandler()", utils.FirstLetter(application), utils.SingularForm(name), utils.Capitalize(utils.SingularForm(name))))
		data.WriteString(separator)
		data.WriteString(separator)

		data.WriteString(fmt.Sprintf("\t%s.Run(handler)", implementing))
		data.WriteString(separator)
	case "nats_subscribe":
		data.WriteString(fmt.Sprintf("\tcontroller := %s.%sProvider.Get%sController()", utils.FirstLetter(application), utils.SingularForm(name), utils.Capitalize(utils.SingularForm(name))))
		data.WriteString(separator)
		data.WriteString(separator)

		data.WriteString(fmt.Sprintf("\t%s.Run(controller)", implementing))
		data.WriteString(separator)
	case "cron_schedule":
		data.WriteString(fmt.Sprintf("\tservice  := %s.%sProvider.Get%sService()", utils.FirstLetter(application), utils.SingularForm(name), utils.Capitalize(utils.SingularForm(name))))
		data.WriteString(separator)
		data.WriteString(separator)

		data.WriteString(fmt.Sprintf("\t%s.Run(service)", implementing))
		data.WriteString(separator)
	default:
		data.WriteString(fmt.Sprintf("\t%s.Run()", implementing))
		data.WriteString(separator)
	}

	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)

	return data.Bytes()
}

func GetNatsSubscriberTemplate(module string, name string) []byte {
	data := bytes.Buffer{}
	separator := utils.GetSeparator()

	imports := []string{
		"\"log\"",
		"\"os\"",
		"\"os/signal\"",
		"\"syscall\"",
		fmt.Sprintf("\"%s/internal/controller\"", module),
		fmt.Sprintf("\"github.com/nats-io/stan.go\""),
	}

	sort.Strings(imports)

	data.WriteString(fmt.Sprintf("package nats_subscribe"))
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

	data.WriteString(fmt.Sprintf("func Run(_ controller.%sControllerInterface) {", utils.Capitalize(utils.SingularForm(name))))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tsc := connect()"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\tdefer sc.Close()"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("func connect() stan.Conn {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\turl := os.Getenv(\"NATS_URL\")"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\tif url == \"\" {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\tlog.Panicf(\"specify nats url\")"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\tcluster := os.Getenv(\"NATS_CLUSTER_ID\")"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\tif cluster == \"\" {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\tlog.Panicf(\"specify the cluster id\")"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\tsc, err := stan.Connect(cluster, \"subscriber-1\", stan.NatsURL(url))"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\tif err != nil {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprint("\t\tlog.Panicf(\"error %v\\n\", err)"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\treturn sc"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("func subscribe(sc stan.Conn, subject string, queue string, handler stan.MsgHandler) {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tsub, err := sc.QueueSubscribe(subject, queue, handler)"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\tif err != nil {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprint("\t\tlog.Panicf(\"error %v\\n\", err)"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\tdefer sub.Unsubscribe()"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprint("\tlog.Printf(\"subscribed to the message queue %s\\n\", subject)"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\tchannel := make(chan os.Signal, 1)"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tsignal.Notify(channel, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t<-channel"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)

	return data.Bytes()
}

func GetCronScheduleTemplate(module string, name string) []byte {
	data := bytes.Buffer{}
	separator := utils.GetSeparator()

	imports := []string{
		"\"log\"",
		"\"os\"",
		"\"os/signal\"",
		"\"syscall\"",
		"\"github.com/robfig/cron/v3\"",
		fmt.Sprintf("\"%s/internal/service\"", module),
	}

	sort.Strings(imports)

	data.WriteString(fmt.Sprintf("package cron_schedule"))
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

	data.WriteString(fmt.Sprintf("func Run(_ service.%sServiceInterface) {", utils.Capitalize(utils.SingularForm(name))))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tc := cron.New()"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\tlog.Printf(\"cron started\\n\")"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\tc.Start()"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\tdefer c.Stop()"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\texit := make(chan os.Signal)"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tsignal.Notify(exit, syscall.SIGINT)"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t<-exit"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\tlog.Printf(\"cron stopped\\n\")"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)

	return data.Bytes()
}

func GetMainTemplate(module string, application string) []byte {
	data := bytes.Buffer{}
	separator := utils.GetSeparator()

	imports := []string{
		"\"log\"",
		"\"golang.org/x/net/context\"",
		fmt.Sprintf("\"%s/cmd/%s\"", module, application),
	}

	sort.Strings(imports)

	data.WriteString(fmt.Sprintf("package main"))
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

	data.WriteString(fmt.Sprintf("func main() {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tctx := context.Background()"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\t%s, err := %s.New%s(ctx)", utils.FirstLetter(application), application, utils.Capitalize(application)))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\tif err != nil {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprint("\t\tlog.Panicf(\"an error occurred while starting the application %v\", err)"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\t%s.Run()", utils.FirstLetter(application)))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)

	return data.Bytes()
}

func GetProtoTemplate(module string, kind string, name string) []byte {
	data := bytes.Buffer{}
	separator := utils.GetSeparator()

	data.WriteString(fmt.Sprintf("syntax = \"proto3\";"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("package %s_v1;", name))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("option go_package = \"%s/%s/%s_v1\";", module, kind, name))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("service %sV1 {}", utils.Capitalize(name)))
	data.WriteString(separator)

	return data.Bytes()
}

func GetGoTemplate(module string, version string) []byte {
	data := bytes.Buffer{}
	separator := utils.GetSeparator()

	data.WriteString(fmt.Sprintf("module %s", module))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("go %s", version))
	data.WriteString(separator)

	return data.Bytes()
}

func GetUtilsConvertErrorTemplate() []byte {
	data := bytes.Buffer{}
	separator := utils.GetSeparator()

	imports := []string{
		"\"encoding/json\"",
	}

	sort.Strings(imports)

	data.WriteString(fmt.Sprintf("package utils"))
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

	data.WriteString(fmt.Sprintf("type ConverterError struct {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tError string `json:\"error\"`"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("func ConvertError(message string) []byte {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tconvert := &ConverterError{"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\tError: message,"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\tresult, err := json.Marshal(convert)"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\tif err != nil {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\treturn []byte(err.Error())"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\treturn result"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)

	return data.Bytes()
}

func GetUtilsResponseTemplate() []byte {
	data := bytes.Buffer{}
	separator := utils.GetSeparator()

	imports := []string{
		"\"net/http\"",
	}

	sort.Strings(imports)

	data.WriteString(fmt.Sprintf("package utils"))
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

	data.WriteString(fmt.Sprintf("func ResponseBadRequest(response http.ResponseWriter, message string) {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tresponse.Header().Set(\"Content-Type\", \"application/json\")"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tresponse.WriteHeader(http.StatusBadRequest)"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tresponse.Write(ConvertError(message))"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("func ResponseNotFound(response http.ResponseWriter, request *http.Request) {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tresponse.Header().Set(\"Content-Type\", \"application/json\")"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tresponse.WriteHeader(http.StatusNotFound)"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tresponse.Write(ConvertError(\"not found\"))"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("func ResponseMethodNotAllowed(response http.ResponseWriter, request *http.Request) {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tresponse.Header().Set(\"Content-Type\", \"application/json\")"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tresponse.WriteHeader(http.StatusMethodNotAllowed)"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tresponse.Write(ConvertError(\"method not allowed\"))"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)

	return data.Bytes()
}

func GetUtilsDatabaseTemplate() []byte {
	data := bytes.Buffer{}
	separator := utils.GetSeparator()

	imports := []string{
		"\"os\"",
		"\"fmt\"",
		"\"log\"",
		"\"context\"",
		"\"github.com/jackc/pgx/v4/pgxpool\"",
	}

	sort.Strings(imports)

	data.WriteString(fmt.Sprintf("package utils"))
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

	data.WriteString(fmt.Sprintf("type DatabaseInterface interface {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tInitialize()"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tGetPool() *pgxpool.Pool"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("type Database struct {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tcredentials string"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("var _ DatabaseInterface = (*Database)(nil)"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("func NewDatabase() *Database {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tdb := &Database{}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\tdb.Initialize()"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\treturn db"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("func (d *Database) Initialize() {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tusername := os.Getenv(\"POSTGRESQL_USERNAME\")"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\tif username == \"\" {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\tlog.Panicf(\"specify the database user\")"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\tpassword := os.Getenv(\"POSTGRESQL_PASSWORD\")"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\tif password == \"\" {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\tlog.Panicf(\"specify the database user password\")"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\tdb := os.Getenv(\"POSTGRESQL_DATABASE\")"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\tif db == \"\" {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\tlog.Panicf(\"ndicate the name of the database\")"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\thostname := os.Getenv(\"POSTGRESQL_HOSTNAME\")"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\tif hostname == \"\" {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\tlog.Panicf(\"specify the database hostname\")"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\tport := os.Getenv(\"POSTGRESQL_PORT\")"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\tif port == \"\" {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\tlog.Panicf(\"specify the database port\")"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\tsslmode := os.Getenv(\"POSTGRESQL_SSLMODE\")"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\tif sslmode == \"\" {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\tlog.Panicf(\"specify the ssl mode of the database\")"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprint("\td.credentials = fmt.Sprintf(\"user=%s password=%s dbname=%s host=%s port=%s sslmode=%s\", username, password, db, hostname, port, sslmode)"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("func (d *Database) GetPool() *pgxpool.Pool {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tpool, err := pgxpool.Connect(context.Background(), d.credentials)"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\tif err != nil {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprint("\t\tlog.Panicf(\"error %v\\n\", err)"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\treturn pool"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)

	return data.Bytes()
}

func GetUtilsNatsPublisherTemplate() []byte {
	data := bytes.Buffer{}
	separator := utils.GetSeparator()

	imports := []string{
		"\"os\"",
		"\"log\"",
		"\"github.com/nats-io/stan.go\"",
	}

	sort.Strings(imports)

	data.WriteString(fmt.Sprintf("package utils"))
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

	data.WriteString(fmt.Sprintf("type NatsPublisherInterface interface {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tInitialize()"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tGetConnect() stan.Conn"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("type NatsPublisher struct {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\turl string"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tcluster string"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("var _ NatsPublisherInterface = (*NatsPublisher)(nil)"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("func NewNatsPublisher() *NatsPublisher {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tpb := &NatsPublisher{}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\tpb.Initialize()"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\treturn pb"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("func (p *NatsPublisher) Initialize() {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\turl := os.Getenv(\"NATS_URL\")"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\tif url == \"\" {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\tlog.Panicf(\"specify nats url\")"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\tp.url = url"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\tcluster := os.Getenv(\"NATS_CLUSTER_ID\")"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\tif cluster == \"\" {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\tlog.Panicf(\"specify the cluster id\")"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\tp.cluster = cluster"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("func (p *NatsPublisher) GetConnect() stan.Conn {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tsc, err := stan.Connect(p.cluster, \"publisher-1\", stan.NatsURL(p.url))"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\tif err != nil {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprint("\t\tlog.Panicf(\"error %v\\n\", err)"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\treturn sc"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))

	return data.Bytes()
}
