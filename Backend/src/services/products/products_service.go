package products

import (
	"github.com/Nistagram-Organization/Agent-Application/src/models/products"
	"github.com/Nistagram-Organization/Agent-Application/src/utils/image_utils"
	"github.com/Nistagram-Organization/Agent-Application/src/utils/rest_errors"
)

const (
	TEMP_FOLDER = "temp"
)

var (
	ProductsService productsServiceInterface = &productsService{}
)

type productsServiceInterface interface {
	Get(uint) (*products.Product, rest_errors.RestErr)
	GetAll() []products.Product
	Create(*products.Product) (*products.Product, rest_errors.RestErr)
}

type productsService struct{}

func (s *productsService) Get(id uint) (*products.Product, rest_errors.RestErr) {
	product := products.Product{ID: id}

	if err := product.Get(); err != nil {
		return nil, err
	}
	product.Image, _ = image_utils.LoadImage(product.Image)

	return &product, nil
}

func (s *productsService) GetAll() []products.Product {
	dao := products.Product{}
	products := dao.GetAll()
	for _, product := range products {
		product.Image, _ = image_utils.LoadImage(product.Image)
	}
	return products
}

func (s *productsService) Create(product *products.Product) (*products.Product, rest_errors.RestErr) {
	if product.Get() == nil {
		return nil, rest_errors.NewBadRequestError("Product with given id already exists")
	}

	if err := product.Validate(); err != nil {
		return nil, err
	}

	imagePath, err := image_utils.SaveImage(product.Image, TEMP_FOLDER)
	if err != nil {
		return nil, err
	}
	product.Image = imagePath

	if err := product.Create(); err != nil {
		return nil, err
	}
	return product, nil
}
