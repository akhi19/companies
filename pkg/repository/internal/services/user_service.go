package services

import (
	"context"

	"github.com/akhi19/companies/pkg/domain"
	"github.com/akhi19/companies/pkg/repository/internal/adaptors"
	"github.com/akhi19/companies/pkg/repository/internal/models"
)

type UserService struct {
	adaptor *adaptors.UserAdaptor
}

func NewUserService(
	adaptor *adaptors.UserAdaptor,
) *UserService {
	return &UserService{
		adaptor: adaptor,
	}
}

func (service *UserService) Add(
	ctx context.Context,
	userDTO domain.UserDTO,
) error {

	userModel := models.UserModel{}
	userModel.FromUserDTO(userDTO)

	return service.adaptor.Add(
		ctx,
		userModel,
	)
}

func (service *UserService) Delete(
	ctx context.Context,
	name string,
) error {
	return service.adaptor.Delete(
		ctx,
		name,
	)
}

func (service *UserService) Update(
	ctx context.Context,
	name string,
	updateDTO domain.UpdateUserDTO,
) error {
	updateUserModel := models.UpdateUserModel{}
	updateUserModel.FromUpdateUserDTO(updateDTO)

	return service.adaptor.Update(
		ctx,
		name,
		updateUserModel,
	)
}

func (service *UserService) GetUserByEmail(
	ctx context.Context,
	email string,
) (*domain.UserDTO, error) {
	user, err := service.adaptor.GetUserByEmail(
		ctx,
		email,
	)
	if user == nil {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	userDTO := user.ToUserDTO()
	return &userDTO, nil
}
