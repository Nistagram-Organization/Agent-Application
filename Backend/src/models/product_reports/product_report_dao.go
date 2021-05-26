package product_reports

import (
	"github.com/Nistagram-Organization/Agent-Application/src/datasources/mysql/agent_application_db"
)

func (p *ProductReport) GetAll() []ProductReport {
	var productReports []ProductReport
	if err := agent_application_db.Client.GetClient().Raw("SELECT A.`name`, SUM(B.`quantity`) as `sold`, SUM(A.`price` * B.`quantity`) as `income` FROM `products` AS A LEFT JOIN `invoice_items` AS B ON B.`product_id` = A.`id` GROUP BY A.`id`").Scan(&productReports).Error; err != nil {
		return []ProductReport{}
	}
	return productReports
}
