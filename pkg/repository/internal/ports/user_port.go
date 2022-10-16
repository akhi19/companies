package ports

import (
	"context"

	"github.com/akhi19/companies/pkg/domain"
)

type IUser interface {
	Add(
		ctx context.Context,
		userDTO domain.UserDTO,
	) error

	Delete(
		ctx context.Context,
		name string,
	) error

	Update(
		ctx context.Context,
		name string,
		updateUserDTO domain.UpdateUserDTO,
	) error

	GetUserByEmail(
		ctx context.Context,
		email string,
	) (*domain.UserDTO, error)
}
