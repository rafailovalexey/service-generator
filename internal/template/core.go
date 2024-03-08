package template

import (
	"bytes"
	"fmt"
	"github.com/rafailovalexey/service-generator/internal/dto"
	"github.com/rafailovalexey/service-generator/internal/util"
	"sort"
)

func GetApplicationTemplate(module string, application string, database string, name *dto.NameDto) []byte {
	data := bytes.Buffer{}
	separator := util.GetSeparator()

	t := ""

	switch application {
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
		"\"fmt\"",
		"\"os\"",
		"\"strconv\"",
		"\"time\"",
		"\"github.com/sirupsen/logrus\"",
		fmt.Sprintf("\"%s/database/%s\"", module, database),
		fmt.Sprintf("\"%s/config\"", module),
		fmt.Sprintf("\"%s/cmd/%s\"", module, t),
		fmt.Sprintf("\"%s/cmd/migration\"", module),
		fmt.Sprintf("\"%s/internal/provider\"", module),
		fmt.Sprintf("%sProvider \"%s/internal/provider/%s\"", name.LowerCamelCaseSingular, module, name.SnakeCasePlural),
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
	data.WriteString(fmt.Sprintf("\tInitializeMigration(context.Context) error"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tInitializeProvider(context.Context) error"))
	data.WriteString(separator)
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
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tdatabase *sql.DB"))
	data.WriteString(separator)
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t%sProvider provider.%sProviderInterface", name.LowerCamelCaseSingular, name.CamelCaseSingular))
	data.WriteString(separator)
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
	data.WriteString(fmt.Sprintf("\t\ta.InitializeMigration,"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\ta.InitializeProvider,"))
	data.WriteString(separator)
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
	data.WriteString(fmt.Sprintf("\t\treturn err"))
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
	data.WriteString(fmt.Sprintf("\tdebugString := os.Getenv(\"DEBUG\")"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\tif debugString == \"\" {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\treturn fmt.Errorf(\"specify the debug value\")"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\tdebug, err := strconv.ParseBool(debugString)"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\tif err != nil {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\treturn fmt.Errorf(\"error parsing debug value: %s\", err)"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t}"))
	data.WriteString(separator)
	data.WriteString(separator)

	switch database {
	case "postgres":
		data.WriteString(fmt.Sprintf("\tusername := os.Getenv(\"POSTGRES_USERNAME\")"))
		data.WriteString(separator)
		data.WriteString(separator)

		data.WriteString(fmt.Sprintf("\tif username == \"\" {"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t\treturn fmt.Errorf(\"specify the username for the database\")"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t}"))
		data.WriteString(separator)
		data.WriteString(separator)

		data.WriteString(fmt.Sprintf("\tpassword := os.Getenv(\"POSTGRES_PASSWORD\")"))
		data.WriteString(separator)
		data.WriteString(separator)

		data.WriteString(fmt.Sprintf("\tif password == \"\" {"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t\treturn fmt.Errorf(\"specify the password for the database\")"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t}"))
		data.WriteString(separator)
		data.WriteString(separator)

		data.WriteString(fmt.Sprintf("\thostname := os.Getenv(\"POSTGRES_HOSTNAME\")"))
		data.WriteString(separator)
		data.WriteString(separator)

		data.WriteString(fmt.Sprintf("\tif hostname == \"\" {"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t\treturn fmt.Errorf(\"specify the hostname for the database\")"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t}"))
		data.WriteString(separator)
		data.WriteString(separator)

		data.WriteString(fmt.Sprintf("\tport := os.Getenv(\"POSTGRES_PORT\")"))
		data.WriteString(separator)
		data.WriteString(separator)

		data.WriteString(fmt.Sprintf("\tif port == \"\" {"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t\treturn fmt.Errorf(\"specify the port for the database\")"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t}"))
		data.WriteString(separator)
		data.WriteString(separator)

		data.WriteString(fmt.Sprintf("\tdb := os.Getenv(\"POSTGRES_DATABASE\")"))
		data.WriteString(separator)
		data.WriteString(separator)

		data.WriteString(fmt.Sprintf("\tif db == \"\" {"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t\treturn fmt.Errorf(\"specify the database for the database\")"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t}"))
		data.WriteString(separator)
		data.WriteString(separator)

		data.WriteString(fmt.Sprintf("\tmaxIdleConnectionsString := os.Getenv(\"POSTGRES_MAX_IDLE_CONNECTIONS\")"))
		data.WriteString(separator)
		data.WriteString(separator)

		data.WriteString(fmt.Sprintf("\tif maxIdleConnectionsString == \"\" {"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t\treturn fmt.Errorf(\"specify the max idle connections for the database\")"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t}"))
		data.WriteString(separator)
		data.WriteString(separator)

		data.WriteString(fmt.Sprintf("\tmaxIdleConnections, err := strconv.Atoi(maxIdleConnectionsString)"))
		data.WriteString(separator)
		data.WriteString(separator)

		data.WriteString(fmt.Sprintf("\tif err != nil {"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t\treturn fmt.Errorf(\"error parsing max idle connections value: %s\", err)"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t}"))
		data.WriteString(separator)
		data.WriteString(separator)

		data.WriteString(fmt.Sprintf("\tmaxOpenConnectionsString := os.Getenv(\"POSTGRES_MAX_OPEN_CONNECTIONS\")"))
		data.WriteString(separator)
		data.WriteString(separator)

		data.WriteString(fmt.Sprintf("\tif maxIdleConnectionsString == \"\" {"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t\treturn fmt.Errorf(\"specify the max open connections for the database\")"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t}"))
		data.WriteString(separator)
		data.WriteString(separator)

		data.WriteString(fmt.Sprintf("\tmaxOpenConnections, err := strconv.Atoi(maxOpenConnectionsString)"))
		data.WriteString(separator)
		data.WriteString(separator)

		data.WriteString(fmt.Sprintf("\tif err != nil {"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t\treturn fmt.Errorf(\"error parsing max open connections value: %s\", err)"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t}"))
		data.WriteString(separator)
		data.WriteString(separator)

		data.WriteString(fmt.Sprintf("\tconnectionMaxIdleTimeString := os.Getenv(\"POSTGRES_CONNECTION_MAX_IDLE_TIME\")"))
		data.WriteString(separator)
		data.WriteString(separator)

		data.WriteString(fmt.Sprintf("\tif connectionMaxIdleTimeString == \"\" {"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t\treturn fmt.Errorf(\"specify the connection max idle time for the database\")"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t}"))
		data.WriteString(separator)
		data.WriteString(separator)

		data.WriteString(fmt.Sprintf("\tconnectionMaxIdleTime, err := strconv.Atoi(connectionMaxIdleTimeString)"))
		data.WriteString(separator)
		data.WriteString(separator)

		data.WriteString(fmt.Sprintf("\tif err != nil {"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t\treturn fmt.Errorf(\"error parsing connection max idle time value: %s\", err)"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t}"))
		data.WriteString(separator)
		data.WriteString(separator)

		data.WriteString(fmt.Sprintf("\tconnectionMaxLifeTimeString := os.Getenv(\"POSTGRES_CONNECTION_MAX_LIFE_TIME\")"))
		data.WriteString(separator)
		data.WriteString(separator)

		data.WriteString(fmt.Sprintf("\tif connectionMaxLifeTimeString == \"\" {"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t\treturn fmt.Errorf(\"specify the connection max life time for the database\")"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t}"))
		data.WriteString(separator)
		data.WriteString(separator)

		data.WriteString(fmt.Sprintf("\tconnectionMaxLifeTime, err := strconv.Atoi(connectionMaxLifeTimeString)"))
		data.WriteString(separator)
		data.WriteString(separator)

		data.WriteString(fmt.Sprintf("\tif err != nil {"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t\treturn fmt.Errorf(\"error parsing connection max life time value: %s\", err)"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t}"))
		data.WriteString(separator)
		data.WriteString(separator)

		data.WriteString(fmt.Sprintf("\tp := &postgres.PostgresDatabaseConfig{"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t\tHostname:\thostname,"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t\tPort:\tport,"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t\tUsername:\tusername,"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t\tPassword:\tpassword,"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t\tDatabase:\tdb,"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t\tMaxIdleConnections:\tmaxIdleConnections,"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t\tMaxOpenConnections:\tmaxOpenConnections,"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t\tConnectionMaxLifeTime:\ttime.Duration(connectionMaxLifeTime),"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t\tConnectionMaxIdleTime:\ttime.Duration(connectionMaxIdleTime),"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t}"))
		data.WriteString(separator)
		data.WriteString(separator)
	case "mysql":
		data.WriteString(fmt.Sprintf("\tusername := os.Getenv(\"MYSQL_USERNAME\")"))
		data.WriteString(separator)
		data.WriteString(separator)

		data.WriteString(fmt.Sprintf("\tif username == \"\" {"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t\treturn fmt.Errorf(\"specify the username for the database\")"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t}"))
		data.WriteString(separator)
		data.WriteString(separator)

		data.WriteString(fmt.Sprintf("\tpassword := os.Getenv(\"MYSQL_PASSWORD\")"))
		data.WriteString(separator)
		data.WriteString(separator)

		data.WriteString(fmt.Sprintf("\tif password == \"\" {"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t\treturn fmt.Errorf(\"specify the password for the database\")"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t}"))
		data.WriteString(separator)
		data.WriteString(separator)

		data.WriteString(fmt.Sprintf("\thostname := os.Getenv(\"MYSQL_HOSTNAME\")"))
		data.WriteString(separator)
		data.WriteString(separator)

		data.WriteString(fmt.Sprintf("\tif hostname == \"\" {"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t\treturn fmt.Errorf(\"specify the hostname for the database\")"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t}"))
		data.WriteString(separator)
		data.WriteString(separator)

		data.WriteString(fmt.Sprintf("\tport := os.Getenv(\"MYSQL_PORT\")"))
		data.WriteString(separator)
		data.WriteString(separator)

		data.WriteString(fmt.Sprintf("\tif port == \"\" {"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t\treturn fmt.Errorf(\"specify the port for the database\")"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t}"))
		data.WriteString(separator)
		data.WriteString(separator)

		data.WriteString(fmt.Sprintf("\tdb := os.Getenv(\"MYSQL_DATABASE\")"))
		data.WriteString(separator)
		data.WriteString(separator)

		data.WriteString(fmt.Sprintf("\tif db == \"\" {"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t\treturn fmt.Errorf(\"specify the database for the database\")"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t}"))
		data.WriteString(separator)
		data.WriteString(separator)

		data.WriteString(fmt.Sprintf("\tmaxIdleConnectionsString := os.Getenv(\"MYSQL_MAX_IDLE_CONNECTIONS\")"))
		data.WriteString(separator)
		data.WriteString(separator)

		data.WriteString(fmt.Sprintf("\tif maxIdleConnectionsString == \"\" {"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t\treturn fmt.Errorf(\"specify the max idle connections for the database\")"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t}"))
		data.WriteString(separator)
		data.WriteString(separator)

		data.WriteString(fmt.Sprintf("\tmaxIdleConnections, err := strconv.Atoi(maxIdleConnectionsString)"))
		data.WriteString(separator)
		data.WriteString(separator)

		data.WriteString(fmt.Sprintf("\tif err != nil {"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t\treturn fmt.Errorf(\"error parsing max idle connections value: %s\", err)"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t}"))
		data.WriteString(separator)
		data.WriteString(separator)

		data.WriteString(fmt.Sprintf("\tmaxOpenConnectionsString := os.Getenv(\"MYSQL_MAX_OPEN_CONNECTIONS\")"))
		data.WriteString(separator)
		data.WriteString(separator)

		data.WriteString(fmt.Sprintf("\tif maxIdleConnectionsString == \"\" {"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t\treturn fmt.Errorf(\"specify the max open connections for the database\")"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t}"))
		data.WriteString(separator)
		data.WriteString(separator)

		data.WriteString(fmt.Sprintf("\tmaxOpenConnections, err := strconv.Atoi(maxOpenConnectionsString)"))
		data.WriteString(separator)
		data.WriteString(separator)

		data.WriteString(fmt.Sprintf("\tif err != nil {"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t\treturn fmt.Errorf(\"error parsing max open connections value: %s\", err)"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t}"))
		data.WriteString(separator)
		data.WriteString(separator)

		data.WriteString(fmt.Sprintf("\tconnectionMaxIdleTimeString := os.Getenv(\"MYSQL_CONNECTION_MAX_IDLE_TIME\")"))
		data.WriteString(separator)
		data.WriteString(separator)

		data.WriteString(fmt.Sprintf("\tif connectionMaxIdleTimeString == \"\" {"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t\treturn fmt.Errorf(\"specify the connection max idle time for the database\")"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t}"))
		data.WriteString(separator)
		data.WriteString(separator)

		data.WriteString(fmt.Sprintf("\tconnectionMaxIdleTime, err := strconv.Atoi(connectionMaxIdleTimeString)"))
		data.WriteString(separator)
		data.WriteString(separator)

		data.WriteString(fmt.Sprintf("\tif err != nil {"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t\treturn fmt.Errorf(\"error parsing connection max idle time value: %s\", err)"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t}"))
		data.WriteString(separator)
		data.WriteString(separator)

		data.WriteString(fmt.Sprintf("\tconnectionMaxLifeTimeString := os.Getenv(\"MYSQL_CONNECTION_MAX_LIFE_TIME\")"))
		data.WriteString(separator)
		data.WriteString(separator)

		data.WriteString(fmt.Sprintf("\tif connectionMaxLifeTimeString == \"\" {"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t\treturn fmt.Errorf(\"specify the connection max life time for the database\")"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t}"))
		data.WriteString(separator)
		data.WriteString(separator)

		data.WriteString(fmt.Sprintf("\tconnectionMaxLifeTime, err := strconv.Atoi(connectionMaxLifeTimeString)"))
		data.WriteString(separator)
		data.WriteString(separator)

		data.WriteString(fmt.Sprintf("\tif err != nil {"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t\treturn fmt.Errorf(\"error parsing connection max life time value: %s\", err)"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t}"))
		data.WriteString(separator)
		data.WriteString(separator)

		data.WriteString(fmt.Sprintf("\tm := &mysql.MySqlDatabaseConfig{"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t\tHostname:\thostname,"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t\tPort:\tport,"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t\tUsername:\tusername,"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t\tPassword:\tpassword,"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t\tDatabase:\tdb,"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t\tMaxIdleConnections:\tmaxIdleConnections,"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t\tMaxOpenConnections:\tmaxOpenConnections,"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t\tConnectionMaxLifeTime:\ttime.Duration(connectionMaxLifeTime),"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t\tConnectionMaxIdleTime:\ttime.Duration(connectionMaxIdleTime),"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t}"))
		data.WriteString(separator)
		data.WriteString(separator)
	}

	data.WriteString(fmt.Sprintf("\tc := &config.Config{"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\tDebug:\tdebug,"))
	data.WriteString(separator)
	switch database {
	case "postgres":
		data.WriteString(fmt.Sprintf("\t\tDatabase:\tp,"))
		data.WriteString(separator)
	case "mysql":
		data.WriteString(fmt.Sprintf("\t\tDatabase:\tm,"))
		data.WriteString(separator)
	}
	data.WriteString(fmt.Sprintf("\t}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\ta.config = c"))
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

	switch database {
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

	data.WriteString(fmt.Sprintf("func (a *Application) InitializeProvider(_ context.Context) error {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\ta.%sProvider = %sProvider.New%sProvider()", name.LowerCamelCaseSingular, name.LowerCamelCaseSingular, name.CamelCaseSingular))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\treturn nil"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("func (a *Application) Run() error {"))
	data.WriteString(separator)

	switch application {
	case "grpc":
		data.WriteString(fmt.Sprintf("\timplementation := a.%sProvider.Get%sImplementation()", name.LowerCamelCaseSingular, name.CamelCaseSingular))
		data.WriteString(separator)
		data.WriteString(separator)

		data.WriteString(fmt.Sprintf("\terr := %s.Run(a.config, implementation)", t))
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
		data.WriteString(fmt.Sprintf("\thandler := a.%sProvider.Get%sHandler()", name.LowerCamelCaseSingular, name.CamelCaseSingular))
		data.WriteString(separator)
		data.WriteString(separator)

		data.WriteString(fmt.Sprintf("\terr := %s.Run(handler)", t))
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
		data.WriteString(fmt.Sprintf("\tservice := a.%sProvider.Get%sService()", name.LowerCamelCaseSingular, name.CamelCaseSingular))
		data.WriteString(separator)
		data.WriteString(separator)

		data.WriteString(fmt.Sprintf("\terr := %s.Run(service)", t))
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
	default:
		data.WriteString(fmt.Sprintf("\terr := %s.Run()", t))
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

func GetMainTemplate(module string) []byte {
	data := bytes.Buffer{}
	separator := util.GetSeparator()

	imports := []string{
		"\"log\"",
		"\"context\"",
		fmt.Sprintf("\"%s/cmd/application\"", module),
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
	data.WriteString(fmt.Sprintf("\t\tlog.Panicf(\"an application error occurred %v\\n\", err)"))
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

func GetDockerTemplate(organization string, version string, port bool) []byte {
	data := bytes.Buffer{}
	separator := util.GetSeparator()

	data.WriteString(fmt.Sprintf("FROM golang:%s-alpine", version))
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
	data.WriteString(fmt.Sprintf("RUN ssh-keyscan %s >> ~/.ssh/known_hosts", organization))
	data.WriteString(separator)
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("RUN git config --global url.\"git@%s:\".insteadOf \"https://%s/\"", organization, organization))
	data.WriteString(separator)
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("ENV PATH=\"$PATH:$GOROOT/bin:$GOPATH/bin\""))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("ENV GOPRIVATE=%s", organization))
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

func GetEnvironmentTemplate(application string, database string) []byte {
	data := bytes.Buffer{}
	separator := util.GetSeparator()

	data.WriteString(fmt.Sprintf("# DEBUG"))
	data.WriteString(separator)
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("DEBUG=true"))
	data.WriteString(separator)
	data.WriteString(separator)

	switch application {
	case "grpc":
		data.WriteString(fmt.Sprintf("# GRPC"))
		data.WriteString(separator)
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("HOSTNAME=localhost"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("PORT=3000"))
		data.WriteString(separator)
		data.WriteString(separator)
	case "http":
		data.WriteString(fmt.Sprintf("# HTTP"))
		data.WriteString(separator)
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("HOSTNAME=localhost"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("PORT=3000"))
		data.WriteString(separator)
		data.WriteString(separator)
	}

	switch database {
	case "mysql":
		data.WriteString(fmt.Sprintf("# MySQL"))
		data.WriteString(separator)
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("MYSQL_USERNAME=mysql"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("MYSQL_PASSWORD=mysql"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("MYSQL_HOSTNAME=localhost"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("MYSQL_PORT=3306"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("MYSQL_DATABASE=database"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("MYSQL_MAX_IDLE_CONNECTIONS=10"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("MYSQL_MAX_OPEN_CONNECTIONS=10"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("MYSQL_CONNECTION_MAX_IDLE_TIME=600000000000"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("MYSQL_CONNECTION_MAX_LIFE_TIME=600000000000"))
		data.WriteString(separator)
	case "postgres":
		data.WriteString(fmt.Sprintf("# Postgres"))
		data.WriteString(separator)
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("POSTGRES_USERNAME=mysql"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("POSTGRES_PASSWORD=mysql"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("POSTGRES_HOSTNAME=localhost"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("POSTGRES_PORT=5432"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("POSTGRES_DATABASE=database"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("POSTGRES_MAX_IDLE_CONNECTIONS=10"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("POSTGRES_MAX_OPEN_CONNECTIONS=10"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("POSTGRES_CONNECTION_MAX_IDLE_TIME=600000000000"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("POSTGRES_CONNECTION_MAX_LIFE_TIME=600000000000"))
		data.WriteString(separator)
	}

	return data.Bytes()
}

func GetGoTemplate(module string, version string) []byte {
	data := bytes.Buffer{}
	separator := util.GetSeparator()

	data.WriteString(fmt.Sprintf("module %s", module))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("go %s", version))
	data.WriteString(separator)

	return data.Bytes()
}

func GetMakefileTemplate(module string, application string, name *dto.NameDto) []byte {
	data := bytes.Buffer{}
	separator := util.GetSeparator()

	data.WriteString(fmt.Sprintf("# Variables"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("DOCKERFILE=\"application.dockerfile\""))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("CONTAINER_TAG=\"registry.%s:latest\"", module))
	data.WriteString(separator)
	data.WriteString(separator)

	switch application {
	case "grpc":
		data.WriteString(fmt.Sprintf("PROTO_SOURCE_DIRECTORY = api"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("PROTO_OUTPUT_DIRECTORY = pkg"))
		data.WriteString(separator)
		data.WriteString(separator)

		data.WriteString(fmt.Sprintf("PROTO_FILES = \\"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t%s_v1/%s.proto", name.SnakeCasePlural, name.SnakeCasePlural))
		data.WriteString(separator)
		data.WriteString(separator)
	}

	data.WriteString(fmt.Sprintf("MOCKS_OUTPUT_DIRECTORY = mocks"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("MOCKS_FILES = \\"))
	data.WriteString(separator)
	data.WriteString(separator)

	switch application {
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

	switch application {
	case "grpc":
		data.WriteString(fmt.Sprintf("generate: grpc-generate mock-generate"))
	default:
		data.WriteString(fmt.Sprintf("generate: mock-generate"))
	}

	data.WriteString(separator)
	data.WriteString(separator)

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

	return data.Bytes()
}

func GetReadmeTemplate(module string) []byte {
	data := bytes.Buffer{}
	separator := util.GetSeparator()

	data.WriteString(fmt.Sprintf("%s", module))
	data.WriteString(separator)

	return data.Bytes()
}
