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
	args := i.Called(productId)
	if args.Get(1) == nil {
		return args.Get(0).(*invoice_items.InvoiceItem), nil
	}
	return nil, args.Get(1).(rest_errors.RestErr)
}
