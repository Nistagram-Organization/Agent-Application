package delivery_informations

import (
	"github.com/Nistagram-Organization/Agent-Application/src/utils/rest_errors"
	"strings"
)

type DeliveryInformation struct {
	ID        uint
	Name      string
	Surname   string
	Phone     string
	Address   string
	City      string
	ZipCode   uint
	InvoiceID uint
}

func (di DeliveryInformation) Validate() rest_errors.RestErr {
	if strings.TrimSpace(di.Name) == "" {
		return rest_errors.NewBadRequestError("Customer name cannot be empty")
	}
	if strings.TrimSpace(di.Surname) == "" {
		return rest_errors.NewBadRequestError("Customer surname cannot be empty")
	}
	if strings.TrimSpace(di.Phone) == "" {
		return rest_errors.NewBadRequestError("Customer phone cannot be empty")
	}
	if strings.TrimSpace(di.Address) == "" {
		return rest_errors.NewBadRequestError("Customer address cannot be empty")
	}
	if strings.TrimSpace(di.City) == "" {
		return rest_errors.NewBadRequestError("Customer city cannot be empty")
	}
	if di.ZipCode == 0 {
		return rest_errors.NewBadRequestError("Zip code must be set")
	}

	return nil
}
