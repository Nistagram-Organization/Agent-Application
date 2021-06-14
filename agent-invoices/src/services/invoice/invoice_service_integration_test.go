package invoice

import (
	"github.com/Nistagram-Organization/Agent-Application/agent-invoices/src/datasources/mysql"
	"github.com/Nistagram-Organization/agent-shared/src/model/delivery_information"
	"github.com/Nistagram-Organization/agent-shared/src/model/invoice"
	"github.com/Nistagram-Organization/agent-shared/src/model/invoice_item"
	"github.com/Nistagram-Organization/Agent-Application/agent-invoices/src/repositories/product"
	invoices "github.com/Nistagram-Organization/Agent-Application/agent-invoices/src/repositories/invoice"
	model "github.com/Nistagram-Organization/agent-shared/src/model/product"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"testing"
)

type InvoiceServiceIntegrationTestsSuite struct {
	suite.Suite
	service InvoicesService
	db      *gorm.DB
}

func TestIntegrationInvoicesServiceIntegrationTestsSuite(t *testing.T) {
	suite.Run(t, new(InvoiceServiceIntegrationTestsSuite))
}

func (suite *InvoiceServiceIntegrationTestsSuite) SetupSuite() {
	database := mysql.NewMySqlDatabaseClient()
	if err := database.Init(); err != nil {
		suite.Fail("Failed to initialize database")
	}

	if err := database.Migrate(
		&model.Product{},
		&invoice.Invoice{},
		&invoice_item.InvoiceItem{},
		&delivery_information.DeliveryInformation{},
	); err != nil {
		suite.Fail("Failed to initialize database")
	}

	suite.db = database.GetClient()
	suite.service = NewInvoicesService(
		product.NewProductRepository(database),
		invoices.NewInvoicesRepository(database),
	)
}
