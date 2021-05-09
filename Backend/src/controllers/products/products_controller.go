package products

import (
	"github.com/Nistagram-Organization/Agent-Application/src/services/products"
	"github.com/Nistagram-Organization/Agent-Application/src/utils/rest_errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var (
	ProductsController productsControllerInterface = &productsController{}
)

type productsControllerInterface interface {
	Get(*gin.Context)
}

type productsController struct{}

func getProductId(productIdParam string) (uint, rest_errors.RestErr) {
	productId, err := strconv.ParseUint(productIdParam, 10, 32)
	if err != nil {
		return 0, rest_errors.NewBadRequestError("Product id should be a number")
	}
	return uint(productId), nil
}

func (c *productsController) Get(ctx *gin.Context) {
	productId, idErr := getProductId(ctx.Param("id"))
	if idErr != nil {
		ctx.JSON(idErr.Status(), idErr)
		return
	}

	product, getErr := products.ProductsService.Get(productId)
	if getErr != nil {
		ctx.JSON(getErr.Status(), getErr)
		return
	}

	ctx.JSON(http.StatusOK, product)
}
