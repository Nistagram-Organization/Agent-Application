package invoices

import (
	"github.com/Nistagram-Organization/Agent-Application/src/datasources/mysql/agent_application_db"
	"github.com/Nistagram-Organization/Agent-Application/src/model/invoices"
	"github.com/Nistagram-Organization/Agent-Application/src/utils/rest_errors"
	"gorm.io/gorm"
)

type InvoicesRepository interface {
	Save(*invoices.Invoice) (*invoices.Invoice, rest_errors.RestErr)
}

type invoicesRepository struct {
	db *gorm.DB
}

func NewInvoicesRepository() InvoicesRepository {
	return &invoicesRepository{
		agent_application_db.Client.GetClient(),
	}
}

func (i *invoicesRepository) Save(invoice *invoices.Invoice) (*invoices.Invoice, rest_errors.RestErr) {
	if err := i.db.Create(&invoice).Error; err != nil {
		return nil, rest_errors.NewInternalServerError("Error when trying to save item", err)
	}
	return invoice, nil
}
