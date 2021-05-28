package products

import (
	"github.com/Nistagram-Organization/Agent-Application/src/model/products"
	"github.com/Nistagram-Organization/Agent-Application/src/repositories/invoice_items"
	productsRepo "github.com/Nistagram-Organization/Agent-Application/src/repositories/products"
	"github.com/Nistagram-Organization/Agent-Application/src/utils/image_utils"
	"github.com/Nistagram-Organization/Agent-Application/src/utils/rest_errors"
)

type ProductsService interface {
	Get(uint) (*products.Product, rest_errors.RestErr)
	GetAll() []products.Product
	Create(*products.Product) (*products.Product, rest_errors.RestErr)
	Delete(uint) rest_errors.RestErr
	Edit(*products.Product) (*products.Product, rest_errors.RestErr)
}

type productsService struct {
	productsRepository     productsRepo.ProductsRepository
	invoiceItemsRepository invoice_items.InvoiceItemsRepository
	imageUtilsService      image_utils.ImageUtilsService
	tempFolder             string
}

func NewProductsService(productsRepository productsRepo.ProductsRepository, invoiceItemsRepository invoice_items.InvoiceItemsRepository, imageUtilsService image_utils.ImageUtilsService, tempFolder string) ProductsService {
	return &productsService{
		productsRepository:     productsRepository,
		invoiceItemsRepository: invoiceItemsRepository,
		imageUtilsService:      imageUtilsService,
		tempFolder:             tempFolder,
	}
}

func (s *productsService) Get(id uint) (*products.Product, rest_errors.RestErr) {
	product, err := s.productsRepository.Get(id)
	if err != nil {
		return nil, err
	}
	product.Image, _ = s.imageUtilsService.LoadImage(product.Image)
	return product, nil
}

func (s *productsService) GetAll() []products.Product {
	products := s.productsRepository.GetAll()
	for i := 0; i < len(products); i++ {
		products[i].Image, _ = s.imageUtilsService.LoadImage(products[i].Image)
	}
	return products
}

func (s *productsService) Create(product *products.Product) (*products.Product, rest_errors.RestErr) {
	if err := product.Validate(); err != nil {
		return nil, err
	}

	imagePath, err := s.imageUtilsService.SaveImage(product.Image, s.tempFolder)
	if err != nil {
		return nil, err
	}
	product.Image = imagePath

	return s.productsRepository.Create(product)
}

func (s *productsService) Edit(product *products.Product) (*products.Product, rest_errors.RestErr) {
	if err := product.Validate(); err != nil {
		return nil, err
	}

	_, getErr := s.productsRepository.Get(product.ID)
	if getErr != nil {
		return nil, getErr
	}

	imagePath, err := s.imageUtilsService.SaveImage(product.Image, s.tempFolder)
	if err != nil {
		return nil, err
	}
	product.Image = imagePath

	return s.productsRepository.Update(product)
}

func (s *productsService) Delete(id uint) rest_errors.RestErr {
	product, getErr := s.productsRepository.Get(id)
	if getErr != nil {
		return getErr
	}

	if _, getItemErr := s.invoiceItemsRepository.GetByProduct(id); getItemErr != nil {
		return s.productsRepository.Delete(product)
	}

	return rest_errors.NewBadRequestError("Product can't be deleted because invoice for it exists")
}
