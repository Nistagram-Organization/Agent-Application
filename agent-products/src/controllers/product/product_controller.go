package product

import (
	"github.com/Nistagram-Organization/Agent-Application/agent-products/src/services/product"
	model "github.com/Nistagram-Organization/agent-shared/src/model/product"
	"github.com/Nistagram-Organization/agent-shared/src/utils/rest_error"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ProductsController interface {
	Get(*gin.Context)
	GetAll(*gin.Context)
	Create(*gin.Context)
	Delete(*gin.Context)
	Edit(*gin.Context)
}

type productsController struct {
	productsService product.ProductService
}

func NewProductController(productsService product.ProductService) ProductsController {
	return &productsController{
		productsService: productsService,
	}
}

func getProductId(productIdParam string) (uint, rest_error.RestErr) {
	productId, err := strconv.ParseUint(productIdParam, 10, 32)
	if err != nil {
		return 0, rest_error.NewBadRequestError("Product id should be a number")
	}
	return uint(productId), nil
}

func (c *productsController) Get(ctx *gin.Context) {
	productId, idErr := getProductId(ctx.Param("id"))
	if idErr != nil {
		ctx.JSON(idErr.Status(), idErr)
		return
	}

	product, getErr := c.productsService.Get(productId)
	if getErr != nil {
		ctx.JSON(getErr.Status(), getErr)
		return
	}

	ctx.JSON(http.StatusOK, product)
}

func (c *productsController) Delete(ctx *gin.Context) {
	productId, idErr := getProductId(ctx.Param("id"))
	if idErr != nil {
		ctx.JSON(idErr.Status(), idErr)
		return
	}

	delErr := c.productsService.Delete(productId)
	if delErr != nil {
		ctx.JSON(delErr.Status(), delErr)
		return
	}

	ctx.JSON(http.StatusOK, "Product deleted successfully")
}

func (c *productsController) GetAll(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, c.productsService.GetAll())
}

func (c *productsController) Create(ctx *gin.Context) {
	var product model.Product
	if err := ctx.ShouldBindJSON(&product); err != nil {
		restErr := rest_error.NewBadRequestError("invalid json body")
		ctx.JSON(restErr.Status(), restErr)
		return
	}

	result, saveErr := c.productsService.Create(&product)
	if saveErr != nil {
		ctx.JSON(saveErr.Status(), saveErr)
		return
	}
	ctx.JSON(http.StatusCreated, result)
}

func (c *productsController) Edit(ctx *gin.Context) {
	var product model.Product
	if err := ctx.ShouldBindJSON(&product); err != nil {
		restErr := rest_error.NewBadRequestError("invalid json body")
		ctx.JSON(restErr.Status(), restErr)
		return
	}

	result, editErr := c.productsService.Edit(&product)
	if editErr != nil {
		ctx.JSON(editErr.Status(), editErr)
		return
	}
	ctx.JSON(http.StatusOK, result)
}
