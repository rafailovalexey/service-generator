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
		"\"os\"",
		"\"os/signal\"",
		"\"syscall\"",
		"\"github.com/robfig/cron\"",
		"\"github.com/sirupsen/logrus\"",
		fmt.Sprintf("\"%s/internal/service\"", module),
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

	data.WriteString(fmt.Sprintf("type CronScheduler struct {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tlogger\t*logrus.Logger"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\twhatsappClientService\tservice.WhatsappClientServiceInterface"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("type CronSchedulerInterface interface {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tRun() error"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("var _ CronSchedulerInterface = (*CronScheduler)(nil)"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("func NewCronScheduler("))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tlogger *logrus.Logger,"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\twhatsappClientService service.WhatsappClientServiceInterface,"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf(") *CronScheduler {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\treturn &CronScheduler{"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\tlogger: logger,"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t\twhatsappClientService: whatsappClientService,"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t}"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("func (c *CronScheduler) Run() error {"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tcr := cron.New()"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\tcr.Start()"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\tlogger.Infof(\"application cron started\")"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\texit := make(chan os.Signal)"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\tsignal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("\t<-exit"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\tcr.Stop()"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\tlogger.Infof(\"application cron stopped\")"))
	data.WriteString(separator)
	data.WriteString(separator)

	data.WriteString(fmt.Sprintf("\treturn nil"))
	data.WriteString(separator)
	data.WriteString(fmt.Sprintf("}"))
	data.WriteString(separator)

	return data.Bytes()
}
