package product

import (
	"fmt"
	"github.com/Nistagram-Organization/Agent-Application/agent-shared/src/datasources"
	model "github.com/Nistagram-Organization/Agent-Application/agent-shared/src/model/product"
	"github.com/Nistagram-Organization/Agent-Application/agent-shared/src/repositories/product"
	"github.com/Nistagram-Organization/Agent-Application/agent-shared/src/utils/rest_error"
	"gorm.io/gorm"
)

type productsRepository struct {
	db *gorm.DB
}

func NewProductRepository(databaseClient datasources.DatabaseClient) product.ProductRepository {
	return &productsRepository{
		databaseClient.GetClient(),
	}
}

func (r *productsRepository) Get(id uint) (*model.Product, rest_error.RestErr) {
	product := model.Product{
		ID: id,
	}
	if err := r.db.Take(&product, product.ID).Error; err != nil {
		fmt.Sprintln(err)
		return nil, rest_error.NewNotFoundError(fmt.Sprintf("Error when trying to get product with id %d", product.ID))
	}
	return &product, nil
}

func (r *productsRepository) GetAll() []model.Product {
	var collection []model.Product
	if err := r.db.Find(&collection).Error; err != nil {
		return []model.Product{}
	}
	return collection
}

func (r *productsRepository) Create(product *model.Product) (*model.Product, rest_error.RestErr) {
	if err := r.db.Create(product).Error; err != nil {
		return nil, rest_error.NewInternalServerError("Error when trying to create product", err)
	}
	return product, nil
}

func (r *productsRepository) Update(product *model.Product) (*model.Product, rest_error.RestErr) {
	if err := r.db.Save(product).Error; err != nil {
		return nil, rest_error.NewInternalServerError("Error when trying to update product", err)
	}
	return product, nil
}

func (r *productsRepository) Delete(product *model.Product) rest_error.RestErr {
	if err := r.db.Delete(product).Error; err != nil {
		return rest_error.NewInternalServerError("Error when trying to delete product", err)
	}
	return nil
}
