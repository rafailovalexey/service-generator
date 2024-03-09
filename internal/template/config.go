package template

import (
	"bytes"
	"fmt"
	"github.com/rafailovalexey/service-generator/internal/dto"
	"github.com/rafailovalexey/service-generator/internal/util"
	"sort"
)

func GetConfigTemplate(application *dto.ApplicationDto) []byte {
	data := bytes.Buffer{}
	separator := util.GetSeparator()

	imports := []string{}

	switch application.Database {
	case "postgres":
		imports = append(imports, fmt.Sprintf("\"%s/database/%s\"", application.Module, "postgres"))
	case "mysql":
		imports = append(imports, fmt.Sprintf("\"%s/database/%s\"", application.Module, "mysql"))
	}

	switch application.Type {
	case "grpc":
		imports = append(imports, fmt.Sprintf("\"%s/cmd/grpc_server\"", application.Module))
	case "http":
		imports = append(imports, fmt.Sprintf("\"%s/cmd/http_server\"", application.Module))
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
	data.WriteString(fmt.Sprintf("\tDebug bool `env:\"DEBUG,required\"`"))
	data.WriteString(separator)
	switch application.Type {
	case "grpc":
		data.WriteString(fmt.Sprintf("\tGrpcServer grpc_server.GrpcServerConfig"))
		data.WriteString(separator)
	case "http":
		data.WriteString(fmt.Sprintf("\tHttpServer http_server.HttpServerConfig"))
		data.WriteString(separator)
	}
	switch application.Database {
	case "postgres":
		data.WriteString(fmt.Sprintf("\tDatabase postgres.PostgresDatabaseConfig"))
		data.WriteString(separator)
	case "mysql":
		data.WriteString(fmt.Sprintf("\tDatabase mysql.MySqlDatabaseConfig"))
		data.WriteString(separator)
	}
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)

	return data.Bytes()
}
