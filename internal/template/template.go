package template

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

func GetProtoTemplate(module string, name string) []byte {
	data := bytes.Buffer{}
	separator := utils.GetSeparator()

	data.WriteString(fmt.Sprintf("syntax = \"proto3\";"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("package %s_v1;", name))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("option go_package = \"%s/%s/%s_v1\";", module, "api", name))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("service %sV1 {}", utils.Capitalize(name)))
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

func GetMakefileTemplate(application string, name string) []byte {
	data := bytes.Buffer{}
	separator := utils.GetSeparator()

	data.WriteString(fmt.Sprintf("# Variables"))
	data.WriteString(separator)
	data.WriteString(separator)

	switch application {
	case "grpc":
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
	}

	data.WriteString(fmt.Sprintf("MOCKS_OUTPUT_DIRECTORY = mocks"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("MOCKS_FILES = \\"))
	data.WriteString(separator)
	data.WriteString(separator)

	switch application {
	case "grpc":
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
	}

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
	data.WriteString(fmt.Sprintf("\t@go template -v ./..."))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf(".PHONY: grpc-generate, mock-generate, generate, download, build, tests"))
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

func GetReadmeTemplate(module string) []byte {
	data := bytes.Buffer{}
	separator := utils.GetSeparator()

	data.WriteString(fmt.Sprintf("%s", module))
	data.WriteString(separator)

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

	data.WriteString(fmt.Sprintf("# Visual Studio Code"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf(".vscode"))
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

func GetExampleEnvironmentTemplate() []byte {
	data := bytes.Buffer{}

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

	data.WriteString(fmt.Sprintf("# Visual Studio Code"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf(".vscode"))
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

func GetMainTemplate(module string) []byte {
	data := bytes.Buffer{}
	separator := utils.GetSeparator()

	imports := []string{
		"\"log\"",
		"\"golang.org/x/net/context\"",
		fmt.Sprintf("\"%s/cmd/application\"", module),
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

	data.WriteString(fmt.Sprintf("\ta, err := application.NewApplication(ctx)"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\tif err != nil {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprint("\t\tlog.Panicf(\"an error occurred while starting the application %v\", err)"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\ta.Run()"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
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

	data.WriteString(fmt.Sprint("\t\tlog.Printf(\"%s %s %s - %s %v\\n\", request.Method, request.URL.Name, request.RemoteAddr, request.UserAgent(), duration)"))
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

func GetApplicationTemplate(module string, application string, name string) []byte {
	data := bytes.Buffer{}
	separator := utils.GetSeparator()

	typ := ""

	switch application {
	case "http":
		typ = "http_server"
	case "grpc":
		typ = "grpc_server"
	case "cron":
		typ = "cron_scheduler"
	case "subscriber":
		typ = "subscriber"
	case "publisher":
		typ = "publisher"
	}

	imports := []string{
		"\"context\"",
		"\"github.com/joho/godotenv\"",
		fmt.Sprintf("\"%s/cmd/%s\"", module, typ),
		fmt.Sprintf("\"%s/internal/provider\"", module),
		fmt.Sprintf("%sProvider \"%s/internal/provider/%s\"", utils.SingularForm(name), module, name),
	}

	sort.Strings(imports)

	data.WriteString(fmt.Sprintf("package application"))
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

	data.WriteString(fmt.Sprintf("type ApplicationInterface interface {"))
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

	data.WriteString(fmt.Sprintf("type Application struct {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t%sProvider provider.%sProviderInterface", utils.SingularForm(name), utils.Capitalize(utils.SingularForm(name))))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("var _ ApplicationInterface = (*Application)(nil)"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("func NewApplication(ctx context.Context) (*Application, error) {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\ta := &Application{}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\terr := a.InitializeDependency(ctx)"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\tif err != nil {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\treturn nil, err"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\treturn a, nil"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("func (a *Application) InitializeDependency(ctx context.Context) error {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tinits := []func(context.Context) error{"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\ta.InitializeEnvironment,"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\ta.InitializeProvider,"))
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

	data.WriteString(fmt.Sprintf("func (a *Application) InitializeEnvironment(_ context.Context) error {"))
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

	data.WriteString(fmt.Sprintf("func (a *Application) InitializeProvider(_ context.Context) error {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\ta.%sProvider = %sProvider.New%sProvider()", utils.SingularForm(name), utils.SingularForm(name), utils.Capitalize(utils.SingularForm(name))))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\treturn nil"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("func (a *Application) Run() {"))
	data.WriteString(separator)

	switch application {
	case "grpc":
		data.WriteString(fmt.Sprintf("\timplementation := a.%sProvider.Get%sImplementation()", utils.SingularForm(name), utils.Capitalize(utils.SingularForm(name))))
		data.WriteString(separator)
		data.WriteString(separator)

		data.WriteString(fmt.Sprintf("\t%s.Run(implementation)", typ))
		data.WriteString(separator)
	case "http":
		data.WriteString(fmt.Sprintf("\thandler := a.%sProvider.Get%sHandler()", utils.SingularForm(name), utils.Capitalize(utils.SingularForm(name))))
		data.WriteString(separator)
		data.WriteString(separator)

		data.WriteString(fmt.Sprintf("\t%s.Run(handler)", typ))
		data.WriteString(separator)
	default:
		data.WriteString(fmt.Sprintf("\t%s.Run()", typ))
		data.WriteString(separator)
	}

	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)

	return data.Bytes()
}

func GetDockerTemplate(withPort bool) []byte {
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

	if withPort {
		data.WriteString(fmt.Sprintf("EXPOSE 3000"))
		data.WriteString(separator)
		data.WriteString(separator)
	}

	data.WriteString(fmt.Sprintf("CMD [\"./build/main\"]"))
	data.WriteString(separator)

	return data.Bytes()
}

func GetUtilsConverterErrorTemplate() []byte {
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
