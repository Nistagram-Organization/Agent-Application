package invoices

import (
	"github.com/Nistagram-Organization/Agent-Application/src/models/invoices"
	"github.com/Nistagram-Organization/Agent-Application/src/models/products"
	"github.com/Nistagram-Organization/Agent-Application/src/utils/rest_errors"
	"github.com/Nistagram-Organization/Agent-Application/src/utils/time_utils"
)

var (
	InvoicesService invoicesServiceInterface = &invoicesService{}
)

type invoicesServiceInterface interface {
	BuyProduct(invoice *invoices.Invoice) rest_errors.RestErr
}

type invoicesService struct{}

func (s *invoicesService) BuyProduct(invoice *invoices.Invoice) rest_errors.RestErr {
	if err := invoice.Validate(); err != nil {
		return err
	}

	var total float32
	for _, item := range invoice.InvoiceItems {
		if err := item.Validate(); err != nil {
			return err
		}

		product := products.Product{ID: item.ProductID}

		if err := product.Get(); err != nil {
			return err
		}

		if product.OnStock < item.Quantity {
			return rest_errors.NewBadRequestError("Not enough products in stock")
		}
		product.OnStock -= item.Quantity

		if err := product.Update(); err != nil {
			return err
		}

		item.ID = 0
		total += product.Price * float32(item.Quantity)
	}

	invoice.Date = time_utils.Now()
	invoice.Total = total
	if err := invoice.Save(); err != nil {
		return err
	}

	return nil
}
