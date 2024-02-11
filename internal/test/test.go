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
	data.WriteString(fmt.Sprintf("\t@go test -v ./..."))
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

func GetMainTemplate(module string, application string) []byte {
	data := bytes.Buffer{}
	separator := utils.GetSeparator()

	typ := "application"

	switch application {
	case "http":
		typ = "application"
	case "grpc":
		typ = "application"
	case "cron":
		typ = "cron"
	case "subscriber":
		typ = "subscriber"
	case "publisher":
		typ = "publisher"
	}

	imports := []string{
		"\"log\"",
		"\"golang.org/x/net/context\"",
		fmt.Sprintf("\"%s/cmd/%s\"", module, typ),
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

	data.WriteString(fmt.Sprintf("\t%s, err := %s.New%s(ctx)", utils.FirstLetter(typ), typ, utils.Capitalize(typ)))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\tif err != nil {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprint("\t\tlog.Panicf(\"an error occurred while starting the application %v\", err)"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\t%s.Run()", utils.FirstLetter(typ)))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)

	return data.Bytes()
}
