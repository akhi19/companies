package token

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/akhi19/companies/pkg/common"
	"github.com/akhi19/companies/pkg/domain"
	"github.com/akhi19/companies/pkg/repository"
	"github.com/akhi19/companies/pkg/token/internal/adaptors"
	"github.com/akhi19/companies/pkg/token/internal/ports"
)

func NewHttpServer(
	router *mux.Router,
	userContainer repository.UserContainer,
) {
	repositoryAdaptor := adaptors.NewRepositoryAdaptor(
		userContainer,
	)

	clientPort := ports.NewClientPort(
		repositoryAdaptor,
	)

	router.HandleFunc(
		"/token",
		common.HttpRequestHandler(
			clientPort.Get,
			domain.RoleAny,
		),
	).Methods(http.MethodGet)
}
