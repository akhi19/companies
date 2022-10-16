package ports

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/akhi19/companies/pkg/common"
	"github.com/akhi19/companies/pkg/company/internal"
	"github.com/akhi19/companies/pkg/company/internal/adaptors"
)

type ClientPort struct {
	iCompanyCommandService iCompanyCommandService
	iCompanyQueryService   iCompanyQueryService
}

func NewClientPort(
	repositoryAdaptor *adaptors.RepositoryAdaptor,
) ClientPort {
	return ClientPort{
		iCompanyCommandService: getCompanyCommandService(
			repositoryAdaptor,
		),
		iCompanyQueryService: getCompanyQueryService(
			repositoryAdaptor,
		),
	}
}

func (port *ClientPort) Add(
	responseWriter http.ResponseWriter,
	request *http.Request,
) {
	var addRequestDTO internal.AddCompanyRequestDTO
	err := addRequestDTO.Populate(
		request.Body,
	)
	if err != nil {
		common.SendHttpError(
			request.Context(),
			responseWriter,
			err,
		)
		return
	}

	err = port.iCompanyCommandService.Add(
		request.Context(),
		addRequestDTO,
	)

	if err != nil {
		common.SendHttpError(
			request.Context(),
			responseWriter,
			err,
		)
		return
	}

	common.SendEmptyHttpResponse(
		responseWriter,
		http.StatusOK,
	)
}

func (port *ClientPort) Update(
	responseWriter http.ResponseWriter,
	request *http.Request,
) {
	var updateRequestDTO internal.UpdateCompanyRequestDTO
	err := updateRequestDTO.Populate(
		request.Body,
	)
	if err != nil {
		common.SendHttpError(
			request.Context(),
			responseWriter,
			err,
		)
		return
	}

	urlParams := mux.Vars(request)
	name := urlParams["name"]

	err = port.iCompanyCommandService.Update(
		request.Context(),
		name,
		updateRequestDTO,
	)
	if err != nil {
		common.SendHttpError(
			request.Context(),
			responseWriter,
			err,
		)
		return
	}

	common.SendEmptyHttpResponse(
		responseWriter,
		http.StatusOK,
	)
}

func (port *ClientPort) Delete(
	responseWriter http.ResponseWriter,
	request *http.Request,
) {

	urlParams := mux.Vars(request)
	name := urlParams["name"]

	err := port.iCompanyCommandService.Delete(
		request.Context(),
		name,
	)
	if err != nil {
		common.SendHttpError(
			request.Context(),
			responseWriter,
			err,
		)
		return
	}

	common.SendEmptyHttpResponse(
		responseWriter,
		http.StatusOK,
	)
}

func (port *ClientPort) Get(
	responseWriter http.ResponseWriter,
	request *http.Request,
) {
	urlParams := mux.Vars(request)
	name := urlParams["name"]

	company, err := port.iCompanyQueryService.Get(
		request.Context(),
		name,
	)
	if err != nil {
		common.SendHttpError(
			request.Context(),
			responseWriter,
			err,
		)
		return
	}

	common.SendHttpResponse(
		responseWriter,
		http.StatusOK,
		company,
	)
}
