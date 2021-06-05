package invoice

import (
	"github.com/Nistagram-Organization/Agent-Application/agent-shared/src/model/delivery_information"
	"github.com/Nistagram-Organization/Agent-Application/agent-shared/src/model/invoice_item"
	"github.com/Nistagram-Organization/Agent-Application/agent-shared/src/utils/rest_error"
)

type Invoice struct {
	ID                  uint    `json:"id"`
	Date                int64   `json:"date"`
	Total               float32 `json:"total"`
	InvoiceItems        []invoice_item.InvoiceItem
	DeliveryInformation delivery_information.DeliveryInformation
}

func (i *Invoice) Validate() rest_error.RestErr {
	if i.InvoiceItems == nil || len(i.InvoiceItems) == 0 {
		return rest_error.NewBadRequestError("No items are selected for buying")
	}

	if err := i.DeliveryInformation.Validate(); err != nil {
		return err
	}

	return nil
}
