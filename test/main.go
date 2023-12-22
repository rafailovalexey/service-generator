package test

import (
	"fmt"
	"github.com/rafailovalexey/service-generator/internal/template"
	"github.com/rafailovalexey/service-generator/internal/utils"
)

func main() {
	layer := "repository"
	name := "employees"
	application, _ := utils.GetApplicationName()
	kind := "internal"
	separator := utils.GetSeparator()

	imports := &template.Imports{
		"net/http",
	}

	methods := &template.Methods{
		template.Method{
			Name:          "Test1",
			ArgumentsType: []string{"string", "string"},
			ArgumentsName: []string{"test11", "test12"},
			Outputs:       []string{"string"},
			Code:          []string{"return \"\""},
		},
		template.Method{
			Name:          "Test2",
			ArgumentsType: []string{"string", "string"},
			ArgumentsName: []string{"test21", "test22"},
			Outputs:       []string{"string"},
			Code:          []string{"return \"\""},
		},
	}

	functions := &template.Functions{
		template.Function{
			Name:          "Test1",
			ArgumentsType: []string{"string", "string"},
			ArgumentsName: []string{"test11", "test12"},
			Outputs:       []string{"string"},
			Code:          []string{"return \"\""},
		},
		template.Function{
			Name:          "Test2",
			ArgumentsType: []string{"string", "string"},
			ArgumentsName: []string{"test21", "test22"},
			Outputs:       []string{"string"},
			Code:          []string{"return \"\""},
		},
	}

	fmt.Println(string(template.GetInterfaceTemplate(layer, name, separator, imports, methods)))
	fmt.Println(string(template.GetRealisationInterfaceTemplate(layer, name, application, kind, separator, imports, methods, functions)))
}
