package ports

import (
	"context"

	"github.com/akhi19/companies/pkg/user/internal"
	"github.com/akhi19/companies/pkg/user/internal/adaptors"
	"github.com/akhi19/companies/pkg/user/internal/services"
)

var userCommandService *services.UserCommandService

type iUserCommandService interface {
	Add(
		ctx context.Context,
		addUserRequest internal.AddUserRequestDTO,
	) error

	Update(
		ctx context.Context,
		name string,
		updateUserRequest internal.UpdateUserRequestDTO,
	) error

	Delete(
		ctx context.Context,
		name string,
	) error
}

func getUserCommandService(
	repositoryAdaptor *adaptors.RepositoryAdaptor,
) *services.UserCommandService {
	if userCommandService == nil {
		userCommandService = services.NewUserCommandService(
			repositoryAdaptor,
		)
	}
	return userCommandService
}
