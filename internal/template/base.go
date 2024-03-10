package template

import (
	"bytes"
	"fmt"
	"github.com/rafailovalexey/service-generator/internal/dto"
	"github.com/rafailovalexey/service-generator/internal/util"
	"sort"
)

func GetBaseDefinitionTemplate(application *dto.ApplicationDto, layer string) []byte {
	data := bytes.Buffer{}
	separator := util.GetSeparator()

	data.WriteString(fmt.Sprintf("package %s", layer))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("type %s%sInterface interface {}", application.Names.CamelCaseSingular, util.GetWithUpperCaseFirstLetter(layer)))
	data.WriteString(separator)

	return data.Bytes()
}

func GetBaseImplementationTemplate(application *dto.ApplicationDto, layer string) []byte {
	data := bytes.Buffer{}
	separator := util.GetSeparator()

	imports := []string{
		"\"github.com/sirupsen/logrus\"",
		fmt.Sprintf("\"%s/config\"", application.Module),
		fmt.Sprintf("definition \"%s/%s/%s\"", application.Module, "internal", layer),
	}

	switch layer {
	case "controller":
		imports = append(imports, fmt.Sprintf("\"%s/internal/converter\"", application.Module))
		imports = append(imports, fmt.Sprintf("\"%s/internal/validation\"", application.Module))
		imports = append(imports, fmt.Sprintf("\"%s/internal/service\"", application.Module))
	case "service":
		imports = append(imports, fmt.Sprintf("\"%s/internal/repository\"", application.Module))
	case "repository":
		imports = append(imports, fmt.Sprintf("\"%s/internal/converter\"", application.Module))
		imports = append(imports, fmt.Sprintf("\"database/sql\""))
	}

	sort.Strings(imports)

	data.WriteString(fmt.Sprintf("package %s", application.Names.SnakeCasePlural))
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

	data.WriteString(fmt.Sprintf("type %s%s struct {", application.Names.CamelCaseSingular, util.GetWithUpperCaseFirstLetter(layer)))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tconfig *config.Config"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tlogger *logrus.Logger"))
	data.WriteString(separator)
	switch layer {
	case "controller":
		data.WriteString(fmt.Sprintf("\t%sValidation validation.%sValidationInterface", application.Names.LowerCamelCaseSingular, application.Names.CamelCaseSingular))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t%sConverter converter.%sConverterInterface", application.Names.LowerCamelCaseSingular, application.Names.CamelCaseSingular))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t%sService service.%sServiceInterface", application.Names.LowerCamelCaseSingular, application.Names.CamelCaseSingular))
		data.WriteString(separator)
	case "service":
		data.WriteString(fmt.Sprintf("\t%sRepository repository.%sRepositoryInterface", application.Names.LowerCamelCaseSingular, application.Names.CamelCaseSingular))
		data.WriteString(separator)
	case "repository":
		data.WriteString(fmt.Sprintf("\tdatabase *sql.DB"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t%sConverter converter.%sConverterInterface", application.Names.LowerCamelCaseSingular, application.Names.CamelCaseSingular))
		data.WriteString(separator)
	}
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("var _ definition.%s%sInterface = (*%s%s)(nil)", application.Names.CamelCaseSingular, util.GetWithUpperCaseFirstLetter(layer), application.Names.CamelCaseSingular, util.GetWithUpperCaseFirstLetter(layer)))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("func New%s%s(", application.Names.CamelCaseSingular, util.GetWithUpperCaseFirstLetter(layer)))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tconfig *config.Config,"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tlogger *logrus.Logger,"))
	data.WriteString(separator)
	switch layer {
	case "controller":
		data.WriteString(fmt.Sprintf("\t%sValidation validation.%sValidationInterface,", application.Names.LowerCamelCaseSingular, application.Names.CamelCaseSingular))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t%sConverter converter.%sConverterInterface,", application.Names.LowerCamelCaseSingular, application.Names.CamelCaseSingular))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t%sService service.%sServiceInterface,", application.Names.LowerCamelCaseSingular, application.Names.CamelCaseSingular))
		data.WriteString(separator)
	case "service":
		data.WriteString(fmt.Sprintf("\t%sRepository repository.%sRepositoryInterface,", application.Names.LowerCamelCaseSingular, application.Names.CamelCaseSingular))
		data.WriteString(separator)
	case "repository":
		data.WriteString(fmt.Sprintf("\tdatabase *sql.DB,"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t%sConverter converter.%sConverterInterface,", application.Names.LowerCamelCaseSingular, application.Names.CamelCaseSingular))
		data.WriteString(separator)
	}
	data.WriteString(fmt.Sprintf(") definition.%s%sInterface {", application.Names.CamelCaseSingular, util.GetWithUpperCaseFirstLetter(layer)))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\treturn &%s%s{", application.Names.CamelCaseSingular, util.GetWithUpperCaseFirstLetter(layer)))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\tconfig: config,"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\tlogger: logger,"))
	data.WriteString(separator)
	switch layer {
	case "controller":
		data.WriteString(fmt.Sprintf("\t\t%sValidation: %sValidation,", application.Names.LowerCamelCaseSingular, application.Names.LowerCamelCaseSingular))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t\t%sConverter: %sConverter,", application.Names.LowerCamelCaseSingular, application.Names.LowerCamelCaseSingular))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t\t%sService: %sService,", application.Names.LowerCamelCaseSingular, application.Names.LowerCamelCaseSingular))
		data.WriteString(separator)
	case "service":
		data.WriteString(fmt.Sprintf("\t\t%sRepository: %sRepository,", application.Names.LowerCamelCaseSingular, application.Names.LowerCamelCaseSingular))
		data.WriteString(separator)
	case "repository":
		data.WriteString(fmt.Sprintf("\t\tdatabase: database,"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t\t%sConverter: %sConverter,", application.Names.LowerCamelCaseSingular, application.Names.LowerCamelCaseSingular))
		data.WriteString(separator)
	}
	data.WriteString(fmt.Sprintf("\t}"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)

	return data.Bytes()
}
