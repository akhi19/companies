package user

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/akhi19/companies/pkg/common"
	"github.com/akhi19/companies/pkg/domain"
	"github.com/akhi19/companies/pkg/repository"
	"github.com/akhi19/companies/pkg/user/internal/adaptors"
	"github.com/akhi19/companies/pkg/user/internal/ports"
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
		"/user",
		common.HttpRequestHandler(
			clientPort.Add,
			domain.RoleAdmin,
		),
	).Methods(http.MethodPost)

	router.HandleFunc(
		"/user/{name}",
		common.HttpRequestHandler(
			clientPort.Update,
			domain.RoleAdmin,
		),
	).Methods(http.MethodPatch)

	router.HandleFunc(
		"/user/{name}",
		common.HttpRequestHandler(
			clientPort.Delete,
			domain.RoleAdmin,
		),
	).Methods(http.MethodDelete)
}
