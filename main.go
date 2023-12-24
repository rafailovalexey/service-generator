package main

import "github.com/rafailovalexey/service-generator/internal/facade"

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

	_ = facade.CreateReadme()
	_ = facade.CreateGitIgnore()
	_ = facade.CreateExampleEnvironment()
	_ = facade.CreateGrpcMicroserviceMakefile()
	//_ = facade.CreateDefaultMicroserviceMakefile()
	_ = facade.CreateDockerIgnore()
	_ = facade.CreateDockerWithPort()
	//_ = facade.CreateDockerWithoutPort()
	_ = facade.CreateGrpcGenerateShellScript()
	_ = facade.CreateMockGenerateShellScript()
	_ = facade.CreateGrpcLoggingInterceptor()
	_ = facade.CreateGrpcTracingInterceptor()
	_ = facade.CreateGrpcAuthenticationMiddleware()
	_ = facade.CreateGrpcServer()
	_ = facade.CreateHttpLoggingInterceptor()
	_ = facade.CreateHttpAuthenticationMiddleware()
	_ = facade.CreateHttpCorsMiddleware()
	_ = facade.CreateHttpChainMiddleware()
	_ = facade.CreateHttpServer()
}
