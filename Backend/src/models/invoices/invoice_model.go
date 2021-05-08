package invoices

import (
	"github.com/Nistagram-Organization/Agent-Application/src/models/delivery_informations"
	"github.com/Nistagram-Organization/Agent-Application/src/models/invoice_items"
)

type Invoice struct {
	ID                  uint
	Date                string
	Total               float32
	InvoiceItems        []invoice_items.InvoiceItem
	DeliveryInformation delivery_informations.DeliveryInformation
}
