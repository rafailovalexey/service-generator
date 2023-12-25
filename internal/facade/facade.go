package facade

import (
	"github.com/rafailovalexey/service-generator/internal/template"
	"github.com/rafailovalexey/service-generator/internal/utils"
	"os"
	"path"
	"sort"
)

func CreateInterface(layer string, name string) error {
	current, err := os.Getwd()

	if err != nil {
		return err
	}

	separator := utils.GetSeparator()

	kind := "internal"
	extension := "go"

	filename := utils.GetFilename(layer, extension)
	directory := path.Join(current, kind, layer)
	filepath := path.Join(directory, filename)

	isExist := utils.PathIsExist(directory)

	if !isExist {
		err = utils.CreateDirectory(directory)

		if err != nil {
			return err
		}
	}

	err = utils.CreateFile(filepath)

	if err != nil {
		return err
	}

	data := template.GetInterfaceTemplate(separator, layer, name)

	err = utils.SetFileData(filepath, data)

	if err != nil {
		return err
	}

	return nil
}

func CreateRealisationInterface(layer string, name string) error {
	current, err := os.Getwd()

	if err != nil {
		return err
	}

	module, err := utils.GetModuleName()

	if err != nil {
		return err
	}

	separator := utils.GetSeparator()

	kind := "internal"
	extension := "go"

	if err != nil {
		return err
	}

	filename := utils.GetFilename(layer, extension)
	directory := path.Join(current, kind, layer, name)
	filepath := path.Join(directory, filename)

	isExist := utils.PathIsExist(directory)

	if !isExist {
		err = utils.CreateDirectory(directory)

		if err != nil {
			return err
		}
	}

	err = utils.CreateFile(filepath)

	if err != nil {
		return err
	}

	data := template.GetRealisationInterfaceTemplate(module, separator, kind, layer, name)

	err = utils.SetFileData(filepath, data)

	if err != nil {
		return err
	}

	return nil
}

func CreateDataTransferObject(layer string, name string) error {
	current, err := os.Getwd()

	if err != nil {
		return err
	}

	separator := utils.GetSeparator()

	kind := "internal"
	extension := "go"

	if err != nil {
		return err
	}

	filename := utils.GetFilename(layer, extension)
	directory := path.Join(current, kind, layer, name)
	filepath := path.Join(directory, filename)

	isExist := utils.PathIsExist(directory)

	if !isExist {
		err = utils.CreateDirectory(directory)

		if err != nil {
			return err
		}
	}

	err = utils.CreateFile(filepath)

	if err != nil {
		return err
	}

	data := template.GetDataTransferObjectTemplate(separator, layer, name)

	err = utils.SetFileData(filepath, data)

	if err != nil {
		return err
	}

	return nil
}

func CreateRequestObject(layer string, name string) error {
	current, err := os.Getwd()

	if err != nil {
		return err
	}

	separator := utils.GetSeparator()

	kind := "internal"
	request := "request"
	extension := "go"

	if err != nil {
		return err
	}

	filename := utils.GetFilename(request, extension)
	directory := path.Join(current, kind, layer, name, request)
	filepath := path.Join(directory, filename)

	isExist := utils.PathIsExist(directory)

	if !isExist {
		err = utils.CreateDirectory(directory)

		if err != nil {
			return err
		}
	}

	err = utils.CreateFile(filepath)

	if err != nil {
		return err
	}

	data := template.GetRequestTemplate(separator, name)

	err = utils.SetFileData(filepath, data)

	if err != nil {
		return err
	}

	return nil
}

func CreateResponseObject(layer string, name string) error {
	current, err := os.Getwd()

	if err != nil {
		return err
	}

	separator := utils.GetSeparator()

	kind := "internal"
	response := "response"
	extension := "go"

	if err != nil {
		return err
	}

	filename := utils.GetFilename(response, extension)
	directory := path.Join(current, kind, layer, name, response)
	filepath := path.Join(directory, filename)

	isExist := utils.PathIsExist(directory)

	if !isExist {
		err = utils.CreateDirectory(directory)

		if err != nil {
			return err
		}
	}

	err = utils.CreateFile(filepath)

	if err != nil {
		return err
	}

	data := template.GetResponseTemplate(separator, name)

	err = utils.SetFileData(filepath, data)

	if err != nil {
		return err
	}

	return nil
}

func CreateProvider(layer string, name string) error {
	current, err := os.Getwd()

	if err != nil {
		return err
	}

	module, err := utils.GetModuleName()

	if err != nil {
		return err
	}

	separator := utils.GetSeparator()

	kind := "internal"
	extension := "go"

	if err != nil {
		return err
	}

	filename := utils.GetFilename(layer, extension)
	directory := path.Join(current, kind, layer, name)
	filepath := path.Join(directory, filename)

	isExist := utils.PathIsExist(directory)

	if !isExist {
		err = utils.CreateDirectory(directory)

		if err != nil {
			return err
		}
	}

	err = utils.CreateFile(filepath)

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

	directory = path.Join(current, kind)
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

	data := template.GetProviderRealisationTemplate(module, separator, kind, layers, layer, name)

	err = utils.SetFileData(filepath, data)

	if err != nil {
		return err
	}

	return nil
}

func CreateProviderInterface(layer string, name string) error {
	current, err := os.Getwd()

	if err != nil {
		return err
	}

	module, err := utils.GetModuleName()

	if err != nil {
		return err
	}

	separator := utils.GetSeparator()

	kind := "internal"
	extension := "go"

	filename := utils.GetFilename(layer, extension)
	directory := path.Join(current, kind, layer)
	filepath := path.Join(directory, filename)

	isExist := utils.PathIsExist(directory)

	if !isExist {
		err = utils.CreateDirectory(directory)

		if err != nil {
			return err
		}
	}

	err = utils.CreateFile(filepath)

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

	directory = path.Join(current, kind)
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

	data := template.GetProviderInterfaceTemplate(module, separator, kind, layers, layer, name)

	err = utils.SetFileData(filepath, data)

	if err != nil {
		return err
	}

	return nil
}

func CreateImplementation(layer string, name string) error {
	current, err := os.Getwd()

	if err != nil {
		return err
	}

	separator := utils.GetSeparator()

	kind := "internal"
	extension := "go"

	if err != nil {
		return err
	}

	filename := utils.GetFilename(layer, extension)
	directory := path.Join(current, kind, layer, name)
	filepath := path.Join(directory, filename)

	isExist := utils.PathIsExist(directory)

	if !isExist {
		err = utils.CreateDirectory(directory)

		if err != nil {
			return err
		}
	}

	err = utils.CreateFile(filepath)

	if err != nil {
		return err
	}

	data := template.GetImplementationRealisationTemplate(separator, layer, name)

	err = utils.SetFileData(filepath, data)

	if err != nil {
		return err
	}

	return nil
}

func CreateReadme() error {
	current, err := os.Getwd()

	if err != nil {
		return err
	}

	name := "README"
	extension := "md"

	if err != nil {
		return err
	}

	filename := utils.GetFilename(name, extension)
	directory := path.Join(current, "test12345")
	filepath := path.Join(directory, filename)

	isExist := utils.PathIsExist(directory)

	if !isExist {
		err = utils.CreateDirectory(directory)

		if err != nil {
			return err
		}
	}

	err = utils.CreateFile(filepath)

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

func CreateGitIgnore() error {
	current, err := os.Getwd()

	if err != nil {
		return err
	}

	separator := utils.GetSeparator()

	name := ".gitignore"
	extension := ""

	if err != nil {
		return err
	}

	filename := utils.GetFilename(name, extension)
	directory := path.Join(current, "test12345")
	filepath := path.Join(directory, filename)

	isExist := utils.PathIsExist(directory)

	if !isExist {
		err = utils.CreateDirectory(directory)

		if err != nil {
			return err
		}
	}

	err = utils.CreateFile(filepath)

	if err != nil {
		return err
	}

	data := template.GetGitIgnoreTemplate(separator)

	err = utils.SetFileData(filepath, data)

	if err != nil {
		return err
	}

	return nil
}

func CreateExampleEnvironment() error {
	current, err := os.Getwd()

	if err != nil {
		return err
	}

	name := ".example"
	extension := "env"

	if err != nil {
		return err
	}

	filename := utils.GetFilename(name, extension)
	directory := path.Join(current, "test12345")
	filepath := path.Join(directory, filename)

	isExist := utils.PathIsExist(directory)

	if !isExist {
		err = utils.CreateDirectory(directory)

		if err != nil {
			return err
		}
	}

	err = utils.CreateFile(filepath)

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

func CreateGrpcMicroserviceMakefile() error {
	current, err := os.Getwd()

	if err != nil {
		return err
	}

	separator := utils.GetSeparator()

	name := "Makefile"
	extension := ""

	if err != nil {
		return err
	}

	filename := utils.GetFilename(name, extension)
	directory := path.Join(current, "test12345")
	filepath := path.Join(directory, filename)

	isExist := utils.PathIsExist(directory)

	if !isExist {
		err = utils.CreateDirectory(directory)

		if err != nil {
			return err
		}
	}

	err = utils.CreateFile(filepath)

	if err != nil {
		return err
	}

	data := template.GetGrpcMicroserviceMakefileTemplate(separator)

	err = utils.SetFileData(filepath, data)

	if err != nil {
		return err
	}

	return nil
}

func CreateDefaultMicroserviceMakefile() error {
	current, err := os.Getwd()

	if err != nil {
		return err
	}

	separator := utils.GetSeparator()

	name := "Makefile"
	extension := ""

	if err != nil {
		return err
	}

	filename := utils.GetFilename(name, extension)
	directory := path.Join(current, "test12345")
	filepath := path.Join(directory, filename)

	isExist := utils.PathIsExist(directory)

	if !isExist {
		err = utils.CreateDirectory(directory)

		if err != nil {
			return err
		}
	}

	err = utils.CreateFile(filepath)

	if err != nil {
		return err
	}

	data := template.GetDefaultMicroserviceMakefileTemplate(separator)

	err = utils.SetFileData(filepath, data)

	if err != nil {
		return err
	}

	return nil
}

func CreateDockerIgnore() error {
	current, err := os.Getwd()

	if err != nil {
		return err
	}

	separator := utils.GetSeparator()

	name := ".dockerignore"
	extension := ""

	if err != nil {
		return err
	}

	filename := utils.GetFilename(name, extension)
	directory := path.Join(current, "test12345")
	filepath := path.Join(directory, filename)

	isExist := utils.PathIsExist(directory)

	if !isExist {
		err = utils.CreateDirectory(directory)

		if err != nil {
			return err
		}
	}

	err = utils.CreateFile(filepath)

	if err != nil {
		return err
	}

	data := template.GetDockerIgnoreTemplate(separator)

	err = utils.SetFileData(filepath, data)

	if err != nil {
		return err
	}

	return nil
}

func CreateDockerWithPort() error {
	current, err := os.Getwd()

	if err != nil {
		return err
	}

	separator := utils.GetSeparator()

	name := "application"
	extension := "dockerfile"

	if err != nil {
		return err
	}

	filename := utils.GetFilename(name, extension)
	directory := path.Join(current, "test12345")
	filepath := path.Join(directory, filename)

	isExist := utils.PathIsExist(directory)

	if !isExist {
		err = utils.CreateDirectory(directory)

		if err != nil {
			return err
		}
	}

	err = utils.CreateFile(filepath)

	if err != nil {
		return err
	}

	data := template.GetDockerWithPortTemplate(separator)

	err = utils.SetFileData(filepath, data)

	if err != nil {
		return err
	}

	return nil
}

func CreateDockerWithoutPort() error {
	current, err := os.Getwd()

	if err != nil {
		return err
	}

	separator := utils.GetSeparator()

	name := "application"
	extension := "dockerfile"

	if err != nil {
		return err
	}

	filename := utils.GetFilename(name, extension)
	directory := path.Join(current, "test12345")
	filepath := path.Join(directory, filename)

	isExist := utils.PathIsExist(directory)

	if !isExist {
		err = utils.CreateDirectory(directory)

		if err != nil {
			return err
		}
	}

	err = utils.CreateFile(filepath)

	if err != nil {
		return err
	}

	data := template.GetDockerWithoutPortTemplate(separator)

	err = utils.SetFileData(filepath, data)

	if err != nil {
		return err
	}

	return nil
}

func CreateGrpcGenerateShellScript() error {
	current, err := os.Getwd()

	if err != nil {
		return err
	}

	separator := utils.GetSeparator()

	kind := "bin"
	name := "grpc-generate"
	extension := "sh"

	if err != nil {
		return err
	}

	filename := utils.GetFilename(name, extension)
	directory := path.Join(current, "test12345", kind)
	filepath := path.Join(directory, filename)

	isExist := utils.PathIsExist(directory)

	if !isExist {
		err = utils.CreateDirectory(directory)

		if err != nil {
			return err
		}
	}

	err = utils.CreateFile(filepath)

	if err != nil {
		return err
	}

	data := template.GetGrpcGenerateShellScriptTemplate(separator)

	err = utils.SetFileData(filepath, data)

	if err != nil {
		return err
	}

	return nil
}

func CreateMockGenerateShellScript() error {
	current, err := os.Getwd()

	if err != nil {
		return err
	}

	separator := utils.GetSeparator()

	kind := "bin"
	name := "mock-generate"
	extension := "sh"

	if err != nil {
		return err
	}

	filename := utils.GetFilename(name, extension)
	directory := path.Join(current, "test12345", kind)
	filepath := path.Join(directory, filename)

	isExist := utils.PathIsExist(directory)

	if !isExist {
		err = utils.CreateDirectory(directory)

		if err != nil {
			return err
		}
	}

	err = utils.CreateFile(filepath)

	if err != nil {
		return err
	}

	data := template.GetMockGenerateShellScriptTemplate(separator)

	err = utils.SetFileData(filepath, data)

	if err != nil {
		return err
	}

	return nil
}

func CreateGrpcLoggingInterceptor() error {
	current, err := os.Getwd()

	if err != nil {
		return err
	}

	separator := utils.GetSeparator()

	kind := "cmd"
	subfolder := "grpc_server"
	folder := "interceptor"
	name := "logging"
	extension := "go"

	if err != nil {
		return err
	}

	filename := utils.GetFilename(name, extension)
	directory := path.Join(current, "test12345", kind, subfolder, folder)
	filepath := path.Join(directory, filename)

	isExist := utils.PathIsExist(directory)

	if !isExist {
		err = utils.CreateDirectory(directory)

		if err != nil {
			return err
		}
	}

	err = utils.CreateFile(filepath)

	if err != nil {
		return err
	}

	data := template.GetGrpcLoggingInterceptorTemplate(separator)

	err = utils.SetFileData(filepath, data)

	if err != nil {
		return err
	}

	return nil
}

func CreateGrpcTracingInterceptor() error {
	current, err := os.Getwd()

	if err != nil {
		return err
	}

	separator := utils.GetSeparator()

	kind := "cmd"
	subfolder := "grpc_server"
	folder := "interceptor"
	name := "tracing"
	extension := "go"

	if err != nil {
		return err
	}

	filename := utils.GetFilename(name, extension)
	directory := path.Join(current, "test12345", kind, subfolder, folder)
	filepath := path.Join(directory, filename)

	isExist := utils.PathIsExist(directory)

	if !isExist {
		err = utils.CreateDirectory(directory)

		if err != nil {
			return err
		}
	}

	err = utils.CreateFile(filepath)

	if err != nil {
		return err
	}

	data := template.GetGrpcTracingInterceptorTemplate(separator)

	err = utils.SetFileData(filepath, data)

	if err != nil {
		return err
	}

	return nil
}

func CreateGrpcAuthenticationMiddleware() error {
	current, err := os.Getwd()

	if err != nil {
		return err
	}

	separator := utils.GetSeparator()

	kind := "cmd"
	subfolder := "grpc_server"
	folder := "middleware"
	name := "authentication"
	extension := "go"

	if err != nil {
		return err
	}

	filename := utils.GetFilename(name, extension)
	directory := path.Join(current, "test12345", kind, subfolder, folder)
	filepath := path.Join(directory, filename)

	isExist := utils.PathIsExist(directory)

	if !isExist {
		err = utils.CreateDirectory(directory)

		if err != nil {
			return err
		}
	}

	err = utils.CreateFile(filepath)

	if err != nil {
		return err
	}

	data := template.GetGrpcAuthenticationMiddlewareTemplate(separator)

	err = utils.SetFileData(filepath, data)

	if err != nil {
		return err
	}

	return nil
}

func CreateGrpcServer() error {
	current, err := os.Getwd()

	if err != nil {
		return err
	}

	module, err := utils.GetModuleName()

	if err != nil {
		return err
	}

	separator := utils.GetSeparator()

	kind := "cmd"
	subfolder := "grpc_server"
	name := "grpc_server"
	extension := "go"

	if err != nil {
		return err
	}

	filename := utils.GetFilename(name, extension)
	directory := path.Join(current, "test12345", kind, subfolder)
	filepath := path.Join(directory, filename)

	isExist := utils.PathIsExist(directory)

	if !isExist {
		err = utils.CreateDirectory(directory)

		if err != nil {
			return err
		}
	}

	err = utils.CreateFile(filepath)

	if err != nil {
		return err
	}

	data := template.GetGrpcServerTemplate(module, separator)

	err = utils.SetFileData(filepath, data)

	if err != nil {
		return err
	}

	return nil
}

func CreateHttpLoggingInterceptor() error {
	current, err := os.Getwd()

	if err != nil {
		return err
	}

	separator := utils.GetSeparator()

	kind := "cmd"
	subfolder := "http_server"
	folder := "interceptor"
	name := "logging"
	extension := "go"

	if err != nil {
		return err
	}

	filename := utils.GetFilename(name, extension)
	directory := path.Join(current, "test12345", kind, subfolder, folder)
	filepath := path.Join(directory, filename)

	isExist := utils.PathIsExist(directory)

	if !isExist {
		err = utils.CreateDirectory(directory)

		if err != nil {
			return err
		}
	}

	err = utils.CreateFile(filepath)

	if err != nil {
		return err
	}

	data := template.GetHttpLoggingInterceptorTemplate(separator)

	err = utils.SetFileData(filepath, data)

	if err != nil {
		return err
	}

	return nil
}

func CreateHttpAuthenticationMiddleware() error {
	current, err := os.Getwd()

	if err != nil {
		return err
	}

	separator := utils.GetSeparator()

	kind := "cmd"
	subfolder := "http_server"
	folder := "middleware"
	name := "authentication"
	extension := "go"

	if err != nil {
		return err
	}

	filename := utils.GetFilename(name, extension)
	directory := path.Join(current, "test12345", kind, subfolder, folder)
	filepath := path.Join(directory, filename)

	isExist := utils.PathIsExist(directory)

	if !isExist {
		err = utils.CreateDirectory(directory)

		if err != nil {
			return err
		}
	}

	err = utils.CreateFile(filepath)

	if err != nil {
		return err
	}

	data := template.GetHttpAuthenticationMiddlewareTemplate(separator)

	err = utils.SetFileData(filepath, data)

	if err != nil {
		return err
	}

	return nil
}

func CreateHttpCorsMiddleware() error {
	current, err := os.Getwd()

	if err != nil {
		return err
	}

	separator := utils.GetSeparator()

	kind := "cmd"
	subfolder := "http_server"
	folder := "middleware"
	name := "cors"
	extension := "go"

	if err != nil {
		return err
	}

	filename := utils.GetFilename(name, extension)
	directory := path.Join(current, "test12345", kind, subfolder, folder)
	filepath := path.Join(directory, filename)

	isExist := utils.PathIsExist(directory)

	if !isExist {
		err = utils.CreateDirectory(directory)

		if err != nil {
			return err
		}
	}

	err = utils.CreateFile(filepath)

	if err != nil {
		return err
	}

	data := template.GetHttpCorsMiddlewareTemplate(separator)

	err = utils.SetFileData(filepath, data)

	if err != nil {
		return err
	}

	return nil
}

func CreateHttpChainMiddleware() error {
	current, err := os.Getwd()

	if err != nil {
		return err
	}

	separator := utils.GetSeparator()

	kind := "cmd"
	subfolder := "http_server"
	folder := "middleware"
	name := "chain"
	extension := "go"

	if err != nil {
		return err
	}

	filename := utils.GetFilename(name, extension)
	directory := path.Join(current, "test12345", kind, subfolder, folder)
	filepath := path.Join(directory, filename)

	isExist := utils.PathIsExist(directory)

	if !isExist {
		err = utils.CreateDirectory(directory)

		if err != nil {
			return err
		}
	}

	err = utils.CreateFile(filepath)

	if err != nil {
		return err
	}

	data := template.GetHttpChainMiddlewareTemplate(separator)

	err = utils.SetFileData(filepath, data)

	if err != nil {
		return err
	}

	return nil
}

func CreateHttpServer() error {
	current, err := os.Getwd()

	if err != nil {
		return err
	}

	module, err := utils.GetModuleName()

	if err != nil {
		return err
	}

	separator := utils.GetSeparator()

	kind := "cmd"
	folder := "http_server"
	name := "http_server"
	extension := "go"

	if err != nil {
		return err
	}

	filename := utils.GetFilename(name, extension)
	directory := path.Join(current, "test12345", kind, folder)
	filepath := path.Join(directory, filename)

	isExist := utils.PathIsExist(directory)

	if !isExist {
		err = utils.CreateDirectory(directory)

		if err != nil {
			return err
		}
	}

	err = utils.CreateFile(filepath)

	if err != nil {
		return err
	}

	data := template.GetHttpServerTemplate(module, separator)

	err = utils.SetFileData(filepath, data)

	if err != nil {
		return err
	}

	return nil
}

func CreateApplication(application string, name string, implementing string) error {
	current, err := os.Getwd()

	if err != nil {
		return err
	}

	module, err := utils.GetModuleName()

	if err != nil {
		return err
	}

	separator := utils.GetSeparator()

	kind := "cmd"
	extension := "go"

	if err != nil {
		return err
	}

	filename := utils.GetFilename(implementing, extension)
	directory := path.Join(current, "test12345", kind, implementing)
	filepath := path.Join(directory, filename)

	isExist := utils.PathIsExist(directory)

	if !isExist {
		err = utils.CreateDirectory(directory)

		if err != nil {
			return err
		}
	}

	err = utils.CreateFile(filepath)

	if err != nil {
		return err
	}

	data := template.GetApplicationTemplate(application, module, separator, name, implementing)

	err = utils.SetFileData(filepath, data)

	if err != nil {
		return err
	}

	return nil
}

func CreateNatsSubscriber() error {
	current, err := os.Getwd()

	if err != nil {
		return err
	}

	separator := utils.GetSeparator()

	kind := "cmd"
	folder := "nats_subscriber"
	name := "nats_subscriber"
	extension := "go"

	if err != nil {
		return err
	}

	filename := utils.GetFilename(name, extension)
	directory := path.Join(current, "test12345", kind, folder)
	filepath := path.Join(directory, filename)

	isExist := utils.PathIsExist(directory)

	if !isExist {
		err = utils.CreateDirectory(directory)

		if err != nil {
			return err
		}
	}

	err = utils.CreateFile(filepath)

	if err != nil {
		return err
	}

	data := template.GetNatsSubscriberTemplate(separator)

	err = utils.SetFileData(filepath, data)

	if err != nil {
		return err
	}

	return nil
}
