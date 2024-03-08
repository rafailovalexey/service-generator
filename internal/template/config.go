package template

import (
	"bytes"
	"fmt"
	"github.com/rafailovalexey/service-generator/internal/util"
	"sort"
)

func GetConfigTemplate(module string, application string, database string) []byte {
	data := bytes.Buffer{}
	separator := util.GetSeparator()

	imports := []string{}

	switch database {
	case "postgres":
		imports = append(imports, fmt.Sprintf("\"%s/database/%s\"", module, "postgres"))
	case "mysql":
		imports = append(imports, fmt.Sprintf("\"%s/database/%s\"", module, "mysql"))
	}

	switch application {
	case "grpc":
		imports = append(imports, fmt.Sprintf("\"%s/cmd/grpc_server\"", module))
	case "http":
		imports = append(imports, fmt.Sprintf("\"%s/cmd/http_server\"", module))
	}

	sort.Strings(imports)

	data.WriteString(fmt.Sprintf("package config"))
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

	data.WriteString(fmt.Sprintf("type Config struct {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tDebug\tbool"))
	data.WriteString(separator)
	switch application {
	case "grpc":
		data.WriteString(fmt.Sprintf("\tGrpcServer\t*grpc_server.GrpcServerConfig"))
		data.WriteString(separator)
	case "http":
		data.WriteString(fmt.Sprintf("\tHttpServer\t*http_server.HttpServerConfig"))
		data.WriteString(separator)
	}
	switch database {
	case "postgres":
		data.WriteString(fmt.Sprintf("\tDatabase\t*postgres.PostgresDatabaseConfig"))
		data.WriteString(separator)
	case "mysql":
		data.WriteString(fmt.Sprintf("\tDatabase\t*mysql.MySqlDatabaseConfig"))
		data.WriteString(separator)
	}
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)

	return data.Bytes()
}
