package repositories

import (
	"github.com/Nistagram-Organization/Agent-Application/src/model/invoice_items"
	"github.com/Nistagram-Organization/Agent-Application/src/utils/rest_errors"
	"github.com/stretchr/testify/mock"
)

type InvoiceItemsRepositoryMock struct {
	mock.Mock
}

func (i *InvoiceItemsRepositoryMock) GetByProduct(productId uint) (*invoice_items.InvoiceItem, rest_errors.RestErr) {
	return nil, nil
}
