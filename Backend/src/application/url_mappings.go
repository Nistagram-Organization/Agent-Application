package application

import (
	"github.com/Nistagram-Organization/Agent-Application/src/controllers/authorization"
	"github.com/Nistagram-Organization/Agent-Application/src/controllers/invoices"
	"github.com/Nistagram-Organization/Agent-Application/src/controllers/ping"
	"github.com/Nistagram-Organization/Agent-Application/src/controllers/products"
)

func mapUrls() {
	router.GET("/ping", ping.PingController.Ping)

	router.GET("/products", products.ProductsController.GetAll)
	router.GET("/products/:id", products.ProductsController.Get)

	router.POST("/invoices", invoices.InvoicesController.BuyProduct)

	router.POST("/login", authorization.AuthorizationController.Login)
}
