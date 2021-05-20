package products

import (
	"github.com/Nistagram-Organization/Agent-Application/src/models/invoices"
	"github.com/Nistagram-Organization/Agent-Application/src/models/products"
	"github.com/Nistagram-Organization/Agent-Application/src/utils/rest_errors"
	"time"
)

var (
	ProductsService productsServiceInterface = &productsService{}
)

type productsServiceInterface interface {
	Get(uint) (*products.Product, rest_errors.RestErr)
	GetAll() []products.Product
	Buy(invoice *invoices.Invoice) rest_errors.RestErr
}

type productsService struct{}

func (s *productsService) Get(id uint) (*products.Product, rest_errors.RestErr) {
	product := products.Product{ID: id}

	if err := product.Get(); err != nil {
		return nil, err
	}

	return &product, nil
}

func (s *productsService) GetAll() []products.Product {
	dao := products.Product{}
	return dao.GetAll()
}

func (s *productsService) Buy(invoice *invoices.Invoice) rest_errors.RestErr {
	if invoice.InvoiceItems == nil || len(invoice.InvoiceItems) == 0 {
		return rest_errors.NewBadRequestError("No items are selected for buying")
	}

	var total float32
	for _, item := range invoice.InvoiceItems {
		product := products.Product{ID: item.ProductID}

		if err := product.Get(); err != nil {
			return err
		}

		if product.OnStock < item.Quantity {
			return rest_errors.NewBadRequestError("Not enough products in stock")
		}
		product.OnStock -= item.Quantity

		if err := product.Update(); err != nil {
			return err
		}

		item.ID = 0
		total += product.Price * float32(item.Quantity)
	}

	invoice.Date = time.Now().Format("02.01.2006. 15:04")
	invoice.Total = total
	if err := invoice.Save(); err != nil {
		return err
	}

	return nil
}