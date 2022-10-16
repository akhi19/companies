package services

import (
	"context"

	"github.com/akhi19/companies/pkg/common"
	"github.com/akhi19/companies/pkg/company/internal"
	"github.com/akhi19/companies/pkg/company/internal/adaptors"
	"github.com/sirupsen/logrus"
)

type CompanyCommandService struct {
	repositoryAdaptor *adaptors.RepositoryAdaptor
}

func NewComapnyCommandService(
	repository *adaptors.RepositoryAdaptor,
) *CompanyCommandService {
	return &CompanyCommandService{
		repositoryAdaptor: repository,
	}
}

func (service *CompanyCommandService) Add(
	ctx context.Context,
	addCompanyRequest internal.AddCompanyRequestDTO,
) error {
	log := common.GetLogger().WithFields(logrus.Fields{"function": "AddCompany"})
	companyDTO := addCompanyRequest.ToCompanyDTO()
	err := service.repositoryAdaptor.CompanyContainer().ICompany.Add(
		ctx,
		companyDTO,
	)
	if err != nil {
		log.Error(err.Error())
		return common.InternalServerError()
	}
	return nil
}

func (service *CompanyCommandService) Update(
	ctx context.Context,
	name string,
	updateCompanyRequest internal.UpdateCompanyRequestDTO,
) error {
	log := common.GetLogger().WithFields(logrus.Fields{"function": "UpdateComapny"})
	updateCompanyDTO := updateCompanyRequest.ToUpdateCompanyDTO()
	err := service.repositoryAdaptor.CompanyContainer().ICompany.Update(
		ctx,
		name,
		updateCompanyDTO,
	)
	if err != nil {
		log.Error(err.Error())
		return common.InternalServerError()
	}
	return nil
}

func (service *CompanyCommandService) Delete(
	ctx context.Context,
	name string,
) error {
	log := common.GetLogger().WithFields(logrus.Fields{"function": "DeleteCompany"})
	err := service.repositoryAdaptor.CompanyContainer().ICompany.Delete(
		ctx,
		name,
	)
	if err != nil {
		log.Error(err.Error())
		return common.InternalServerError()
	}

	return nil
}
