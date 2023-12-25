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
	wd = path.Join(wd, "test")

	module := "github.com/rafailovalexey/service-test"
	version := "1.19"

	name := "employees"

	application := "application"
	//application := "subscriber"
	//application := "cron"

	implementing := "grpc_server"
	//implementing := "http_server"
	//implementing := "cron_scheduler"
	//implementing := "nats_subscriber"

	_ = facade.CreateGo(wd, module, version)
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
	_ = facade.CreateGrpcServer(wd, module)
	_ = facade.CreateProvider(wd, module, name)
	_ = facade.CreateProviderInterface(wd, module, name)
	_ = facade.CreateImplementation(wd, name)
	_ = facade.CreateApplication(wd, module, application, name, implementing)
}
