package credentials

import (
	"fmt"
	"github.com/Nistagram-Organization/Agent-Application/src/datasources/mysql/agent_application_db"
	"github.com/Nistagram-Organization/Agent-Application/src/utils/rest_errors"
)

func (c *Credentials) GetByUsername() rest_errors.RestErr {
	if err := agent_application_db.Client.GetClient().Where("username = ?", c.Username).Take(&c).Error; err != nil {
		fmt.Sprintln(err)
		return rest_errors.NewNotFoundError(fmt.Sprintf("Error when trying to get credentials with username %s", c.Username))
	}
	return nil
}
