package product_report

import (
	"github.com/Nistagram-Organization/agent-shared/src/datasources"
	model "github.com/Nistagram-Organization/agent-shared/src/model/product_report"
	"github.com/Nistagram-Organization/agent-shared/src/repositories/product_report"
	"gorm.io/gorm"
)

type productReportRepository struct {
	db *gorm.DB
}

func NewProductReportRepository(databaseClient datasources.DatabaseClient) product_report.ProductReportRepository {
	return &productReportRepository{
		databaseClient.GetClient(),
	}
}

func (p *productReportRepository) GetAll() []model.ProductReport {
	var productReports []model.ProductReport
	if err := p.db.Raw("SELECT A.`name`, SUM(B.`quantity`) as `sold`, SUM(A.`price` * B.`quantity`) as `income` FROM `products` AS A LEFT JOIN `invoice_items` AS B ON B.`product_id` = A.`id` GROUP BY A.`id`").Scan(&productReports).Error; err != nil {
		return []model.ProductReport{}
	}
	return productReports
}
