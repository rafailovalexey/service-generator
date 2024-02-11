package template

//func GetDockerWithPortTemplate() []byte {
//	data := bytes.Buffer{}
//	separator := utils.GetSeparator()
//
//	data.WriteString(fmt.Sprintf("FROM golang:latest"))
//	data.WriteString(separator)
//	data.WriteString(separator)
//
//	data.WriteString(fmt.Sprintf("WORKDIR /usr/local/application"))
//	data.WriteString(separator)
//	data.WriteString(separator)
//
//	data.WriteString(fmt.Sprintf("COPY . ."))
//	data.WriteString(separator)
//	data.WriteString(separator)
//
//	data.WriteString(fmt.Sprintf("RUN apt-get update --yes"))
//	data.WriteString(separator)
//	data.WriteString(fmt.Sprintf("RUN apt-get upgrade --yes"))
//	data.WriteString(separator)
//	data.WriteString(separator)
//
//	data.WriteString(fmt.Sprintf("RUN apt-get install --yes make"))
//	data.WriteString(separator)
//	data.WriteString(separator)
//
//	data.WriteString(fmt.Sprintf("RUN export PATH=\"$PATH:$(go env GOPATH)/bin\""))
//	data.WriteString(separator)
//	data.WriteString(separator)
//
//	data.WriteString(fmt.Sprintf("RUN make download"))
//	data.WriteString(separator)
//	data.WriteString(fmt.Sprintf("RUN make build"))
//	data.WriteString(separator)
//	data.WriteString(separator)
//
//	data.WriteString(fmt.Sprintf("EXPOSE 3000"))
//	data.WriteString(separator)
//	data.WriteString(separator)
//
//	data.WriteString(fmt.Sprintf("CMD [\"./build/main\"]"))
//	data.WriteString(separator)
//
//	return data.Bytes()
//}

//func GetDockerWithoutPortTemplate() []byte {
//	data := bytes.Buffer{}
//	separator := utils.GetSeparator()
//
//	data.WriteString(fmt.Sprintf("FROM golang:latest"))
//	data.WriteString(separator)
//	data.WriteString(separator)
//
//	data.WriteString(fmt.Sprintf("WORKDIR /usr/local/application"))
//	data.WriteString(separator)
//	data.WriteString(separator)
//
//	data.WriteString(fmt.Sprintf("COPY . ."))
//	data.WriteString(separator)
//	data.WriteString(separator)
//
//	data.WriteString(fmt.Sprintf("RUN apt-get update --yes"))
//	data.WriteString(separator)
//	data.WriteString(fmt.Sprintf("RUN apt-get upgrade --yes"))
//	data.WriteString(separator)
//	data.WriteString(separator)
//
//	data.WriteString(fmt.Sprintf("RUN apt-get install --yes make"))
//	data.WriteString(separator)
//	data.WriteString(separator)
//
//	data.WriteString(fmt.Sprintf("RUN export PATH=\"$PATH:$(go env GOPATH)/bin\""))
//	data.WriteString(separator)
//	data.WriteString(separator)
//
//	data.WriteString(fmt.Sprintf("RUN make download"))
//	data.WriteString(separator)
//	data.WriteString(fmt.Sprintf("RUN make build"))
//	data.WriteString(separator)
//	data.WriteString(separator)
//
//	data.WriteString(fmt.Sprintf("CMD [\"./build/main\"]"))
//	data.WriteString(separator)
//
//	return data.Bytes()
//}

//func GetApplicationTemplate(module string, application string, name string, implementing string) []byte {
//	data := bytes.Buffer{}
//	separator := utils.GetSeparator()
//
//	imports := []string{
//		"\"context\"",
//		"\"github.com/joho/godotenv\"",
//		fmt.Sprintf("\"%s/cmd/%s\"", module, implementing),
//		fmt.Sprintf("\"%s/internal/provider\"", module),
//		fmt.Sprintf("%sProvider \"%s/internal/provider/%s\"", utils.SingularForm(name), module, name),
//	}
//
//	sort.Strings(imports)
//
//	data.WriteString(fmt.Sprintf("package %s", application))
//	data.WriteString(separator)
//	data.WriteString(separator)
//
//	data.WriteString(fmt.Sprintf("import ("))
//	data.WriteString(separator)
//
//	for _, i := range imports {
//		data.WriteString(fmt.Sprintf("\t%s", i))
//		data.WriteString(separator)
//	}
//
//	data.WriteString(fmt.Sprintf(")"))
//	data.WriteString(separator)
//	data.WriteString(separator)
//
//	data.WriteString(fmt.Sprintf("type %sInterface interface {", utils.Capitalize(application)))
//	data.WriteString(separator)
//	data.WriteString(fmt.Sprintf("\tInitializeDependency(context.Context) error"))
//	data.WriteString(separator)
//	data.WriteString(fmt.Sprintf("\tInitializeEnvironment(context.Context) error"))
//	data.WriteString(separator)
//	data.WriteString(fmt.Sprintf("\tInitializeProvider(context.Context) error"))
//	data.WriteString(separator)
//	data.WriteString(fmt.Sprintf("\tRun()"))
//	data.WriteString(separator)
//	data.WriteString(fmt.Sprintf("}"))
//	data.WriteString(separator)
//	data.WriteString(separator)
//
//	data.WriteString(fmt.Sprintf("type %s struct {", utils.Capitalize(application)))
//	data.WriteString(separator)
//	data.WriteString(fmt.Sprintf("\t%sProvider provider.%sProviderInterface", utils.SingularForm(name), utils.Capitalize(utils.SingularForm(name))))
//	data.WriteString(separator)
//	data.WriteString(fmt.Sprintf("}"))
//	data.WriteString(separator)
//	data.WriteString(separator)
//
//	data.WriteString(fmt.Sprintf("var _ %sInterface = (*%s)(nil)", utils.Capitalize(application), utils.Capitalize(application)))
//	data.WriteString(separator)
//	data.WriteString(separator)
//
//	data.WriteString(fmt.Sprintf("func New%s(ctx context.Context) (*%s, error) {", utils.Capitalize(application), utils.Capitalize(application)))
//	data.WriteString(separator)
//	data.WriteString(fmt.Sprintf("\t%s := &%s{}", utils.FirstLetter(application), utils.Capitalize(application)))
//	data.WriteString(separator)
//	data.WriteString(separator)
//
//	data.WriteString(fmt.Sprintf("\terr := %s.InitializeDependency(ctx)", utils.FirstLetter(application)))
//	data.WriteString(separator)
//	data.WriteString(separator)
//
//	data.WriteString(fmt.Sprintf("\tif err != nil {"))
//	data.WriteString(separator)
//	data.WriteString(fmt.Sprintf("\t\treturn nil, err"))
//	data.WriteString(separator)
//	data.WriteString(fmt.Sprintf("\t}"))
//	data.WriteString(separator)
//	data.WriteString(separator)
//
//	data.WriteString(fmt.Sprintf("\treturn %s, nil", utils.FirstLetter(application)))
//	data.WriteString(separator)
//	data.WriteString(fmt.Sprintf("}"))
//	data.WriteString(separator)
//	data.WriteString(separator)
//
//	data.WriteString(fmt.Sprintf("func (%s *%s) InitializeDependency(ctx context.Context) error {", utils.FirstLetter(application), utils.Capitalize(application)))
//	data.WriteString(separator)
//	data.WriteString(fmt.Sprintf("\tinits := []func(context.Context) error{"))
//	data.WriteString(separator)
//	data.WriteString(fmt.Sprintf("\t\t%s.InitializeEnvironment,", utils.FirstLetter(application)))
//	data.WriteString(separator)
//	data.WriteString(fmt.Sprintf("\t\t%s.InitializeProvider,", utils.FirstLetter(application)))
//	data.WriteString(separator)
//	data.WriteString(fmt.Sprintf("\t}"))
//	data.WriteString(separator)
//	data.WriteString(separator)
//
//	data.WriteString(fmt.Sprintf("\tfor _, function := range inits {"))
//	data.WriteString(separator)
//	data.WriteString(fmt.Sprintf("\t\terr := function(ctx)"))
//	data.WriteString(separator)
//	data.WriteString(separator)
//
//	data.WriteString(fmt.Sprintf("\t\tif err != nil {"))
//	data.WriteString(separator)
//	data.WriteString(fmt.Sprintf("\t\treturn err"))
//	data.WriteString(separator)
//	data.WriteString(fmt.Sprintf("\t\t}"))
//	data.WriteString(separator)
//	data.WriteString(fmt.Sprintf("\t}"))
//	data.WriteString(separator)
//	data.WriteString(separator)
//
//	data.WriteString(fmt.Sprintf("\treturn nil"))
//	data.WriteString(separator)
//	data.WriteString(fmt.Sprintf("}"))
//	data.WriteString(separator)
//	data.WriteString(separator)
//
//	data.WriteString(fmt.Sprintf("func (%s *%s) InitializeEnvironment(_ context.Context) error {", utils.FirstLetter(application), utils.Capitalize(application)))
//	data.WriteString(separator)
//	data.WriteString(fmt.Sprintf("\terr := godotenv.Load(\".env\")"))
//	data.WriteString(separator)
//	data.WriteString(separator)
//
//	data.WriteString(fmt.Sprintf("\tif err != nil {"))
//	data.WriteString(separator)
//	data.WriteString(fmt.Sprintf("\t\treturn err"))
//	data.WriteString(separator)
//	data.WriteString(fmt.Sprintf("\t}"))
//	data.WriteString(separator)
//	data.WriteString(separator)
//
//	data.WriteString(fmt.Sprintf("\treturn nil"))
//	data.WriteString(separator)
//	data.WriteString(fmt.Sprintf("}"))
//	data.WriteString(separator)
//	data.WriteString(separator)
//
//	data.WriteString(fmt.Sprintf("func (%s *%s) InitializeProvider(_ context.Context) error {", utils.FirstLetter(application), utils.Capitalize(application)))
//	data.WriteString(separator)
//	data.WriteString(fmt.Sprintf("\t%s.%sProvider = %sProvider.New%sProvider()", utils.FirstLetter(application), utils.SingularForm(name), utils.SingularForm(name), utils.Capitalize(utils.SingularForm(name))))
//	data.WriteString(separator)
//	data.WriteString(separator)
//
//	data.WriteString(fmt.Sprintf("\treturn nil"))
//	data.WriteString(separator)
//	data.WriteString(fmt.Sprintf("}"))
//	data.WriteString(separator)
//	data.WriteString(separator)
//
//	data.WriteString(fmt.Sprintf("func (%s *%s) Run() {", utils.FirstLetter(application), utils.Capitalize(application)))
//	data.WriteString(separator)
//
//	switch implementing {
//	case "grpc_server":
//		data.WriteString(fmt.Sprintf("\timplementation := %s.%sProvider.Get%sImplementation()", utils.FirstLetter(application), utils.SingularForm(name), utils.Capitalize(utils.SingularForm(name))))
//		data.WriteString(separator)
//		data.WriteString(separator)
//
//		data.WriteString(fmt.Sprintf("\t%s.Run(implementation)", implementing))
//		data.WriteString(separator)
//	case "http_server":
//		data.WriteString(fmt.Sprintf("\thandler := %s.%sProvider.Get%sHandler()", utils.FirstLetter(application), utils.SingularForm(name), utils.Capitalize(utils.SingularForm(name))))
//		data.WriteString(separator)
//		data.WriteString(separator)
//
//		data.WriteString(fmt.Sprintf("\t%s.Run(handler)", implementing))
//		data.WriteString(separator)
//	case "nats_subscribe":
//		data.WriteString(fmt.Sprintf("\tcontroller := %s.%sProvider.Get%sController()", utils.FirstLetter(application), utils.SingularForm(name), utils.Capitalize(utils.SingularForm(name))))
//		data.WriteString(separator)
//		data.WriteString(separator)
//
//		data.WriteString(fmt.Sprintf("\t%s.Run(controller)", implementing))
//		data.WriteString(separator)
//	case "cron_schedule":
//		data.WriteString(fmt.Sprintf("\tservice  := %s.%sProvider.Get%sService()", utils.FirstLetter(application), utils.SingularForm(name), utils.Capitalize(utils.SingularForm(name))))
//		data.WriteString(separator)
//		data.WriteString(separator)
//
//		data.WriteString(fmt.Sprintf("\t%s.Run(service)", implementing))
//		data.WriteString(separator)
//	default:
//		data.WriteString(fmt.Sprintf("\t%s.Run()", implementing))
//		data.WriteString(separator)
//	}
//
//	data.WriteString(fmt.Sprintf("}"))
//	data.WriteString(separator)
//
//	return data.Bytes()
//}

//func GetCronScheduleTemplate(module string, name string) []byte {
//	data := bytes.Buffer{}
//	separator := utils.GetSeparator()
//
//	imports := []string{
//		"\"log\"",
//		"\"os\"",
//		"\"os/signal\"",
//		"\"syscall\"",
//		"\"github.com/robfig/cron/v3\"",
//		fmt.Sprintf("\"%s/internal/service\"", module),
//	}
//
//	sort.Strings(imports)
//
//	data.WriteString(fmt.Sprintf("package cron_schedule"))
//	data.WriteString(separator)
//	data.WriteString(separator)
//
//	data.WriteString(fmt.Sprintf("import ("))
//	data.WriteString(separator)
//
//	for _, i := range imports {
//		data.WriteString(fmt.Sprintf("\t%s", i))
//		data.WriteString(separator)
//	}
//
//	data.WriteString(fmt.Sprintf(")"))
//	data.WriteString(separator)
//	data.WriteString(separator)
//
//	data.WriteString(fmt.Sprintf("func Run(_ service.%sServiceInterface) {", utils.Capitalize(utils.SingularForm(name))))
//	data.WriteString(separator)
//	data.WriteString(fmt.Sprintf("\tc := cron.New()"))
//	data.WriteString(separator)
//	data.WriteString(separator)
//
//	data.WriteString(fmt.Sprintf("\tlog.Printf(\"cron started\\n\")"))
//	data.WriteString(separator)
//	data.WriteString(separator)
//
//	data.WriteString(fmt.Sprintf("\tc.Start()"))
//	data.WriteString(separator)
//	data.WriteString(separator)
//
//	data.WriteString(fmt.Sprintf("\tdefer c.Stop()"))
//	data.WriteString(separator)
//	data.WriteString(separator)
//
//	data.WriteString(fmt.Sprintf("\texit := make(chan os.Signal)"))
//	data.WriteString(separator)
//	data.WriteString(fmt.Sprintf("\tsignal.Notify(exit, syscall.SIGINT)"))
//	data.WriteString(separator)
//	data.WriteString(fmt.Sprintf("\t<-exit"))
//	data.WriteString(separator)
//	data.WriteString(separator)
//
//	data.WriteString(fmt.Sprintf("\tlog.Printf(\"cron stopped\\n\")"))
//	data.WriteString(separator)
//	data.WriteString(fmt.Sprintf("}"))
//	data.WriteString(separator)
//
//	return data.Bytes()
//}

//func GetUtilsConvertErrorTemplate() []byte {
//	data := bytes.Buffer{}
//	separator := utils.GetSeparator()
//
//	imports := []string{
//		"\"encoding/json\"",
//	}
//
//	sort.Strings(imports)
//
//	data.WriteString(fmt.Sprintf("package utils"))
//	data.WriteString(separator)
//	data.WriteString(separator)
//
//	data.WriteString(fmt.Sprintf("import ("))
//	data.WriteString(separator)
//
//	for _, i := range imports {
//		data.WriteString(fmt.Sprintf("\t%s", i))
//		data.WriteString(separator)
//	}
//
//	data.WriteString(fmt.Sprintf(")"))
//	data.WriteString(separator)
//	data.WriteString(separator)
//
//	data.WriteString(fmt.Sprintf("type ConverterError struct {"))
//	data.WriteString(separator)
//	data.WriteString(fmt.Sprintf("\tError string `json:\"error\"`"))
//	data.WriteString(separator)
//	data.WriteString(fmt.Sprintf("}"))
//	data.WriteString(separator)
//	data.WriteString(separator)
//
//	data.WriteString(fmt.Sprintf("func ConvertError(message string) []byte {"))
//	data.WriteString(separator)
//	data.WriteString(fmt.Sprintf("\tconvert := &ConverterError{"))
//	data.WriteString(separator)
//	data.WriteString(fmt.Sprintf("\t\tError: message,"))
//	data.WriteString(separator)
//	data.WriteString(fmt.Sprintf("\t}"))
//	data.WriteString(separator)
//	data.WriteString(separator)
//
//	data.WriteString(fmt.Sprintf("\tresult, err := json.Marshal(convert)"))
//	data.WriteString(separator)
//	data.WriteString(separator)
//
//	data.WriteString(fmt.Sprintf("\tif err != nil {"))
//	data.WriteString(separator)
//	data.WriteString(fmt.Sprintf("\t\treturn []byte(err.Error())"))
//	data.WriteString(separator)
//	data.WriteString(fmt.Sprintf("\t}"))
//	data.WriteString(separator)
//	data.WriteString(separator)
//
//	data.WriteString(fmt.Sprintf("\treturn result"))
//	data.WriteString(separator)
//	data.WriteString(fmt.Sprintf("}"))
//	data.WriteString(separator)
//
//	return data.Bytes()
//}

//func GetUtilsResponseTemplate() []byte {
//	data := bytes.Buffer{}
//	separator := utils.GetSeparator()
//
//	imports := []string{
//		"\"net/http\"",
//	}
//
//	sort.Strings(imports)
//
//	data.WriteString(fmt.Sprintf("package utils"))
//	data.WriteString(separator)
//	data.WriteString(separator)
//
//	data.WriteString(fmt.Sprintf("import ("))
//	data.WriteString(separator)
//
//	for _, i := range imports {
//		data.WriteString(fmt.Sprintf("\t%s", i))
//		data.WriteString(separator)
//	}
//
//	data.WriteString(fmt.Sprintf(")"))
//	data.WriteString(separator)
//	data.WriteString(separator)
//
//	data.WriteString(fmt.Sprintf("func ResponseBadRequest(response http.ResponseWriter, message string) {"))
//	data.WriteString(separator)
//	data.WriteString(fmt.Sprintf("\tresponse.Header().Set(\"Content-Type\", \"application/json\")"))
//	data.WriteString(separator)
//	data.WriteString(fmt.Sprintf("\tresponse.WriteHeader(http.StatusBadRequest)"))
//	data.WriteString(separator)
//	data.WriteString(fmt.Sprintf("\tresponse.Write(ConvertError(message))"))
//	data.WriteString(separator)
//	data.WriteString(fmt.Sprintf("}"))
//	data.WriteString(separator)
//	data.WriteString(separator)
//
//	data.WriteString(fmt.Sprintf("func ResponseNotFound(response http.ResponseWriter, request *http.Request) {"))
//	data.WriteString(separator)
//	data.WriteString(fmt.Sprintf("\tresponse.Header().Set(\"Content-Type\", \"application/json\")"))
//	data.WriteString(separator)
//	data.WriteString(fmt.Sprintf("\tresponse.WriteHeader(http.StatusNotFound)"))
//	data.WriteString(separator)
//	data.WriteString(fmt.Sprintf("\tresponse.Write(ConvertError(\"not found\"))"))
//	data.WriteString(separator)
//	data.WriteString(fmt.Sprintf("}"))
//	data.WriteString(separator)
//	data.WriteString(separator)
//
//	data.WriteString(fmt.Sprintf("func ResponseMethodNotAllowed(response http.ResponseWriter, request *http.Request) {"))
//	data.WriteString(separator)
//	data.WriteString(fmt.Sprintf("\tresponse.Header().Set(\"Content-Type\", \"application/json\")"))
//	data.WriteString(separator)
//	data.WriteString(fmt.Sprintf("\tresponse.WriteHeader(http.StatusMethodNotAllowed)"))
//	data.WriteString(separator)
//	data.WriteString(fmt.Sprintf("\tresponse.Write(ConvertError(\"method not allowed\"))"))
//	data.WriteString(separator)
//	data.WriteString(fmt.Sprintf("}"))
//	data.WriteString(separator)
//
//	return data.Bytes()
//}

//func GetUtilsDatabaseTemplate() []byte {
//	data := bytes.Buffer{}
//	separator := utils.GetSeparator()
//
//	imports := []string{
//		"\"os\"",
//		"\"fmt\"",
//		"\"log\"",
//		"\"context\"",
//		"\"github.com/jackc/pgx/v4/pgxpool\"",
//	}
//
//	sort.Strings(imports)
//
//	data.WriteString(fmt.Sprintf("package utils"))
//	data.WriteString(separator)
//	data.WriteString(separator)
//
//	data.WriteString(fmt.Sprintf("import ("))
//	data.WriteString(separator)
//
//	for _, i := range imports {
//		data.WriteString(fmt.Sprintf("\t%s", i))
//		data.WriteString(separator)
//	}
//
//	data.WriteString(fmt.Sprintf(")"))
//	data.WriteString(separator)
//	data.WriteString(separator)
//
//	data.WriteString(fmt.Sprintf("type DatabaseInterface interface {"))
//	data.WriteString(separator)
//	data.WriteString(fmt.Sprintf("\tInitialize()"))
//	data.WriteString(separator)
//	data.WriteString(fmt.Sprintf("\tGetPool() *pgxpool.Pool"))
//	data.WriteString(separator)
//	data.WriteString(fmt.Sprintf("}"))
//	data.WriteString(separator)
//	data.WriteString(separator)
//
//	data.WriteString(fmt.Sprintf("type Database struct {"))
//	data.WriteString(separator)
//	data.WriteString(fmt.Sprintf("\tcredentials string"))
//	data.WriteString(separator)
//	data.WriteString(fmt.Sprintf("}"))
//	data.WriteString(separator)
//	data.WriteString(separator)
//
//	data.WriteString(fmt.Sprintf("var _ DatabaseInterface = (*Database)(nil)"))
//	data.WriteString(separator)
//	data.WriteString(separator)
//
//	data.WriteString(fmt.Sprintf("func NewDatabase() *Database {"))
//	data.WriteString(separator)
//	data.WriteString(fmt.Sprintf("\tdb := &Database{}"))
//	data.WriteString(separator)
//	data.WriteString(separator)
//
//	data.WriteString(fmt.Sprintf("\tdb.Initialize()"))
//	data.WriteString(separator)
//	data.WriteString(separator)
//
//	data.WriteString(fmt.Sprintf("\treturn db"))
//	data.WriteString(separator)
//	data.WriteString(fmt.Sprintf("}"))
//	data.WriteString(separator)
//	data.WriteString(separator)
//
//	data.WriteString(fmt.Sprintf("func (d *Database) Initialize() {"))
//	data.WriteString(separator)
//	data.WriteString(fmt.Sprintf("\tusername := os.Getenv(\"POSTGRESQL_USERNAME\")"))
//	data.WriteString(separator)
//	data.WriteString(separator)
//
//	data.WriteString(fmt.Sprintf("\tif username == \"\" {"))
//	data.WriteString(separator)
//	data.WriteString(fmt.Sprintf("\t\tlog.Panicf(\"specify the database user\")"))
//	data.WriteString(separator)
//	data.WriteString(fmt.Sprintf("\t}"))
//	data.WriteString(separator)
//	data.WriteString(separator)
//
//	data.WriteString(fmt.Sprintf("\tpassword := os.Getenv(\"POSTGRESQL_PASSWORD\")"))
//	data.WriteString(separator)
//	data.WriteString(separator)
//
//	data.WriteString(fmt.Sprintf("\tif password == \"\" {"))
//	data.WriteString(separator)
//	data.WriteString(fmt.Sprintf("\t\tlog.Panicf(\"specify the database user password\")"))
//	data.WriteString(separator)
//	data.WriteString(fmt.Sprintf("\t}"))
//	data.WriteString(separator)
//	data.WriteString(separator)
//
//	data.WriteString(fmt.Sprintf("\tdb := os.Getenv(\"POSTGRESQL_DATABASE\")"))
//	data.WriteString(separator)
//	data.WriteString(separator)
//
//	data.WriteString(fmt.Sprintf("\tif db == \"\" {"))
//	data.WriteString(separator)
//	data.WriteString(fmt.Sprintf("\t\tlog.Panicf(\"ndicate the name of the database\")"))
//	data.WriteString(separator)
//	data.WriteString(fmt.Sprintf("\t}"))
//	data.WriteString(separator)
//	data.WriteString(separator)
//
//	data.WriteString(fmt.Sprintf("\thostname := os.Getenv(\"POSTGRESQL_HOSTNAME\")"))
//	data.WriteString(separator)
//	data.WriteString(separator)
//
//	data.WriteString(fmt.Sprintf("\tif hostname == \"\" {"))
//	data.WriteString(separator)
//	data.WriteString(fmt.Sprintf("\t\tlog.Panicf(\"specify the database hostname\")"))
//	data.WriteString(separator)
//	data.WriteString(fmt.Sprintf("\t}"))
//	data.WriteString(separator)
//	data.WriteString(separator)
//
//	data.WriteString(fmt.Sprintf("\tport := os.Getenv(\"POSTGRESQL_PORT\")"))
//	data.WriteString(separator)
//	data.WriteString(separator)
//
//	data.WriteString(fmt.Sprintf("\tif port == \"\" {"))
//	data.WriteString(separator)
//	data.WriteString(fmt.Sprintf("\t\tlog.Panicf(\"specify the database port\")"))
//	data.WriteString(separator)
//	data.WriteString(fmt.Sprintf("\t}"))
//	data.WriteString(separator)
//	data.WriteString(separator)
//
//	data.WriteString(fmt.Sprintf("\tsslmode := os.Getenv(\"POSTGRESQL_SSLMODE\")"))
//	data.WriteString(separator)
//	data.WriteString(separator)
//
//	data.WriteString(fmt.Sprintf("\tif sslmode == \"\" {"))
//	data.WriteString(separator)
//	data.WriteString(fmt.Sprintf("\t\tlog.Panicf(\"specify the ssl mode of the database\")"))
//	data.WriteString(separator)
//	data.WriteString(fmt.Sprintf("\t}"))
//	data.WriteString(separator)
//	data.WriteString(separator)
//
//	data.WriteString(fmt.Sprint("\td.credentials = fmt.Sprintf(\"user=%s password=%s dbname=%s host=%s port=%s sslmode=%s\", username, password, db, hostname, port, sslmode)"))
//	data.WriteString(separator)
//	data.WriteString(fmt.Sprintf("}"))
//	data.WriteString(separator)
//	data.WriteString(separator)
//
//	data.WriteString(fmt.Sprintf("func (d *Database) GetPool() *pgxpool.Pool {"))
//	data.WriteString(separator)
//	data.WriteString(fmt.Sprintf("\tpool, err := pgxpool.Connect(context.Background(), d.credentials)"))
//	data.WriteString(separator)
//	data.WriteString(separator)
//
//	data.WriteString(fmt.Sprintf("\tif err != nil {"))
//	data.WriteString(separator)
//	data.WriteString(fmt.Sprint("\t\tlog.Panicf(\"error %v\\n\", err)"))
//	data.WriteString(separator)
//	data.WriteString(fmt.Sprintf("\t}"))
//	data.WriteString(separator)
//	data.WriteString(separator)
//
//	data.WriteString(fmt.Sprintf("\treturn pool"))
//	data.WriteString(separator)
//	data.WriteString(fmt.Sprintf("}"))
//	data.WriteString(separator)
//
//	return data.Bytes()
//}
