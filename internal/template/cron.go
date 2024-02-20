package template

import (
	"bytes"
	"fmt"
	"github.com/rafailovalexey/service-generator/internal/dto"
	"github.com/rafailovalexey/service-generator/internal/util"
	"sort"
)

func GetCronSchedulerTemplate(module string, name *dto.NameDto) []byte {
	data := bytes.Buffer{}
	separator := util.GetSeparator()

	imports := []string{
		"\"log\"",
		"\"os\"",
		"\"os/signal\"",
		"\"syscall\"",
		"\"github.com/robfig/cron\"",
		fmt.Sprintf("\"%s/internal/controller\"", module),
	}

	sort.Strings(imports)

	data.WriteString(fmt.Sprintf("package cron_scheduler"))
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

	data.WriteString(fmt.Sprintf("func Run(%sController controller.%sControllerInterface) {", name.LowerCamelCaseSingular, name.CamelCaseSingular))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tc := cron.New()"))
	data.WriteString(separator)
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tc.Start()"))
	data.WriteString(separator)
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tlog.Println(\"Application cron started\")"))
	data.WriteString(separator)
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tdone := make(chan os.Signal)"))
	data.WriteString(separator)
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tsignal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)"))
	data.WriteString(separator)
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t<-done"))
	data.WriteString(separator)
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tc.Stop()"))
	data.WriteString(separator)
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tlog.Println(\"Application cron stopped\")"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(fmt.Sprintf(""))

	return data.Bytes()
}
