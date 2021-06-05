package product

import (
	"github.com/Nistagram-Organization/Agent-Application/agent-products/src/utils/image_utils"
	model "github.com/Nistagram-Organization/Agent-Application/agent-shared/src/model/product"
	"github.com/Nistagram-Organization/Agent-Application/agent-shared/src/repositories/invoice_item"
	"github.com/Nistagram-Organization/Agent-Application/agent-shared/src/repositories/product"
	"github.com/Nistagram-Organization/Agent-Application/agent-shared/src/utils/rest_error"
)

type ProductService interface {
	Get(uint) (*model.Product, rest_error.RestErr)
	GetAll() []model.Product
	Create(*model.Product) (*model.Product, rest_error.RestErr)
	Delete(uint) rest_error.RestErr
	Edit(*model.Product) (*model.Product, rest_error.RestErr)
}

type productsService struct {
	productsRepository     product.ProductRepository
	invoiceItemsRepository invoice_item.InvoiceItemRepository
	imageUtilsService      image_utils.ImageUtilsService
	tempFolder             string
}

func NewProductService(productsRepository product.ProductRepository, invoiceItemsRepository invoice_item.InvoiceItemRepository, imageUtilsService image_utils.ImageUtilsService, tempFolder string) ProductService {
	return &productsService{
		productsRepository:     productsRepository,
		invoiceItemsRepository: invoiceItemsRepository,
		imageUtilsService:      imageUtilsService,
		tempFolder:             tempFolder,
	}
}

func (s *productsService) Get(id uint) (*model.Product, rest_error.RestErr) {
	product, err := s.productsRepository.Get(id)
	if err != nil {
		return nil, err
	}
	product.Image, _ = s.imageUtilsService.LoadImage(product.Image)
	return product, nil
}

func (s *productsService) GetAll() []model.Product {
	products := s.productsRepository.GetAll()
	for i := 0; i < len(products); i++ {
		products[i].Image, _ = s.imageUtilsService.LoadImage(products[i].Image)
	}
	return products
}

func (s *productsService) Create(product *model.Product) (*model.Product, rest_error.RestErr) {
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

func (s *productsService) Edit(product *model.Product) (*model.Product, rest_error.RestErr) {
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

func (s *productsService) Delete(id uint) rest_error.RestErr {
	product, getErr := s.productsRepository.Get(id)
	if getErr != nil {
		return getErr
	}

	if _, getItemErr := s.invoiceItemsRepository.GetByProduct(id); getItemErr != nil {
		return s.productsRepository.Delete(product)
	}

	return rest_error.NewBadRequestError("Product can't be deleted because invoice for it exists")
}
