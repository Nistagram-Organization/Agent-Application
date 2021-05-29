package repositories

import (
	"github.com/Nistagram-Organization/Agent-Application/src/model/invoices"
	"github.com/Nistagram-Organization/Agent-Application/src/utils/rest_errors"
	"github.com/stretchr/testify/mock"
)

type InvoiceRepositoryMock struct {
	mock.Mock
}

func (i *InvoiceRepositoryMock) Save(invoice *invoices.Invoice) (*invoices.Invoice, rest_errors.RestErr) {
	args := i.Called(invoice)
	if args.Get(1) == nil {
		return args.Get(0).(*invoices.Invoice), nil
	}
	return nil, args.Get(1).(rest_errors.RestErr)
}