package ports

import (
	"context"

	"github.com/akhi19/companies/pkg/company/internal"
	"github.com/akhi19/companies/pkg/company/internal/adaptors"
	"github.com/akhi19/companies/pkg/company/internal/services"
	"github.com/akhi19/companies/pkg/domain"
)

var companyCommandService *services.CompanyCommandService
var companyQueryService *services.CompanyQueryService

type iCompanyCommandService interface {
	Add(
		ctx context.Context,
		addCompanyRequest internal.AddCompanyRequestDTO,
	) error

	Update(
		ctx context.Context,
		name string,
		updateCompanyRequest internal.UpdateCompanyRequestDTO,
	) error

	Delete(
		ctx context.Context,
		name string,
	) error
}

type iCompanyQueryService interface {
	Get(
		ctx context.Context,
		name string,
	) (*domain.CompanyDTO, error)
}

func getCompanyCommandService(
	repositoryAdaptor *adaptors.RepositoryAdaptor,
) *services.CompanyCommandService {
	if companyCommandService == nil {
		companyCommandService = services.NewComapnyCommandService(
			repositoryAdaptor,
		)
	}
	return companyCommandService
}

func getCompanyQueryService(
	repositoryAdaptor *adaptors.RepositoryAdaptor,
) *services.CompanyQueryService {
	if companyQueryService == nil {
		companyQueryService = services.NewCompanyQueryService(
			repositoryAdaptor,
		)
	}
	return companyQueryService
}
