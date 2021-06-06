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
	"os"
	"testing"
)

var (
	productRepositoryMock *repositories.ProductRepositoryMock
	invoiceRepositoryMock *repositories.InvoiceRepositoryMock
	service               InvoicesService
)

func setup() {
	productRepositoryMock = new(repositories.ProductRepositoryMock)
	invoiceRepositoryMock = new(repositories.InvoiceRepositoryMock)
	service = NewInvoicesService(productRepositoryMock, invoiceRepositoryMock)
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}

func TestNewInvoicesService(t *testing.T) {
	assert.NotNil(t, service, "Service is nil")
}

func TestInvoicesService_BuyProduct_NoInvoiceItems(t *testing.T) {
	invoice := invoice.Invoice{}
	err := rest_error.NewBadRequestError("No items are selected for buying")

	buyErr := service.BuyProduct(&invoice)

	assert.Equal(t, err, buyErr)
}

func TestInvoicesService_BuyProduct_DeliveryInformationInvalid(t *testing.T) {
	invoiceItem := invoice_item.InvoiceItem{}
	invoice := invoice.Invoice{
		DeliveryInformation: delivery_information.DeliveryInformation{},
	}
	invoice.InvoiceItems = append(invoice.InvoiceItems, invoiceItem)

	err := rest_error.NewBadRequestError("Customer name cannot be empty")

	buyErr := service.BuyProduct(&invoice)

	assert.Equal(t, err, buyErr)
}

func TestInvoicesService_BuyProduct_InvoiceItemInvalid(t *testing.T) {
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

	buyErr := service.BuyProduct(&invoice)

	assert.Equal(t, err, buyErr)
}

func TestInvoicesService_BuyProduct_ProductNotFound(t *testing.T) {
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

	productRepositoryMock.On("Get", invoiceItem.ProductID).Return(nil, err).Once()

	buyErr := service.BuyProduct(&invoice)

	assert.Equal(t, err, buyErr)
}

func TestInvoicesService_BuyProduct_NotEnoughProducts(t *testing.T) {
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

	productRepositoryMock.On("Get", invoiceItem.ProductID).Return(&product, nil).Once()

	buyErr := service.BuyProduct(&invoice)

	assert.Equal(t, err, buyErr)
}

func TestInvoicesService_BuyProduct_ProductsRepositoryError(t *testing.T) {
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

	productRepositoryMock.On("Get", invoiceItem.ProductID).Return(&product, nil).Once()
	productRepositoryMock.On("Update", &product).Return(nil, err).Once()
	buyErr := service.BuyProduct(&invoice)

	assert.Equal(t, err, buyErr)
}

func TestInvoicesService_BuyProduct_InvoicesRepositoryError(t *testing.T) {
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

	productRepositoryMock.On("Get", invoiceItem.ProductID).Return(&product, nil).Once()
	productRepositoryMock.On("Update", &product).Return(&product, nil).Once()
	invoiceRepositoryMock.On("Save", &invoice).Return(nil, err).Once()
	buyErr := service.BuyProduct(&invoice)

	assert.Equal(t, err, buyErr)
}

func TestInvoicesService_BuyProduct(t *testing.T) {
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

	productRepositoryMock.On("Get", invoiceItem.ProductID).Return(&product, nil).Once()
	productRepositoryMock.On("Update", &product).Return(&product, nil).Once()
	invoiceRepositoryMock.On("Save", &invoice).Return(&invoice, nil).Once()
	buyErr := service.BuyProduct(&invoice)

	assert.Equal(t, nil, buyErr)
}
