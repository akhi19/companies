package company

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/akhi19/companies/pkg/common"
	"github.com/akhi19/companies/pkg/company/internal/adaptors"
	"github.com/akhi19/companies/pkg/company/internal/ports"
	"github.com/akhi19/companies/pkg/domain"
	"github.com/akhi19/companies/pkg/repository"
)

func NewHttpServer(
	router *mux.Router,
	companyContainer repository.CompanyContainer,
) {
	repositoryAdaptor := adaptors.NewRepositoryAdaptor(
		companyContainer,
	)

	clientPort := ports.NewClientPort(
		repositoryAdaptor,
	)

	router.HandleFunc(
		"/company",
		common.HttpRequestHandler(
			clientPort.Add,
			domain.RoleAdmin,
		),
	).Methods(http.MethodPost)

	router.HandleFunc(
		"/company/{name}",
		common.HttpRequestHandler(
			clientPort.Update,
			domain.RoleAdmin,
		),
	).Methods(http.MethodPatch)

	router.HandleFunc(
		"/company/{name}",
		common.HttpRequestHandler(
			clientPort.Delete,
			domain.RoleAdmin,
		),
	).Methods(http.MethodDelete)

	router.HandleFunc(
		"/company/{name}",
		common.HttpRequestHandler(
			clientPort.Get,
			domain.RoleViewer,
		),
	).Methods(http.MethodGet)
}
