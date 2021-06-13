package product

import (
	"errors"
	"fmt"
	"github.com/Nistagram-Organization/agent-shared/src/mocks/repositories"
	"github.com/Nistagram-Organization/agent-shared/src/mocks/services"
	"github.com/Nistagram-Organization/agent-shared/src/model/invoice_item"
	model "github.com/Nistagram-Organization/agent-shared/src/model/product"
	"github.com/Nistagram-Organization/agent-shared/src/utils/rest_error"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ProductServiceUnitTestsSuite struct {
	suite.Suite
	productRepositoryMock      *repositories.ProductRepositoryMock
	invoiceItemsRepositoryMock *repositories.InvoiceItemsRepositoryMock
	imageUtilsServiceMock      *services.ImageUtilsServiceMock
	temp                       string
	service                    ProductService
}

func TestProductServiceUnitTestsSuite(t *testing.T) {
	suite.Run(t, new(ProductServiceUnitTestsSuite))
}

func (suite *ProductServiceUnitTestsSuite) SetupSuite() {
	suite.productRepositoryMock = new(repositories.ProductRepositoryMock)
	suite.invoiceItemsRepositoryMock = new(repositories.InvoiceItemsRepositoryMock)
	suite.imageUtilsServiceMock = new(services.ImageUtilsServiceMock)
	suite.temp = "/temp"
	suite.service = NewProductService(
		suite.productRepositoryMock,
		suite.invoiceItemsRepositoryMock,
		suite.imageUtilsServiceMock,
		suite.temp,
	)
}

func (suite *ProductServiceUnitTestsSuite) TestNewProductsService() {
	assert.NotNil(suite.T(), suite.service, "Service is nil")
}

func (suite *ProductServiceUnitTestsSuite) TestProductsService_Get_ProductNotFound() {
	id := uint(3)
	err := rest_error.NewNotFoundError(fmt.Sprintf("Error when trying to get product with id %d", id))

	suite.productRepositoryMock.On("Get", id).Return(nil, err).Once()

	product, getErr := suite.service.Get(id)

	suite.productRepositoryMock.AssertCalled(suite.T(), "Get", id)
	assert.Nil(suite.T(), product)
	assert.Equal(suite.T(), err, getErr)
}

func (suite *ProductServiceUnitTestsSuite) TestProductsService_Get() {
	id := uint(3)
	product := model.Product{
		ID:          id,
		Name:        "proizvod",
		Description: "opis",
		Price:       300,
		OnStock:     4,
		Image:       "temp/img.jpg",
	}
	base64Img := "data:image/jpg;base64,R0lGODlhAQABAAAAACH5BAEKAAEALAAAAAABAAEAAAICTAEAOw=="

	suite.productRepositoryMock.On("Get", id).Return(&product, nil).Once()
	suite.imageUtilsServiceMock.On("LoadImage", product.Image).Return(base64Img, nil).Once()

	returned, getErr := suite.service.Get(id)

	suite.productRepositoryMock.AssertCalled(suite.T(), "Get", id)
	assert.Nil(suite.T(), getErr)
	assert.NotNil(suite.T(), returned)
	assert.Equal(suite.T(), product.ID, returned.ID)
	assert.Equal(suite.T(), product.Name, returned.Name)
	assert.Equal(suite.T(), product.Description, returned.Description)
	assert.Equal(suite.T(), product.Price, returned.Price)
	assert.Equal(suite.T(), product.OnStock, returned.OnStock)
	assert.Equal(suite.T(), base64Img, returned.Image)
}

func (suite *ProductServiceUnitTestsSuite) TestProductsService_GetAll() {
	all_products := []model.Product{
		{
			ID:          2,
			Name:        "p1",
			Description: "d1",
			Price:       30,
			OnStock:     6,
			Image:       "temp/img1.jpg",
		},
		{
			ID:          3,
			Name:        "p2",
			Description: "d2",
			Price:       50,
			OnStock:     0,
			Image:       "temp/img2.jpg",
		},
	}
	base64Img := "data:image/jpg;base64,R0lGODlhAQABAAAAACH5BAEKAAEALAAAAAABAAEAAAICTAEAOw=="

	suite.productRepositoryMock.On("GetAll").Return(all_products).Once()
	for _, product := range all_products {
		suite.imageUtilsServiceMock.On("LoadImage", product.Image).Return(base64Img, nil).Once()
	}

	returned := suite.service.GetAll()

	assert.NotNil(suite.T(), returned)
	assert.NotEmpty(suite.T(), returned)
	assert.Equal(suite.T(), 2, len(returned))
	for i := 0; i < len(returned); i++ {
		assert.Equal(suite.T(), all_products[i].ID, returned[i].ID)
		assert.Equal(suite.T(), all_products[i].Name, returned[i].Name)
		assert.Equal(suite.T(), all_products[i].Description, returned[i].Description)
		assert.Equal(suite.T(), all_products[i].Price, returned[i].Price)
		assert.Equal(suite.T(), all_products[i].OnStock, returned[i].OnStock)
		assert.Equal(suite.T(), base64Img, returned[i].Image)
	}
}

func (suite *ProductServiceUnitTestsSuite) TestProductsService_Create_ProductInvalid() {
	product := model.Product{
		Name:        "",
		Description: "opis",
		Price:       300,
		OnStock:     4,
		Image:       "data:image/jpg;base64,R0lGODlhAQABAAAAACH5BAEKAAEALAAAAAABAAEAAAICTAEAOw==",
	}
	err := rest_error.NewBadRequestError("Product name cannot be empty")

	returned, createErr := suite.service.Create(&product)

	assert.Nil(suite.T(), returned)
	assert.NotNil(suite.T(), createErr)
	assert.Equal(suite.T(), err, createErr)
}

func (suite *ProductServiceUnitTestsSuite) TestProductsService_Create_ImageInvalid() {
	product := model.Product{
		Name:        "p1",
		Description: "opis",
		Price:       300,
		OnStock:     4,
		Image:       "data:image/jpg;base64,R0lGODlhAQABAAAAACH5BAEKAAEALAAAAAABAAEAAAICTAEAOw==",
	}
	err := rest_error.NewBadRequestError("Error when decoding image")

	suite.imageUtilsServiceMock.On("SaveImage", product.Image, suite.temp).Return("", err).Once()

	returned, createErr := suite.service.Create(&product)

	assert.Nil(suite.T(), returned)
	assert.NotNil(suite.T(), createErr)
	assert.Equal(suite.T(), err, createErr)
}

func (suite *ProductServiceUnitTestsSuite) TestProductsService_Create_RepositoryError() {
	product := model.Product{
		Name:        "p1",
		Description: "opis",
		Price:       300,
		OnStock:     4,
		Image:       "data:image/jpg;base64,R0lGODlhAQABAAAAACH5BAEKAAEALAAAAAABAAEAAAICTAEAOw==",
	}
	err := rest_error.NewInternalServerError("Error when trying to create product", errors.New("test"))

	suite.imageUtilsServiceMock.On("SaveImage", product.Image, suite.temp).Return("temp/img.jpg", nil).Once()
	suite.productRepositoryMock.On("Create", &product).Return(nil, err).Once()

	returned, createErr := suite.service.Create(&product)

	assert.Nil(suite.T(), returned)
	assert.NotNil(suite.T(), createErr)
	assert.Equal(suite.T(), err, createErr)
}

func (suite *ProductServiceUnitTestsSuite) TestProductsService_Create() {
	product := model.Product{
		Name:        "p1",
		Description: "opis",
		Price:       300,
		OnStock:     4,
		Image:       "data:image/jpg;base64,R0lGODlhAQABAAAAACH5BAEKAAEALAAAAAABAAEAAAICTAEAOw==",
	}
	expectedProduct := model.Product{
		ID:          1,
		Name:        "p1",
		Description: "opis",
		Price:       300,
		OnStock:     4,
		Image:       "temp/img.jpg",
	}

	suite.imageUtilsServiceMock.On("SaveImage", product.Image, suite.temp).Return("temp/img.jpg", nil).Once()
	suite.productRepositoryMock.On("Create", &product).Return(&expectedProduct, nil).Once()

	returned, createErr := suite.service.Create(&product)

	assert.Nil(suite.T(), createErr)
	assert.NotNil(suite.T(), returned)
	assert.Equal(suite.T(), expectedProduct.ID, returned.ID)
	assert.Equal(suite.T(), expectedProduct.Name, returned.Name)
	assert.Equal(suite.T(), expectedProduct.Description, returned.Description)
	assert.Equal(suite.T(), expectedProduct.Price, returned.Price)
	assert.Equal(suite.T(), expectedProduct.OnStock, returned.OnStock)
	assert.Equal(suite.T(), expectedProduct.Image, returned.Image)
}

func (suite *ProductServiceUnitTestsSuite) TestProductsService_Delete_ProductNotFound() {
	id := uint(3)
	err := rest_error.NewNotFoundError(fmt.Sprintf("Error when trying to get product with id %d", id))

	suite.productRepositoryMock.On("Get", id).Return(nil, err).Once()

	getErr := suite.service.Delete(id)

	suite.productRepositoryMock.AssertCalled(suite.T(), "Get", id)
	assert.Equal(suite.T(), err, getErr)
}

func (suite *ProductServiceUnitTestsSuite) TestProductsService_Delete_ProductCantBeDeleted() {
	id := uint(3)
	err := rest_error.NewBadRequestError("Product can't be deleted because invoice for it exists")

	suite.productRepositoryMock.On("Get", id).Return(&model.Product{}, nil).Once()
	suite.invoiceItemsRepositoryMock.On("GetByProduct", id).Return(&invoice_item.InvoiceItem{}, nil).Once()

	getItemErr := suite.service.Delete(id)

	suite.productRepositoryMock.AssertCalled(suite.T(), "Get", id)
	suite.invoiceItemsRepositoryMock.AssertCalled(suite.T(), "GetByProduct", id)
	assert.Equal(suite.T(), err, getItemErr)
}

func (suite *ProductServiceUnitTestsSuite) TestProductsService_Delete_RepositoryError() {
	id := uint(3)
	product := model.Product{}
	getErr := rest_error.NewNotFoundError("Error when trying to get invoice item by product")
	err := rest_error.NewInternalServerError("Error when trying to delete product", errors.New("test"))

	suite.productRepositoryMock.On("Get", id).Return(&product, nil).Once()
	suite.invoiceItemsRepositoryMock.On("GetByProduct", id).Return(nil, getErr).Once()
	suite.productRepositoryMock.On("Delete", &product).Return(err).Once()

	repErr := suite.service.Delete(id)

	suite.productRepositoryMock.AssertCalled(suite.T(), "Get", id)
	suite.productRepositoryMock.AssertCalled(suite.T(), "Delete", &product)
	assert.Equal(suite.T(), err, repErr)
}

func (suite *ProductServiceUnitTestsSuite) TestProductsService_Delete() {
	id := uint(3)
	product := model.Product{}
	getErr := rest_error.NewNotFoundError("Error when trying to get invoice item by product")

	suite.productRepositoryMock.On("Get", id).Return(&product, nil).Once()
	suite.invoiceItemsRepositoryMock.On("GetByProduct", id).Return(nil, getErr).Once()
	suite.productRepositoryMock.On("Delete", &product).Return(nil).Once()

	result := suite.service.Delete(id)

	suite.productRepositoryMock.AssertCalled(suite.T(), "Get", id)
	suite.productRepositoryMock.AssertCalled(suite.T(), "Delete", &product)
	assert.Equal(suite.T(), nil, result)
}

func (suite *ProductServiceUnitTestsSuite) TestProductsService_Edit_ProductInvalid() {
	product := model.Product{
		Name:        "",
		Description: "opis",
		Price:       300,
		OnStock:     4,
		Image:       "data:image/jpg;base64,R0lGODlhAQABAAAAACH5BAEKAAEALAAAAAABAAEAAAICTAEAOw==",
	}
	err := rest_error.NewBadRequestError("Product name cannot be empty")

	returned, editErr := suite.service.Edit(&product)

	assert.Nil(suite.T(), returned)
	assert.NotNil(suite.T(), editErr)
	assert.Equal(suite.T(), err, editErr)
}

func (suite *ProductServiceUnitTestsSuite) TestProductsService_Edit_ProductNotFound() {
	id := uint(3)
	product := model.Product{
		ID:          id,
		Name:        "Proizvod",
		Description: "opis",
		Price:       300,
		OnStock:     4,
		Image:       "data:image/jpg;base64,R0lGODlhAQABAAAAACH5BAEKAAEALAAAAAABAAEAAAICTAEAOw==",
	}
	err := rest_error.NewNotFoundError(fmt.Sprintf("Error when trying to get product with id %d", id))

	suite.productRepositoryMock.On("Get", id).Return(nil, err).Once()

	returned, editErr := suite.service.Edit(&product)

	suite.productRepositoryMock.AssertCalled(suite.T(), "Get", id)
	assert.Nil(suite.T(), returned)
	assert.NotNil(suite.T(), editErr)
	assert.Equal(suite.T(), err, editErr)
}

func (suite *ProductServiceUnitTestsSuite) TestProductsService_Edit_ImageInvalid() {
	id := uint(3)
	product := model.Product{
		ID:          id,
		Name:        "p1",
		Description: "opis",
		Price:       300,
		OnStock:     4,
		Image:       "data:image/jpg;base64,R0lGODlhAQABAAAAACH5BAEKAAEALAAAAAABAAEAAAICTAEAOw==",
	}
	err := rest_error.NewBadRequestError("Error when decoding image")

	suite.productRepositoryMock.On("Get", id).Return(&product, nil).Once()
	suite.imageUtilsServiceMock.On("SaveImage", product.Image, suite.temp).Return("", err).Once()

	returned, editErr := suite.service.Edit(&product)

	suite.productRepositoryMock.AssertCalled(suite.T(), "Get", id)
	assert.Nil(suite.T(), returned)
	assert.NotNil(suite.T(), editErr)
	assert.Equal(suite.T(), err, editErr)
}

func (suite *ProductServiceUnitTestsSuite) TestProductsService_Edit_RepositoryError() {
	id := uint(3)
	product := model.Product{
		ID:          id,
		Name:        "p1",
		Description: "opis",
		Price:       300,
		OnStock:     4,
		Image:       "data:image/jpg;base64,R0lGODlhAQABAAAAACH5BAEKAAEALAAAAAABAAEAAAICTAEAOw==",
	}
	err := rest_error.NewInternalServerError("Error when trying to update product", errors.New("test"))

	suite.productRepositoryMock.On("Get", id).Return(&product, nil).Once()
	suite.imageUtilsServiceMock.On("SaveImage", product.Image, suite.temp).Return("temp/img.jpg", nil).Once()
	suite.productRepositoryMock.On("Update", &product).Return(nil, err).Once()

	returned, editErr := suite.service.Edit(&product)

	suite.productRepositoryMock.AssertCalled(suite.T(), "Get", id)
	assert.Nil(suite.T(), returned)
	assert.NotNil(suite.T(), editErr)
	assert.Equal(suite.T(), err, editErr)
}

func (suite *ProductServiceUnitTestsSuite) TestProductsService_Edit() {
	id := uint(1)
	product := model.Product{
		ID:          id,
		Name:        "p1",
		Description: "opis",
		Price:       300,
		OnStock:     4,
		Image:       "data:image/jpg;base64,R0lGODlhAQABAAAAACH5BAEKAAEALAAAAAABAAEAAAICTAEAOw==",
	}
	expectedProduct := model.Product{
		ID:          1,
		Name:        "p1",
		Description: "opis",
		Price:       300,
		OnStock:     4,
		Image:       "temp/img.jpg",
	}

	suite.productRepositoryMock.On("Get", id).Return(&product, nil).Once()
	suite.imageUtilsServiceMock.On("SaveImage", product.Image, suite.temp).Return("temp/img.jpg", nil).Once()
	suite.productRepositoryMock.On("Update", &product).Return(&expectedProduct, nil).Once()

	returned, editErr := suite.service.Edit(&product)

	suite.productRepositoryMock.AssertCalled(suite.T(), "Get", id)
	assert.Nil(suite.T(), editErr)
	assert.NotNil(suite.T(), returned)
	assert.Equal(suite.T(), expectedProduct.ID, returned.ID)
	assert.Equal(suite.T(), expectedProduct.Name, returned.Name)
	assert.Equal(suite.T(), expectedProduct.Description, returned.Description)
	assert.Equal(suite.T(), expectedProduct.Price, returned.Price)
	assert.Equal(suite.T(), expectedProduct.OnStock, returned.OnStock)
	assert.Equal(suite.T(), expectedProduct.Image, returned.Image)
}
