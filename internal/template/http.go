package template

import (
	"bytes"
	"fmt"
	"github.com/rafailovalexey/service-generator/internal/dto"
	"github.com/rafailovalexey/service-generator/internal/util"
	"sort"
)

func GetHttpHandlerDefinitionTemplate(application *dto.ApplicationDto) []byte {
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

	data.WriteString(fmt.Sprintf("type %s%sInterface interface {", application.Names.CamelCaseSingular, util.GetWithUpperCaseFirstLetter("handler")))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tHandle(response http.ResponseWriter, request *http.Request)"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)

	return data.Bytes()
}

func GetHttpHandlerImplementationTemplate(application *dto.ApplicationDto) []byte {
	data := bytes.Buffer{}
	separator := util.GetSeparator()

	imports := []string{
		"\"github.com/sirupsen/logrus\"",
		fmt.Sprintf("\"%s/config\"", application.Module),
		fmt.Sprintf("\"%s/util\"", application.Module),
		fmt.Sprintf("\"net/http\""),
		fmt.Sprintf("\"%s/internal/controller\"", application.Module),
		fmt.Sprintf("definition \"%s/%s/%s\"", application.Module, "internal", "handler"),
	}

	sort.Strings(imports)

	data.WriteString(fmt.Sprintf("package %s", application.Names.SnakeCasePlural))
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

	data.WriteString(fmt.Sprintf("type %s%s struct {", application.Names.CamelCaseSingular, util.GetWithUpperCaseFirstLetter("handler")))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tconfig *config.Config"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tlogger *logrus.Logger"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t%sController controller.%sControllerInterface", application.Names.LowerCamelCaseSingular, application.Names.CamelCaseSingular))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("var _ definition.%s%sInterface = (*%s%s)(nil)", application.Names.CamelCaseSingular, util.GetWithUpperCaseFirstLetter("handler"), application.Names.CamelCaseSingular, util.GetWithUpperCaseFirstLetter("handler")))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("func New%s%s(", application.Names.CamelCaseSingular, util.GetWithUpperCaseFirstLetter("handler")))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tconfig *config.Config,"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tlogger *logrus.Logger,"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t%sController controller.%sControllerInterface,", application.Names.LowerCamelCaseSingular, application.Names.CamelCaseSingular))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf(") *%s%s {", application.Names.CamelCaseSingular, util.GetWithUpperCaseFirstLetter("handler")))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\treturn &%s%s{", application.Names.CamelCaseSingular, util.GetWithUpperCaseFirstLetter("handler")))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\tconfig: config,"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\tlogger: logger,"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\t%sController: %sController,", application.Names.LowerCamelCaseSingular, application.Names.LowerCamelCaseSingular))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t}"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("func (%s *%s%s) Handle(response http.ResponseWriter, request *http.Request) {", application.Names.LowerCaseFirstLetter, application.Names.CamelCaseSingular, util.GetWithUpperCaseFirstLetter("handler")))
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

func GetHttpInterceptorTemplate() []byte {
	data := bytes.Buffer{}
	separator := util.GetSeparator()

	data.WriteString(fmt.Sprintf("package interceptor"))
	data.WriteString(separator)
	data.WriteString(separator)

	imports := []string{
		"\"net/http\"",
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

	data.WriteString(fmt.Sprintf("type InterceptorInterface interface {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tApply(http.Handler) http.Handler"))
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
		"\"net/http\"",
		"\"time\"",
		"\"github.com/sirupsen/logrus\"",
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

	data.WriteString(fmt.Sprintf("type LoggingInterceptor struct {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tlogger *logrus.Logger"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("var _ InterceptorInterface = (*LoggingInterceptor)(nil)"))
	data.WriteString(separator)
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("func NewLoggingInterceptor(logger *logrus.Logger) *LoggingInterceptor {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\treturn &LoggingInterceptor{"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\tlogger: logger,"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t}"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("func (l *LoggingInterceptor) Apply(next http.Handler) http.Handler {"))
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

	data.WriteString(fmt.Sprint("\t\tl.logger.Debugf(\"%s %s %s - %s %v\", request.Method, request.URL.Host, request.RemoteAddr, request.UserAgent(), duration)"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t})"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)

	return data.Bytes()
}

func GetHttpMiddlewareTemplate() []byte {
	data := bytes.Buffer{}
	separator := util.GetSeparator()

	data.WriteString(fmt.Sprintf("package middleware"))
	data.WriteString(separator)
	data.WriteString(separator)

	imports := []string{
		"\"net/http\"",
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

	data.WriteString(fmt.Sprintf("type MiddlewareInterface interface {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tApply(http.Handler) http.Handler"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)

	return data.Bytes()
}

func GetHttpAuthenticationMiddlewareTemplate(application *dto.ApplicationDto) []byte {
	data := bytes.Buffer{}
	separator := util.GetSeparator()

	imports := []string{
		"\"net/http\"",
		fmt.Sprintf("\"%s/util\"", application.Module),
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

	data.WriteString(fmt.Sprintf("type HttpAuthenticationConfig struct {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tHeader string `env:\"HTTP_SERVER_AUTHENTICATION_TOKEN_HEADER,required\"`"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tToken string `env:\"HTTP_SERVER_AUTHENTICATION_TOKEN,required\"`"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("type AuthenticationMiddleware struct {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tconfig HttpAuthenticationConfig"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("var _ MiddlewareInterface = (*AuthenticationMiddleware)(nil)"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("func NewAuthenticationMiddleware(config HttpAuthenticationConfig) *AuthenticationMiddleware {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\treturn &AuthenticationMiddleware{"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\tconfig: config,"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t}"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("func (a *AuthenticationMiddleware) Apply(next http.Handler) http.Handler {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\treturn http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\tkey := request.Header.Get(a.config.Header)"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\t\tif key != a.config.Token {"))
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

	data.WriteString(fmt.Sprintf("type CorsMiddleware struct{}"))
	data.WriteString(separator)
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("var _ MiddlewareInterface = (*CorsMiddleware)(nil)"))
	data.WriteString(separator)
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("func NewCorsMiddleware() *CorsMiddleware {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\treturn &CorsMiddleware{}"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("func (c *CorsMiddleware) Apply(next http.Handler) http.Handler {"))
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

func GetHttpServerTemplate(application *dto.ApplicationDto) []byte {
	data := bytes.Buffer{}
	separator := util.GetSeparator()

	imports := []string{
		"\"fmt\"",
		"\"net/http\"",
		"\"github.com/sirupsen/logrus\"",
		fmt.Sprintf("\"%s/internal/handler\"", application.Module),
		fmt.Sprintf("\"%s/cmd/http_server/interceptor\"", application.Module),
		fmt.Sprintf("\"%s/cmd/http_server/middleware\"", application.Module),
		fmt.Sprintf("\"%s/util\"", application.Module),
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

	data.WriteString(fmt.Sprintf("type HttpServerConfig struct {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tHostname string `env:\"HTTP_SERVER_HOSTNAME,required\"`"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tPort string `env:\"HTTP_SERVER_PORT,required\"`"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tAuthentication middleware.HttpAuthenticationConfig"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("type HttpServer struct {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tconfig HttpServerConfig"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tlogger *logrus.Logger"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t%sHandler handler.%sHandlerInterface", application.Names.LowerCamelCaseSingular, application.Names.CamelCaseSingular))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("type HttpServerInterface interface {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tRun() error"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("var _ HttpServerInterface = (*HttpServer)(nil)"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("func NewHttpServer("))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tconfig HttpServerConfig,"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tlogger *logrus.Logger,"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t%sHandler handler.%sHandlerInterface,", application.Names.LowerCamelCaseSingular, application.Names.CamelCaseSingular))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf(") *HttpServer {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\treturn &HttpServer{"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\tconfig: config,"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\tlogger: logger,"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\t%sHandler: %sHandler,", application.Names.LowerCamelCaseSingular, application.Names.LowerCamelCaseSingular))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t}"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("func (h *HttpServer) Run() error {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\trouter := http.NewServeMux()"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\tmiddlewares := middleware.ChainMiddleware("))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\tinterceptor.NewLoggingInterceptor(h.logger).Apply,"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\tmiddleware.NewCorsMiddleware().Apply,"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\tmiddleware.NewAuthenticationMiddleware(h.config.Authentication).Apply,"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t)"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\troute := middlewares(http.HandlerFunc(h.%sHandler.Handle))", application.Names.LowerCamelCaseSingular))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\trouter.Handle(\"/v1/%s\", route)", application.Names.SnakeCasePlural))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\trouter.Handle(\"/\", middlewares(http.HandlerFunc(util.ResponseNotFound)))"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprint("\taddress := fmt.Sprintf(\"%s:%s\", h.config.Hostname, h.config.Port)"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprint("\th.logger.Infof(\"http server starts at address %s\", address)"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\terr := http.ListenAndServe(address, router)"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\tif err != nil {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprint("\t\treturn fmt.Errorf(\"error when starting the http server %v\", err)"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t}"))
	data.WriteString(separator)
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\treturn nil"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)

	return data.Bytes()
}
