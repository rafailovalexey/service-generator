package template

import (
	"bytes"
	"fmt"
	"github.com/rafailovalexey/service-generator/internal/util"
	"sort"
)

func GetMigrationTemplate(database string, module string) []byte {
	data := bytes.Buffer{}
	separator := util.GetSeparator()

	imports := []string{
		"\"fmt\"",
		"\"log\"",
		"\"os\"",
		fmt.Sprintf("\"%s/database\"", module),
		"\"github.com/pressly/goose\"",
	}

	sort.Strings(imports)

	data.WriteString(fmt.Sprintf("package migration"))
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

	data.WriteString(fmt.Sprintf("func Run(db database.DatabaseInterface) error {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tdirectory := os.Getenv(\"MIGRATION_DIRECTORY\")"))
	data.WriteString(separator)
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tif directory == \"\" {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\treturn fmt.Errorf(\"specify the directory for migration\\n\")"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t}"))
	data.WriteString(separator)
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tinstance, err := db.Connect()"))
	data.WriteString(separator)
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tif err != nil {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprint("\t\treturn fmt.Errorf(\"failed to connect %s\\n\", err.Error())"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t}"))
	data.WriteString(separator)

	switch database {
	case "mysql":
		data.WriteString(fmt.Sprintf("\terr = goose.SetDialect(\"mysql\")"))
		data.WriteString(separator)
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\tif err != nil {"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t\treturn fmt.Errorf(\"failed to set dialect %s\\n\", err.Error())"))
		data.WriteString(separator)
		data.WriteString(fmt.Sprintf("\t}"))
		data.WriteString(separator)
	}

	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\terr = goose.Up(instance, directory)"))
	data.WriteString(separator)
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tif err != nil {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprint("\t\treturn fmt.Errorf(\"migration failed %v\\n\", err)"))
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
