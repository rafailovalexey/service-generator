package template

import (
	"bytes"
	"fmt"
	"github.com/rafailovalexey/service-generator/internal/dto"
	"github.com/rafailovalexey/service-generator/internal/util"
	"sort"
)

func GetGrpcServerImplementationTemplate(application *dto.ApplicationDto) []byte {
	data := bytes.Buffer{}
	separator := util.GetSeparator()

	imports := []string{
		fmt.Sprintf("\"%s/pkg/%s_v1\"", application.Module, application.Names.SnakeCasePlural),
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

	data.WriteString(fmt.Sprintf("type %s%s struct {", application.Names.CamelCaseSingular, util.GetWithUpperCaseFirstLetter("implementation")))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t%s_v1.Unimplemented%sV1Server", application.Names.SnakeCasePlural, application.Names.CamelCasePlural))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("func New%s%s() *%s%s {", application.Names.CamelCaseSingular, util.GetWithUpperCaseFirstLetter("implementation"), application.Names.CamelCaseSingular, util.GetWithUpperCaseFirstLetter("implementation")))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\treturn &%s%s{}", application.Names.CamelCaseSingular, util.GetWithUpperCaseFirstLetter("implementation")))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("func (%s *%s%s) mustEmbedUnimplemented%sV1Server() {", application.Names.LowerCaseFirstLetter, application.Names.CamelCaseSingular, util.GetWithUpperCaseFirstLetter("implementation"), application.Names.CamelCaseSingular))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\treturn"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)

	return data.Bytes()
}

func GetGrpcLoggingInterceptorTemplate() []byte {
	data := bytes.Buffer{}
	separator := util.GetSeparator()

	imports := []string{
		"\"context\"",
		"\"google.golang.org/grpc\"",
		"\"google.golang.org/grpc/codes\"",
		"\"google.golang.org/grpc/metadata\"",
		"\"google.golang.org/grpc/status\"",
		"\"github.com/sirupsen/logrus\"",
	}

	sort.Strings(imports)

	data.WriteString(fmt.Sprintf("package interceptor"))
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

	data.WriteString(fmt.Sprintf("func LoggingInterceptor(logger *logrus.Logger) grpc.UnaryServerInterceptor {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\treturn func(ctx context.Context, request any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\tmd, isExist := metadata.FromIncomingContext(ctx)"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\t\tif !isExist {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\t\treturn nil, status.Errorf(codes.Internal, \"failed to read metadata\")"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\t}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\t\ttrace := md[\"trace\"][0]"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprint("\t\tlogger.Debugf(\"incoming grpc request: %s (%s)\", info.FullMethod, trace)"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\t\tresponse, err := handler(ctx, request)"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\t\tif err != nil {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprint("\t\t\tlogger.Debugf(\"error in grpc request %s (%s) %v\", info.FullMethod, trace, err)"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\t}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\t\tif err == nil {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprint("\t\t\tlogger.Debugf(\"outgoing grpc response %s (%s)\", info.FullMethod, trace)"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\t}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\t\treturn response, err"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t}"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)

	return data.Bytes()
}

func GetGrpcTracingInterceptorTemplate() []byte {
	data := bytes.Buffer{}
	separator := util.GetSeparator()

	imports := []string{
		"\"context\"",
		"\"crypto/rand\"",
		"\"encoding/hex\"",
		"\"google.golang.org/grpc\"",
		"\"google.golang.org/grpc/codes\"",
		"\"google.golang.org/grpc/metadata\"",
		"\"google.golang.org/grpc/status\"",
	}

	sort.Strings(imports)

	data.WriteString(fmt.Sprintf("package interceptor"))
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

	data.WriteString(fmt.Sprintf("func TracingInterceptor() grpc.UnaryServerInterceptor {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\treturn func(ctx context.Context, request any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\tmd, isExist := metadata.FromIncomingContext(ctx)"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\t\tif !isExist {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\t\treturn nil, status.Errorf(codes.Internal, \"failed to read metadata\")"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\t}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\t\tif len(md[\"trace\"]) != 0 {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\t\treturn handler(ctx, request)"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\t}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\t\ttrace, err := GenerateTrace()"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\t\tif err != nil {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\t\treturn nil, status.Errorf(codes.Internal, \"failed to generate trace\")"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\t}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\t\tmd = metadata.Join(md, metadata.New(map[string]string{\"trace\": trace}))"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\tctx = metadata.NewIncomingContext(ctx, md)"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\t\treturn handler(ctx, request)"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t}"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("func GenerateTrace() (string, error) {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\ttrace := make([]byte, 16)"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\tif _, err := rand.Read(trace); err != nil {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\treturn \"\", err"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\treturn hex.EncodeToString(trace), nil"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)

	return data.Bytes()
}

func GetGrpcAuthenticationMiddlewareTemplate() []byte {
	data := bytes.Buffer{}
	separator := util.GetSeparator()

	imports := []string{
		"\"context\"",
		"\"google.golang.org/grpc\"",
		"\"google.golang.org/grpc/codes\"",
		"\"google.golang.org/grpc/metadata\"",
		"\"google.golang.org/grpc/status\"",
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

	data.WriteString(fmt.Sprintf("type GrpcAuthenticationConfig struct {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tHeader\tstring"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tToken\tstring"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("func AuthenticationMiddleware(config *GrpcAuthenticationConfig) grpc.UnaryServerInterceptor {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\treturn func(ctx context.Context, request any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\tmd, isExist := metadata.FromIncomingContext(ctx)"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\t\tif !isExist {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\t\treturn nil, status.Errorf(codes.Unauthenticated, \"authentication token not found\")"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\t}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\t\tlist := md[config.Header]"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\t\tif len(list) == 0 {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\t\treturn nil, status.Errorf(codes.Unauthenticated, \"authentication token not found\")"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\t}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\t\tkey := list[0]"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\t\tif config.Token != key {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\t\treturn nil, status.Errorf(codes.PermissionDenied, \"invalid authentication token\")"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\t}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\t\treturn handler(ctx, request)"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t}"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)

	return data.Bytes()
}

func GetGrpcServerTemplate(application *dto.ApplicationDto) []byte {
	data := bytes.Buffer{}
	separator := util.GetSeparator()

	imports := []string{
		"\"fmt\"",
		"\"net\"",
		"\"google.golang.org/grpc\"",
		"\"google.golang.org/grpc/reflection\"",
		"\"github.com/sirupsen/logrus\"",
		fmt.Sprintf("\"%s/cmd/grpc_server/interceptor\"", application.Module),
		fmt.Sprintf("\"%s/cmd/grpc_server/middleware\"", application.Module),
		fmt.Sprintf("\"%s/pkg/%s_v1\"", application.Module, application.Names.SnakeCasePlural),
	}

	sort.Strings(imports)

	data.WriteString(fmt.Sprintf("package grpc_server"))
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

	data.WriteString(fmt.Sprintf("type GrpcServerConfig struct {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tHostname\tstring"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tPort\tstring"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tAuthentication\t*middleware.GrpcAuthenticationConfig"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("type GrpcServer struct {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tconfig\t*GrpcServerConfig"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tlogger\t*logrus.Logger"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tapi\t%s_v1.%sV1Server", application.Names.SnakeCasePlural, application.Names.CamelCasePlural))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("type GrpcServerInterface interface {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tRun() error"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("var _ GrpcServerInterface = (*GrpcServer)(nil)"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("func NewGrpcServer("))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tconfig *GrpcServerConfig,"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tlogger *logrus.Logger,"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tapi %s_v1.%sV1Server,", application.Names.SnakeCasePlural, application.Names.CamelCasePlural))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf(") *GrpcServer {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\treturn &GrpcServer{"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\tconfig: config,"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\tlogger: logger,"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\tapi: api,"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t}"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("func (g *GrpcServer) Run() error {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprint("\taddress := fmt.Sprintf(\"%s:%s\", g.config.Hostname, g.config.Port)"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprint("\tg.logger.Infof(fmt.Sprintf(\"grpc server starts at address %s\", address))"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\tlistener, err := net.Listen(\"tcp\", address)"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\tif err != nil {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprint("\t\treturn fmt.Errorf(\"grpc server startup error %v\", err)"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\tserver := grpc.NewServer("))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\tgrpc.ChainUnaryInterceptor("))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\t\tinterceptor.TracingInterceptor(),"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\t\tinterceptor.LoggingInterceptor(g.logger),"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\t\tmiddleware.AuthenticationMiddleware(g.config.Authentication),"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\t),"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t)"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\treflection.Register(server)"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\t%s_v1.Register%sV1Server(server, g.api)", application.Names.SnakeCasePlural, application.Names.CamelCasePlural))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprint("\tg.logger.Infof(fmt.Sprintf(\"grpc server is running at %s\", address))"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\terr = server.Serve(listener)"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\tif err != nil {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprint("\t\treturn fmt.Errorf(\"grpc server startup error %v\", err)"))
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

func GetProtoTemplate(application *dto.ApplicationDto) []byte {
	data := bytes.Buffer{}
	separator := util.GetSeparator()

	data.WriteString(fmt.Sprintf("syntax = \"proto3\";"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("package %s_v1;", application.Names.SnakeCasePlural))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("option go_package = \"%s/%s/%s_v1\";", application.Module, "api", application.Names.SnakeCasePlural))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("service %sV1 {}", application.Names.CamelCasePlural))
	data.WriteString(separator)

	return data.Bytes()
}

func GetGrpcGenerateShellScriptTemplate() []byte {
	data := bytes.Buffer{}
	separator := util.GetSeparator()

	data.WriteString(fmt.Sprintf("#!/bin/bash"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("if [ \"$#\" -lt 2 ]; then"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\techo \"Usage: $0 <PROTO_SOURCE_DIRECTORY> <PROTO_OUTPUT_DIRECTORY> <PROTO_FILES>\""))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\texit 1"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("fi"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("PROTO_SOURCE_DIRECTORY=\"$1\""))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("PROTO_OUTPUT_DIRECTORY=\"$2\""))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("PROTO_FILES=\"${*:3}\""))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("for proto_file in $PROTO_FILES; do"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tPROTO_FILE_DIRECTORY=$(dirname \"$proto_file\")"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\tmkdir -p \"$PROTO_OUTPUT_DIRECTORY/$PROTO_FILE_DIRECTORY\""))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\techo \"Generating proto file for $proto_file...\""))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\tprotoc \\"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\t--proto_path=\"$PROTO_SOURCE_DIRECTORY\" \\"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\t--go_out=\"$PROTO_OUTPUT_DIRECTORY\" \\"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\t--go_opt=paths=source_relative \"$proto_file\" \\"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\t--go-grpc_out=\"$PROTO_OUTPUT_DIRECTORY\" \\"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\t--go-grpc_opt=paths=source_relative \"$proto_file\""))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("done"))
	data.WriteString(separator)

	return data.Bytes()
}
