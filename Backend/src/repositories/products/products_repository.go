package products

import (
	"fmt"
	"github.com/Nistagram-Organization/Agent-Application/src/datasources/mysql/agent_application_db"
	"github.com/Nistagram-Organization/Agent-Application/src/model/products"
	"github.com/Nistagram-Organization/Agent-Application/src/utils/rest_errors"
	"gorm.io/gorm"
)

type ProductsRepository interface {
	Get(uint) (*products.Product, rest_errors.RestErr)
	GetAll() []products.Product
	Create(*products.Product) (*products.Product, rest_errors.RestErr)
	Update(*products.Product) (*products.Product, rest_errors.RestErr)
	Delete(*products.Product) rest_errors.RestErr
}
type productsRepository struct {
	db *gorm.DB
}

func NewProductsRepository() ProductsRepository {
	return &productsRepository{
		db: agent_application_db.Client.GetClient(),
	}
}

func (r *productsRepository) Get(id uint) (*products.Product, rest_errors.RestErr) {
	product := products.Product{
		ID: id,
	}
	if err := r.db.Take(&product, product.ID).Error; err != nil {
		fmt.Sprintln(err)
		return nil, rest_errors.NewNotFoundError(fmt.Sprintf("Error when trying to get product with id %d", product.ID))
	}
	return &product, nil
}

func (r *productsRepository) GetAll() []products.Product {
	var collection []products.Product
	if err := r.db.Find(&collection).Error; err != nil {
		return []products.Product{}
	}
	return collection
}

func (r *productsRepository) Create(product *products.Product) (*products.Product, rest_errors.RestErr) {
	if err := r.db.Create(product).Error; err != nil {
		return nil, rest_errors.NewInternalServerError("Error when trying to create product", err)
	}
	return product, nil
}

func (r *productsRepository) Update(product *products.Product) (*products.Product, rest_errors.RestErr) {
	if err := r.db.Save(product).Error; err != nil {
		return nil, rest_errors.NewInternalServerError("Error when trying to update product", err)
	}
	return product, nil
}

func (r *productsRepository) Delete(product *products.Product) rest_errors.RestErr {
	if err := r.db.Delete(product).Error; err != nil {
		return rest_errors.NewInternalServerError("Error when trying to delete product", err)
	}
	return nil
}
