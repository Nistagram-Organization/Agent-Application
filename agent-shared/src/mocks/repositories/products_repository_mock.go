package repositories

import (
	model "github.com/Nistagram-Organization/Agent-Application/agent-shared/src/model/product"
	"github.com/Nistagram-Organization/Agent-Application/agent-shared/src/utils/rest_error"
	"github.com/stretchr/testify/mock"
)

type ProductRepositoryMock struct {
	mock.Mock
}

func (p *ProductRepositoryMock) Get(id uint) (*model.Product, rest_error.RestErr) {
	args := p.Called(id)
	if args.Get(0) == nil {
		return nil, args.Get(1).(rest_error.RestErr)
	}
	return args.Get(0).(*model.Product), nil
}

func (p *ProductRepositoryMock) GetAll() []model.Product {
	args := p.Called()
	return args.Get(0).([]model.Product)
}

func (p *ProductRepositoryMock) Create(product *model.Product) (*model.Product, rest_error.RestErr) {
	args := p.Called(product)
	if args.Get(1) == nil {
		return args.Get(0).(*model.Product), nil
	}
	return nil, args.Get(1).(rest_error.RestErr)
}

func (p *ProductRepositoryMock) Update(product *model.Product) (*model.Product, rest_error.RestErr) {
	args := p.Called(product)
	if args.Get(1) == nil {
		return args.Get(0).(*model.Product), nil
	}
	return nil, args.Get(1).(rest_error.RestErr)
}

func (p *ProductRepositoryMock) Delete(product *model.Product) rest_error.RestErr {
	args := p.Called(product)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(rest_error.RestErr)
}
