package adaptors

import (
	"github.com/akhi19/companies/pkg/repository"
)

type RepositoryAdaptor struct {
	companyContainer repository.CompanyContainer
}

func NewRepositoryAdaptor(
	companyContainer repository.CompanyContainer,
) *RepositoryAdaptor {
	return &RepositoryAdaptor{
		companyContainer: companyContainer,
	}
}

func (adaptor *RepositoryAdaptor) CompanyContainer() repository.CompanyContainer {
	return adaptor.companyContainer
}
