package invoices

import (
	"github.com/Nistagram-Organization/Agent-Application/src/models/delivery_informations"
	"github.com/Nistagram-Organization/Agent-Application/src/models/invoice_items"
)

type Invoice struct {
	ID                  uint    `json:"id"`
	Date                string  `json:"date"`
	Total               float32 `json:"total"`
	InvoiceItems        []invoice_items.InvoiceItem
	DeliveryInformation delivery_informations.DeliveryInformation
}
