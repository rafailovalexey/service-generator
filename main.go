package main

import (
	"context"
	"github.com/rafailovalexey/service-generator/cmd/application"
	"log"
)

func main() {
	ctx := context.Background()

	a, err := application.NewApplication(ctx)

	if err != nil {
		log.Panicf("an error occurred while starting the utils %v\n", err)
	}

	a.Run()

	//layer := "provider"
	//name := "employees"
	//application, _ := utils.GetApplicationName()
	//kind := "internal"
	//separator := utils.GetSeparator()

	//imports := &template.Imports{
	//	"net/http",
	//}
	//
	//methods := &template.Methods{
	//	template.Method{
	//		Name:          "Test1",
	//		ArgumentsType: []string{"string", "string"},
	//		ArgumentsName: []string{"test11", "test12"},
	//		Outputs:       []string{"string"},
	//		Code:          []string{"return \"\""},
	//	},
	//	template.Method{
	//		Name:          "Test2",
	//		ArgumentsType: []string{"string", "string"},
	//		ArgumentsName: []string{"test21", "test22"},
	//		Outputs:       []string{"string"},
	//		Code:          []string{"return \"\""},
	//	},
	//}
	//
	//functions := &template.Functions{
	//	template.Function{
	//		Name:          "Test1",
	//		ArgumentsType: []string{"string", "string"},
	//		ArgumentsName: []string{"test11", "test12"},
	//		Outputs:       []string{"string"},
	//		Code:          []string{"return \"\""},
	//	},
	//	template.Function{
	//		Name:          "Test2",
	//		ArgumentsType: []string{"string", "string"},
	//		ArgumentsName: []string{"test21", "test22"},
	//		Outputs:       []string{"string"},
	//		Code:          []string{"return \"\""},
	//	},
	//}

	//layers := []string{
	//	"api",
	//	"controller",
	//	"validation",
	//	"converter",
	//	"service",
	//	"repository",
	//}

	//fmt.Printf("%s", string(template.GetProviderInterfaceTemplate(application, separator, kind, layers, layer, name)))
	//fmt.Printf("%s", string(template.GetProviderRealisationTemplate(application, separator, kind, layers, layer, name)))

	//imports := &template.Imports{}
	//methods := &template.Methods{}
	//functions := &template.Functions{}
	//
	//_ = facade.CreateInterface("repository", "employees", imports, methods)
	//_ = facade.CreateRealisationInterface("repository", "employees", imports, methods, functions)
	//
	//_ = facade.CreateProviderInterface("provider", "employees")
	//_ = facade.CreateProvider("provider", "employees")
}
