package facade

import (
	"github.com/rafailovalexey/service-generator/internal/template"
	"github.com/rafailovalexey/service-generator/internal/utils"
	"path"
	"sort"
)

func CreateInterface(wd string, layer string, name string) error {
	directory := path.Join(wd, "internal", layer)
	filename := utils.GetFilename(layer, "go")
	filepath := path.Join(directory, filename)

	isExist := utils.PathIsExist(directory)

	if !isExist {
		err := utils.CreateDirectory(directory)

		if err != nil {
			return err
		}
	}

	err := utils.CreateFile(filepath)

	if err != nil {
		return err
	}

	data := template.GetInterfaceTemplate(layer, name)

	err = utils.SetFileData(filepath, data)

	if err != nil {
		return err
	}

	return nil
}

func CreateRealisationInterface(wd string, module string, layer string, name string) error {
	kind := "internal"

	directory := path.Join(wd, kind, layer, name)
	filename := utils.GetFilename(layer, "go")
	filepath := path.Join(directory, filename)

	isExist := utils.PathIsExist(directory)

	if !isExist {
		err := utils.CreateDirectory(directory)

		if err != nil {
			return err
		}
	}

	err := utils.CreateFile(filepath)

	if err != nil {
		return err
	}

	data := template.GetRealisationInterfaceTemplate(module, kind, layer, name)

	err = utils.SetFileData(filepath, data)

	if err != nil {
		return err
	}

	return nil
}

func CreateDataTransferObject(wd string, layer string, name string) error {
	directory := path.Join(wd, "internal", layer, name)
	filename := utils.GetFilename(layer, "go")
	filepath := path.Join(directory, filename)

	isExist := utils.PathIsExist(directory)

	if !isExist {
		err := utils.CreateDirectory(directory)

		if err != nil {
			return err
		}
	}

	err := utils.CreateFile(filepath)

	if err != nil {
		return err
	}

	data := template.GetDataTransferObjectTemplate(layer, name)

	err = utils.SetFileData(filepath, data)

	if err != nil {
		return err
	}

	return nil
}

func CreateRequestObject(wd string, layer string, name string) error {
	directory := path.Join(wd, "internal", layer, name, "request")
	filename := utils.GetFilename("request", "go")
	filepath := path.Join(directory, filename)

	isExist := utils.PathIsExist(directory)

	if !isExist {
		err := utils.CreateDirectory(directory)

		if err != nil {
			return err
		}
	}

	err := utils.CreateFile(filepath)

	if err != nil {
		return err
	}

	data := template.GetRequestTemplate(name)

	err = utils.SetFileData(filepath, data)

	if err != nil {
		return err
	}

	return nil
}

func CreateResponseObject(wd string, layer string, name string) error {
	directory := path.Join(wd, "internal", layer, name, "response")
	filename := utils.GetFilename("response", "go")
	filepath := path.Join(directory, filename)

	isExist := utils.PathIsExist(directory)

	if !isExist {
		err := utils.CreateDirectory(directory)

		if err != nil {
			return err
		}
	}

	err := utils.CreateFile(filepath)

	if err != nil {
		return err
	}

	data := template.GetResponseTemplate(name)

	err = utils.SetFileData(filepath, data)

	if err != nil {
		return err
	}

	return nil
}

func CreateProvider(wd string, module string, name string) error {
	layer := "provider"
	kind := "internal"

	directory := path.Join(wd, kind, layer, name)
	filename := utils.GetFilename(layer, "go")
	filepath := path.Join(directory, filename)

	isExist := utils.PathIsExist(directory)

	if !isExist {
		err := utils.CreateDirectory(directory)

		if err != nil {
			return err
		}
	}

	err := utils.CreateFile(filepath)

	if err != nil {
		return err
	}

	available := make(map[string]struct{}, 10)

	available["api"] = struct{}{}
	available["controller"] = struct{}{}
	available["implementation"] = struct{}{}
	available["handler"] = struct{}{}
	available["client"] = struct{}{}
	available["validation"] = struct{}{}
	available["converter"] = struct{}{}
	available["repository"] = struct{}{}
	available["service"] = struct{}{}

	directory = path.Join(wd, kind)
	directories, err := utils.GetDirectories(directory)

	if err != nil {
		return err
	}

	layers := make([]string, 0, 10)

	for _, d := range directories {
		if _, isExist = available[d]; isExist {
			layers = append(layers, d)
		}
	}

	sort.Strings(layers)

	data := template.GetProviderRealisationTemplate(module, kind, layer, name, layers)

	err = utils.SetFileData(filepath, data)

	if err != nil {
		return err
	}

	return nil
}

func CreateProviderInterface(wd string, module string, name string) error {
	layer := "provider"
	kind := "internal"

	directory := path.Join(wd, kind, layer)
	filename := utils.GetFilename(layer, "go")
	filepath := path.Join(directory, filename)

	isExist := utils.PathIsExist(directory)

	if !isExist {
		err := utils.CreateDirectory(directory)

		if err != nil {
			return err
		}
	}

	err := utils.CreateFile(filepath)

	if err != nil {
		return err
	}

	available := make(map[string]struct{}, 10)

	available["api"] = struct{}{}
	available["controller"] = struct{}{}
	available["implementation"] = struct{}{}
	available["handler"] = struct{}{}
	available["client"] = struct{}{}
	available["validation"] = struct{}{}
	available["converter"] = struct{}{}
	available["repository"] = struct{}{}
	available["service"] = struct{}{}

	directory = path.Join(wd, kind)
	directories, err := utils.GetDirectories(directory)

	if err != nil {
		return err
	}

	layers := make([]string, 0, 10)

	for _, d := range directories {
		if _, isExist = available[d]; isExist {
			layers = append(layers, d)
		}
	}

	sort.Strings(layers)

	data := template.GetProviderInterfaceTemplate(module, kind, layer, name, layers)

	err = utils.SetFileData(filepath, data)

	if err != nil {
		return err
	}

	return nil
}

func CreateImplementation(wd string, name string) error {
	layer := "implementation"

	directory := path.Join(wd, "internal", layer, name)
	filename := utils.GetFilename(layer, "go")
	filepath := path.Join(directory, filename)

	isExist := utils.PathIsExist(directory)

	if !isExist {
		err := utils.CreateDirectory(directory)

		if err != nil {
			return err
		}
	}

	err := utils.CreateFile(filepath)

	if err != nil {
		return err
	}

	data := template.GetImplementationRealisationTemplate(layer, name)

	err = utils.SetFileData(filepath, data)

	if err != nil {
		return err
	}

	return nil
}

func CreateHandler(wd string, module string, name string) error {
	layer := "handler"
	kind := "internal"

	directory := path.Join(wd, kind, layer, name)
	filename := utils.GetFilename(layer, "go")
	filepath := path.Join(directory, filename)

	isExist := utils.PathIsExist(directory)

	if !isExist {
		err := utils.CreateDirectory(directory)

		if err != nil {
			return err
		}
	}

	err := utils.CreateFile(filepath)

	if err != nil {
		return err
	}

	data := template.GetHandlerRealisationTemplate(module, kind, layer, name)

	err = utils.SetFileData(filepath, data)

	if err != nil {
		return err
	}

	return nil
}

func CreateHandlerInterface(wd string, name string) error {
	layer := "handler"
	kind := "internal"

	directory := path.Join(wd, kind, layer)
	filename := utils.GetFilename(layer, "go")
	filepath := path.Join(directory, filename)

	isExist := utils.PathIsExist(directory)

	if !isExist {
		err := utils.CreateDirectory(directory)

		if err != nil {
			return err
		}
	}

	err := utils.CreateFile(filepath)

	if err != nil {
		return err
	}

	data := template.GetHandlerInterfaceTemplate(layer, name)

	err = utils.SetFileData(filepath, data)

	if err != nil {
		return err
	}

	return nil
}

func CreateReadme(wd string) error {
	directory := path.Join(wd)
	filename := utils.GetFilename("README", "md")
	filepath := path.Join(directory, filename)

	isExist := utils.PathIsExist(directory)

	if !isExist {
		err := utils.CreateDirectory(directory)

		if err != nil {
			return err
		}
	}

	err := utils.CreateFile(filepath)

	if err != nil {
		return err
	}

	data := template.GetReadmeTemplate()

	err = utils.SetFileData(filepath, data)

	if err != nil {
		return err
	}

	return nil
}

func CreateGitIgnore(wd string) error {
	directory := path.Join(wd)
	filename := utils.GetFilename(".gitignore", "")
	filepath := path.Join(directory, filename)

	isExist := utils.PathIsExist(directory)

	if !isExist {
		err := utils.CreateDirectory(directory)

		if err != nil {
			return err
		}
	}

	err := utils.CreateFile(filepath)

	if err != nil {
		return err
	}

	data := template.GetGitIgnoreTemplate()

	err = utils.SetFileData(filepath, data)

	if err != nil {
		return err
	}

	return nil
}

func CreateEnvironment(wd string) error {
	directory := path.Join(wd)
	filename := utils.GetFilename("", "env")
	filepath := path.Join(directory, filename)

	isExist := utils.PathIsExist(directory)

	if !isExist {
		err := utils.CreateDirectory(directory)

		if err != nil {
			return err
		}
	}

	err := utils.CreateFile(filepath)

	if err != nil {
		return err
	}

	data := template.GetEnvironmentTemplate()

	err = utils.SetFileData(filepath, data)

	if err != nil {
		return err
	}

	return nil
}

func CreateExampleEnvironment(wd string, application string) error {
	directory := path.Join(wd)
	filename := utils.GetFilename(".example", "env")
	filepath := path.Join(directory, filename)

	isExist := utils.PathIsExist(directory)

	if !isExist {
		err := utils.CreateDirectory(directory)

		if err != nil {
			return err
		}
	}

	err := utils.CreateFile(filepath)

	if err != nil {
		return err
	}

	data := template.GetExampleEnvironmentTemplate(application)

	err = utils.SetFileData(filepath, data)

	if err != nil {
		return err
	}

	return nil
}

func CreateGrpcMicroserviceMakefile(wd string, name string) error {
	directory := path.Join(wd)
	filename := utils.GetFilename("Makefile", "")
	filepath := path.Join(directory, filename)

	isExist := utils.PathIsExist(directory)

	if !isExist {
		err := utils.CreateDirectory(directory)

		if err != nil {
			return err
		}
	}

	err := utils.CreateFile(filepath)

	if err != nil {
		return err
	}

	data := template.GetGrpcMicroserviceMakefileTemplate(name)

	err = utils.SetFileData(filepath, data)

	if err != nil {
		return err
	}

	return nil
}

func CreateDefaultMicroserviceMakefile(wd string) error {
	directory := path.Join(wd)
	filename := utils.GetFilename("Makefile", "")
	filepath := path.Join(directory, filename)

	isExist := utils.PathIsExist(directory)

	if !isExist {
		err := utils.CreateDirectory(directory)

		if err != nil {
			return err
		}
	}

	err := utils.CreateFile(filepath)

	if err != nil {
		return err
	}

	data := template.GetDefaultMicroserviceMakefileTemplate()

	err = utils.SetFileData(filepath, data)

	if err != nil {
		return err
	}

	return nil
}

func CreateDockerIgnore(wd string) error {
	directory := path.Join(wd)
	filename := utils.GetFilename(".dockerignore", "")
	filepath := path.Join(directory, filename)

	isExist := utils.PathIsExist(directory)

	if !isExist {
		err := utils.CreateDirectory(directory)

		if err != nil {
			return err
		}
	}

	err := utils.CreateFile(filepath)

	if err != nil {
		return err
	}

	data := template.GetDockerIgnoreTemplate()

	err = utils.SetFileData(filepath, data)

	if err != nil {
		return err
	}

	return nil
}

func CreateDockerWithPort(wd string, application string) error {
	directory := path.Join(wd)
	filename := utils.GetFilename(application, "dockerfile")
	filepath := path.Join(directory, filename)

	isExist := utils.PathIsExist(directory)

	if !isExist {
		err := utils.CreateDirectory(directory)

		if err != nil {
			return err
		}
	}

	err := utils.CreateFile(filepath)

	if err != nil {
		return err
	}

	data := template.GetDockerWithPortTemplate()

	err = utils.SetFileData(filepath, data)

	if err != nil {
		return err
	}

	return nil
}

func CreateDockerWithoutPort(wd string, application string) error {
	directory := path.Join(wd)
	filename := utils.GetFilename(application, "dockerfile")
	filepath := path.Join(directory, filename)

	isExist := utils.PathIsExist(directory)

	if !isExist {
		err := utils.CreateDirectory(directory)

		if err != nil {
			return err
		}
	}

	err := utils.CreateFile(filepath)

	if err != nil {
		return err
	}

	data := template.GetDockerWithoutPortTemplate()

	err = utils.SetFileData(filepath, data)

	if err != nil {
		return err
	}

	return nil
}

func CreateGrpcGenerateShellScript(wd string) error {
	directory := path.Join(wd, "bin")
	filename := utils.GetFilename("grpc-generate", "sh")
	filepath := path.Join(directory, filename)

	isExist := utils.PathIsExist(directory)

	if !isExist {
		err := utils.CreateDirectory(directory)

		if err != nil {
			return err
		}
	}

	err := utils.CreateFile(filepath)

	if err != nil {
		return err
	}

	data := template.GetGrpcGenerateShellScriptTemplate()

	err = utils.SetFileData(filepath, data)

	if err != nil {
		return err
	}

	return nil
}

func CreateMockGenerateShellScript(wd string) error {
	directory := path.Join(wd, "bin")
	filename := utils.GetFilename("mock-generate", "sh")
	filepath := path.Join(directory, filename)

	isExist := utils.PathIsExist(directory)

	if !isExist {
		err := utils.CreateDirectory(directory)

		if err != nil {
			return err
		}
	}

	err := utils.CreateFile(filepath)

	if err != nil {
		return err
	}

	data := template.GetMockGenerateShellScriptTemplate()

	err = utils.SetFileData(filepath, data)

	if err != nil {
		return err
	}

	return nil
}

func CreateGrpcLoggingInterceptor(wd string) error {
	directory := path.Join(wd, "cmd", "grpc_server", "interceptor")
	filename := utils.GetFilename("logging", "go")
	filepath := path.Join(directory, filename)

	isExist := utils.PathIsExist(directory)

	if !isExist {
		err := utils.CreateDirectory(directory)

		if err != nil {
			return err
		}
	}

	err := utils.CreateFile(filepath)

	if err != nil {
		return err
	}

	data := template.GetGrpcLoggingInterceptorTemplate()

	err = utils.SetFileData(filepath, data)

	if err != nil {
		return err
	}

	return nil
}

func CreateGrpcTracingInterceptor(wd string) error {
	directory := path.Join(wd, "cmd", "grpc_server", "interceptor")
	filename := utils.GetFilename("tracing", "go")
	filepath := path.Join(directory, filename)

	isExist := utils.PathIsExist(directory)

	if !isExist {
		err := utils.CreateDirectory(directory)

		if err != nil {
			return err
		}
	}

	err := utils.CreateFile(filepath)

	if err != nil {
		return err
	}

	data := template.GetGrpcTracingInterceptorTemplate()

	err = utils.SetFileData(filepath, data)

	if err != nil {
		return err
	}

	return nil
}

func CreateGrpcAuthenticationMiddleware(wd string) error {
	directory := path.Join(wd, "cmd", "grpc_server", "middleware")
	filename := utils.GetFilename("authentication", "go")
	filepath := path.Join(directory, filename)

	isExist := utils.PathIsExist(directory)

	if !isExist {
		err := utils.CreateDirectory(directory)

		if err != nil {
			return err
		}
	}

	err := utils.CreateFile(filepath)

	if err != nil {
		return err
	}

	data := template.GetGrpcAuthenticationMiddlewareTemplate()

	err = utils.SetFileData(filepath, data)

	if err != nil {
		return err
	}

	return nil
}

func CreateGrpcServer(wd string, module string, name string) error {
	directory := path.Join(wd, "cmd", "grpc_server")
	filename := utils.GetFilename("grpc_server", "go")
	filepath := path.Join(directory, filename)

	isExist := utils.PathIsExist(directory)

	if !isExist {
		err := utils.CreateDirectory(directory)

		if err != nil {
			return err
		}
	}

	err := utils.CreateFile(filepath)

	if err != nil {
		return err
	}

	data := template.GetGrpcServerTemplate(module, name)

	err = utils.SetFileData(filepath, data)

	if err != nil {
		return err
	}

	return nil
}

func CreateHttpLoggingInterceptor(wd string) error {
	directory := path.Join(wd, "cmd", "http_server", "interceptor")
	filename := utils.GetFilename("logging", "go")
	filepath := path.Join(directory, filename)

	isExist := utils.PathIsExist(directory)

	if !isExist {
		err := utils.CreateDirectory(directory)

		if err != nil {
			return err
		}
	}

	err := utils.CreateFile(filepath)

	if err != nil {
		return err
	}

	data := template.GetHttpLoggingInterceptorTemplate()

	err = utils.SetFileData(filepath, data)

	if err != nil {
		return err
	}

	return nil
}

func CreateHttpAuthenticationMiddleware(wd string, module string) error {
	directory := path.Join(wd, "cmd", "http_server", "middleware")
	filename := utils.GetFilename("authentication", "go")
	filepath := path.Join(directory, filename)

	isExist := utils.PathIsExist(directory)

	if !isExist {
		err := utils.CreateDirectory(directory)

		if err != nil {
			return err
		}
	}

	err := utils.CreateFile(filepath)

	if err != nil {
		return err
	}

	data := template.GetHttpAuthenticationMiddlewareTemplate(module)

	err = utils.SetFileData(filepath, data)

	if err != nil {
		return err
	}

	return nil
}

func CreateHttpCorsMiddleware(wd string) error {
	directory := path.Join(wd, "cmd", "http_server", "middleware")
	filename := utils.GetFilename("cors", "go")
	filepath := path.Join(directory, filename)

	isExist := utils.PathIsExist(directory)

	if !isExist {
		err := utils.CreateDirectory(directory)

		if err != nil {
			return err
		}
	}

	err := utils.CreateFile(filepath)

	if err != nil {
		return err
	}

	data := template.GetHttpCorsMiddlewareTemplate()

	err = utils.SetFileData(filepath, data)

	if err != nil {
		return err
	}

	return nil
}

func CreateHttpChainMiddleware(wd string) error {
	directory := path.Join(wd, "cmd", "http_server", "middleware")
	filename := utils.GetFilename("chain", "go")
	filepath := path.Join(directory, filename)

	isExist := utils.PathIsExist(directory)

	if !isExist {
		err := utils.CreateDirectory(directory)

		if err != nil {
			return err
		}
	}

	err := utils.CreateFile(filepath)

	if err != nil {
		return err
	}

	data := template.GetHttpChainMiddlewareTemplate()

	err = utils.SetFileData(filepath, data)

	if err != nil {
		return err
	}

	return nil
}

func CreateHttpServer(wd string, module string, name string) error {
	directory := path.Join(wd, "cmd", "http_server")
	filename := utils.GetFilename("http_server", "go")
	filepath := path.Join(directory, filename)

	isExist := utils.PathIsExist(directory)

	if !isExist {
		err := utils.CreateDirectory(directory)

		if err != nil {
			return err
		}
	}

	err := utils.CreateFile(filepath)

	if err != nil {
		return err
	}

	data := template.GetHttpServerTemplate(module, name)

	err = utils.SetFileData(filepath, data)

	if err != nil {
		return err
	}

	return nil
}

func CreateApplication(wd string, module string, application string, name string, implementing string) error {
	directory := path.Join(wd, "cmd", application)
	filename := utils.GetFilename(application, "go")
	filepath := path.Join(directory, filename)

	isExist := utils.PathIsExist(directory)

	if !isExist {
		err := utils.CreateDirectory(directory)

		if err != nil {
			return err
		}
	}

	err := utils.CreateFile(filepath)

	if err != nil {
		return err
	}

	data := template.GetApplicationTemplate(module, application, name, implementing)

	err = utils.SetFileData(filepath, data)

	if err != nil {
		return err
	}

	return nil
}

func CreateNatsSubscriber(wd string) error {
	directory := path.Join(wd, "cmd", "nats_subscribe")
	filename := utils.GetFilename("nats_subscribe", "go")
	filepath := path.Join(directory, filename)

	isExist := utils.PathIsExist(directory)

	if !isExist {
		err := utils.CreateDirectory(directory)

		if err != nil {
			return err
		}
	}

	err := utils.CreateFile(filepath)

	if err != nil {
		return err
	}

	data := template.GetNatsSubscriberTemplate()

	err = utils.SetFileData(filepath, data)

	if err != nil {
		return err
	}

	return nil
}

func CreateCronScheduler(wd string) error {
	directory := path.Join(wd, "cmd", "cron_scheduler")
	filename := utils.GetFilename("cron_scheduler", "go")
	filepath := path.Join(directory, filename)

	isExist := utils.PathIsExist(directory)

	if !isExist {
		err := utils.CreateDirectory(directory)

		if err != nil {
			return err
		}
	}

	err := utils.CreateFile(filepath)

	if err != nil {
		return err
	}

	data := template.GetCronSchedulerTemplate()

	err = utils.SetFileData(filepath, data)

	if err != nil {
		return err
	}

	return nil
}

func CreateProto(wd string, module string, name string) error {
	kind := "api"

	directory := path.Join(wd, kind, name+"_"+"v1")
	filename := utils.GetFilename(name, "proto")
	filepath := path.Join(directory, filename)

	isExist := utils.PathIsExist(directory)

	if !isExist {
		err := utils.CreateDirectory(directory)

		if err != nil {
			return err
		}
	}

	err := utils.CreateFile(filepath)

	if err != nil {
		return err
	}

	data := template.GetProtoTemplate(module, kind, name)

	err = utils.SetFileData(filepath, data)

	if err != nil {
		return err
	}

	return nil
}

func CreateMain(wd string, module string, application string) error {
	directory := path.Join(wd)
	filename := utils.GetFilename("main", "go")
	filepath := path.Join(directory, filename)

	isExist := utils.PathIsExist(directory)

	if !isExist {
		err := utils.CreateDirectory(directory)

		if err != nil {
			return err
		}
	}

	err := utils.CreateFile(filepath)

	if err != nil {
		return err
	}

	data := template.GetMainTemplate(module, application)

	err = utils.SetFileData(filepath, data)

	if err != nil {
		return err
	}

	return nil
}

func CreateGo(wd string, module string, version string) error {
	directory := path.Join(wd)
	filename := utils.GetFilename("go", "mod")
	filepath := path.Join(directory, filename)

	isExist := utils.PathIsExist(directory)

	if !isExist {
		err := utils.CreateDirectory(directory)

		if err != nil {
			return err
		}
	}

	err := utils.CreateFile(filepath)

	if err != nil {
		return err
	}

	data := template.GetGoTemplate(module, version)

	err = utils.SetFileData(filepath, data)

	if err != nil {
		return err
	}

	return nil
}

func CreateUtilsConvertError(wd string) error {
	directory := path.Join(wd, "utils")
	filename := utils.GetFilename("error", "go")
	filepath := path.Join(directory, filename)

	isExist := utils.PathIsExist(directory)

	if !isExist {
		err := utils.CreateDirectory(directory)

		if err != nil {
			return err
		}
	}

	err := utils.CreateFile(filepath)

	if err != nil {
		return err
	}

	data := template.GetUtilsConvertErrorTemplate()

	err = utils.SetFileData(filepath, data)

	if err != nil {
		return err
	}

	return nil
}

func CreateUtilsResponse(wd string) error {
	directory := path.Join(wd, "utils")
	filename := utils.GetFilename("response", "go")
	filepath := path.Join(directory, filename)

	isExist := utils.PathIsExist(directory)

	if !isExist {
		err := utils.CreateDirectory(directory)

		if err != nil {
			return err
		}
	}

	err := utils.CreateFile(filepath)

	if err != nil {
		return err
	}

	data := template.GetUtilsResponseTemplate()

	err = utils.SetFileData(filepath, data)

	if err != nil {
		return err
	}

	return nil
}
