package services

import (
	"context"

	"github.com/akhi19/companies/pkg/common"
	"github.com/akhi19/companies/pkg/company/internal/adaptors"
	"github.com/akhi19/companies/pkg/domain"
	"github.com/sirupsen/logrus"
)

type CompanyQueryService struct {
	repositoryAdaptor *adaptors.RepositoryAdaptor
}

func NewCompanyQueryService(
	repository *adaptors.RepositoryAdaptor,
) *CompanyQueryService {
	return &CompanyQueryService{
		repositoryAdaptor: repository,
	}
}

func (service *CompanyQueryService) Get(
	ctx context.Context,
	name string,
) (*domain.CompanyDTO, error) {
	log := common.GetLogger().WithFields(logrus.Fields{"function": "GetCompany"})
	company, err := service.repositoryAdaptor.CompanyContainer().ICompany.GetCompanyByName(
		ctx,
		name,
	)
	if err != nil {
		log.Error(err.Error())
		return nil, common.InternalServerError()
	}
	return company, nil
}
