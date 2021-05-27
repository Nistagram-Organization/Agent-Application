package invoice_items

import "github.com/Nistagram-Organization/Agent-Application/src/utils/rest_errors"

type InvoiceItem struct {
	ID        uint `json:"id"`
	Quantity  uint `json:"quantity"`
	ProductID uint
	InvoiceID uint
}

func (item *InvoiceItem) Validate() rest_errors.RestErr {
	if item.ProductID == 0 {
		return rest_errors.NewBadRequestError("Product must be selected")
	}

	if item.Quantity == 0 {
		return rest_errors.NewBadRequestError("Quantity must be greater than zero")
	}

	return nil
}
