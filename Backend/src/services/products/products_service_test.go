package products

import (
	"errors"
	"fmt"
	"github.com/Nistagram-Organization/Agent-Application/src/mocks/repositories"
	"github.com/Nistagram-Organization/Agent-Application/src/mocks/services"
	"github.com/Nistagram-Organization/Agent-Application/src/model/products"
	"github.com/Nistagram-Organization/Agent-Application/src/utils/rest_errors"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

const (
	temp = "/temp"
)

var (
	productRepositoryMock      *repositories.ProductRepositoryMock
	invoiceItemsRepositoryMock *repositories.InvoiceItemsRepositoryMock
	imageUtilsServiceMock      *services.ImageUtilsServiceMock
	service                    ProductsService
)

func setup() {
	productRepositoryMock = new(repositories.ProductRepositoryMock)
	invoiceItemsRepositoryMock = new(repositories.InvoiceItemsRepositoryMock)
	imageUtilsServiceMock = new(services.ImageUtilsServiceMock)
	service = NewProductsService(productRepositoryMock, invoiceItemsRepositoryMock, imageUtilsServiceMock, temp)
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}

func TestNewProductsService(t *testing.T) {
	assert.NotNil(t, service, "Service is nil")
}

func TestProductsService_Get_ProductNotFound(t *testing.T) {
	id := uint(3)
	err := rest_errors.NewNotFoundError(fmt.Sprintf("Error when trying to get product with id %d", id))

	productRepositoryMock.On("Get", id).Return(nil, err).Once()

	product, getErr := service.Get(id)

	productRepositoryMock.AssertCalled(t, "Get", id)
	assert.Nil(t, product)
	assert.Equal(t, err, getErr)
}

func TestProductsService_Get(t *testing.T) {
	id := uint(3)
	product := products.Product{
		ID:          id,
		Name:        "proizvod",
		Description: "opis",
		Price:       300,
		OnStock:     4,
		Image:       "temp/img.jpg",
	}
	base64Img := "data:image/jpg;base64,R0lGODlhAQABAAAAACH5BAEKAAEALAAAAAABAAEAAAICTAEAOw=="

	productRepositoryMock.On("Get", id).Return(&product, nil).Once()
	imageUtilsServiceMock.On("LoadImage", product.Image).Return(base64Img, nil).Once()

	returned, getErr := service.Get(id)

	productRepositoryMock.AssertCalled(t, "Get", id)
	assert.Nil(t, getErr)
	assert.NotNil(t, returned)
	assert.Equal(t, product.ID, returned.ID)
	assert.Equal(t, product.Name, returned.Name)
	assert.Equal(t, product.Description, returned.Description)
	assert.Equal(t, product.Price, returned.Price)
	assert.Equal(t, product.OnStock, returned.OnStock)
	assert.Equal(t, base64Img, returned.Image)
}

func TestProductsService_GetAll(t *testing.T) {
	products := []products.Product{
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

	productRepositoryMock.On("GetAll").Return(products).Once()
	for _, product := range products {
		imageUtilsServiceMock.On("LoadImage", product.Image).Return(base64Img, nil).Once()
	}

	returned := service.GetAll()

	assert.NotNil(t, returned)
	assert.NotEmpty(t, returned)
	assert.Equal(t, 2, len(returned))
	for i := 0; i < len(returned); i++ {
		assert.Equal(t, products[i].ID, returned[i].ID)
		assert.Equal(t, products[i].Name, returned[i].Name)
		assert.Equal(t, products[i].Description, returned[i].Description)
		assert.Equal(t, products[i].Price, returned[i].Price)
		assert.Equal(t, products[i].OnStock, returned[i].OnStock)
		assert.Equal(t, base64Img, returned[i].Image)
	}
}

func TestProductsService_Create_ProductInvalid(t *testing.T) {
	product := products.Product{
		Name:        "",
		Description: "opis",
		Price:       300,
		OnStock:     4,
		Image:       "data:image/jpg;base64,R0lGODlhAQABAAAAACH5BAEKAAEALAAAAAABAAEAAAICTAEAOw==",
	}
	err := rest_errors.NewBadRequestError("Product name cannot be empty")

	returned, createErr := service.Create(&product)

	assert.Nil(t, returned)
	assert.NotNil(t, createErr)
	assert.Equal(t, err, createErr)
}

func TestProductsService_Create_ImageInvalid(t *testing.T) {
	product := products.Product{
		Name:        "p1",
		Description: "opis",
		Price:       300,
		OnStock:     4,
		Image:       "data:image/jpg;base64,R0lGODlhAQABAAAAACH5BAEKAAEALAAAAAABAAEAAAICTAEAOw==",
	}
	err := rest_errors.NewBadRequestError("Error when decoding image")

	imageUtilsServiceMock.On("SaveImage", product.Image, temp).Return("", err).Once()

	returned, createErr := service.Create(&product)

	assert.Nil(t, returned)
	assert.NotNil(t, createErr)
	assert.Equal(t, err, createErr)
}

func TestProductsService_Create_RepositoryError(t *testing.T) {
	product := products.Product{
		Name:        "p1",
		Description: "opis",
		Price:       300,
		OnStock:     4,
		Image:       "data:image/jpg;base64,R0lGODlhAQABAAAAACH5BAEKAAEALAAAAAABAAEAAAICTAEAOw==",
	}
	err := rest_errors.NewInternalServerError("Error when trying to create product", errors.New("test"))

	imageUtilsServiceMock.On("SaveImage", product.Image, temp).Return("temp/img.jpg", nil).Once()
	productRepositoryMock.On("Create", &product).Return(nil, err).Once()

	returned, createErr := service.Create(&product)

	assert.Nil(t, returned)
	assert.NotNil(t, createErr)
	assert.Equal(t, err, createErr)
}

func TestProductsService_Create(t *testing.T) {
	product := products.Product{
		Name:        "p1",
		Description: "opis",
		Price:       300,
		OnStock:     4,
		Image:       "data:image/jpg;base64,R0lGODlhAQABAAAAACH5BAEKAAEALAAAAAABAAEAAAICTAEAOw==",
	}
	expectedProduct := products.Product{
		ID:          1,
		Name:        "p1",
		Description: "opis",
		Price:       300,
		OnStock:     4,
		Image:       "temp/img.jpg",
	}

	imageUtilsServiceMock.On("SaveImage", product.Image, temp).Return("temp/img.jpg", nil).Once()
	productRepositoryMock.On("Create", &product).Return(&expectedProduct, nil).Once()

	returned, createErr := service.Create(&product)

	assert.Nil(t, createErr)
	assert.NotNil(t, returned)
	assert.Equal(t, expectedProduct.ID, returned.ID)
	assert.Equal(t, expectedProduct.Name, returned.Name)
	assert.Equal(t, expectedProduct.Description, returned.Description)
	assert.Equal(t, expectedProduct.Price, returned.Price)
	assert.Equal(t, expectedProduct.OnStock, returned.OnStock)
	assert.Equal(t, expectedProduct.Image, returned.Image)
}
