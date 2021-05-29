package invoices

import (
	"github.com/Nistagram-Organization/Agent-Application/src/model/invoices"
	invoicesRepo "github.com/Nistagram-Organization/Agent-Application/src/repositories/invoices"
	productsRepo "github.com/Nistagram-Organization/Agent-Application/src/repositories/products"
	"github.com/Nistagram-Organization/Agent-Application/src/utils/rest_errors"
	"github.com/Nistagram-Organization/Agent-Application/src/utils/time_utils"
)

type InvoicesService interface {
	BuyProduct(*invoices.Invoice) rest_errors.RestErr
}

type invoicesService struct {
	productsRepository productsRepo.ProductsRepository
	invoicesRepository invoicesRepo.InvoicesRepository
}

func NewInvoicesService(productsRepository productsRepo.ProductsRepository, invoicesRepository invoicesRepo.InvoicesRepository) InvoicesService {
	return &invoicesService{
		productsRepository: productsRepository,
		invoicesRepository: invoicesRepository,
	}
}

func (s *invoicesService) BuyProduct(invoice *invoices.Invoice) rest_errors.RestErr {
	if err := invoice.Validate(); err != nil {
		return err
	}

	var total float32
	for _, item := range invoice.InvoiceItems {
		if err := item.Validate(); err != nil {
			return err
		}

		product, err := s.productsRepository.Get(item.ProductID)

		if err != nil {
			return err
		}

		if product.OnStock < item.Quantity {
			return rest_errors.NewBadRequestError("Not enough products in stock")
		}
		product.OnStock -= item.Quantity

		product, err = s.productsRepository.Update(product)
		if err != nil {
			return err
		}

		item.ID = 0
		total += product.Price * float32(item.Quantity)
	}

	invoice.Date = time_utils.Now()
	invoice.Total = total

	_, err := s.invoicesRepository.Save(invoice)
	if err != nil {
		return err
	}
	return nil
}
