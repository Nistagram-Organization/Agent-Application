package application

import (
	"github.com/Nistagram-Organization/Agent-Application/src/controllers/authorization"
	"github.com/Nistagram-Organization/Agent-Application/src/controllers/invoices"
	"github.com/Nistagram-Organization/Agent-Application/src/controllers/product_reports"
	"github.com/Nistagram-Organization/Agent-Application/src/controllers/products"
	"github.com/Nistagram-Organization/Agent-Application/src/repositories/credentials"
	"github.com/Nistagram-Organization/Agent-Application/src/repositories/invoice_items"
	invoicesRepo "github.com/Nistagram-Organization/Agent-Application/src/repositories/invoices"
	productReportsRepo "github.com/Nistagram-Organization/Agent-Application/src/repositories/product_reports"
	productsRepo "github.com/Nistagram-Organization/Agent-Application/src/repositories/products"
	authorizationServ "github.com/Nistagram-Organization/Agent-Application/src/services/authorization"
	invoicesServ "github.com/Nistagram-Organization/Agent-Application/src/services/invoices"
	productReportsServ "github.com/Nistagram-Organization/Agent-Application/src/services/product_reports"
	productsServ "github.com/Nistagram-Organization/Agent-Application/src/services/products"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	initDbs()
	router.Use(cors.Default())

	productsRepository := productsRepo.NewProductsRepository()
	invoiceItemsRepository := invoice_items.NewInvoiceItemsRepository()
	productsService := productsServ.NewProductsService(productsRepository, invoiceItemsRepository, "temp")
	productsController := products.NewProductsController(productsService)

	productReportsController := product_reports.NewProductReportsController(
		productReportsServ.NewProductReportsService(productReportsRepo.NewProductReportsRepository()),
	)

	invoicesController := invoices.NewInvoicesController(
		invoicesServ.NewInvoicesService(productsRepository, invoicesRepo.NewInvoicesRepository()),
	)

	authorizationController := authorization.NewAuthorizationController(
		authorizationServ.NewAuthorizationService(credentials.NewCredentialsRepository()),
	)

	router.GET("/products", productsController.GetAll)
	router.GET("/products/:id", productsController.Get)
	router.POST("/products", productsController.Create)
	router.PUT("/products", productsController.Edit)
	router.DELETE("/products/:id", productsController.Delete)

	router.GET("/reports", productReportsController.GenerateReport)

	router.POST("/invoices", invoicesController.BuyProduct)

	router.POST("/login", authorizationController.Login)

	router.Run(":8080")
}
