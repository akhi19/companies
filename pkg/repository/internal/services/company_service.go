package services

import (
	"context"

	"github.com/akhi19/companies/pkg/domain"
	"github.com/akhi19/companies/pkg/repository/internal/adaptors"
	"github.com/akhi19/companies/pkg/repository/internal/models"
)

type CompanyService struct {
	adaptor *adaptors.CompanyAdaptor
}

func NewCompanyService(
	adaptor *adaptors.CompanyAdaptor,
) *CompanyService {
	return &CompanyService{
		adaptor: adaptor,
	}
}

func (service *CompanyService) Add(
	ctx context.Context,
	companyDTO domain.CompanyDTO,
) error {

	companyModel := models.CompanyModel{}
	companyModel.FromCompanyDTO(companyDTO)

	return service.adaptor.Add(
		ctx,
		companyModel,
	)
}

func (service *CompanyService) Delete(
	ctx context.Context,
	name string,
) error {
	return service.adaptor.Delete(
		ctx,
		name,
	)
}

func (service *CompanyService) Update(
	ctx context.Context,
	name string,
	updateDTO domain.UpdateCompanyDTO,
) error {
	updateCompanyModel := models.UpdateCompanyModel{}
	updateCompanyModel.FromUpdateCompanyDTO(updateDTO)

	return service.adaptor.Update(
		ctx,
		name,
		updateCompanyModel,
	)
}

func (service *CompanyService) GetCompanyByName(
	ctx context.Context,
	name string,
) (*domain.CompanyDTO, error) {
	company, err := service.adaptor.GetCompanyByName(
		ctx,
		name,
	)
	if company == nil {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	companyDTO := company.ToCompanyDTO()
	return &companyDTO, nil
}
