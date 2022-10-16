package models

import (
	"github.com/akhi19/companies/pkg/domain"
)

type UserModel struct {
	ID       string
	Name     string
	Password string
	Role     domain.Role
	Email    string
	Status   domain.EntityStatus
}

func (entity *UserModel) FromUserDTO(userDTO domain.UserDTO) {
	entity.ID = string(userDTO.ID)
	entity.Name = userDTO.Name
	entity.Password = userDTO.Password
	entity.Role = userDTO.Role
	entity.Email = userDTO.Email
	entity.Status = userDTO.Status
}

func (entity *UserModel) ToUserDTO() domain.UserDTO {
	return domain.UserDTO{
		ID:       domain.UUID(entity.ID),
		Name:     entity.Name,
		Password: entity.Password,
		Role:     entity.Role,
		Email:    entity.Email,
		Status:   entity.Status,
	}
}

type UpdateUserModel struct {
	Name     *string
	Password *string
	Role     *domain.Role
}

func (entity *UpdateUserModel) FromUpdateUserDTO(userDTO domain.UpdateUserDTO) {
	entity.Name = userDTO.Name
	entity.Password = userDTO.Password
	entity.Role = userDTO.Role
}
