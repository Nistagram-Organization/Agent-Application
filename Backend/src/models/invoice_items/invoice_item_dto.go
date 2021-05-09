package invoice_items

type InvoiceItem struct {
	ID        uint `json:"id"`
	Quantity  uint `json:"quantity"`
	ProductID uint
	InvoiceID uint
}
