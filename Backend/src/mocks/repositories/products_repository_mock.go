package repositories

import (
	"github.com/Nistagram-Organization/Agent-Application/src/model/products"
	"github.com/Nistagram-Organization/Agent-Application/src/utils/rest_errors"
	"github.com/stretchr/testify/mock"
)

type ProductRepositoryMock struct {
	mock.Mock
}

func (p *ProductRepositoryMock) Get(id uint) (*products.Product, rest_errors.RestErr) {
	args := p.Called(id)
	if args.Get(0) == nil {
		return nil, args.Get(1).(rest_errors.RestErr)
	}
	return args.Get(0).(*products.Product), nil
}

func (p *ProductRepositoryMock) GetAll() []products.Product {
	args := p.Called()
	return args.Get(0).([]products.Product)
}

func (p *ProductRepositoryMock) Create(product *products.Product) (*products.Product, rest_errors.RestErr) {
	args := p.Called(product)
	if args.Get(1) == nil {
		return args.Get(0).(*products.Product), nil
	}
	return nil, args.Get(1).(rest_errors.RestErr)
}

func (p *ProductRepositoryMock) Update(product *products.Product) (*products.Product, rest_errors.RestErr) {
	return nil, nil
}

func (p *ProductRepositoryMock) Delete(product *products.Product) rest_errors.RestErr {
	return nil
}
