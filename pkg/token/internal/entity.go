package internal

import (
	"encoding/json"
	"io"
	"regexp"

	"github.com/akhi19/companies/pkg/common"
)

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

type GetTokenRequestDTO struct {
	Password string `json:"password" validate:"required"`
	Email    string `json:"email" validate:"required"`
}

func (entity *GetTokenRequestDTO) Populate(
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
	return nil
}

type TokenDTO struct {
	Token string `json:"token"`
}
