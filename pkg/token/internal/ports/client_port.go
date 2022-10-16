package ports

import (
	"net/http"

	"github.com/akhi19/companies/pkg/common"
	"github.com/akhi19/companies/pkg/token/internal"
	"github.com/akhi19/companies/pkg/token/internal/adaptors"
)

type ClientPort struct {
	iTokenQueryService iTokenQueryService
}

func NewClientPort(
	repositoryAdaptor *adaptors.RepositoryAdaptor,
) ClientPort {
	return ClientPort{
		iTokenQueryService: getTokenQueryService(
			repositoryAdaptor,
		),
	}
}

func (port *ClientPort) Get(
	responseWriter http.ResponseWriter,
	request *http.Request,
) {
	var getTokenRequestDTO internal.GetTokenRequestDTO
	err := getTokenRequestDTO.Populate(
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

	response, err := port.iTokenQueryService.Get(
		request.Context(),
		getTokenRequestDTO,
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
		response,
	)
}
