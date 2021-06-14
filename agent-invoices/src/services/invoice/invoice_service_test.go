package invoice

import (
	"errors"
	"fmt"
	"github.com/Nistagram-Organization/agent-shared/src/mocks/repositories"
	"github.com/Nistagram-Organization/agent-shared/src/model/delivery_information"
	"github.com/Nistagram-Organization/agent-shared/src/model/invoice"
	"github.com/Nistagram-Organization/agent-shared/src/model/invoice_item"
	"github.com/Nistagram-Organization/agent-shared/src/model/product"
	"github.com/Nistagram-Organization/agent-shared/src/utils/rest_error"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type InvoiceServiceUnitTestsSuite struct {
	suite.Suite
	productRepositoryMock *repositories.ProductRepositoryMock
	invoiceRepositoryMock *repositories.InvoiceRepositoryMock
	service               InvoicesService
}

func TestProductServiceUnitTestsSuite(t *testing.T) {
	suite.Run(t, new(InvoiceServiceUnitTestsSuite))
}

func (suite *InvoiceServiceUnitTestsSuite) SetupSuite() {
	suite.productRepositoryMock = new(repositories.ProductRepositoryMock)
	suite.invoiceRepositoryMock = new(repositories.InvoiceRepositoryMock)
	suite.service = NewInvoicesService(suite.productRepositoryMock, suite.invoiceRepositoryMock)
}

func (suite *InvoiceServiceUnitTestsSuite) TestNewInvoicesService() {
	assert.NotNil(suite.T(), suite.service, "Service is nil")
}

func (suite *InvoiceServiceUnitTestsSuite) TestInvoicesService_BuyProduct_NoInvoiceItems() {
	invoice := invoice.Invoice{}
	err := rest_error.NewBadRequestError("No items are selected for buying")

	buyErr := suite.service.BuyProduct(&invoice)

	assert.Equal(suite.T(), err, buyErr)
}

func (suite *InvoiceServiceUnitTestsSuite) TestInvoicesService_BuyProduct_DeliveryInformationInvalid() {
	invoiceItem := invoice_item.InvoiceItem{}
	invoice := invoice.Invoice{
		DeliveryInformation: delivery_information.DeliveryInformation{},
	}
	invoice.InvoiceItems = append(invoice.InvoiceItems, invoiceItem)

	err := rest_error.NewBadRequestError("Customer name cannot be empty")

	buyErr := suite.service.BuyProduct(&invoice)

	assert.Equal(suite.T(), err, buyErr)
}

func (suite *InvoiceServiceUnitTestsSuite) TestInvoicesService_BuyProduct_InvoiceItemInvalid() {
	invoiceItem := invoice_item.InvoiceItem{}
	invoice := invoice.Invoice{
		DeliveryInformation: delivery_information.DeliveryInformation{
			Name:    "Mujo",
			Surname: "Alen",
			Phone:   "+381698886969",
			Address: "Partizanska 69",
			City:    "Novi Sad",
			ZipCode: 21000,
		},
	}
	invoice.InvoiceItems = append(invoice.InvoiceItems, invoiceItem)

	err := rest_error.NewBadRequestError("Product must be selected")

	buyErr := suite.service.BuyProduct(&invoice)

	assert.Equal(suite.T(), err, buyErr)
}

func (suite *InvoiceServiceUnitTestsSuite) TestInvoicesService_BuyProduct_ProductNotFound() {
	invoiceItem := invoice_item.InvoiceItem{
		ProductID: 1,
		Quantity:  1,
	}
	invoice := invoice.Invoice{
		DeliveryInformation: delivery_information.DeliveryInformation{
			Name:    "Mujo",
			Surname: "Alen",
			Phone:   "+381698886969",
			Address: "Partizanska 69",
			City:    "Novi Sad",
			ZipCode: 21000,
		},
	}
	invoice.InvoiceItems = append(invoice.InvoiceItems, invoiceItem)
	err := rest_error.NewNotFoundError(fmt.Sprintf("Error when trying to get product with id %d", invoiceItem.ProductID))

	suite.productRepositoryMock.On("Get", invoiceItem.ProductID).Return(nil, err).Once()

	buyErr := suite.service.BuyProduct(&invoice)

	assert.Equal(suite.T(), err, buyErr)
}

func (suite *InvoiceServiceUnitTestsSuite) TestInvoicesService_BuyProduct_NotEnoughProducts() {
	invoiceItem := invoice_item.InvoiceItem{
		ProductID: 1,
		Quantity:  2,
	}
	invoice := invoice.Invoice{
		DeliveryInformation: delivery_information.DeliveryInformation{
			Name:    "Mujo",
			Surname: "Alen",
			Phone:   "+381698886969",
			Address: "Partizanska 69",
			City:    "Novi Sad",
			ZipCode: 21000,
		},
	}
	product := product.Product{
		OnStock: 1,
	}
	invoice.InvoiceItems = append(invoice.InvoiceItems, invoiceItem)
	err := rest_error.NewBadRequestError("Not enough products in stock")

	suite.productRepositoryMock.On("Get", invoiceItem.ProductID).Return(&product, nil).Once()

	buyErr := suite.service.BuyProduct(&invoice)

	assert.Equal(suite.T(), err, buyErr)
}

func (suite *InvoiceServiceUnitTestsSuite) TestInvoicesService_BuyProduct_ProductsRepositoryError() {
	invoiceItem := invoice_item.InvoiceItem{
		ProductID: 1,
		Quantity:  2,
	}
	invoice := invoice.Invoice{
		DeliveryInformation: delivery_information.DeliveryInformation{
			Name:    "Mujo",
			Surname: "Alen",
			Phone:   "+381698886969",
			Address: "Partizanska 69",
			City:    "Novi Sad",
			ZipCode: 21000,
		},
	}
	product := product.Product{
		OnStock: 5,
	}
	invoice.InvoiceItems = append(invoice.InvoiceItems, invoiceItem)
	err := rest_error.NewInternalServerError("Error when trying to delete product", errors.New("test"))

	suite.productRepositoryMock.On("Get", invoiceItem.ProductID).Return(&product, nil).Once()
	suite.productRepositoryMock.On("Update", &product).Return(nil, err).Once()
	buyErr := suite.service.BuyProduct(&invoice)

	assert.Equal(suite.T(), err, buyErr)
}

func (suite *InvoiceServiceUnitTestsSuite) TestInvoicesService_BuyProduct_InvoicesRepositoryError() {
	invoiceItem := invoice_item.InvoiceItem{
		ProductID: 1,
		Quantity:  2,
	}
	invoice := invoice.Invoice{
		DeliveryInformation: delivery_information.DeliveryInformation{
			Name:    "Mujo",
			Surname: "Alen",
			Phone:   "+381698886969",
			Address: "Partizanska 69",
			City:    "Novi Sad",
			ZipCode: 21000,
		},
	}
	product := product.Product{
		OnStock: 5,
	}
	invoice.InvoiceItems = append(invoice.InvoiceItems, invoiceItem)
	err := rest_error.NewInternalServerError("Error when trying to save item", errors.New("test"))

	suite.productRepositoryMock.On("Get", invoiceItem.ProductID).Return(&product, nil).Once()
	suite.productRepositoryMock.On("Update", &product).Return(&product, nil).Once()
	suite.invoiceRepositoryMock.On("Save", &invoice).Return(nil, err).Once()
	buyErr := suite.service.BuyProduct(&invoice)

	assert.Equal(suite.T(), err, buyErr)
}

func (suite *InvoiceServiceUnitTestsSuite) TestInvoicesService_BuyProduct() {
	invoiceItem := invoice_item.InvoiceItem{
		ProductID: 1,
		Quantity:  2,
	}
	invoice := invoice.Invoice{
		DeliveryInformation: delivery_information.DeliveryInformation{
			Name:    "Mujo",
			Surname: "Alen",
			Phone:   "+381698886969",
			Address: "Partizanska 69",
			City:    "Novi Sad",
			ZipCode: 21000,
		},
	}
	product := product.Product{
		OnStock: 5,
	}
	invoice.InvoiceItems = append(invoice.InvoiceItems, invoiceItem)

	suite.productRepositoryMock.On("Get", invoiceItem.ProductID).Return(&product, nil).Once()
	suite.productRepositoryMock.On("Update", &product).Return(&product, nil).Once()
	suite.invoiceRepositoryMock.On("Save", &invoice).Return(&invoice, nil).Once()
	buyErr := suite.service.BuyProduct(&invoice)

	assert.Equal(suite.T(), nil, buyErr)
}
