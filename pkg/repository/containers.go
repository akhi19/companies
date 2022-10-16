package repository

import (
	"github.com/akhi19/companies/pkg/repository/internal/adaptors"
	"github.com/akhi19/companies/pkg/repository/internal/ports"
	"github.com/akhi19/companies/pkg/repository/internal/services"
)

type UserContainer struct {
	IUser ports.IUser
}

type CompanyContainer struct {
	ICompany ports.ICompany
}

func (container *UserContainer) Build() {
	userAdaptor := adaptors.NewUserAdaptor()
	container.IUser = services.NewUserService(
		userAdaptor,
	)
}

func (container *CompanyContainer) Build() {
	companyAdaptor := adaptors.NewCompanyAdaptor()
	container.ICompany = services.NewCompanyService(
		companyAdaptor,
	)
}
