package repositories

import (
	model "github.com/Nistagram-Organization/Agent-Application/agent-shared/src/model/invoice"
	"github.com/Nistagram-Organization/Agent-Application/agent-shared/src/utils/rest_error"
	"github.com/stretchr/testify/mock"
)

type InvoiceRepositoryMock struct {
	mock.Mock
}

func (i *InvoiceRepositoryMock) Save(invoice *model.Invoice) (*model.Invoice, rest_error.RestErr) {
	args := i.Called(invoice)
	if args.Get(1) == nil {
		return args.Get(0).(*model.Invoice), nil
	}
	return nil, args.Get(1).(rest_error.RestErr)
}
