package repositories

import (
	"github.com/Nistagram-Organization/Agent-Application/agent-shared/src/model/invoice_item"
	"github.com/Nistagram-Organization/Agent-Application/agent-shared/src/utils/rest_error"
	"github.com/stretchr/testify/mock"
)

type InvoiceItemsRepositoryMock struct {
	mock.Mock
}

func (i *InvoiceItemsRepositoryMock) GetByProduct(productId uint) (*invoice_item.InvoiceItem, rest_error.RestErr) {
	args := i.Called(productId)
	if args.Get(1) == nil {
		return args.Get(0).(*invoice_item.InvoiceItem), nil
	}
	return nil, args.Get(1).(rest_error.RestErr)
}
