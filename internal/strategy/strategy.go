package strategy

import "github.com/rafailovalexey/service-generator/internal/facade"

func GenerateDataTransferObject(wd string, layer string, name string) error {
	err := facade.CreateDataTransferObject(wd, layer, name)

	if err != nil {
		return err
	}

	return nil
}

func GenerateRealisation(wd string, module string, layer string, name string) error {
	err := facade.CreateInterface(wd, layer, name)

	if err != nil {
		return err
	}

	err = facade.CreateRealisationInterface(wd, module, layer, name)

	if err != nil {
		return err
	}

	return nil
}

func Generate(wd string, module string, layer string, name string) error {
	err := facade.CreateInterface(wd, layer, name)

	if err != nil {
		return err
	}

	err = facade.CreateRealisationInterface(wd, module, layer, name)

	if err != nil {
		return err
	}

	err = facade.CreateRequestObject(wd, layer, name)

	if err != nil {
		return err
	}

	err = facade.CreateResponseObject(wd, layer, name)

	if err != nil {
		return err
	}

	return nil
}

func GenerateProvider(wd string, module string, name string) error {
	err := facade.CreateProviderInterface(wd, module, name)

	if err != nil {
		return err
	}

	err = facade.CreateProvider(wd, module, name)

	if err != nil {
		return err
	}

	return nil
}

func GenerateImplementation(wd string, name string) error {
	err := facade.CreateImplementation(wd, name)

	if err != nil {
		return err
	}

	return nil
}

func GenerateGrpcApplication(wd string, module string, version string, application string, name string, implementing string) error {
	err := facade.CreateGo(wd, module, version)

	if err != nil {
		return err
	}

	err = facade.CreateReadme(wd)

	if err != nil {
		return err
	}

	err = facade.CreateGitIgnore(wd)

	if err != nil {
		return err
	}

	err = facade.CreateEnvironment(wd)

	if err != nil {
		return err
	}

	err = facade.CreateExampleEnvironment(wd, application)

	if err != nil {
		return err
	}

	err = facade.CreateGrpcMicroserviceMakefile(wd, name)

	if err != nil {
		return err
	}

	err = facade.CreateDockerIgnore(wd)

	if err != nil {
		return err
	}

	err = facade.CreateDockerWithPort(wd, application)

	if err != nil {
		return err
	}

	err = facade.CreateGrpcGenerateShellScript(wd)

	if err != nil {
		return err
	}

	err = facade.CreateMockGenerateShellScript(wd)

	if err != nil {
		return err
	}

	err = facade.CreateProto(wd, module, name)

	if err != nil {
		return err
	}

	err = facade.CreateGrpcLoggingInterceptor(wd)

	if err != nil {
		return err
	}

	err = facade.CreateGrpcTracingInterceptor(wd)

	if err != nil {
		return err
	}

	err = facade.CreateGrpcAuthenticationMiddleware(wd)

	if err != nil {
		return err
	}

	err = facade.CreateGrpcServer(wd, module, name)

	if err != nil {
		return err
	}

	err = facade.CreateImplementation(wd, name)

	if err != nil {
		return err
	}

	err = facade.CreateProvider(wd, module, name)

	if err != nil {
		return err
	}

	err = facade.CreateProviderInterface(wd, module, name)

	if err != nil {
		return err
	}

	err = facade.CreateApplication(wd, module, application, name, implementing)

	if err != nil {
		return err
	}

	err = facade.CreateMain(wd, module, application)

	if err != nil {
		return err
	}

	return nil
}

func GenerateHttpApplication(wd string, module string, version string, application string, name string, implementing string) error {
	err := facade.CreateGo(wd, module, version)

	if err != nil {
		return err
	}

	err = facade.CreateReadme(wd)

	if err != nil {
		return err
	}

	err = facade.CreateGitIgnore(wd)

	if err != nil {
		return err
	}

	err = facade.CreateEnvironment(wd)

	if err != nil {
		return err
	}

	err = facade.CreateExampleEnvironment(wd, application)

	if err != nil {
		return err
	}

	err = facade.CreateDefaultMicroserviceMakefile(wd)

	if err != nil {
		return err
	}

	err = facade.CreateDockerIgnore(wd)

	if err != nil {
		return err
	}

	err = facade.CreateDockerWithPort(wd, application)

	if err != nil {
		return err
	}

	err = facade.CreateMockGenerateShellScript(wd)

	if err != nil {
		return err
	}

	err = facade.CreateUtilsConvertError(wd)

	if err != nil {
		return err
	}

	err = facade.CreateUtilsResponse(wd)

	if err != nil {
		return err
	}

	err = facade.CreateHttpLoggingInterceptor(wd)

	if err != nil {
		return err
	}

	err = facade.CreateHttpAuthenticationMiddleware(wd, module)

	if err != nil {
		return err
	}

	err = facade.CreateHttpCorsMiddleware(wd)

	if err != nil {
		return err
	}

	err = facade.CreateHttpChainMiddleware(wd)

	if err != nil {
		return err
	}

	err = facade.CreateHttpServer(wd, module, name)

	if err != nil {
		return err
	}

	err = facade.CreateHandlerInterface(wd, name)

	if err != nil {
		return err
	}

	err = facade.CreateHandler(wd, module, name)

	if err != nil {
		return err
	}

	err = facade.CreateProvider(wd, module, name)

	if err != nil {
		return err
	}

	err = facade.CreateProviderInterface(wd, module, name)

	if err != nil {
		return err
	}

	err = facade.CreateApplication(wd, module, application, name, implementing)

	if err != nil {
		return err
	}

	err = facade.CreateMain(wd, module, application)

	if err != nil {
		return err
	}

	return nil
}
