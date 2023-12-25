package main

import (
	"github.com/rafailovalexey/service-generator/internal/facade"
	"github.com/rafailovalexey/service-generator/internal/utils"
	"path"
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

	wd, _ := utils.GetWorkDirectory()
	wd = path.Join(wd, "test12345")

	_ = facade.CreateReadme(wd)
	_ = facade.CreateGitIgnore(wd)
	_ = facade.CreateExampleEnvironment(wd)
	_ = facade.CreateGrpcMicroserviceMakefile(wd)
	_ = facade.CreateDockerIgnore(wd)
	_ = facade.CreateDockerWithPort(wd)
	_ = facade.CreateGrpcGenerateShellScript(wd)
	_ = facade.CreateMockGenerateShellScript(wd)
	_ = facade.CreateGrpcLoggingInterceptor(wd)
	_ = facade.CreateGrpcTracingInterceptor(wd)
	_ = facade.CreateGrpcAuthenticationMiddleware(wd)
	_ = facade.CreateGrpcServer(wd)
	_ = facade.CreateProvider(wd, "employees")
	_ = facade.CreateProviderInterface(wd, "employees")
	_ = facade.CreateImplementation(wd, "employees")
	_ = facade.CreateApplication(wd, "application", "employees", "grpc_server")

	//_ = facade.CreateDefaultMicroserviceMakefile()
	//_ = facade.CreateDockerWithoutPort()
	//_ = facade.CreateHttpLoggingInterceptor()
	//_ = facade.CreateHttpAuthenticationMiddleware()
	//_ = facade.CreateHttpCorsMiddleware()
	//_ = facade.CreateHttpChainMiddleware()
	//_ = facade.CreateHttpServer()
}
