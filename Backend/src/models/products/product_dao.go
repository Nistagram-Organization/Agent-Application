package products

import (
	"fmt"
	"github.com/Nistagram-Organization/Agent-Application/src/datasources/mysql/agent_application_db"
	"github.com/Nistagram-Organization/Agent-Application/src/utils/rest_errors"
)

func (p *Product) Get() rest_errors.RestErr {
	if err := agent_application_db.Client.GetClient().Take(&p, p.ID).Error; err != nil {
		fmt.Sprintln(err)
		return rest_errors.NewNotFoundError(fmt.Sprintf("Error when trying to get product with id %d", p.ID))
	}
	return nil
}

func (p *Product) GetAll() []Product {
	var products []Product

	if err := agent_application_db.Client.GetClient().Find(&products).Error; err != nil {
		return []Product{}
	}
	return products
}
