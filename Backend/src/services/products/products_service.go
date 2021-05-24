package products

import (
	"github.com/Nistagram-Organization/Agent-Application/src/models/invoice_items"
	"github.com/Nistagram-Organization/Agent-Application/src/models/products"
	"github.com/Nistagram-Organization/Agent-Application/src/utils/rest_errors"
)

var (
	ProductsService productsServiceInterface = &productsService{}
)

type productsServiceInterface interface {
	Get(uint) (*products.Product, rest_errors.RestErr)
	GetAll() []products.Product
	Delete(uint) rest_errors.RestErr
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

func (s *productsService) Delete(id uint) rest_errors.RestErr {
	product := products.Product{ID: id}
	if getErr := product.Get(); getErr != nil {
		return getErr
	}

	invoiceItem := invoice_items.InvoiceItem{}
	if getItemErr := invoiceItem.GetByProduct(id); getItemErr != nil {
		if delErr := product.Delete(); delErr != nil {
			return delErr
		}
		return nil
	}

	return rest_errors.NewBadRequestError("Product can't be deleted because invoice for it exists")
}
