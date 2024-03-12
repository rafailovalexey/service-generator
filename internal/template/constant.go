package template

import (
	"bytes"
	"fmt"
	"github.com/rafailovalexey/service-generator/internal/dto"
	"github.com/rafailovalexey/service-generator/internal/util"
)

func GetConstantTemplate(application *dto.ApplicationDto) []byte {
	data := bytes.Buffer{}
	separator := util.GetSeparator()

	data.WriteString(fmt.Sprintf("package %s", application.Names.SnakeCasePlural))
	data.WriteString(separator)

	return data.Bytes()
}
