package template

import (
	"bytes"
	"fmt"
	"github.com/rafailovalexey/service-generator/internal/dto"
	"github.com/rafailovalexey/service-generator/internal/util"
	"sort"
)

func GetProviderDefinitionTemplate(application *dto.ApplicationDto, layers []string) []byte {
	data := bytes.Buffer{}
	separator := util.GetSeparator()

	imports := make([]string, 0, 10)

	for _, l := range layers {
		switch l {
		case "implementation":
			imports = append(imports, fmt.Sprintf("%s%s \"%s/%s/%s/%s\"", application.Names.LowerCamelCaseSingular, util.GetWithUpperCaseFirstLetter(l), application.Module, "internal", l, application.Names.SnakeCasePlural))
		default:
			imports = append(imports, fmt.Sprintf("\"%s/%s/%s\"", application.Module, "internal", l))
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
		data.WriteString(fmt.Sprintf("type %s%sInterface any", application.Names.CamelCaseSingular, util.GetWithUpperCaseFirstLetter("provider")))
		data.WriteString(separator)
	default:
		data.WriteString(fmt.Sprintf("type %s%sInterface interface {", application.Names.CamelCaseSingular, util.GetWithUpperCaseFirstLetter("provider")))
		data.WriteString(separator)

		for _, layer := range layers {
			switch layer {
			case "implementation":
				data.WriteString(fmt.Sprintf("\tGet%s%s() *%s%s.%sImplementation", application.Names.CamelCaseSingular, util.GetWithUpperCaseFirstLetter(layer), application.Names.LowerCamelCaseSingular, util.GetWithUpperCaseFirstLetter(layer), application.Names.CamelCaseSingular))
				data.WriteString(separator)
			default:
				data.WriteString(fmt.Sprintf("\tGet%s%s() %s.%s%sInterface", application.Names.CamelCaseSingular, util.GetWithUpperCaseFirstLetter(layer), layer, application.Names.CamelCaseSingular, util.GetWithUpperCaseFirstLetter(layer)))
				data.WriteString(separator)
			}
		}

		data.WriteString(fmt.Sprintf("}"))
		data.WriteString(separator)
	}

	return data.Bytes()
}

func GetProviderImplementationTemplate(application *dto.ApplicationDto, layers []string) []byte {
	data := bytes.Buffer{}
	separator := util.GetSeparator()

	imports := []string{
		"\"database/sql\"",
		"\"github.com/sirupsen/logrus\"",
		fmt.Sprintf("\"%s/config\"", application.Module),
		fmt.Sprintf("definition \"%s/%s/%s\"", application.Module, "internal", "provider"),
	}

	for _, layer := range layers {
		switch layer {
		case "implementation":
			imports = append(imports, fmt.Sprintf("%s%s \"%s/%s/implementation/%s\"", application.Names.LowerCamelCaseSingular, util.GetWithUpperCaseFirstLetter(layer), application.Module, "internal", application.Names.SnakeCasePlural))
		default:
			imports = append(imports, fmt.Sprintf("\"%s/%s/%s\"", application.Module, "internal", layer))
			imports = append(imports, fmt.Sprintf("%s%s \"%s/%s/%s/%s\"", application.Names.LowerCamelCaseSingular, util.GetWithUpperCaseFirstLetter(layer), application.Module, "internal", layer, application.Names.SnakeCasePlural))
		}
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

	switch len(layers) {
	case 0:
		data.WriteString(fmt.Sprintf("type %s%s struct {", application.Names.CamelCaseSingular, util.GetWithUpperCaseFirstLetter("provider")))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\tconfig *config.Config"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\tlogger *logrus.Logger"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\tdatabase *sql.DB"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("}"))
		data.WriteString(separator)
		data.WriteString(separator)
	default:
		data.WriteString(fmt.Sprintf("type %s%s struct {", application.Names.CamelCaseSingular, util.GetWithUpperCaseFirstLetter("provider")))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\tconfig *config.Config"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\tlogger *logrus.Logger"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\tdatabase *sql.DB"))
		data.WriteString(separator)

		for _, layer := range layers {
			switch layer {
			case "implementation":
				data.WriteString(fmt.Sprintf("\t%s%s *%s%s.%s%s", application.Names.LowerCamelCaseSingular, util.GetWithUpperCaseFirstLetter(layer), application.Names.LowerCamelCaseSingular, util.GetWithUpperCaseFirstLetter(layer), application.Names.CamelCaseSingular, util.GetWithUpperCaseFirstLetter(layer)))
				data.WriteString(separator)
			default:
				data.WriteString(fmt.Sprintf("\t%s%s %s.%s%sInterface", application.Names.LowerCamelCaseSingular, util.GetWithUpperCaseFirstLetter(layer), layer, application.Names.CamelCaseSingular, util.GetWithUpperCaseFirstLetter(layer)))
				data.WriteString(separator)
			}
		}

		data.WriteString(fmt.Sprintf("}"))
		data.WriteString(separator)
		data.WriteString(separator)
	}

	data.WriteString(fmt.Sprintf("var _ definition.%s%sInterface = (*%s%s)(nil)", application.Names.CamelCaseSingular, util.GetWithUpperCaseFirstLetter("provider"), application.Names.CamelCaseSingular, util.GetWithUpperCaseFirstLetter("provider")))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("func New%s%s(", application.Names.CamelCaseSingular, util.GetWithUpperCaseFirstLetter("provider")))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tconfig *config.Config,"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tlogger *logrus.Logger,"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tdatabase *sql.DB,"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf(") definition.%s%sInterface {", application.Names.CamelCaseSingular, util.GetWithUpperCaseFirstLetter("provider")))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\treturn &%s%s{", application.Names.CamelCaseSingular, util.GetWithUpperCaseFirstLetter("provider")))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\tconfig: config,"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\tlogger: logger,"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\tdatabase: database,"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t}"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)

	for _, l := range layers {
		switch l {
		case "implementation":
			data.WriteString(separator)
			data.WriteString(fmt.Sprintf("func (%s *%s%s) Get%s%s() *%s%s.%sImplementation {", util.GetFirstLetterLowerCase("provider"), application.Names.CamelCaseSingular, util.GetWithUpperCaseFirstLetter("provider"), application.Names.CamelCaseSingular, util.GetWithUpperCaseFirstLetter(l), application.Names.LowerCamelCaseSingular, util.GetWithUpperCaseFirstLetter(l), application.Names.CamelCaseSingular))
			data.WriteString(separator)
			data.WriteString(fmt.Sprintf("\tif %s.%s%s == nil {", util.GetFirstLetterLowerCase("provider"), application.Names.LowerCamelCaseSingular, util.GetWithUpperCaseFirstLetter(l)))
			data.WriteString(separator)
			data.WriteString(fmt.Sprintf("\t\t%s.%s%s = %s%s.New%s%s(", util.GetFirstLetterLowerCase("provider"), application.Names.LowerCamelCaseSingular, util.GetWithUpperCaseFirstLetter(l), application.Names.LowerCamelCaseSingular, util.GetWithUpperCaseFirstLetter(l), application.Names.CamelCaseSingular, util.GetWithUpperCaseFirstLetter(l)))
			data.WriteString(separator)
			data.WriteString(fmt.Sprintf("\t\t\tp.config,"))
			data.WriteString(separator)
			data.WriteString(fmt.Sprintf("\t\t\tp.logger,"))
			data.WriteString(separator)
			data.WriteString(fmt.Sprintf("\t\t\tp.%sController,", application.Names.LowerCamelCaseSingular))
			data.WriteString(separator)
			data.WriteString(fmt.Sprintf("\t\t)"))
			data.WriteString(separator)
			data.WriteString(fmt.Sprintf("\t}"))
			data.WriteString(separator)
			data.WriteString(separator)

			data.WriteString(fmt.Sprintf("\treturn %s.%s%s", util.GetFirstLetterLowerCase("provider"), application.Names.LowerCamelCaseSingular, util.GetWithUpperCaseFirstLetter(l)))
			data.WriteString(separator)
			data.WriteString(fmt.Sprintf("}"))
			data.WriteString(separator)
		default:
			data.WriteString(separator)
			data.WriteString(fmt.Sprintf("func (%s *%s%s) Get%s%s() %s.%s%sInterface {", util.GetFirstLetterLowerCase("provider"), application.Names.CamelCaseSingular, util.GetWithUpperCaseFirstLetter("provider"), application.Names.CamelCaseSingular, util.GetWithUpperCaseFirstLetter(l), l, application.Names.CamelCaseSingular, util.GetWithUpperCaseFirstLetter(l)))
			data.WriteString(separator)
			data.WriteString(fmt.Sprintf("\tif %s.%s%s == nil {", util.GetFirstLetterLowerCase("provider"), application.Names.LowerCamelCaseSingular, util.GetWithUpperCaseFirstLetter(l)))
			data.WriteString(separator)
			data.WriteString(fmt.Sprintf("\t\t%s.%s%s = %s%s.New%s%s(", util.GetFirstLetterLowerCase("provider"), application.Names.LowerCamelCaseSingular, util.GetWithUpperCaseFirstLetter(l), application.Names.LowerCamelCaseSingular, util.GetWithUpperCaseFirstLetter(l), application.Names.CamelCaseSingular, util.GetWithUpperCaseFirstLetter(l)))
			data.WriteString(separator)
			data.WriteString(fmt.Sprintf("\t\t\tp.config,"))
			data.WriteString(separator)
			data.WriteString(fmt.Sprintf("\t\t\tp.logger,"))
			data.WriteString(separator)
			switch l {
			case "handler":
				data.WriteString(fmt.Sprintf("\t\t\tp.%sController,", application.Names.LowerCamelCaseSingular))
				data.WriteString(separator)
			case "controller":
				data.WriteString(fmt.Sprintf("\t\t\tp.%sValidation,", application.Names.LowerCamelCaseSingular))
				data.WriteString(separator)
				data.WriteString(fmt.Sprintf("\t\t\tp.%sConverter,", application.Names.LowerCamelCaseSingular))
				data.WriteString(separator)
				data.WriteString(fmt.Sprintf("\t\t\tp.%sService,", application.Names.LowerCamelCaseSingular))
				data.WriteString(separator)
			case "service":
				data.WriteString(fmt.Sprintf("\t\t\tp.%sRepository,", application.Names.LowerCamelCaseSingular))
				data.WriteString(separator)
			case "repository":
				data.WriteString(fmt.Sprintf("\t\t\tp.database,"))
				data.WriteString(separator)
				data.WriteString(fmt.Sprintf("\t\t\tp.%sConverter,", application.Names.LowerCamelCaseSingular))
				data.WriteString(separator)
			}
			data.WriteString(fmt.Sprintf("\t\t)"))
			data.WriteString(separator)
			data.WriteString(fmt.Sprintf("\t}"))
			data.WriteString(separator)
			data.WriteString(separator)

			data.WriteString(fmt.Sprintf("\treturn %s.%s%s", util.GetFirstLetterLowerCase("provider"), application.Names.LowerCamelCaseSingular, util.GetWithUpperCaseFirstLetter(l)))
			data.WriteString(separator)
			data.WriteString(fmt.Sprintf("}"))
			data.WriteString(separator)
		}
	}

	return data.Bytes()
}
