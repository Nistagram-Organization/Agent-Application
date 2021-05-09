package application

import (
	"github.com/Nistagram-Organization/Agent-Application/src/controllers/ping"
	"github.com/Nistagram-Organization/Agent-Application/src/controllers/products"
)

func mapUrls() {
	router.GET("/ping", ping.PingController.Ping)

	router.GET("/products/:id", products.ProductsController.Get)
}
