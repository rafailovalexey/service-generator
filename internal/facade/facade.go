package facade

import (
	"github.com/rafailovalexey/service-generator/internal/template"
	"github.com/rafailovalexey/service-generator/internal/utils"
	"path"
	"sort"
)

func CreateInterface(wd string, layer string, name string) error {
	kind := "internal"
	extension := "go"

	directory := path.Join(wd, kind, layer)
	filename := utils.GetFilename(layer, extension)
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
	extension := "go"

	directory := path.Join(wd, kind, layer, name)
	filename := utils.GetFilename(layer, extension)
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
	kind := "internal"
	extension := "go"

	directory := path.Join(wd, kind, layer, name)
	filename := utils.GetFilename(layer, extension)
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
	kind := "internal"
	request := "request"
	extension := "go"

	directory := path.Join(wd, kind, layer, name, request)
	filename := utils.GetFilename(request, extension)
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
	kind := "internal"
	response := "response"
	extension := "go"

	directory := path.Join(wd, kind, layer, name, response)
	filename := utils.GetFilename(response, extension)
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
	extension := "go"

	directory := path.Join(wd, kind, layer, name)
	filename := utils.GetFilename(layer, extension)
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
	extension := "go"

	directory := path.Join(wd, kind, layer)
	filename := utils.GetFilename(layer, extension)
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
	kind := "internal"
	layer := "implementation"
	extension := "go"

	directory := path.Join(wd, kind, layer, name)
	filename := utils.GetFilename(layer, extension)
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

func CreateReadme(wd string) error {
	name := "README"
	extension := "md"

	directory := path.Join(wd)
	filename := utils.GetFilename(name, extension)
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
	name := ".gitignore"
	extension := ""

	directory := path.Join(wd)
	filename := utils.GetFilename(name, extension)
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

func CreateExampleEnvironment(wd string) error {
	name := ".example"
	extension := "env"

	directory := path.Join(wd)
	filename := utils.GetFilename(name, extension)
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

	data := template.GetExampleEnvironmentTemplate()

	err = utils.SetFileData(filepath, data)

	if err != nil {
		return err
	}

	return nil
}

func CreateGrpcMicroserviceMakefile(wd string) error {
	name := "Makefile"
	extension := ""

	directory := path.Join(wd)
	filename := utils.GetFilename(name, extension)
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

	data := template.GetGrpcMicroserviceMakefileTemplate()

	err = utils.SetFileData(filepath, data)

	if err != nil {
		return err
	}

	return nil
}

func CreateDefaultMicroserviceMakefile(wd string) error {
	name := "Makefile"
	extension := ""

	directory := path.Join(wd)
	filename := utils.GetFilename(name, extension)
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
	name := ".dockerignore"
	extension := ""

	directory := path.Join(wd)
	filename := utils.GetFilename(name, extension)
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

func CreateDockerWithPort(wd string) error {
	name := "application"
	extension := "dockerfile"

	directory := path.Join(wd)
	filename := utils.GetFilename(name, extension)
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

func CreateDockerWithoutPort(wd string) error {
	name := "application"
	extension := "dockerfile"

	directory := path.Join(wd)
	filename := utils.GetFilename(name, extension)
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
	kind := "bin"
	name := "grpc-generate"
	extension := "sh"

	directory := path.Join(wd, kind)
	filename := utils.GetFilename(name, extension)
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
	kind := "bin"
	name := "mock-generate"
	extension := "sh"

	directory := path.Join(wd, kind)
	filename := utils.GetFilename(name, extension)
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
	kind := "cmd"
	subfolder := "grpc_server"
	folder := "interceptor"
	name := "logging"
	extension := "go"

	directory := path.Join(wd, kind, subfolder, folder)
	filename := utils.GetFilename(name, extension)
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
	kind := "cmd"
	subfolder := "grpc_server"
	folder := "interceptor"
	name := "tracing"
	extension := "go"

	directory := path.Join(wd, kind, subfolder, folder)
	filename := utils.GetFilename(name, extension)
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
	kind := "cmd"
	subfolder := "grpc_server"
	folder := "middleware"
	name := "authentication"
	extension := "go"

	directory := path.Join(wd, kind, subfolder, folder)
	filename := utils.GetFilename(name, extension)
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

func CreateGrpcServer(wd string, module string) error {
	kind := "cmd"
	subfolder := "grpc_server"
	name := "grpc_server"
	extension := "go"

	directory := path.Join(wd, kind, subfolder)
	filename := utils.GetFilename(name, extension)
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

	data := template.GetGrpcServerTemplate(module)

	err = utils.SetFileData(filepath, data)

	if err != nil {
		return err
	}

	return nil
}

func CreateHttpLoggingInterceptor(wd string) error {
	kind := "cmd"
	subfolder := "http_server"
	folder := "interceptor"
	name := "logging"
	extension := "go"

	directory := path.Join(wd, kind, subfolder, folder)
	filename := utils.GetFilename(name, extension)
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

func CreateHttpAuthenticationMiddleware(wd string) error {
	kind := "cmd"
	subfolder := "http_server"
	folder := "middleware"
	name := "authentication"
	extension := "go"

	directory := path.Join(wd, kind, subfolder, folder)
	filename := utils.GetFilename(name, extension)
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

	data := template.GetHttpAuthenticationMiddlewareTemplate()

	err = utils.SetFileData(filepath, data)

	if err != nil {
		return err
	}

	return nil
}

func CreateHttpCorsMiddleware(wd string) error {
	kind := "cmd"
	subfolder := "http_server"
	folder := "middleware"
	name := "cors"
	extension := "go"

	directory := path.Join(wd, kind, subfolder, folder)
	filename := utils.GetFilename(name, extension)
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
	kind := "cmd"
	subfolder := "http_server"
	folder := "middleware"
	name := "chain"
	extension := "go"

	directory := path.Join(wd, kind, subfolder, folder)
	filename := utils.GetFilename(name, extension)
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

func CreateHttpServer(wd string, module string) error {
	kind := "cmd"
	folder := "http_server"
	name := "http_server"
	extension := "go"

	directory := path.Join(wd, kind, folder)
	filename := utils.GetFilename(name, extension)
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

	data := template.GetHttpServerTemplate(module)

	err = utils.SetFileData(filepath, data)

	if err != nil {
		return err
	}

	return nil
}

func CreateApplication(wd string, module string, application string, name string, implementing string) error {
	kind := "cmd"
	extension := "go"

	directory := path.Join(wd, kind, application)
	filename := utils.GetFilename(application, extension)
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
	kind := "cmd"
	folder := "nats_subscriber"
	name := "nats_subscriber"
	extension := "go"

	directory := path.Join(wd, kind, folder)
	filename := utils.GetFilename(name, extension)
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
	kind := "cmd"
	folder := "cron_scheduler"
	name := "cron_scheduler"
	extension := "go"

	directory := path.Join(wd, kind, folder)
	filename := utils.GetFilename(name, extension)
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

func CreateGo(wd string, module string, version string) error {
	name := "go"
	extension := "mod"

	directory := path.Join(wd)
	filename := utils.GetFilename(name, extension)
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
