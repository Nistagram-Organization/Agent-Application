package invoice

import (
	"github.com/Nistagram-Organization/Agent-Application/agent-shared/src/model/invoice"
	"github.com/Nistagram-Organization/Agent-Application/agent-shared/src/utils/rest_error"
)

type InvoiceRepository interface {
	Save(*invoice.Invoice) (*invoice.Invoice, rest_error.RestErr)
}
