package product_reports

import "github.com/Nistagram-Organization/Agent-Application/src/models/product_reports"

var (
	ProductReportsService productReportsServiceInterface = &productReportsService{}
)

type productReportsServiceInterface interface {
	GenerateReport() []product_reports.ProductReport
}

type productReportsService struct{}

func (s *productReportsService) GenerateReport() []product_reports.ProductReport {
	dao := product_reports.ProductReport{}
	return dao.GetAll()
}
