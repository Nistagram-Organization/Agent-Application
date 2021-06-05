package delivery_information

import (
	"github.com/Nistagram-Organization/Agent-Application/agent-shared/src/utils/rest_error"
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

func (di DeliveryInformation) Validate() rest_error.RestErr {
	if strings.TrimSpace(di.Name) == "" {
		return rest_error.NewBadRequestError("Customer name cannot be empty")
	}
	if strings.TrimSpace(di.Surname) == "" {
		return rest_error.NewBadRequestError("Customer surname cannot be empty")
	}
	if strings.TrimSpace(di.Phone) == "" {
		return rest_error.NewBadRequestError("Customer phone cannot be empty")
	}
	if strings.TrimSpace(di.Address) == "" {
		return rest_error.NewBadRequestError("Customer address cannot be empty")
	}
	if strings.TrimSpace(di.City) == "" {
		return rest_error.NewBadRequestError("Customer city cannot be empty")
	}
	if di.ZipCode == 0 {
		return rest_error.NewBadRequestError("Zip code must be set")
	}

	return nil
}
