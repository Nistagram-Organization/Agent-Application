package application

import (
	invoice2 "github.com/Nistagram-Organization/Agent-Application/agent-invoices/src/controllers/invoice"
	"github.com/Nistagram-Organization/Agent-Application/agent-invoices/src/datasources/mysql"
	invoice4 "github.com/Nistagram-Organization/Agent-Application/agent-invoices/src/repositories/invoice"
	product2 "github.com/Nistagram-Organization/Agent-Application/agent-invoices/src/repositories/product"
	invoice3 "github.com/Nistagram-Organization/Agent-Application/agent-invoices/src/services/invoice"
	"github.com/Nistagram-Organization/agent-shared/src/model/delivery_information"
	"github.com/Nistagram-Organization/agent-shared/src/model/invoice"
	"github.com/Nistagram-Organization/agent-shared/src/model/invoice_item"
	"github.com/Nistagram-Organization/agent-shared/src/model/product"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AddAllowHeaders("Authorization")
	router.Use(cors.New(corsConfig))

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

	invoicesController := invoice2.NewInvoicesController(
		invoice3.NewInvoicesService(
			product2.NewProductRepository(database),
			invoice4.NewInvoicesRepository(database),
		),
	)

	router.POST("/invoices", invoicesController.BuyProduct)

	router.Run(":8083")
}
