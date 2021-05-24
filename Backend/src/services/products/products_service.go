package products

import (
	"github.com/Nistagram-Organization/Agent-Application/src/models/products"
	"github.com/Nistagram-Organization/Agent-Application/src/utils/rest_errors"
)

var (
	ProductsService productsServiceInterface = &productsService{}
)

type productsServiceInterface interface {
	Get(uint) (*products.Product, rest_errors.RestErr)
	GetAll() []products.Product
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
