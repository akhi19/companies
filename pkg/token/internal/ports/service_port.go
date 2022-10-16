package ports

import (
	"context"

	"github.com/akhi19/companies/pkg/token/internal"
	"github.com/akhi19/companies/pkg/token/internal/adaptors"
	"github.com/akhi19/companies/pkg/token/internal/services"
)

var tokenQueryService *services.TokenQueryService

type iTokenQueryService interface {
	Get(
		ctx context.Context,
		getTokenRequestDTO internal.GetTokenRequestDTO,
	) (*internal.TokenDTO, error)
}

func getTokenQueryService(
	repositoryAdaptor *adaptors.RepositoryAdaptor,
) *services.TokenQueryService {
	if tokenQueryService == nil {
		tokenQueryService = services.NewTokenQueryService(
			repositoryAdaptor,
		)
	}
	return tokenQueryService
}
