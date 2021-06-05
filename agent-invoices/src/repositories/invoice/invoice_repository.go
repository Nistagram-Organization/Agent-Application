package invoice

import (
	"github.com/Nistagram-Organization/Agent-Application/agent-shared/src/datasources"
	"github.com/Nistagram-Organization/Agent-Application/agent-shared/src/model/invoice"
	repo "github.com/Nistagram-Organization/Agent-Application/agent-shared/src/repositories/invoice"
	"github.com/Nistagram-Organization/Agent-Application/agent-shared/src/utils/rest_error"
	"gorm.io/gorm"
)

type invoicesRepository struct {
	db *gorm.DB
}

func NewInvoicesRepository(databaseClient datasources.DatabaseClient) repo.InvoiceRepository {
	return &invoicesRepository{
		databaseClient.GetClient(),
	}
}

func (i *invoicesRepository) Save(invoice *invoice.Invoice) (*invoice.Invoice, rest_error.RestErr) {
	if err := i.db.Create(&invoice).Error; err != nil {
		return nil, rest_error.NewInternalServerError("Error when trying to save item", err)
	}
	return invoice, nil
}
