package template

import (
	"bytes"
	"fmt"
	"github.com/rafailovalexey/service-generator/internal/util"
	"sort"
)

func GetUtilConverterErrorTemplate() []byte {
	data := bytes.Buffer{}
	separator := util.GetSeparator()

	imports := []string{
		"\"encoding/json\"",
	}

	sort.Strings(imports)

	data.WriteString(fmt.Sprintf("package util"))
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

func GetUtilResponseTemplate() []byte {
	data := bytes.Buffer{}
	separator := util.GetSeparator()

	imports := []string{
		"\"net/http\"",
	}

	sort.Strings(imports)

	data.WriteString(fmt.Sprintf("package util"))
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

	data.WriteString(fmt.Sprintf("func ResponseUnauthorized(response http.ResponseWriter, request *http.Request) {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tresponse.Header().Set(\"Content-Type\", \"application/json\")"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tresponse.WriteHeader(http.StatusUnauthorized)"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tresponse.Write(ConvertError(\"unauthorized\"))"))
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
