package invoices

import (
	"fmt"
	"github.com/Nistagram-Organization/Agent-Application/src/datasources/mysql/agent_application_db"
	"github.com/Nistagram-Organization/Agent-Application/src/utils/rest_errors"
)

func (i *Invoice) Save() rest_errors.RestErr {
	if err := agent_application_db.Client.GetClient().Create(&i).Error; err != nil {
		fmt.Sprintln(err)
		return rest_errors.NewInternalServerError("Error when trying to save item", err)
	}
	return nil
}
