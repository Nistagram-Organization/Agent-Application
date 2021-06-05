package product_report

import (
	"github.com/Nistagram-Organization/Agent-Application/agent-shared/src/model/product_report"
	repo "github.com/Nistagram-Organization/Agent-Application/agent-shared/src/repositories/product_report"
)

type ProductReportsService interface {
	GenerateReport() []product_report.ProductReport
}

type productReportsService struct {
	productReportsRepository repo.ProductReportRepository
}

func NewProductReportService(productReportsRepository repo.ProductReportRepository) ProductReportsService {
	return &productReportsService{
		productReportsRepository: productReportsRepository,
	}
}

func (s *productReportsService) GenerateReport() []product_report.ProductReport {
	return s.productReportsRepository.GetAll()
}
