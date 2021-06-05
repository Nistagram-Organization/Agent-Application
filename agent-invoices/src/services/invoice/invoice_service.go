package invoice

import (
	"github.com/Nistagram-Organization/Agent-Application/agent-shared/src/model/invoice"
	invoiceRepo "github.com/Nistagram-Organization/Agent-Application/agent-shared/src/repositories/invoice"
	productRepo "github.com/Nistagram-Organization/Agent-Application/agent-shared/src/repositories/product"
	"github.com/Nistagram-Organization/Agent-Application/agent-shared/src/utils/rest_error"
	"github.com/Nistagram-Organization/Agent-Application/agent-shared/src/utils/time_utils"
)

type InvoicesService interface {
	BuyProduct(*invoice.Invoice) rest_error.RestErr
}

type invoicesService struct {
	productsRepository productRepo.ProductRepository
	invoicesRepository invoiceRepo.InvoiceRepository
}

func NewInvoicesService(productsRepository productRepo.ProductRepository, invoicesRepository invoiceRepo.InvoiceRepository) InvoicesService {
	return &invoicesService{
		productsRepository: productsRepository,
		invoicesRepository: invoicesRepository,
	}
}

func (s *invoicesService) BuyProduct(invoice *invoice.Invoice) rest_error.RestErr {
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
			return rest_error.NewBadRequestError("Not enough products in stock")
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
