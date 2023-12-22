package employees

import (
	"github.com/emptyhopes/employees_proxy/internal/api"
	employeeApi "github.com/emptyhopes/employees_proxy/internal/api/employees"
	"github.com/emptyhopes/employees_proxy/internal/client"
	employeeClient "github.com/emptyhopes/employees_proxy/internal/client/employees"
	"github.com/emptyhopes/employees_proxy/internal/converter"
	employeeConverter "github.com/emptyhopes/employees_proxy/internal/converter/employees"
	definition "github.com/emptyhopes/employees_proxy/internal/provider"
	"github.com/emptyhopes/employees_proxy/internal/service"
	employeeService "github.com/emptyhopes/employees_proxy/internal/service/employees"
	"github.com/emptyhopes/employees_proxy/internal/validation"
	employeeValidation "github.com/emptyhopes/employees_proxy/internal/validation/employees"
)

type provider struct {
	employeeApi        api.InterfaceEmployeeApi
	employeeService    service.InterfaceEmployeeService
	employeeClient     client.InterfaceClientEmployees
	employeeConverter  converter.InterfaceEmployeeConverter
	employeeValidation validation.InterfaceEmployeeValidation
}

var _ definition.InterfaceEmployeeProvider = (*provider)(nil)

func NewEmployeeProvider() *provider {
	return &provider{}
}

func (p *provider) GetEmployeeApi() api.InterfaceEmployeeApi {
	if p.employeeApi == nil {
		p.employeeApi = employeeApi.NewEmployeeApi(
			p.GetEmployeeValidation(),
			p.GetEmployeeConverter(),
			p.GetEmployeeService(),
		)
	}

	return p.employeeApi
}

func (p *provider) GetEmployeeService() service.InterfaceEmployeeService {
	if p.employeeService == nil {
		p.employeeService = employeeService.NewEmployeeService(
			p.GetEmployeeClient(),
		)
	}

	return p.employeeService
}

func (p *provider) GetEmployeeClient() client.InterfaceClientEmployees {
	if p.employeeClient == nil {
		p.employeeClient = employeeClient.NewEmployeeClient(
			p.GetEmployeeConverter(),
		)
	}

	return p.employeeClient
}

func (p *provider) GetEmployeeConverter() converter.InterfaceEmployeeConverter {
	if p.employeeConverter == nil {
		p.employeeConverter = employeeConverter.NewEmployeeConverter()
	}

	return p.employeeConverter
}

func (p *provider) GetEmployeeValidation() validation.InterfaceEmployeeValidation {
	if p.employeeValidation == nil {
		p.employeeValidation = employeeValidation.NewEmployeeValidation()
	}

	return p.employeeValidation
}
