package product

import (
	"github.com/Nistagram-Organization/Agent-Application/agent-products/src/datasources/mysql"
	invoice_item2 "github.com/Nistagram-Organization/Agent-Application/agent-products/src/repositories/invoice_item"
	product3 "github.com/Nistagram-Organization/Agent-Application/agent-products/src/repositories/product"
	"github.com/Nistagram-Organization/Agent-Application/agent-products/src/utils/image_utils"
	"github.com/Nistagram-Organization/agent-shared/src/model/delivery_information"
	"github.com/Nistagram-Organization/agent-shared/src/model/invoice"
	"github.com/Nistagram-Organization/agent-shared/src/model/invoice_item"
	model "github.com/Nistagram-Organization/agent-shared/src/model/product"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ProductServiceIntegrationTestsSuite struct {
	suite.Suite
	service ProductService
}

func (suite *ProductServiceIntegrationTestsSuite) SetupSuite() {
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

	suite.service = NewProductService(
		product3.NewProductRepository(database),
		invoice_item2.NewInvoiceItemRepository(database),
		image_utils.NewImageUtilsService(),
		"temp",
	)
}

func (suite *ProductServiceIntegrationTestsSuite) SetupTest() {
	// insert data in database
}

func (suite *ProductServiceIntegrationTestsSuite) TearDownTest() {
	// clear data from database
}

func TestProductServiceIntegrationTestsSuite(t *testing.T) {
	suite.Run(t, new(ProductServiceIntegrationTestsSuite))
}
