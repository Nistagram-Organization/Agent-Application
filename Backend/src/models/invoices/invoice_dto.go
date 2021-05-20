package invoices

import (
	"github.com/Nistagram-Organization/Agent-Application/src/models/delivery_informations"
	"github.com/Nistagram-Organization/Agent-Application/src/models/invoice_items"
	"github.com/Nistagram-Organization/Agent-Application/src/utils/rest_errors"
)

type Invoice struct {
	ID                  uint    `json:"id"`
	Date                string  `json:"date"`
	Total               float32 `json:"total"`
	InvoiceItems        []invoice_items.InvoiceItem
	DeliveryInformation delivery_informations.DeliveryInformation
}

func (i *Invoice) Validate() rest_errors.RestErr {
	if i.InvoiceItems == nil || len(i.InvoiceItems) == 0 {
		return rest_errors.NewBadRequestError("No items are selected for buying")
	}

	if err := i.DeliveryInformation.Validate(); err != nil {
		return err
	}

	return nil
}
