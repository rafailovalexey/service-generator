package template

import (
	"bytes"
	"fmt"
	"github.com/rafailovalexey/service-generator/internal/utils"
)

func GetInterfaceTemplate(separator string, layer string, name string) []byte {
	data := bytes.Buffer{}

	data.WriteString(fmt.Sprintf("package %s", utils.Lowercase(layer)))
	data.WriteString(separator)
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("type %s%sInterface interface {}", utils.Capitalize(name), utils.Capitalize(layer)))
	data.WriteString(separator)

	return data.Bytes()
}

func GetRealisationInterfaceTemplate(application string, separator string, kind string, layer string, name string) []byte {
	data := bytes.Buffer{}

	data.WriteString(fmt.Sprintf("package %s", utils.Lowercase(name)))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("import ("))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tdefinition \"%s/%s/%s\"", application, kind, utils.Lowercase(layer)))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf(")"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("type %s%s struct {}", utils.Capitalize(name), utils.Capitalize(layer)))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("var _ definition.%s%sInterface = (*%s%s)(nil)", utils.Capitalize(name), utils.Capitalize(layer), utils.Capitalize(name), utils.Capitalize(layer)))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("func New%s%s() definition.%s%sInterface {", utils.Capitalize(name), utils.Capitalize(layer), utils.Capitalize(name), utils.Capitalize(layer)))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\treturn &%s%s{}", utils.Capitalize(name), utils.Capitalize(layer)))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)

	return data.Bytes()
}

func GetDataTransferObjectTemplate(separator string, layer string, name string) []byte {
	data := bytes.Buffer{}

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

func GetRequestTemplate(separator string, name string) []byte {
	data := bytes.Buffer{}

	data.WriteString(fmt.Sprintf("package %s", utils.Lowercase(name)))
	data.WriteString(separator)

	return data.Bytes()
}

func GetResponseTemplate(separator string, name string) []byte {
	data := bytes.Buffer{}

	data.WriteString(fmt.Sprintf("package %s", utils.Lowercase(name)))
	data.WriteString(separator)

	return data.Bytes()
}

func GetProviderInterfaceTemplate(application string, separator string, kind string, layers []string, layer string, name string) []byte {
	data := bytes.Buffer{}

	data.WriteString(fmt.Sprintf("package %s", utils.Lowercase(layer)))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("import ("))
	data.WriteString(separator)

	for _, l := range layers {
		switch {
		case l == "implementation":
			data.WriteString(fmt.Sprintf("\t\"%s/%s/%s/%s\"", application, kind, l, name))
			data.WriteString(separator)
		default:
			data.WriteString(fmt.Sprintf("\t\"%s/%s/%s\"", application, kind, l))
			data.WriteString(separator)
		}
	}

	data.WriteString(fmt.Sprintf(")"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("type %s%sInterface interface {", utils.Capitalize(name), utils.Capitalize(layer)))
	data.WriteString(separator)

	for _, l := range layers {
		switch {
		case l == "implementation":
			data.WriteString(fmt.Sprintf("\tGet%s%s() *%s.%sImplementation", utils.Capitalize(name), utils.Capitalize(l), name, utils.Capitalize(name)))
			data.WriteString(separator)
		default:
			data.WriteString(fmt.Sprintf("\tGet%s%s() %s.%s%sInterface", utils.Capitalize(name), utils.Capitalize(l), l, utils.Capitalize(name), utils.Capitalize(l)))
			data.WriteString(separator)
		}
	}

	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)

	return data.Bytes()
}

func GetProviderRealisationTemplate(application string, separator string, kind string, layers []string, layer string, name string) []byte {
	data := bytes.Buffer{}

	data.WriteString(fmt.Sprintf("package %s", utils.Lowercase(name)))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("import ("))
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\tdefinition \"%s/%s/%s\"", application, kind, utils.Lowercase(layer)))
	data.WriteString(separator)

	for _, l := range layers {
		switch {
		case l == "implementation":
			data.WriteString(fmt.Sprintf("\t\"%s/%s/implementation/%s\"", application, kind, name))
			data.WriteString(separator)
		default:
			data.WriteString(fmt.Sprintf("\t\"%s/%s/%s\"", application, kind, l))
			data.WriteString(separator)
			data.WriteString(fmt.Sprintf("\t%s%s \"%s/%s/%s/%s\"", name, utils.Capitalize(l), application, kind, l, name))
			data.WriteString(separator)
		}
	}

	data.WriteString(fmt.Sprintf(")"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("type %s%s struct {", utils.Capitalize(name), utils.Capitalize(layer)))
	data.WriteString(separator)

	for _, l := range layers {
		switch {
		case l == "implementation":
			data.WriteString(fmt.Sprintf("\t%s%s *%s.%sImplementation", name, utils.Capitalize(l), name, utils.Capitalize(name)))
			data.WriteString(separator)
		default:
			data.WriteString(fmt.Sprintf("\t%s%s %s.%s%sInterface", name, utils.Capitalize(l), l, utils.Capitalize(name), utils.Capitalize(l)))
			data.WriteString(separator)
		}
	}

	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("var _ definition.%s%sInterface = (*%s%s)(nil)", utils.Capitalize(name), utils.Capitalize(layer), utils.Capitalize(name), utils.Capitalize(layer)))
	data.WriteString(separator)
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("func New%s%s() definition.%s%sInterface {", utils.Capitalize(name), utils.Capitalize(layer), utils.Capitalize(name), utils.Capitalize(layer)))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\treturn &%s%s{}", utils.Capitalize(name), utils.Capitalize(layer)))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)

	for _, l := range layers {
		switch {
		case l == "implementation":
			data.WriteString(separator)
			data.WriteString(fmt.Sprintf("func (%s *%s%s) Get%s%s() *%s.%sImplementation {", utils.FirstLetter(layer), utils.Capitalize(name), utils.Capitalize(layer), utils.Capitalize(name), utils.Capitalize(l), name, utils.Capitalize(name)))
			data.WriteString(separator)
			data.WriteString(fmt.Sprintf("\tif %s.%s%s == nil {", utils.FirstLetter(layer), name, utils.Capitalize(l)))
			data.WriteString(separator)
			data.WriteString(fmt.Sprintf("\t\t%s.%s%s = %s.New%s%s()", utils.FirstLetter(layer), name, utils.Capitalize(l), name, utils.Capitalize(name), utils.Capitalize(l)))
			data.WriteString(separator)
			data.WriteString(fmt.Sprintf("\t}"))
			data.WriteString(separator)
			data.WriteString(separator)
			data.WriteString(fmt.Sprintf("\treturn %s.%s%s", utils.FirstLetter(layer), name, utils.Capitalize(l)))
			data.WriteString(separator)
			data.WriteString(fmt.Sprintf("}"))
			data.WriteString(separator)
		default:
			data.WriteString(separator)
			data.WriteString(fmt.Sprintf("func (%s *%s%s) Get%s%s() %s.%s%sInterface {", utils.FirstLetter(layer), utils.Capitalize(name), utils.Capitalize(layer), utils.Capitalize(name), utils.Capitalize(l), l, utils.Capitalize(name), utils.Capitalize(l)))
			data.WriteString(separator)
			data.WriteString(fmt.Sprintf("\tif %s.%s%s == nil {", utils.FirstLetter(layer), name, utils.Capitalize(l)))
			data.WriteString(separator)
			data.WriteString(fmt.Sprintf("\t\t%s.%s%s = %s%s.New%s%s()", utils.FirstLetter(layer), name, utils.Capitalize(l), name, utils.Capitalize(l), utils.Capitalize(name), utils.Capitalize(l)))
			data.WriteString(separator)
			data.WriteString(fmt.Sprintf("\t}"))
			data.WriteString(separator)
			data.WriteString(separator)
			data.WriteString(fmt.Sprintf("\treturn %s.%s%s", utils.FirstLetter(layer), name, utils.Capitalize(l)))
			data.WriteString(separator)
			data.WriteString(fmt.Sprintf("}"))
			data.WriteString(separator)
		}
	}

	return data.Bytes()
}

func GetImplementationRealisationTemplate(separator string, layer string, name string) []byte {
	data := bytes.Buffer{}

	data.WriteString(fmt.Sprintf("package %s", utils.Lowercase(name)))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("type %s%s struct {}", utils.Capitalize(name), utils.Capitalize(layer)))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("func New%s%s() *%s%s {", utils.Capitalize(name), utils.Capitalize(layer), utils.Capitalize(name), utils.Capitalize(layer)))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\treturn &%s%s{}", utils.Capitalize(name), utils.Capitalize(layer)))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("func (%s *%s%s) mustEmbedUnimplemented%sV1Server() {", utils.FirstLetter(name), utils.Capitalize(name), utils.Capitalize(layer), utils.Capitalize(name)))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\treturn"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)

	return data.Bytes()
}

func GetReadmeTemplate(separator string) []byte {
	data := bytes.Buffer{}

	data.WriteString(separator)

	return data.Bytes()
}

func GetGitIgnoreTemplate(separator string) []byte {
	data := bytes.Buffer{}

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

func GetExampleEnvironmentTemplate(separator string) []byte {
	data := bytes.Buffer{}

	data.WriteString(separator)

	return data.Bytes()
}

func GetGrpcMicroserviceMakefileTemplate(separator string) []byte {
	data := bytes.Buffer{}

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
	data.WriteString(fmt.Sprintf("\tbin/grpc-generate.sh $(PROTO_SOURCE_DIRECTORY) $(PROTO_OUTPUT_DIRECTORY) $(PROTO_FILES)"))
	data.WriteString(separator)
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("# Mocks"))
	data.WriteString(separator)
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("mocks-generate:"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tbin/mocks-generate.sh $(MOCKS_OUTPUT_DIRECTORY) $(MOCKS_FILES)"))
	data.WriteString(separator)
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("# Generate"))
	data.WriteString(separator)
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("generate:"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tmake grpc-generate"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tmake mocks-generate"))
	data.WriteString(separator)
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("# Download"))
	data.WriteString(separator)
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("download:"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tgo mod download"))
	data.WriteString(separator)
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("# Build"))
	data.WriteString(separator)
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("build:"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tgo build -o build/main main.go"))
	data.WriteString(separator)
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("# Tests"))
	data.WriteString(separator)
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("tests:"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tgo test -v ./..."))
	data.WriteString(separator)
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf(".PHONY: grpc-generate, mocks-generate, generate, download, build, tests"))
	data.WriteString(separator)

	return data.Bytes()
}

func GetDefaultMicroserviceMakefileTemplate(separator string) []byte {
	data := bytes.Buffer{}

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
	data.WriteString(fmt.Sprintf("mocks-generate:"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tbin/mocks-generate.sh $(MOCKS_OUTPUT_DIRECTORY) $(MOCKS_FILES)"))
	data.WriteString(separator)
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("# Generate"))
	data.WriteString(separator)
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("generate:"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tmake mocks-generate"))
	data.WriteString(separator)
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("# Download"))
	data.WriteString(separator)
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("download:"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tgo mod download"))
	data.WriteString(separator)
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("# Build"))
	data.WriteString(separator)
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("build:"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tgo build -o build/main main.go"))
	data.WriteString(separator)
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("# Tests"))
	data.WriteString(separator)
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("tests:"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tgo test -v ./..."))
	data.WriteString(separator)
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf(".PHONY: mocks-generate, generate, download, build, tests"))
	data.WriteString(separator)

	return data.Bytes()
}

func GetDockerIgnoreTemplate(separator string) []byte {
	data := bytes.Buffer{}

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

func GetDockerWithPortTemplate(separator string) []byte {
	data := bytes.Buffer{}

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

func GetDockerWithoutPortTemplate(separator string) []byte {
	data := bytes.Buffer{}

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

func GetGrpcGenerateShellScriptTemplate(separator string) []byte {
	data := bytes.Buffer{}

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

func GetMockGenerateShellScriptTemplate(separator string) []byte {
	data := bytes.Buffer{}

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
	data.WriteString(fmt.Sprintf("\tFILENAME_WITHOUT_EXTENSIONS=\"${FILENAME%.*}\""))
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

	//#!/bin/bash
	//
	//if [ "$#" -lt 1 ]; then
	//echo "Usage: $0 <MOCKS_OUTPUT_DIRECTORY> <MOCKS_FILES>"
	//exit 1
	//fi
	//
	//MOCKS_OUTPUT_DIRECTORY="$1"
	//MOCKS_FILES="${*:2}"
	//
	//for mock_file in $MOCKS_FILES; do
	//
	//DIRECTORY=$(dirname "$mock_file")
	//FILENAME=$(basename "$mock_file")
	//
	//EXTENSION="${FILENAME##*.}"
	//FILENAME_WITHOUT_EXTENSIONS="${FILENAME%.*}"
	//
	//OUTPUT_PATH="$DIRECTORY/$MOCKS_OUTPUT_DIRECTORY/${FILENAME_WITHOUT_EXTENSIONS}_mock.$EXTENSION"
	//
	//mkdir -p "$DIRECTORY/$MOCKS_OUTPUT_DIRECTORY"
	//
	//echo "Generating mock file for $mock_file..."
	//
	//mockgen -source="$mock_file" -destination="$OUTPUT_PATH"
	//
	//done

	return data.Bytes()
}

func GetGrpcLoggingInterceptorTemplate(separator string) []byte {
	data := bytes.Buffer{}

	//package interceptor
	//
	//import (
	//	"context"
	//"google.golang.org/grpc"
	//"google.golang.org/grpc/codes"
	//"google.golang.org/grpc/metadata"
	//"google.golang.org/grpc/status"
	//"log"
	//)
	//
	//func LoggingInterceptor() grpc.UnaryServerInterceptor {
	//	return func(ctx context.Context, request interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	//	md, isExist := metadata.FromIncomingContext(ctx)
	//
	//	if !isExist {
	//	return nil, status.Errorf(codes.Internal, "failed to read metadata")
	//}
	//
	//	tracecode := md["tracecode"][0]
	//
	//	log.Printf("incoming grpc request: %s (%s)", info.FullMethod, tracecode)
	//
	//	response, err := handler(ctx, request)
	//
	//	if err != nil {
	//	log.Printf("error in grpc request %s (%s) \n %v", info.FullMethod, tracecode, err)
	//}
	//
	//	if err == nil {
	//	log.Printf("outgoing grpc response %s (%s)", info.FullMethod, tracecode)
	//}
	//
	//	return response, err
	//}
	//}

	data.WriteString(separator)
	data.WriteString(separator)

	return data.Bytes()
}

func GetGrpcTraceCodeInterceptorTemplate(separator string) []byte {
	data := bytes.Buffer{}

	//package interceptor
	//
	//import (
	//	"context"
	//"crypto/rand"
	//"encoding/hex"
	//"google.golang.org/grpc"
	//"google.golang.org/grpc/codes"
	//"google.golang.org/grpc/metadata"
	//"google.golang.org/grpc/status"
	//"log"
	//)
	//
	//func TracecodeInterceptor() grpc.UnaryServerInterceptor {
	//	return func(ctx context.Context, request interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	//	md, isExist := metadata.FromIncomingContext(ctx)
	//
	//	if !isExist {
	//	log.Printf("metadata not found in the request context\n")
	//
	//	return nil, status.Errorf(codes.Internal, "failed to read metadata")
	//}
	//
	//	if len(md["tracecode"]) != 0 {
	//	return handler(ctx, request)
	//}
	//
	//	tracecode, err := generateTracecode()
	//
	//	if err != nil {
	//	return nil, status.Errorf(codes.Internal, "failed to generate tracecode")
	//}
	//
	//	md = metadata.Join(md, metadata.New(map[string]string{"tracecode": tracecode}))
	//	ctx = metadata.NewIncomingContext(ctx, md)
	//
	//	return handler(ctx, request)
	//}
	//}
	//
	//func generateTracecode() (string, error) {
	//	tracecode := make([]byte, 16)
	//
	//	if _, err := rand.Read(tracecode); err != nil {
	//		return "", err
	//	}
	//
	//	return hex.EncodeToString(tracecode), nil
	//}

	data.WriteString(separator)
	data.WriteString(separator)

	return data.Bytes()
}

func GetGrpcAuthenticationMiddlewareTemplate(separator string) []byte {
	data := bytes.Buffer{}

	//package middleware
	//
	//import (
	//	"context"
	//"google.golang.org/grpc"
	//"google.golang.org/grpc/codes"
	//"google.golang.org/grpc/metadata"
	//"google.golang.org/grpc/status"
	//"log"
	//"os"
	//)
	//
	//func AuthenticationTokenMiddleware() grpc.UnaryServerInterceptor {
	//	return func(ctx context.Context, request interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	//	md, isExist := metadata.FromIncomingContext(ctx)
	//
	//	if !isExist {
	//	return nil, status.Errorf(codes.Unauthenticated, "authentication token not found")
	//}
	//
	//	header := os.Getenv("AUTHENTICATION_TOKEN_HEADER")
	//
	//	if header == "" {
	//	log.Panicf("not found authentication token header in environment")
	//}
	//
	//	list := md[header]
	//
	//	if len(list) == 0 {
	//	return nil, status.Errorf(codes.Unauthenticated, "authentication token not found")
	//}
	//
	//	key := list[0]
	//
	//	token := os.Getenv("AUTHENTICATION_TOKEN")
	//
	//	if token == "" {
	//	log.Panicf("not found authentication token in environment")
	//}
	//
	//	if token != key {
	//	log.Printf("invalid authentication token: %s", key)
	//
	//	return nil, status.Errorf(codes.PermissionDenied, "invalid authentication token")
	//}
	//
	//	return handler(ctx, request)
	//}
	//}

	data.WriteString(separator)
	data.WriteString(separator)

	return data.Bytes()
}

func GetGrpcServerTemplate(separator string) []byte {
	data := bytes.Buffer{}

	//package grpc_server
	//
	//import (
	//	"fmt"
	//interceptor "github.com/emptyhopes/employees/cmd/grpc_server/intereptor"
	//"github.com/emptyhopes/employees/cmd/grpc_server/middleware"
	//"github.com/emptyhopes/employees/pkg/employees_v1"
	//"google.golang.org/grpc"
	//"google.golang.org/grpc/reflection"
	//"log"
	//"net"
	//"os"
	//)
	//
	//func Run(api employees_v1.EmployeesV1Server) {
	//	hostname := os.Getenv("HOSTNAME")
	//
	//	port := os.Getenv("PORT")
	//
	//	if port == "" {
	//		log.Panicf("specify the port")
	//	}
	//
	//	address := fmt.Sprintf("%s:%s", hostname, port)
	//
	//	log.Printf("%s\n", fmt.Sprintf("grpc server starts at address %s", address))
	//
	//	listener, err := net.Listen("tcp", address)
	//
	//	if err != nil {
	//		log.Panicf("grpc server startup error %v", err)
	//	}
	//
	//	server := grpc.NewServer(
	//		grpc.ChainUnaryInterceptor(
	//			interceptor.TracecodeInterceptor(),
	//			interceptor.LoggingInterceptor(),
	//			middleware.AuthenticationTokenMiddleware(),
	//		),
	//	)
	//
	//	reflection.Register(server)
	//
	//	employees_v1.RegisterEmployeesV1Server(server, api)
	//
	//	log.Printf("%s\n", fmt.Sprintf("grpc server is running at %s", address))
	//
	//	err = server.Serve(listener)
	//
	//	if err != nil {
	//		log.Panicf("grpc server startup error %v", err)
	//	}
	//}

	data.WriteString(separator)
	data.WriteString(separator)

	return data.Bytes()
}

func GetHttpLoggingInterceptorTemplate(separator string) []byte {
	data := bytes.Buffer{}

	//package interceptor
	//
	//import (
	//	"log"
	//"net/http"
	//"time"
	//)
	//
	//func LoggingInterceptor(next http.Handler) http.Handler {
	//	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
	//	start := time.Now()
	//
	//	next.ServeHTTP(response, request)
	//
	//	duration := time.Since(start)
	//
	//	log.Printf("%s %s %s - %s %v\n", request.Method, request.URL.Path, request.RemoteAddr, request.UserAgent(), duration)
	//})
	//}

	data.WriteString(separator)
	data.WriteString(separator)

	return data.Bytes()
}

func GetHttpAuthenticationMiddlewareTemplate(separator string) []byte {
	data := bytes.Buffer{}

	//package middleware
	//
	//import (
	//	"encoding/json"
	//"log"
	//"net/http"
	//"os"
	//)
	//
	//func AuthenticationMiddleware(next http.Handler) http.Handler {
	//	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
	//	header := os.Getenv("AUTHENTICATION_TOKEN_HEADER")
	//
	//	if header == "" {
	//	log.Panicf("specify the name of the authentication token")
	//}
	//
	//	token := os.Getenv("AUTHENTICATION_TOKEN")
	//
	//	if token == "" {
	//	log.Panicf("specify the value of the authentication token")
	//}
	//
	//	key := request.Header.Get(header)
	//
	//	if key != token {
	//	response.Header().Set("Content-Type", "application/json")
	//	response.WriteHeader(http.StatusUnauthorized)
	//	response.Write(getErrorInJson("unauthorized"))
	//
	//	return
	//}
	//
	//	next.ServeHTTP(response, request)
	//})
	//}
	//
	//func getErrorInJson(message string) []byte {
	//	type ErrorStruct struct {
	//	Error string `json:"error"`
	//}
	//
	//	errorStruct := &ErrorStruct{
	//	Error: message,
	//}
	//
	//	errJson, err := json.Marshal(errorStruct)
	//
	//	if err != nil {
	//	return []byte(err.Error())
	//}
	//
	//	return errJson
	//}

	data.WriteString(separator)
	data.WriteString(separator)

	return data.Bytes()
}

func GetHttpCorsMiddlewareTemplate(separator string) []byte {
	data := bytes.Buffer{}

	//package middleware
	//
	//import (
	//	"net/http"
	//)
	//
	//func CorsMiddleware(next http.Handler) http.Handler {
	//	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
	//	response.Header().Add("Access-Control-Allow-Origin", "*")
	//	response.Header().Add("Access-Control-Allow-Headers", "*")
	//	response.Header().Add("Access-Control-Allow-Methods", "*")
	//	response.Header().Add("Access-Control-Allow-Credentials", "true")
	//
	//	if request.Method == "OPTIONS" {
	//	response.WriteHeader(http.StatusOK)
	//
	//	return
	//}
	//
	//	next.ServeHTTP(response, request)
	//})
	//}

	data.WriteString(separator)
	data.WriteString(separator)

	return data.Bytes()
}

func GetHttpChainMiddlewareTemplate(separator string) []byte {
	data := bytes.Buffer{}

	//package middleware
	//
	//import (
	//	"net/http"
	//)
	//
	//func ChainMiddleware(middlewares ...func(http.Handler) http.Handler) func(http.Handler) http.Handler {
	//	return func(next http.Handler) http.Handler {
	//		for index := len(middlewares) - 1; index >= 0; index-- {
	//			next = middlewares[index](next)
	//		}
	//
	//		return next
	//	}
	//}

	data.WriteString(separator)
	data.WriteString(separator)

	return data.Bytes()
}

func GetHttpServerTemplate(separator string) []byte {
	data := bytes.Buffer{}

	//package http_server
	//
	//import (
	//	"fmt"
	//"github.com/emptyhopes/employees_proxy/cmd/http_server/interceptor"
	//"github.com/emptyhopes/employees_proxy/cmd/http_server/middleware"
	//"github.com/emptyhopes/employees_proxy/internal/api"
	//"github.com/gorilla/mux"
	//"log"
	//"net/http"
	//"os"
	//)
	//
	//func Run(employeeApi api.InterfaceEmployeeApi) {
	//	router := mux.NewRouter()
	//
	//	middlewares := middleware.ChainMiddleware(
	//		interceptor.LoggingInterceptor,
	//		middleware.CorsMiddleware,
	//		middleware.AuthenticationMiddleware,
	//	)
	//
	//	router.Use(middlewares)
	//
	//	router.NotFoundHandler = http.HandlerFunc(employeeApi.NotFound)
	//	router.MethodNotAllowedHandler = http.HandlerFunc(employeeApi.MethodNotAllowed)
	//
	//	router.HandleFunc("/v1/employees/{id:[a-zA-Z0-9-]+}", employeeApi.GetEmployeeById).Methods("GET")
	//	router.HandleFunc("/v1/employees", employeeApi.CreateEmployee).Methods("POST")
	//
	//	hostname := os.Getenv("HOSTNAME")
	//
	//	port := os.Getenv("PORT")
	//
	//	if port == "" {
	//		log.Panicf("specify the port")
	//	}
	//
	//	address := fmt.Sprintf("%s:%s", hostname, port)
	//
	//	log.Printf("http server starts at address %s\n", address)
	//
	//	err := http.ListenAndServe(address, router)
	//
	//	if err != nil {
	//		log.Panicf("error when starting the grpc server %v\n", err)
	//	}
	//}

	data.WriteString(separator)
	data.WriteString(separator)

	return data.Bytes()
}
