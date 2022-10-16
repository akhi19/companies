package adaptors

import (
	"github.com/akhi19/companies/pkg/repository"
)

type RepositoryAdaptor struct {
	userContainer repository.UserContainer
}

func NewRepositoryAdaptor(
	userContainer repository.UserContainer,
) *RepositoryAdaptor {
	return &RepositoryAdaptor{
		userContainer: userContainer,
	}
}

func (adaptor *RepositoryAdaptor) UserContainer() repository.UserContainer {
	return adaptor.userContainer
}
