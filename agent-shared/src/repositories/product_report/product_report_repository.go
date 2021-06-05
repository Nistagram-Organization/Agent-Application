package product_report

import "github.com/Nistagram-Organization/Agent-Application/agent-shared/src/model/product_report"

type ProductReportRepository interface {
	GetAll() []product_report.ProductReport
}
