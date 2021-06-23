package application

import (
	"fmt"
	controller "github.com/Nistagram-Organization/Agent-Application/agent-reports/src/controllers/product_report"
	"github.com/Nistagram-Organization/Agent-Application/agent-reports/src/datasources/postgre"
	"github.com/Nistagram-Organization/Agent-Application/agent-reports/src/repositories/product_report"
	service "github.com/Nistagram-Organization/Agent-Application/agent-reports/src/services/product_report"
	"github.com/Nistagram-Organization/agent-shared/src/model/delivery_information"
	"github.com/Nistagram-Organization/agent-shared/src/model/invoice"
	"github.com/Nistagram-Organization/agent-shared/src/model/invoice_item"
	"github.com/Nistagram-Organization/agent-shared/src/model/product"
	"github.com/Nistagram-Organization/agent-shared/src/utils/jwt_utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"os"
)

var (
	router = gin.Default()
)

func StartApplication() {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AddAllowHeaders("Authorization")
	router.Use(cors.New(corsConfig))

	database := postgre.NewPostgreSqlDatabaseClient()
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

	router.GET("/reports", jwt_utils.GetJwtMiddleware(), jwt_utils.CheckScope("create:report"), productReportController.GenerateReport)

	if port, exists := os.LookupEnv("PORT"); exists {
		router.Run(fmt.Sprintf(":%s", port))
	} else {
		router.Run(":8082")
	}
}
