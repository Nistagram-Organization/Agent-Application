package product_reports

import (
	"github.com/Nistagram-Organization/Agent-Application/src/model/product_reports"
	productReportsRepo "github.com/Nistagram-Organization/Agent-Application/src/repositories/product_reports"
)

type ProductReportsService interface {
	GenerateReport() []product_reports.ProductReport
}

type productReportsService struct {
	productReportsRepository productReportsRepo.ProductReportsRepository
}

func NewProductReportsService(productReportsRepository productReportsRepo.ProductReportsRepository) ProductReportsService {
	return &productReportsService{
		productReportsRepository: productReportsRepository,
	}
}

func (s *productReportsService) GenerateReport() []product_reports.ProductReport {
	return s.productReportsRepository.GetAll()
}
