package product

import (
	"github.com/Nistagram-Organization/Agent-Application/agent-shared/src/utils/rest_error"
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

func (p *Product) Validate() rest_error.RestErr {
	if strings.TrimSpace(p.Name) == "" {
		return rest_error.NewBadRequestError("Product name cannot be empty")
	}
	if strings.TrimSpace(p.Description) == "" {
		return rest_error.NewBadRequestError("Product description cannot be empty")
	}
	if p.Price <= 0 {
		return rest_error.NewBadRequestError("Product price must be greater than zero")
	}
	if strings.TrimSpace(p.Image) == "" {
		return rest_error.NewBadRequestError("Product image cannot be empty")
	}
	return nil
}
