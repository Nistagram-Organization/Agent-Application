package invoice_items

import (
	"fmt"
	"github.com/Nistagram-Organization/Agent-Application/src/datasources/mysql/agent_application_db"
	"github.com/Nistagram-Organization/Agent-Application/src/model/invoice_items"
	"github.com/Nistagram-Organization/Agent-Application/src/utils/rest_errors"
	"gorm.io/gorm"
)

type InvoiceItemsRepository interface {
	GetByProduct(uint) (*invoice_items.InvoiceItem, rest_errors.RestErr)
}

type invoiceItemsRepository struct {
	db *gorm.DB
}

func NewInvoiceItemsRepository() InvoiceItemsRepository {
	return &invoiceItemsRepository{
		agent_application_db.Client.GetClient(),
	}
}

func (i *invoiceItemsRepository) GetByProduct(productId uint) (*invoice_items.InvoiceItem, rest_errors.RestErr) {
	var item invoice_items.InvoiceItem
	if err := i.db.Where("product_id = ?", productId).Take(&item).Error; err != nil {
		return nil, rest_errors.NewNotFoundError(fmt.Sprintf("Error when trying to get invoice item by product"))
	}
	return &item, nil
}
