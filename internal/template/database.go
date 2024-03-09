package template

import (
	"bytes"
	"fmt"
	"github.com/rafailovalexey/service-generator/internal/util"
	"sort"
)

func GetDatabasePostgresTemplate() []byte {
	data := bytes.Buffer{}
	separator := util.GetSeparator()

	imports := []string{
		"\"database/sql\"",
		"\"fmt\"",
		"\"time\"",
		"_ \"github.com/lib/pq\"",
	}

	sort.Strings(imports)

	data.WriteString(fmt.Sprintf("package postgres"))
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

	data.WriteString(fmt.Sprintf("type PostgresDatabaseConfig struct {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tHostname string `mapstructure:\"POSTGRES_HOSTNAME\"`"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tPort int `mapstructure:\"POSTGRES_PORT\"`"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tUsername string `mapstructure:\"POSTGRES_USERNAME\"`"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tPassword string `mapstructure:\"POSTGRES_PASSWORD\"`"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tDatabase string `mapstructure:\"POSTGRES_DATABASE\"`"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tMaxOpenConnections int `mapstructure:\"POSTGRES_MAX_IDLE_CONNECTIONS\"`"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tMaxIdleConnections int `mapstructure:\"POSTGRES_MAX_OPEN_CONNECTIONS\"`"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tConnectionMaxLifeTime int `mapstructure:\"POSTGRES_CONNECTION_MAX_LIFE_TIME\"`"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tConnectionMaxIdleTime int `mapstructure:\"POSTGRES_CONNECTION_MAX_IDLE_TIME\"`"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("func NewPostgresDatabase(config *PostgresDatabaseConfig) (*sql.DB, error) {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tdsn := fmt.Sprintf("))
	data.WriteString(separator)
	data.WriteString(fmt.Sprint("\t\t\"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable\","))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\tconfig.Hostname,"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\tconfig.Port,"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\tconfig.Username,"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\tconfig.Password,"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\tconfig.Database,"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t)"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\tdb, err := sql.Open(\"postgres\", dsn)"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\tif err != nil {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\treturn nil, err"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\tdb.SetMaxOpenConns(config.MaxOpenConnections)"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tdb.SetMaxIdleConns(config.MaxIdleConnections)"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tdb.SetConnMaxLifetime(time.Duration(config.ConnectionMaxLifeTime))"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tdb.SetConnMaxIdleTime(time.Duration(config.ConnectionMaxIdleTime))"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\terr = db.Ping()"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\tif err != nil {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\treturn nil, err"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\treturn db, nil"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)

	return data.Bytes()
}

func GetDatabaseMySqlTemplate() []byte {
	data := bytes.Buffer{}
	separator := util.GetSeparator()

	imports := []string{
		"\"database/sql\"",
		"\"fmt\"",
		"\"time\"",
		"_ \"github.com/go-sql-driver/mysql\"",
	}

	sort.Strings(imports)

	data.WriteString(fmt.Sprintf("package mysql"))
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

	data.WriteString(fmt.Sprintf("type MySqlDatabaseConfig struct {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tHostname string `mapstructure:\"MYSQL_HOSTNAME\"`"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tPort int `mapstructure:\"MYSQL_PORT\"`"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tUsername string `mapstructure:\"MYSQL_USERNAME\"`"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tPassword string `mapstructure:\"MYSQL_PASSWORD\"`"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tDatabase string `mapstructure:\"MYSQL_DATABASE\"`"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tMaxOpenConnections int `mapstructure:\"MYSQL_MAX_IDLE_CONNECTIONS\"`"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tMaxIdleConnections int `mapstructure:\"MYSQL_MAX_OPEN_CONNECTIONS\"`"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tConnectionMaxLifeTime int `mapstructure:\"MYSQL_CONNECTION_MAX_LIFE_TIME\"`"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tConnectionMaxIdleTime int `mapstructure:\"MYSQL_CONNECTION_MAX_IDLE_TIME\"`"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("func NewMySqlDatabase(config *MySqlDatabaseConfig) (*sql.DB, error) {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tdsn := fmt.Sprintf("))
	data.WriteString(separator)
	data.WriteString(fmt.Sprint("\t\t\"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable\","))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\tconfig.Hostname,"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\tconfig.Port,"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\tconfig.Username,"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\tconfig.Password,"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\tconfig.Database,"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t)"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\tdb, err := sql.Open(\"mysql\", dsn)"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\tif err != nil {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\treturn nil, err"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\tdb.SetMaxOpenConns(config.MaxOpenConnections)"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tdb.SetMaxIdleConns(config.MaxIdleConnections)"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tdb.SetConnMaxLifetime(time.Duration(config.ConnectionMaxLifeTime))"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tdb.SetConnMaxIdleTime(time.Duration(config.ConnectionMaxIdleTime))"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\terr = db.Ping()"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\tif err != nil {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\treturn nil, err"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\treturn db, nil"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)

	return data.Bytes()
}
