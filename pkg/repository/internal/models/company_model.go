package models

import (
	"github.com/akhi19/companies/pkg/domain"
)

type CompanyModel struct {
	ID          string
	Name        string
	Description string
	Employees   int64
	Registered  bool
	Type        domain.CompanyType
	Status      domain.EntityStatus
}

func (entity *CompanyModel) FromCompanyDTO(companyDTO domain.CompanyDTO) {
	entity.ID = string(companyDTO.ID)
	entity.Name = companyDTO.Name
	entity.Description = companyDTO.Description
	entity.Employees = companyDTO.Employees
	entity.Registered = companyDTO.Registered
	entity.Type = companyDTO.Type
	entity.Status = companyDTO.Status
}

func (entity *CompanyModel) ToCompanyDTO() domain.CompanyDTO {
	return domain.CompanyDTO{
		ID:          domain.UUID(entity.ID),
		Name:        entity.Name,
		Description: entity.Description,
		Employees:   entity.Employees,
		Registered:  entity.Registered,
		Type:        entity.Type,
		Status:      entity.Status,
	}
}

type UpdateCompanyModel struct {
	Description *string
	Employees   *int64
	Registered  *bool
	Type        *domain.CompanyType
}

func (entity *UpdateCompanyModel) FromUpdateCompanyDTO(companyDTO domain.UpdateCompanyDTO) {
	entity.Description = companyDTO.Description
	entity.Employees = companyDTO.Employees
	entity.Registered = companyDTO.Registered
	entity.Type = companyDTO.Type
}
