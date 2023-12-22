package provider

import (
	"github.com/emptyhopes/employees_proxy/internal/api"
	"github.com/emptyhopes/employees_proxy/internal/client"
	"github.com/emptyhopes/employees_proxy/internal/converter"
	"github.com/emptyhopes/employees_proxy/internal/service"
)

type InterfaceEmployeeProvider interface {
	GetEmployeeApi() api.InterfaceEmployeeApi
	GetEmployeeService() service.InterfaceEmployeeService
	GetEmployeeClient() client.InterfaceClientEmployees
	GetEmployeeConverter() converter.InterfaceEmployeeConverter
}
