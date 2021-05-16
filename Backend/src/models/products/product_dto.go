package products

import (
	"github.com/Nistagram-Organization/Agent-Application/src/utils/rest_errors"
	"strings"
)

type Product struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	OnStock     uint    `json:"on_stock"`
	Image       string  `json:"image"`
}

func (p *Product) Validate() rest_errors.RestErr {
	if strings.TrimSpace(p.Name) == "" {
		return rest_errors.NewBadRequestError("Product name cannot be empty")
	}
	if strings.TrimSpace(p.Description) == "" {
		return rest_errors.NewBadRequestError("Product description cannot be empty")
	}
	if p.Price <= 0 {
		return rest_errors.NewBadRequestError("Product price must be greater than zero")
	}
	if p.OnStock < 0 {
		return rest_errors.NewBadRequestError("Product stock must be equal or greater than zero")
	}
	if strings.TrimSpace(p.Image) == "" {
		return rest_errors.NewBadRequestError("Product image cannot be empty")
	}
	return nil
}
