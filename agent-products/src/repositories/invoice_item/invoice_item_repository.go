package invoice_item

import (
	"fmt"
	"github.com/Nistagram-Organization/agent-shared/src/datasources"
	model "github.com/Nistagram-Organization/agent-shared/src/model/invoice_item"
	"github.com/Nistagram-Organization/agent-shared/src/repositories/invoice_item"
	"github.com/Nistagram-Organization/agent-shared/src/utils/rest_error"
	"gorm.io/gorm"
)

type invoiceItemsRepository struct {
	db *gorm.DB
}

func NewInvoiceItemRepository(databaseClient datasources.DatabaseClient) invoice_item.InvoiceItemRepository {
	return &invoiceItemsRepository{
		databaseClient.GetClient(),
	}
}

func (i *invoiceItemsRepository) GetByProduct(productId uint) (*model.InvoiceItem, rest_error.RestErr) {
	var item model.InvoiceItem
	if err := i.db.Where("product_id = ?", productId).Take(&item).Error; err != nil {
		return nil, rest_error.NewNotFoundError(fmt.Sprintf("Error when trying to get invoice item by product"))
	}
	return &item, nil
}
