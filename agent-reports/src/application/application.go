package application

import (
	controller "github.com/Nistagram-Organization/Agent-Application/agent-reports/src/controllers/product_report"
	"github.com/Nistagram-Organization/Agent-Application/agent-reports/src/datasources/mysql"
	"github.com/Nistagram-Organization/Agent-Application/agent-reports/src/repositories/product_report"
	service "github.com/Nistagram-Organization/Agent-Application/agent-reports/src/services/product_report"
	"github.com/Nistagram-Organization/Agent-Application/agent-shared/src/model/delivery_information"
	"github.com/Nistagram-Organization/Agent-Application/agent-shared/src/model/invoice"
	"github.com/Nistagram-Organization/Agent-Application/agent-shared/src/model/invoice_item"
	"github.com/Nistagram-Organization/Agent-Application/agent-shared/src/model/product"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	router.Use(cors.Default())

	database := mysql.NewMySqlDatabaseClient()
	if err := database.Init(); err != nil {
		panic(err)
	}

	if err := database.Migrate(
		&product.Product{},
		&invoice.Invoice{},
		&invoice_item.InvoiceItem{},
		&delivery_information.DeliveryInformation{},
	); err != nil {
		panic(err)
	}

	productReportController := controller.NewProductReportController(
		service.NewProductReportService(product_report.NewProductReportRepository(database)),
	)

	router.GET("/reports", productReportController.GenerateReport)

	router.Run(":8082")
}
