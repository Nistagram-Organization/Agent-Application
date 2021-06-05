package product

import (
	"github.com/Nistagram-Organization/Agent-Application/agent-shared/src/model/product"
	"github.com/Nistagram-Organization/Agent-Application/agent-shared/src/utils/rest_error"
)

type ProductRepository interface {
	Get(uint) (*product.Product, rest_error.RestErr)
	GetAll() []product.Product
	Create(*product.Product) (*product.Product, rest_error.RestErr)
	Update(*product.Product) (*product.Product, rest_error.RestErr)
	Delete(*product.Product) rest_error.RestErr
}
