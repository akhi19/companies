package ports

import (
	"context"

	"github.com/akhi19/companies/pkg/domain"
)

type ICompany interface {
	Add(
		ctx context.Context,
		companyDTO domain.CompanyDTO,
	) error

	Delete(
		ctx context.Context,
		name string,
	) error

	Update(
		ctx context.Context,
		name string,
		updateCompanyDTO domain.UpdateCompanyDTO,
	) error

	GetCompanyByName(
		ctx context.Context,
		name string,
	) (*domain.CompanyDTO, error)
}
