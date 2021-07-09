package application

import (
	"fmt"
	controller "github.com/Nistagram-Organization/Agent-Application/agent-products/src/controllers/product"
	"github.com/Nistagram-Organization/Agent-Application/agent-products/src/datasources/postgre"
	invoice_item2 "github.com/Nistagram-Organization/Agent-Application/agent-products/src/repositories/invoice_item"
	product3 "github.com/Nistagram-Organization/Agent-Application/agent-products/src/repositories/product"
	product2 "github.com/Nistagram-Organization/Agent-Application/agent-products/src/services/product"
	"github.com/Nistagram-Organization/Agent-Application/agent-products/src/utils/image_utils"
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

	productController := controller.NewProductController(
		product2.NewProductService(
			product3.NewProductRepository(database),
			invoice_item2.NewInvoiceItemRepository(database),
			image_utils.NewImageUtilsService(),
			"temp",
		),
	)

	router.GET("/products", productController.GetAll)
	router.GET("/products/:id", productController.Get)
	router.POST("/products", jwt_utils.GetJwtMiddleware(), jwt_utils.CheckScope("create:product"), productController.Create)
	router.PUT("/products", jwt_utils.GetJwtMiddleware(), jwt_utils.CheckScope("edit:product"), productController.Edit)
	router.DELETE("/products/:id", jwt_utils.GetJwtMiddleware(), jwt_utils.CheckScope("delete:product"), productController.Delete)

	if port, exists := os.LookupEnv("PORT"); exists {
		router.Run(fmt.Sprintf(":%s", port))
	} else {
		router.Run(":8081")
	}
}
