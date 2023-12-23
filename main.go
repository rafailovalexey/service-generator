package main

import (
	"fmt"
	"github.com/rafailovalexey/service-generator/internal/template"
	"github.com/rafailovalexey/service-generator/internal/utils"
)

func main() {
	//ctx := context.Background()
	//
	//a, err := application.NewApplication(ctx)
	//
	//if err != nil {
	//	log.Panicf("an error occurred while starting the utils %v\n", err)
	//}
	//
	//a.Run()

	separator := utils.GetSeparator()

	fmt.Printf("%s\n", string(template.GetReadmeTemplate(separator)))
	fmt.Printf("%s\n", string(template.GetGitIgnoreTemplate(separator)))
	fmt.Printf("%s\n", string(template.GetExampleEnvironmentTemplate(separator)))
	fmt.Printf("%s\n", string(template.GetGrpcMicroserviceMakefileTemplate(separator)))
	fmt.Printf("%s\n", string(template.GetDefaultMicroserviceMakefileTemplate(separator)))
	fmt.Printf("%s\n", string(template.GetDockerIgnoreTemplate(separator)))
	fmt.Printf("%s\n", string(template.GetDockerWithPortTemplate(separator)))
	fmt.Printf("%s\n", string(template.GetDockerWithoutPortTemplate(separator)))
	fmt.Printf("%s\n", string(template.GetGrpcGenerateShellScriptTemplate(separator)))
	fmt.Printf("%s\n", string(template.GetMockGenerateShellScriptTemplate(separator)))
	fmt.Printf("%s\n", string(template.GetGrpcLoggingInterceptorTemplate(separator)))
	fmt.Printf("%s\n", string(template.GetGrpcTraceCodeInterceptorTemplate(separator)))
	fmt.Printf("%s\n", string(template.GetGrpcAuthenticationMiddlewareTemplate(separator)))
	fmt.Printf("%s\n", string(template.GetGrpcServerTemplate(separator)))
	fmt.Printf("%s\n", string(template.GetHttpLoggingInterceptorTemplate(separator)))
	fmt.Printf("%s\n", string(template.GetHttpAuthenticationMiddlewareTemplate(separator)))
	fmt.Printf("%s\n", string(template.GetHttpCorsMiddlewareTemplate(separator)))
	fmt.Printf("%s\n", string(template.GetHttpChainMiddlewareTemplate(separator)))
	fmt.Printf("%s\n", string(template.GetHttpServerTemplate(separator)))
}
