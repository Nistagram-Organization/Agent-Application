package invoice_item

import "github.com/Nistagram-Organization/Agent-Application/agent-shared/src/utils/rest_error"

type InvoiceItem struct {
	ID        uint `json:"id"`
	Quantity  uint `json:"quantity"`
	ProductID uint
	InvoiceID uint
}

func (item *InvoiceItem) Validate() rest_error.RestErr {
	if item.ProductID == 0 {
		return rest_error.NewBadRequestError("Product must be selected")
	}

	if item.Quantity == 0 {
		return rest_error.NewBadRequestError("Quantity must be greater than zero")
	}

	return nil
}
