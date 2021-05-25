package invoices

import (
	model "github.com/Nistagram-Organization/Agent-Application/src/models/invoices"
	"github.com/Nistagram-Organization/Agent-Application/src/services/invoices"
	"github.com/Nistagram-Organization/Agent-Application/src/utils/rest_errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	InvoicesController invoicesControllerInterface = &invoicesController{}
)

type invoicesControllerInterface interface {
	BuyProduct(*gin.Context)
}

type invoicesController struct {
}

func (i *invoicesController) BuyProduct(ctx *gin.Context) {
	var invoice model.Invoice
	if err := ctx.ShouldBindJSON(&invoice); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body")
		ctx.JSON(restErr.Status(), restErr)
		return
	}

	if err := invoices.InvoicesService.BuyProduct(&invoice); err != nil {
		ctx.JSON(err.Status(), err)
		return
	}

	ctx.JSON(http.StatusOK, "Product bought successfully")
}
