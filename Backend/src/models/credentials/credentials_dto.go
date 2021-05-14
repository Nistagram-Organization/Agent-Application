package credentials

import (
	"github.com/Nistagram-Organization/Agent-Application/src/utils/bcrypt_utils"
	"github.com/Nistagram-Organization/Agent-Application/src/utils/rest_errors"
)

type Credentials struct {
	ID       uint
	Username string `json:"username"`
	Password string `json:"password"`
}

func (c *Credentials) Validate() rest_errors.RestErr {
	plainPassword := c.Password
	if err := c.GetByUsername(); err != nil {
		return err
	}
	if !bcrypt_utils.CompareHashAndValue(c.Password, plainPassword) {
		return rest_errors.NewUnauthorizedError("invalid password")
	}
	return nil
}
