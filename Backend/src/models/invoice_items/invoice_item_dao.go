package invoice_items

import (
	"fmt"
	"github.com/Nistagram-Organization/Agent-Application/src/datasources/mysql/agent_application_db"
	"github.com/Nistagram-Organization/Agent-Application/src/utils/rest_errors"
)

func (item *InvoiceItem) GetByProduct(productID uint) rest_errors.RestErr {
	if err := agent_application_db.Client.GetClient().Where("product_id = ?", productID).Take(&item).Error; err != nil {
		fmt.Sprintln(err)
		return rest_errors.NewNotFoundError(fmt.Sprintf("Error when trying to get invoice item by product"))
	}
	return nil
}
