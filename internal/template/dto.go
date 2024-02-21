package template

import (
	"bytes"
	"fmt"
	"github.com/rafailovalexey/service-generator/internal/dto"
	"github.com/rafailovalexey/service-generator/internal/util"
)

func GetDataTransferObjectTemplate(layer string, name *dto.NameDto) []byte {
	data := bytes.Buffer{}
	separator := util.GetSeparator()

	data.WriteString(fmt.Sprintf("package %s", name.SnakeCasePlural))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("type %s%s struct {}", name.CamelCaseSingular, util.GetWithUpperCaseFirstLetter(layer)))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("type %s%s = []%s%s", name.CamelCasePlural, util.GetWithUpperCaseFirstLetter(layer), name.CamelCaseSingular, util.GetWithUpperCaseFirstLetter(layer)))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("func New%s%s() *%s%s {", name.CamelCaseSingular, util.GetWithUpperCaseFirstLetter(layer), name.CamelCaseSingular, util.GetWithUpperCaseFirstLetter(layer)))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\treturn &%s%s{}", name.CamelCaseSingular, util.GetWithUpperCaseFirstLetter(layer)))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("func New%s%s(%s ...%s%s) *%s%s {", name.CamelCasePlural, util.GetWithUpperCaseFirstLetter(layer), name.LowerCamelCasePlural, name.CamelCaseSingular, util.GetWithUpperCaseFirstLetter(layer), name.CamelCasePlural, util.GetWithUpperCaseFirstLetter(layer)))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t%s := make(%s%s, len(%s))", layer, name.CamelCasePlural, util.GetWithUpperCaseFirstLetter(layer), name.LowerCamelCasePlural))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\tfor _, %s := range %s {", name.LowerCamelCaseSingular, name.LowerCamelCasePlural))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\t%s = append(%s, %s)", layer, layer, name.LowerCamelCaseSingular))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\treturn &%s", layer))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)

	return data.Bytes()
}
