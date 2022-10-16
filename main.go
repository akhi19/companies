package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/akhi19/companies/configs"
	"github.com/akhi19/companies/pkg/common"
	"github.com/akhi19/companies/pkg/company"
	"github.com/akhi19/companies/pkg/repository"
	"github.com/akhi19/companies/pkg/token"
	"github.com/akhi19/companies/pkg/user"
)

func main() {
	common.InitializeConnection(configs.GetConfig())

	companyContaier := repository.CompanyContainer{}
	companyContaier.Build()

	userContainer := repository.UserContainer{}
	userContainer.Build()

	router := mux.NewRouter().StrictSlash(false)

	companiesRoute := router.PathPrefix(
		"/companies/v1/",
	).Subrouter()

	company.NewHttpServer(
		companiesRoute,
		companyContaier,
	)

	user.NewHttpServer(
		companiesRoute,
		userContainer,
	)

	token.NewHttpServer(
		companiesRoute,
		userContainer,
	)

	common.GetLogger().Info("Starting server with ports : " + configs.GetConfig().Port)

	http.ListenAndServe(":"+configs.GetConfig().Port, router)
}
