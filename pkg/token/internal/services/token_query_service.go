package services

import (
	"context"

	"github.com/akhi19/companies/pkg/common"
	"github.com/akhi19/companies/pkg/domain"
	"github.com/akhi19/companies/pkg/token/internal"
	"github.com/akhi19/companies/pkg/token/internal/adaptors"
	"github.com/sirupsen/logrus"
)

type TokenQueryService struct {
	repositoryAdaptor *adaptors.RepositoryAdaptor
}

func NewTokenQueryService(
	repository *adaptors.RepositoryAdaptor,
) *TokenQueryService {
	return &TokenQueryService{
		repositoryAdaptor: repository,
	}
}

func (service *TokenQueryService) Get(
	ctx context.Context,
	getTokenDTO internal.GetTokenRequestDTO,
) (*internal.TokenDTO, error) {
	log := common.GetLogger().WithFields(logrus.Fields{"function": "GetToken"})
	user, err := service.repositoryAdaptor.UserContainer().IUser.GetUserByEmail(
		ctx,
		getTokenDTO.Email,
	)
	if err != nil {
		log.Error(err.Error())
		return nil, common.InternalServerError()
	}
	if user == nil || user.Status == domain.EntityStatusInactive {
		log.Error("User does not exist")
		return nil, common.BadRequest(common.BadRequestCode, "User does not exist")
	}
	if !common.CheckPasswordHash(getTokenDTO.Password, user.Password) {
		log.Errorf("wrong password : got %v, expected : %v", getTokenDTO.Password, user.Password)
		return nil, common.BadRequest(common.BadRequestCode, "Incorrect password")
	}

	validToken, err := common.GenerateJWT(getTokenDTO.Email, string(user.Role))
	if err != nil {
		return nil, err
	}

	return &internal.TokenDTO{
		Token: validToken,
	}, nil
}
