package product_reports

import (
	"github.com/Nistagram-Organization/Agent-Application/src/datasources/mysql/agent_application_db"
	"github.com/Nistagram-Organization/Agent-Application/src/model/product_reports"
	"gorm.io/gorm"
)

type ProductReportsRepository interface {
	GetAll() []product_reports.ProductReport
}

type productReportsRepository struct {
	db *gorm.DB
}

func NewProductReportsRepository() ProductReportsRepository {
	return &productReportsRepository{
		agent_application_db.Client.GetClient(),
	}
}

func (p *productReportsRepository) GetAll() []product_reports.ProductReport {
	var productReports []product_reports.ProductReport
	if err := p.db.Raw("SELECT A.`name`, SUM(B.`quantity`) as `sold`, SUM(A.`price` * B.`quantity`) as `income` FROM `products` AS A LEFT JOIN `invoice_items` AS B ON B.`product_id` = A.`id` GROUP BY A.`id`").Scan(&productReports).Error; err != nil {
		return []product_reports.ProductReport{}
	}
	return productReports
}
