package internal

import (
	"encoding/json"
	"io"
	"regexp"

	"github.com/akhi19/companies/pkg/common"
	"github.com/akhi19/companies/pkg/domain"
	"github.com/google/uuid"
)

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

type AddUserRequestDTO struct {
	Name     string      `json:"name" validate:"required"`
	Password string      `json:"password" validate:"required"`
	Email    string      `json:"email" validate:"required"`
	Role     domain.Role `json:"role" validate:"required"`
}

func (entity *AddUserRequestDTO) Populate(
	body io.ReadCloser,
) error {
	decoder := json.NewDecoder(body)
	err := decoder.Decode(entity)
	if err != nil {
		return common.ValidationError(
			common.InvalidRequestMsg,
		)
	}

	err = common.Validator.Struct(entity)
	if err != nil {
		return common.ValidationError(
			common.InvalidRequestMsg,
		)
	}
	if !emailRegex.MatchString(entity.Email) {
		return common.ValidationError(
			"Please provide valid mail",
		)
	}
	hash, err := common.GeneratehashPassword(entity.Password)
	if err != nil {
		return common.ValidationError(
			"Error hashing password",
		)
	}
	entity.Password = hash
	return nil
}

func (entity *AddUserRequestDTO) ToUserDTO() domain.UserDTO {
	return domain.UserDTO{
		ID:       domain.UUID(uuid.New().String()),
		Password: entity.Password,
		Name:     entity.Name,
		Role:     entity.Role,
		Email:    entity.Email,
		Status:   domain.EntityStatusActive,
	}
}

type UpdateUserRequestDTO struct {
	Name     *string      `json:"name"`
	Password *string      `json:"password"`
	Role     *domain.Role `json:"role"`
}

func (entity *UpdateUserRequestDTO) Populate(
	body io.ReadCloser,
) error {
	decoder := json.NewDecoder(body)
	err := decoder.Decode(entity)
	if err != nil {
		return common.ValidationError(
			common.InvalidRequestMsg,
		)
	}

	err = common.Validator.Struct(entity)
	if err != nil {
		return common.ValidationError(
			common.InvalidRequestMsg,
		)
	}
	if entity.Password != nil {
		hash, err := common.GeneratehashPassword(*entity.Password)
		if err != nil {
			return common.ValidationError(
				"Error hashing password",
			)
		}
		entity.Password = &hash
	}
	return nil
}

func (entity *UpdateUserRequestDTO) ToUpdateUserDTO() domain.UpdateUserDTO {
	return domain.UpdateUserDTO{
		Name:     entity.Name,
		Password: entity.Password,
		Role:     entity.Role,
	}
}
