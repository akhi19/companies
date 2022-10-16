package internal

import (
	"encoding/json"
	"io"

	"github.com/akhi19/companies/pkg/common"
	"github.com/akhi19/companies/pkg/domain"
	"github.com/google/uuid"
)

type AddCompanyRequestDTO struct {
	Name        string             `json:"name" validate:"required"`
	Description string             `json:"description"`
	Employees   *int64             `json:"employees" validate:"required"`
	Registered  *bool              `json:"registered" validate:"required"`
	Type        domain.CompanyType `json:"company_type" validate:"required"`
}

func (entity *AddCompanyRequestDTO) Populate(
	body io.ReadCloser,
) error {
	decoder := json.NewDecoder(body)
	err := decoder.Decode(entity)
	if err != nil {
		return common.ValidationError(
			common.InvalidRequestMsg,
		)
	}

	err = common.Validator.Struct(entity)
	if err != nil {
		return common.ValidationError(
			common.InvalidRequestMsg,
		)
	}
	return nil
}

func (entity *AddCompanyRequestDTO) ToCompanyDTO() domain.CompanyDTO {
	return domain.CompanyDTO{
		ID:          domain.UUID(uuid.New().String()),
		Name:        entity.Name,
		Description: entity.Description,
		Employees:   *entity.Employees,
		Registered:  *entity.Registered,
		Type:        entity.Type,
		Status:      domain.EntityStatusActive,
	}
}

type UpdateCompanyRequestDTO struct {
	Description *string             `json:"description"`
	Employees   *int64              `json:"employees"`
	Registered  *bool               `json:"registered"`
	Type        *domain.CompanyType `json:"company_type"`
}

func (entity *UpdateCompanyRequestDTO) Populate(
	body io.ReadCloser,
) error {
	decoder := json.NewDecoder(body)
	err := decoder.Decode(entity)
	if err != nil {
		return common.ValidationError(
			common.InvalidRequestMsg,
		)
	}

	err = common.Validator.Struct(entity)
	if err != nil {
		return common.ValidationError(
			common.InvalidRequestMsg,
		)
	}
	return nil
}

func (entity *UpdateCompanyRequestDTO) ToUpdateCompanyDTO() domain.UpdateCompanyDTO {
	return domain.UpdateCompanyDTO{
		Description: entity.Description,
		Employees:   entity.Employees,
		Registered:  entity.Registered,
		Type:        entity.Type,
	}
}
