package services

import (
	"context"

	"github.com/akhi19/companies/pkg/common"
	"github.com/akhi19/companies/pkg/user/internal"
	"github.com/akhi19/companies/pkg/user/internal/adaptors"
	"github.com/sirupsen/logrus"
)

type UserCommandService struct {
	repositoryAdaptor *adaptors.RepositoryAdaptor
}

func NewUserCommandService(
	repository *adaptors.RepositoryAdaptor,
) *UserCommandService {
	return &UserCommandService{
		repositoryAdaptor: repository,
	}
}

func (service *UserCommandService) Add(
	ctx context.Context,
	addRequestDTO internal.AddUserRequestDTO,
) error {
	log := common.GetLogger().WithFields(logrus.Fields{"function": "AddUser"})
	userDTO := addRequestDTO.ToUserDTO()
	err := service.repositoryAdaptor.UserContainer().IUser.Add(
		ctx,
		userDTO,
	)
	if err != nil {
		log.Error(err.Error())
		return common.InternalServerError()
	}
	return nil
}

func (service *UserCommandService) Delete(
	ctx context.Context,
	name string,
) error {
	log := common.GetLogger().WithFields(logrus.Fields{"function": "DeleteUser"})
	err := service.repositoryAdaptor.UserContainer().IUser.Delete(
		ctx,
		name,
	)
	if err != nil {
		log.Error(err.Error())
		return common.InternalServerError()
	}
	return nil
}

func (service *UserCommandService) Update(
	ctx context.Context,
	name string,
	updateRequest internal.UpdateUserRequestDTO,
) error {
	log := common.GetLogger().WithFields(logrus.Fields{"function": "UpdateUser"})
	updateUserDTO := updateRequest.ToUpdateUserDTO()
	err := service.repositoryAdaptor.UserContainer().IUser.Update(
		ctx,
		name,
		updateUserDTO,
	)
	if err != nil {
		log.Error(err.Error())
		return common.InternalServerError()
	}
	return nil
}
