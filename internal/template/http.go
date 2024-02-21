package template

import (
	"bytes"
	"fmt"
	"github.com/rafailovalexey/service-generator/internal/dto"
	"github.com/rafailovalexey/service-generator/internal/util"
	"sort"
)

func GetHttpHandlerDefinitionTemplate(name *dto.NameDto) []byte {
	data := bytes.Buffer{}
	separator := util.GetSeparator()

	imports := []string{
		"\"net/http\"",
	}

	sort.Strings(imports)

	data.WriteString(fmt.Sprintf("package %s", "handler"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("import ("))
	data.WriteString(separator)

	for _, i := range imports {
		data.WriteString(fmt.Sprintf("\t%s", i))
		data.WriteString(separator)
	}

	data.WriteString(fmt.Sprintf(")"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("type %s%sInterface interface {", name.CamelCaseSingular, util.GetWithUpperCaseFirstLetter("handler")))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tHandle(response http.ResponseWriter, request *http.Request)"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)

	return data.Bytes()
}

func GetHttpHandlerImplementationTemplate(module string, name *dto.NameDto) []byte {
	data := bytes.Buffer{}
	separator := util.GetSeparator()

	imports := []string{
		fmt.Sprintf("\"%s/util\"", module),
		fmt.Sprintf("\"net/http\""),
		fmt.Sprintf("definition \"%s/%s/%s\"", module, "internal", "handler"),
	}

	sort.Strings(imports)

	data.WriteString(fmt.Sprintf("package %s", name.SnakeCasePlural))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("import ("))
	data.WriteString(separator)

	for _, i := range imports {
		data.WriteString(fmt.Sprintf("\t%s", i))
		data.WriteString(separator)
	}

	data.WriteString(fmt.Sprintf(")"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("type %s%s struct {}", name.CamelCaseSingular, util.GetWithUpperCaseFirstLetter("handler")))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("var _ definition.%s%sInterface = (*%s%s)(nil)", name.CamelCaseSingular, util.GetWithUpperCaseFirstLetter("handler"), name.CamelCaseSingular, util.GetWithUpperCaseFirstLetter("handler")))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("func New%s%s() *%s%s {", name.CamelCaseSingular, util.GetWithUpperCaseFirstLetter("handler"), name.CamelCaseSingular, util.GetWithUpperCaseFirstLetter("handler")))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\treturn &%s%s{}", name.CamelCaseSingular, util.GetWithUpperCaseFirstLetter("handler")))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("func (%s *%s%s) Handle(response http.ResponseWriter, request *http.Request) {", name.LowerCaseFirstLetter, name.CamelCaseSingular, util.GetWithUpperCaseFirstLetter("handler")))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tswitch request.Method {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tdefault:"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\tutil.ResponseMethodNotAllowed(response, request)"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\t\treturn"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t}"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)

	return data.Bytes()
}

func GetHttpLoggingInterceptorTemplate() []byte {
	data := bytes.Buffer{}
	separator := util.GetSeparator()

	data.WriteString(fmt.Sprintf("package interceptor"))
	data.WriteString(separator)
	data.WriteString(separator)

	imports := []string{
		"\"log\"",
		"\"net/http\"",
		"\"time\"",
	}

	sort.Strings(imports)

	data.WriteString(fmt.Sprintf("import ("))
	data.WriteString(separator)

	for _, i := range imports {
		data.WriteString(fmt.Sprintf("\t%s", i))
		data.WriteString(separator)
	}

	data.WriteString(fmt.Sprintf(")"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("func LoggingInterceptor(next http.Handler) http.Handler {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\treturn http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\tstart := time.Now()"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\t\tnext.ServeHTTP(response, request)"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\t\tduration := time.Since(start)"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprint("\t\tlog.Printf(\"%s %s %s - %s %v\\n\", request.Method, request.URL.Host, request.RemoteAddr, request.UserAgent(), duration)"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t})"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)

	return data.Bytes()
}

func GetHttpAuthenticationMiddlewareTemplate(module string) []byte {
	data := bytes.Buffer{}
	separator := util.GetSeparator()

	imports := []string{
		fmt.Sprintf("\"%s/util\"", module),
		"\"log\"",
		"\"net/http\"",
		"\"os\"",
	}

	sort.Strings(imports)

	data.WriteString(fmt.Sprintf("package middleware"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("import ("))
	data.WriteString(separator)

	for _, i := range imports {
		data.WriteString(fmt.Sprintf("\t%s", i))
		data.WriteString(separator)
	}

	data.WriteString(fmt.Sprintf(")"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("func AuthenticationMiddleware(next http.Handler) http.Handler {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\treturn http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\theader := os.Getenv(\"AUTHENTICATION_TOKEN_HEADER\")"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\t\tif header == \"\" {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\t\tlog.Panicf(\"specify the name of the authentication token\")"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\t}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\t\ttoken := os.Getenv(\"AUTHENTICATION_TOKEN\")"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\t\tif token == \"\" {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\t\tlog.Panicf(\"specify the value of the authentication token\")"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\t}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\t\tkey := request.Header.Get(header)"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\t\tif key != token {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\t\tutil.ResponseUnauthorized(response, request)"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\t\t\treturn"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\t}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\t\tnext.ServeHTTP(response, request)"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t})"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)

	return data.Bytes()
}

func GetHttpCorsMiddlewareTemplate() []byte {
	data := bytes.Buffer{}
	separator := util.GetSeparator()

	imports := []string{
		"\"net/http\"",
	}

	sort.Strings(imports)

	data.WriteString(fmt.Sprintf("package middleware"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("import ("))
	data.WriteString(separator)

	for _, i := range imports {
		data.WriteString(fmt.Sprintf("\t%s", i))
		data.WriteString(separator)
	}

	data.WriteString(fmt.Sprintf(")"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("func CorsMiddleware(next http.Handler) http.Handler {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\treturn http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\tresponse.Header().Add(\"Access-Control-Allow-Origin\", \"*\")"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\tresponse.Header().Add(\"Access-Control-Allow-Headers\", \"*\")"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\tresponse.Header().Add(\"Access-Control-Allow-Methods\", \"*\")"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\tresponse.Header().Add(\"Access-Control-Allow-Credentials\", \"true\")"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\t\tif request.Method == \"OPTIONS\" {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\t\tresponse.WriteHeader(http.StatusOK)"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\t\t\treturn"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\t}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\t\tnext.ServeHTTP(response, request)"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t})"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)

	return data.Bytes()
}

func GetHttpChainMiddlewareTemplate() []byte {
	data := bytes.Buffer{}
	separator := util.GetSeparator()

	imports := []string{
		"\"net/http\"",
	}

	sort.Strings(imports)

	data.WriteString(fmt.Sprintf("package middleware"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("import ("))
	data.WriteString(separator)

	for _, i := range imports {
		data.WriteString(fmt.Sprintf("\t%s", i))
		data.WriteString(separator)
	}

	data.WriteString(fmt.Sprintf(")"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("func ChainMiddleware(middlewares ...func(http.Handler) http.Handler) func(http.Handler) http.Handler {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\treturn func(next http.Handler) http.Handler {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\tfor index := len(middlewares) - 1; index >= 0; index-- {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\t\tnext = middlewares[index](next)"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\t}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\t\treturn next"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t}"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)

	return data.Bytes()
}

func GetHttpServerTemplate(module string, name *dto.NameDto) []byte {
	data := bytes.Buffer{}
	separator := util.GetSeparator()

	imports := []string{
		"\"fmt\"",
		"\"log\"",
		"\"net/http\"",
		"\"os\"",
		fmt.Sprintf("\"%s/internal/handler\"", module),
		fmt.Sprintf("\"%s/cmd/http_server/interceptor\"", module),
		fmt.Sprintf("\"%s/cmd/http_server/middleware\"", module),
		fmt.Sprintf("\"%s/util\"", module),
	}

	sort.Strings(imports)

	data.WriteString(fmt.Sprintf("package http_server"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("import ("))
	data.WriteString(separator)

	for _, i := range imports {
		data.WriteString(fmt.Sprintf("\t%s", i))
		data.WriteString(separator)
	}

	data.WriteString(fmt.Sprintf(")"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("func Run(%sHandler handler.%sHandlerInterface) {", name.LowerCamelCaseSingular, name.CamelCaseSingular))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\trouter := http.NewServeMux()"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\tmiddlewares := middleware.ChainMiddleware("))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\tinterceptor.LoggingInterceptor,"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\tmiddleware.CorsMiddleware,"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\tmiddleware.AuthenticationMiddleware,"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t)"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\troute := middlewares(http.HandlerFunc(%sHandler.Handle))", name.LowerCamelCaseSingular))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\trouter.Handle(\"/v1/%s\", route)", name.SnakeCasePlural))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\trouter.Handle(\"/\", middlewares(http.HandlerFunc(util.ResponseNotFound)))"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\thostname := os.Getenv(\"HOSTNAME\")"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tport := os.Getenv(\"PORT\")"))
	data.WriteString(separator)
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tif port == \"\" {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\tlog.Panicf(\"specify the port\")"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprint("\taddress := fmt.Sprintf(\"%s:%s\", hostname, port)"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprint("\tlog.Printf(\"http server starts at address %s\\n\", address)"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\terr := http.ListenAndServe(address, router)"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\tif err != nil {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprint("\t\tlog.Panicf(\"error when starting the http server %v\\n\", err)"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t}"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)

	return data.Bytes()
}
