package dto

import "fmt"

type ApplicationDto struct {
	Version      string
	Type         string
	Database     string
	Organization string
	Project      string
	Directory    string
	Module       string
	Names        *NameDto
}

func NewApplicationDto(
	version string,
	t string,
	database string,
	organization string,
	project string,
	directory string,
	names *NameDto,
) *ApplicationDto {
	dto := &ApplicationDto{
		Version:      version,
		Type:         t,
		Database:     database,
		Organization: organization,
		Project:      project,
		Directory:    directory,
		Names:        names,
	}

	dto.Module = fmt.Sprintf("%s/%s/%s", organization, project, directory)

	return dto
}
