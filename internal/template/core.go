package template

import (
	"bytes"
	"fmt"
	"github.com/rafailovalexey/service-generator/internal/dto"
	"github.com/rafailovalexey/service-generator/internal/util"
	"sort"
)

func GetApplicationTemplate(application *dto.ApplicationDto) []byte {
	data := bytes.Buffer{}
	separator := util.GetSeparator()

	t := ""

	switch application.Type {
	case "http":
		t = "http_server"
	case "grpc":
		t = "grpc_server"
	case "cron":
		t = "cron_scheduler"
	}

	imports := []string{
		"\"context\"",
		"\"database/sql\"",
		"\"os\"",
		"\"github.com/caarlos0/env\"",
		"\"github.com/sirupsen/logrus\"",
		fmt.Sprintf("\"%s/database/%s\"", application.Module, application.Database),
		fmt.Sprintf("\"%s/config\"", application.Module),
		fmt.Sprintf("\"%s/cmd/%s\"", application.Module, t),
		fmt.Sprintf("\"%s/internal/provider\"", application.Module),
		fmt.Sprintf("%sProvider \"%s/internal/provider/%s\"", application.Names.LowerCamelCaseSingular, application.Module, application.Names.SnakeCasePlural),
	}

	switch application.Type {
	case "grpc":
		imports = append(imports, fmt.Sprintf("\"%s/cmd/migration\"", application.Module))
	case "http":
		imports = append(imports, fmt.Sprintf("\"%s/cmd/migration\"", application.Module))
	}

	sort.Strings(imports)

	data.WriteString(fmt.Sprintf("package application"))
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

	data.WriteString(fmt.Sprintf("type ApplicationInterface interface {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tInitializeDependency(context.Context) error"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tInitializeConfig(context.Context) error"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tInitializeLogger(context.Context) error"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tInitializeDatabase(context.Context) error"))
	data.WriteString(separator)
	switch application.Type {
	case "grpc":
		data.WriteString(fmt.Sprintf("\tInitializeMigration(context.Context) error"))
		data.WriteString(separator)
	case "http":
		data.WriteString(fmt.Sprintf("\tInitializeMigration(context.Context) error"))
		data.WriteString(separator)
	}
	data.WriteString(fmt.Sprintf("\tInitializeProvider(context.Context) error"))
	data.WriteString(separator)
	switch application.Type {
	case "grpc":
		data.WriteString(fmt.Sprintf("\tInitializeGrpcServer(context.Context) error"))
		data.WriteString(separator)
	case "http":
		data.WriteString(fmt.Sprintf("\tInitializeHttpServer(context.Context) error"))
		data.WriteString(separator)
	case "cron":
		data.WriteString(fmt.Sprintf("\tInitializeCronScheduler(context.Context) error"))
		data.WriteString(separator)
	}
	data.WriteString(fmt.Sprintf("\tRun() error"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("type Application struct {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tconfig *config.Config"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tlogger *logrus.Logger"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tdatabase *sql.DB"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t%sProvider provider.%sProviderInterface", application.Names.LowerCamelCaseSingular, application.Names.CamelCaseSingular))
	data.WriteString(separator)
	switch application.Type {
	case "grpc":
		data.WriteString(fmt.Sprintf("\tgrpcServer %s.GrpcServerInterface", t))
		data.WriteString(separator)
	case "http":
		data.WriteString(fmt.Sprintf("\thttpServer %s.HttpServerInterface", t))
		data.WriteString(separator)
	case "cron":
		data.WriteString(fmt.Sprintf("\tcronScheduler %s.CronSchedulerInterface", t))
		data.WriteString(separator)
	}
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("var _ ApplicationInterface = (*Application)(nil)"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("func NewApplication(ctx context.Context) (*Application, error) {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\ta := &Application{}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\terr := a.InitializeDependency(ctx)"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\tif err != nil {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\treturn nil, err"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\treturn a, nil"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("func (a *Application) InitializeDependency(ctx context.Context) error {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tinits := []func(context.Context) error{"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\ta.InitializeConfig,"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\ta.InitializeLogger,"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\ta.InitializeDatabase,"))
	data.WriteString(separator)
	switch application.Type {
	case "grpc":
		data.WriteString(fmt.Sprintf("\t\ta.InitializeMigration,"))
		data.WriteString(separator)
	case "http":
		data.WriteString(fmt.Sprintf("\t\ta.InitializeMigration,"))
		data.WriteString(separator)
	}
	data.WriteString(fmt.Sprintf("\t\ta.InitializeProvider,"))
	data.WriteString(separator)
	switch application.Type {
	case "grpc":
		data.WriteString(fmt.Sprintf("\t\ta.InitializeGrpcServer,"))
		data.WriteString(separator)
	case "http":
		data.WriteString(fmt.Sprintf("\t\ta.InitializeHttpServer,"))
		data.WriteString(separator)
	case "cron":
		data.WriteString(fmt.Sprintf("\t\ta.InitializeCronScheduler,"))
		data.WriteString(separator)
	}
	data.WriteString(fmt.Sprintf("\t}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\tfor _, function := range inits {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\terr := function(ctx)"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\t\tif err != nil {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\t\treturn err"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\t}"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\treturn nil"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("func (a *Application) InitializeConfig(_ context.Context) error {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tvar c config.Config"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\terr := env.Parse(&c)"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\tif err != nil {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\treturn err"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\terr = env.Parse(&c.Database)"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\tif err != nil {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\treturn err"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t}"))
	data.WriteString(separator)
	data.WriteString(separator)

	switch application.Type {
	case "grpc":
		data.WriteString(fmt.Sprintf("\terr = env.Parse(&c.GrpcServer)"))
		data.WriteString(separator)
		data.WriteString(separator)

		data.WriteString(fmt.Sprintf("\tif err != nil {"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t\treturn err"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t}"))
		data.WriteString(separator)
		data.WriteString(separator)

		data.WriteString(fmt.Sprintf("\terr = env.Parse(&c.GrpcServer.Authentication)"))
		data.WriteString(separator)
		data.WriteString(separator)

		data.WriteString(fmt.Sprintf("\tif err != nil {"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t\treturn err"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t}"))
		data.WriteString(separator)
		data.WriteString(separator)
	case "http":
		data.WriteString(fmt.Sprintf("\terr = env.Parse(&c.HttpServer)"))
		data.WriteString(separator)
		data.WriteString(separator)

		data.WriteString(fmt.Sprintf("\tif err != nil {"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t\treturn err"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t}"))
		data.WriteString(separator)
		data.WriteString(separator)

		data.WriteString(fmt.Sprintf("\terr = env.Parse(&c.HttpServer.Authentication)"))
		data.WriteString(separator)
		data.WriteString(separator)

		data.WriteString(fmt.Sprintf("\tif err != nil {"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t\treturn err"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t}"))
		data.WriteString(separator)
		data.WriteString(separator)
	}

	data.WriteString(fmt.Sprintf("\ta.config = &c"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\treturn nil"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("func (a *Application) InitializeLogger(_ context.Context) error {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tlevel := logrus.InfoLevel"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\tif a.config.Debug {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\tlevel = logrus.DebugLevel"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\tformatter := &logrus.JSONFormatter{"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\tFieldMap: logrus.FieldMap{"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\t\tlogrus.FieldKeyMsg: \"message\","))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\t},"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\tl := logrus.New()"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\tl.Formatter = formatter"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\tl.SetLevel(level)"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tl.SetOutput(os.Stdout)"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tl.SetNoLock()"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\ta.logger = l"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\treturn nil"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("func (a *Application) InitializeDatabase(_ context.Context) error {"))
	data.WriteString(separator)

	switch application.Database {
	case "postgres":
		data.WriteString(fmt.Sprintf("\tdatabase, err := postgres.NewPostgresDatabase(a.config.Database)"))
		data.WriteString(separator)
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\tif err != nil {"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t\treturn err"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t}"))
		data.WriteString(separator)
		data.WriteString(separator)
	case "mysql":
		data.WriteString(fmt.Sprintf("\tdatabase, err := mysql.NewMySqlDatabase(a.config.Database)"))
		data.WriteString(separator)
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\tif err != nil {"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t\treturn err"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t}"))
		data.WriteString(separator)
		data.WriteString(separator)
	}

	data.WriteString(fmt.Sprintf("\ta.database = database"))
	data.WriteString(separator)
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\treturn nil"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)
	data.WriteString(separator)

	switch application.Type {
	case "grpc":
		data.WriteString(fmt.Sprintf("func (a *Application) InitializeMigration(_ context.Context) error {"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\terr := migration.Run(a.database)"))
		data.WriteString(separator)
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\tif err != nil {"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t\treturn err"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t}"))
		data.WriteString(separator)
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\treturn nil"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("}"))
		data.WriteString(separator)
		data.WriteString(separator)
	case "http":
		data.WriteString(fmt.Sprintf("func (a *Application) InitializeMigration(_ context.Context) error {"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\terr := migration.Run(a.database)"))
		data.WriteString(separator)
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\tif err != nil {"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t\treturn err"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t}"))
		data.WriteString(separator)
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\treturn nil"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("}"))
		data.WriteString(separator)
		data.WriteString(separator)
	}

	data.WriteString(fmt.Sprintf("func (a *Application) InitializeProvider(_ context.Context) error {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\ta.%sProvider = %sProvider.New%sProvider(", application.Names.LowerCamelCaseSingular, application.Names.LowerCamelCaseSingular, application.Names.CamelCaseSingular))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\ta.config,"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\ta.logger,"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\ta.database,"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t)"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\treturn nil"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)
	data.WriteString(separator)

	switch application.Type {
	case "grpc":
		data.WriteString(fmt.Sprintf("func (a *Application) InitializeGrpcServer(_ context.Context) error {"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\timplementation := a.%sProvider.Get%sImplementation()", application.Names.LowerCamelCaseSingular, application.Names.CamelCaseSingular))
		data.WriteString(separator)
		data.WriteString(separator)

		data.WriteString(fmt.Sprintf("\ts := %s.NewGrpcServer(", t))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t\ta.config.GrpcServer,"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t\ta.logger,"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t\timplementation,"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t)"))
		data.WriteString(separator)
		data.WriteString(separator)

		data.WriteString(fmt.Sprintf("\ta.grpcServer = s"))
		data.WriteString(separator)
		data.WriteString(separator)

		data.WriteString(fmt.Sprintf("\treturn nil"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("}"))
		data.WriteString(separator)
		data.WriteString(separator)
	case "http":
		data.WriteString(fmt.Sprintf("func (a *Application) InitializeHttpServer(_ context.Context) error {"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\thandler := a.%sProvider.Get%sHandler()", application.Names.LowerCamelCaseSingular, application.Names.CamelCaseSingular))
		data.WriteString(separator)
		data.WriteString(separator)

		data.WriteString(fmt.Sprintf("\ts := %s.NewHttpServer(", t))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t\ta.config.HttpServer,"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t\ta.logger,"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t\thandler,"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t)"))
		data.WriteString(separator)
		data.WriteString(separator)

		data.WriteString(fmt.Sprintf("\ta.httpServer = s"))
		data.WriteString(separator)
		data.WriteString(separator)

		data.WriteString(fmt.Sprintf("\treturn nil"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("}"))
		data.WriteString(separator)
		data.WriteString(separator)
	case "cron":
		data.WriteString(fmt.Sprintf("func (a *Application) InitializeCronScheduler(_ context.Context) error {"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\tservice := a.%sProvider.Get%sService()", application.Names.LowerCamelCaseSingular, application.Names.CamelCaseSingular))
		data.WriteString(separator)
		data.WriteString(separator)

		data.WriteString(fmt.Sprintf("\tc := %s.NewCronScheduler(", t))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t\ta.logger,"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t\tservice,"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t)"))
		data.WriteString(separator)
		data.WriteString(separator)

		data.WriteString(fmt.Sprintf("\ta.cronScheduler = c"))
		data.WriteString(separator)
		data.WriteString(separator)

		data.WriteString(fmt.Sprintf("\treturn nil"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("}"))
		data.WriteString(separator)
		data.WriteString(separator)
	}

	data.WriteString(fmt.Sprintf("func (a *Application) Run() error {"))
	data.WriteString(separator)

	switch application.Type {
	case "grpc":
		data.WriteString(fmt.Sprintf("\terr := a.grpcServer.Run()"))
		data.WriteString(separator)
		data.WriteString(separator)

		data.WriteString(fmt.Sprintf("\tif err != nil {"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t\treturn err"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t}"))
		data.WriteString(separator)
		data.WriteString(separator)

		data.WriteString(fmt.Sprintf("\treturn nil"))
		data.WriteString(separator)
	case "http":
		data.WriteString(fmt.Sprintf("\terr := a.httpServer.Run()"))
		data.WriteString(separator)
		data.WriteString(separator)

		data.WriteString(fmt.Sprintf("\tif err != nil {"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t\treturn err"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t}"))
		data.WriteString(separator)
		data.WriteString(separator)

		data.WriteString(fmt.Sprintf("\treturn nil"))
		data.WriteString(separator)
	case "cron":
		data.WriteString(fmt.Sprintf("\terr := a.cronScheduler.Run()"))
		data.WriteString(separator)
		data.WriteString(separator)

		data.WriteString(fmt.Sprintf("\tif err != nil {"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t\treturn err"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t}"))
		data.WriteString(separator)
		data.WriteString(separator)

		data.WriteString(fmt.Sprintf("\treturn nil"))
		data.WriteString(separator)
	}

	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)

	return data.Bytes()
}

func GetMainTemplate(application *dto.ApplicationDto) []byte {
	data := bytes.Buffer{}
	separator := util.GetSeparator()

	imports := []string{
		"\"log\"",
		"\"context\"",
		fmt.Sprintf("\"%s/cmd/application\"", application.Module),
	}

	sort.Strings(imports)

	data.WriteString(fmt.Sprintf("package main"))
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

	data.WriteString(fmt.Sprintf("func main() {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tctx := context.Background()"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\ta, err := application.NewApplication(ctx)"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\tif err != nil {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprint("\t\tlog.Panicf(\"an error occurred while starting the application %v\", err)"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\terr = a.Run()"))
	data.WriteString(separator)
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tif err != nil {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprint("\t\tlog.Panicf(\"an application error occurred %v\", err)"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t}"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)

	return data.Bytes()
}

func GetMockGenerateShellScriptTemplate() []byte {
	data := bytes.Buffer{}
	separator := util.GetSeparator()

	data.WriteString(fmt.Sprintf("#!/bin/bash"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("if [ \"$#\" -lt 1 ]; then"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\techo \"Usage: $0 <MOCKS_OUTPUT_DIRECTORY> <MOCKS_FILES>\""))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\texit 1"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("fi"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("MOCKS_OUTPUT_DIRECTORY=\"$1\""))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("MOCKS_FILES=\"${*:2}\""))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("for mock_file in $MOCKS_FILES; do"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tDIRECTORY=$(dirname \"$mock_file\")"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tFILENAME=$(basename \"$mock_file\")"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\tEXTENSION=\"${FILENAME##*.}\""))
	data.WriteString(separator)
	data.WriteString(fmt.Sprint("\tFILENAME_WITHOUT_EXTENSIONS=\"${FILENAME%.*}\""))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tOUTPUT_PATH=\"$DIRECTORY/$MOCKS_OUTPUT_DIRECTORY/${FILENAME_WITHOUT_EXTENSIONS}_mock.$EXTENSION\""))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\tmkdir -p \"$DIRECTORY/$MOCKS_OUTPUT_DIRECTORY\""))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\techo \"Generating mock file for $mock_file...\""))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\tmockgen -source=\"$mock_file\" -destination=\"$OUTPUT_PATH\""))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("done"))
	data.WriteString(separator)

	return data.Bytes()
}

func GetGitIgnoreTemplate() []byte {
	data := bytes.Buffer{}
	separator := util.GetSeparator()

	data.WriteString(fmt.Sprintf("# JetBrains"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf(".idea"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("# Visual Studio Code"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf(".vscode"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("# Build"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("build"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("# Environment"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf(".env"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("# Mocks"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("*_mock.go"))
	data.WriteString(separator)

	return data.Bytes()
}

func GetDockerIgnoreTemplate() []byte {
	data := bytes.Buffer{}
	separator := util.GetSeparator()

	data.WriteString(fmt.Sprintf("# JetBrains"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf(".idea"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("# Visual Studio Code"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf(".vscode"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("# Build"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("build"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("# Mocks"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("*_mock.go"))
	data.WriteString(separator)

	return data.Bytes()
}

func GetDockerTemplate(application *dto.ApplicationDto, port bool) []byte {
	data := bytes.Buffer{}
	separator := util.GetSeparator()

	data.WriteString(fmt.Sprintf("FROM golang:%s-alpine", application.Version))
	data.WriteString(separator)
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("WORKDIR /usr/local/application"))
	data.WriteString(separator)
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("RUN apk update"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("RUN apk upgrade"))
	data.WriteString(separator)
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("RUN apk add git openssh gcc libc-dev ca-certificates"))
	data.WriteString(separator)
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("RUN mkdir -p ~/.ssh"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("RUN chmod 600 ~/.ssh"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("RUN ssh-keyscan %s >> ~/.ssh/known_hosts", application.Organization))
	data.WriteString(separator)
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("RUN git config --global url.\"git@%s:\".insteadOf \"https://%s/\"", application.Organization, application.Organization))
	data.WriteString(separator)
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("ENV PATH=\"$PATH:$GOROOT/bin:$GOPATH/bin\""))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("ENV GOPRIVATE=%s", application.Organization))
	data.WriteString(separator)
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("COPY go.mod go.mod"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("COPY go.sum go.sum"))
	data.WriteString(separator)
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("RUN --mount=type=ssh go mod download"))
	data.WriteString(separator)
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("COPY . ."))
	data.WriteString(separator)
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("RUN go build -o build/main main.go"))
	data.WriteString(separator)
	data.WriteString(separator)

	if port {
		data.WriteString(fmt.Sprintf("EXPOSE 3000"))
		data.WriteString(separator)
		data.WriteString(separator)
	}

	data.WriteString(fmt.Sprintf("CMD [\"./build/main\"]"))

	return data.Bytes()
}

func GetDockerComposeTemplate(application *dto.ApplicationDto, port bool) []byte {
	data := bytes.Buffer{}
	separator := util.GetSeparator()

	data.WriteString(fmt.Sprintf("version: \"3.8\""))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("services:"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("  %s:", application.Directory))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("    image: \"registry.%s:latest\"", application.Module))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("    container_name: \"%s\"", application.Directory))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("    hostname: \"%s\"", application.Directory))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("    restart: \"always\""))
	data.WriteString(separator)
	if port {
		data.WriteString(fmt.Sprintf("    ports:"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("      - \"3000:3000\""))
		data.WriteString(separator)
	}
	data.WriteString(fmt.Sprintf("    env_file:"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("      - \".env\""))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("    build:"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("      context: \".\""))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("      dockerfile: \"application.dockerfile\""))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("      ssh: [\"default\"]"))
	data.WriteString(separator)

	return data.Bytes()
}

func GetEnvironmentTemplate(application *dto.ApplicationDto) []byte {
	data := bytes.Buffer{}
	separator := util.GetSeparator()

	data.WriteString(fmt.Sprintf("# Debug"))
	data.WriteString(separator)
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("DEBUG=true"))
	data.WriteString(separator)
	data.WriteString(separator)

	switch application.Type {
	case "grpc":
		data.WriteString(fmt.Sprintf("# Grpc"))
		data.WriteString(separator)
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("GRPC_SERVER_HOSTNAME=localhost"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("GRPC_SERVER_PORT=3000"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("GRPC_SERVER_AUTHENTICATION_TOKEN_HEADER=authentication-header"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("GRPC_SERVER_AUTHENTICATION_TOKEN=authentication-token"))
		data.WriteString(separator)
		data.WriteString(separator)
	case "http":
		data.WriteString(fmt.Sprintf("# Http"))
		data.WriteString(separator)
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("HTTP_SERVER_HOSTNAME=localhost"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("HTTP_SERVER_PORT=3000"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("HTTP_SERVER_AUTHENTICATION_TOKEN_HEADER=authentication-header"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("HTTP_SERVER_AUTHENTICATION_TOKEN=authentication-token"))
		data.WriteString(separator)
		data.WriteString(separator)
	}

	switch application.Database {
	case "mysql":
		data.WriteString(fmt.Sprintf("# MySql"))
		data.WriteString(separator)
		data.WriteString(separator)

		data.WriteString(fmt.Sprintf("MYSQL_HOSTNAME=localhost"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("MYSQL_PORT=3306"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("MYSQL_USERNAME=mysql"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("MYSQL_PASSWORD=mysql"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("MYSQL_DATABASE=database"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("MYSQL_MAX_OPEN_CONNECTIONS=10"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("MYSQL_MAX_IDLE_CONNECTIONS=10"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("MYSQL_CONNECTION_MAX_LIFE_TIME=600000000000"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("MYSQL_CONNECTION_MAX_IDLE_TIME=600000000000"))
		data.WriteString(separator)
	case "postgres":
		data.WriteString(fmt.Sprintf("# Postgres"))
		data.WriteString(separator)
		data.WriteString(separator)

		data.WriteString(fmt.Sprintf("POSTGRES_HOSTNAME=localhost"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("POSTGRES_PORT=5432"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("POSTGRES_USERNAME=postgres"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("POSTGRES_PASSWORD=postgres"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("POSTGRES_DATABASE=database"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("POSTGRES_MAX_OPEN_CONNECTIONS=10"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("POSTGRES_MAX_IDLE_CONNECTIONS=10"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("POSTGRES_CONNECTION_MAX_LIFE_TIME=600000000000"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("POSTGRES_CONNECTION_MAX_IDLE_TIME=600000000000"))
		data.WriteString(separator)
	}

	return data.Bytes()
}

func GetGoTemplate(application *dto.ApplicationDto) []byte {
	data := bytes.Buffer{}
	separator := util.GetSeparator()

	data.WriteString(fmt.Sprintf("module %s", application.Module))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("go %s", application.Version))
	data.WriteString(separator)

	return data.Bytes()
}

func GetMakefileTemplate(application *dto.ApplicationDto) []byte {
	data := bytes.Buffer{}
	separator := util.GetSeparator()

	data.WriteString(fmt.Sprintf("# Variables"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("DOCKERFILE=\"application.dockerfile\""))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("CONTAINER_TAG=\"registry.%s:latest\"", application.Module))
	data.WriteString(separator)
	data.WriteString(separator)

	switch application.Type {
	case "grpc":
		data.WriteString(fmt.Sprintf("PROTO_SOURCE_DIRECTORY = api"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("PROTO_OUTPUT_DIRECTORY = pkg"))
		data.WriteString(separator)
		data.WriteString(separator)

		data.WriteString(fmt.Sprintf("PROTO_FILES = \\"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t%s_v1/%s.proto", application.Names.SnakeCasePlural, application.Names.SnakeCasePlural))
		data.WriteString(separator)
		data.WriteString(separator)
	}

	data.WriteString(fmt.Sprintf("MOCKS_OUTPUT_DIRECTORY = mocks"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("MOCKS_FILES = \\"))
	data.WriteString(separator)
	data.WriteString(separator)

	switch application.Type {
	case "grpc":
		data.WriteString(fmt.Sprintf("# GRPC"))
		data.WriteString(separator)
		data.WriteString(separator)

		data.WriteString(fmt.Sprintf("grpc-generate:"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t@echo \"Generating GRPC...\""))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t@chmod +x bin/grpc-generate.sh"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t@bin/grpc-generate.sh $(PROTO_SOURCE_DIRECTORY) $(PROTO_OUTPUT_DIRECTORY) $(PROTO_FILES)"))
		data.WriteString(separator)
		data.WriteString(separator)
	}

	data.WriteString(fmt.Sprintf("# Mocks"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("mock-generate:"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t@echo \"Generating Mocks...\""))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t@chmod +x bin/mock-generate.sh"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t@bin/mock-generate.sh $(MOCKS_OUTPUT_DIRECTORY) $(MOCKS_FILES)"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("# Generate"))
	data.WriteString(separator)
	data.WriteString(separator)

	switch application.Type {
	case "grpc":
		data.WriteString(fmt.Sprintf("generate: grpc-generate mock-generate"))
	default:
		data.WriteString(fmt.Sprintf("generate: mock-generate"))
	}

	data.WriteString(separator)
	data.WriteString(separator)

	switch application.Type {
	case "grpc":
		data.WriteString(fmt.Sprintf("# Migration create"))
		data.WriteString(separator)
		data.WriteString(separator)

		data.WriteString(fmt.Sprintf("migration-create:"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t@echo \"Migration create...\""))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t@mkdir -p ${MIGRATION_DIRECTORY}"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t@goose -dir ${MIGRATION_DIRECTORY} create $(MIGRATION_NAME) sql"))
		data.WriteString(separator)
		data.WriteString(separator)
	case "http":
		data.WriteString(fmt.Sprintf("# Migration create"))
		data.WriteString(separator)
		data.WriteString(separator)

		data.WriteString(fmt.Sprintf("migration-create:"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t@echo \"Migration create...\""))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t@mkdir -p ${MIGRATION_DIRECTORY}"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t@goose -dir ${MIGRATION_DIRECTORY} create $(MIGRATION_NAME) sql"))
		data.WriteString(separator)
		data.WriteString(separator)
	}

	data.WriteString(fmt.Sprintf("# Tidy"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("tidy:"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t@echo \"Tidy dependencies...\""))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t@go mod tidy"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("# Download"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("download:"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t@echo \"Downloading dependencies...\""))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t@go mod download"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("# Setup"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("setup:"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t@make generate"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t@make tidy"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t@make download"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("# Test"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("test:"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t@echo \"Running test...\""))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t@go template -v ./..."))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("# Docker build"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("docker-build:"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tdocker build --ssh default . --file ${DOCKERFILE} --tag ${CONTAINER_TAG}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("# Docker push"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("docker-push:"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tdocker image push ${CONTAINER_TAG}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("# Docker compose run"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("docker-compose-run:"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tdocker compose --file docker-compose-application.yml up --detach --build"))
	data.WriteString(separator)

	return data.Bytes()
}

func GetReadmeTemplate(application *dto.ApplicationDto) []byte {
	data := bytes.Buffer{}
	separator := util.GetSeparator()

	data.WriteString(fmt.Sprintf("%s", application.Directory))
	data.WriteString(separator)

	return data.Bytes()
}
