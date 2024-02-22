package template

import (
	"bytes"
	"fmt"
	"github.com/rafailovalexey/service-generator/internal/util"
	"sort"
)

func GetDatabaseTemplate() []byte {
	data := bytes.Buffer{}
	separator := util.GetSeparator()

	imports := []string{
		"\"database/sql\"",
	}

	sort.Strings(imports)

	data.WriteString(fmt.Sprintf("package database"))
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

	data.WriteString(fmt.Sprintf("type DatabaseInterface interface {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tConnect() (*sql.DB, error)"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tPing()"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tGetDsn() string"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("type Database struct {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tdsn string"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)

	return data.Bytes()
}

func GetDatabaseMySQLTemplate() []byte {
	data := bytes.Buffer{}
	separator := util.GetSeparator()

	imports := []string{
		"\"fmt\"",
		"\"log\"",
		"\"os\"",
		"\"database/sql\"",
		"_ \"github.com/go-sql-driver/mysql\"",
	}

	sort.Strings(imports)

	data.WriteString(fmt.Sprintf("package database"))
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

	data.WriteString(fmt.Sprintf("var _ DatabaseInterface = (*Database)(nil)"))
	data.WriteString(separator)
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("func NewMySQLDatabase() *Database {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tusername := os.Getenv(\"MYSQL_USERNAME\")"))
	data.WriteString(separator)
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tif username == \"\" {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\tlog.Panicln(\"specify the username for the database\")"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t}"))
	data.WriteString(separator)
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tpassword := os.Getenv(\"MYSQL_PASSWORD\")"))
	data.WriteString(separator)
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tif password == \"\" {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\tlog.Panicln(\"specify the password for the database\")"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t}"))
	data.WriteString(separator)
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\thostname := os.Getenv(\"MYSQL_HOSTNAME\")"))
	data.WriteString(separator)
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tif hostname == \"\" {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\tlog.Panicln(\"specify the hostname for the database\")"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t}"))
	data.WriteString(separator)
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tport := os.Getenv(\"MYSQL_PORT\")"))
	data.WriteString(separator)
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tif port == \"\" {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\tlog.Panicln(\"specify the port for the database\")"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t}"))
	data.WriteString(separator)
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tdatabase := os.Getenv(\"MYSQL_DATABASE\")"))
	data.WriteString(separator)
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tif database == \"\" {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\tlog.Panicln(\"specify the database for the database\")"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t}"))
	data.WriteString(separator)
	data.WriteString(separator)
	data.WriteString(fmt.Sprint("\tdsn := fmt.Sprintf(\"%s:%s@tcp(%s:%s)/%s\", username, password, hostname, port, database)"))
	data.WriteString(separator)
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\td := &Database{"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\tdsn: dsn,"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t}"))
	data.WriteString(separator)
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\td.Ping()"))
	data.WriteString(separator)
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\treturn d"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("func (d *Database) Connect() (*sql.DB, error) {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tdatabase, err := sql.Open(\"mysql\", d.dsn)"))
	data.WriteString(separator)
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tif err != nil {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\treturn nil, err"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t}"))
	data.WriteString(separator)
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\treturn database, nil"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("func (d *Database) Ping() {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tdatabase, err := sql.Open(\"mysql\", d.dsn)"))
	data.WriteString(separator)
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tif err != nil {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\tlog.Panicln(err.Error())"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t}"))
	data.WriteString(separator)
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\terr = database.Ping()"))
	data.WriteString(separator)
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tif err != nil {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\tlog.Panicln(err.Error())"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t}"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("func (d *Database) GetDsn() string {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\treturn d.dsn"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)

	return data.Bytes()
}
