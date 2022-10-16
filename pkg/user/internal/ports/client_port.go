package ports

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/akhi19/companies/pkg/common"
	"github.com/akhi19/companies/pkg/user/internal"
	"github.com/akhi19/companies/pkg/user/internal/adaptors"
)

type ClientPort struct {
	iUserCommandService iUserCommandService
}

func NewClientPort(
	repositoryAdaptor *adaptors.RepositoryAdaptor,
) ClientPort {
	return ClientPort{
		iUserCommandService: getUserCommandService(
			repositoryAdaptor,
		),
	}
}

func (port *ClientPort) Add(
	responseWriter http.ResponseWriter,
	request *http.Request,
) {
	var addUserRequestDTO internal.AddUserRequestDTO
	err := addUserRequestDTO.Populate(
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

	err = port.iUserCommandService.Add(
		request.Context(),
		addUserRequestDTO,
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

	err := port.iUserCommandService.Delete(
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

func (port *ClientPort) Update(
	responseWriter http.ResponseWriter,
	request *http.Request,
) {
	var updateRequestDTO internal.UpdateUserRequestDTO
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

	err = port.iUserCommandService.Update(
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
