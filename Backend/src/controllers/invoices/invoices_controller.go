package invoices

import (
	invoicesMod "github.com/Nistagram-Organization/Agent-Application/src/model/invoices"
	"github.com/Nistagram-Organization/Agent-Application/src/services/invoices"
	"github.com/Nistagram-Organization/Agent-Application/src/utils/rest_errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type InvoicesController interface {
	BuyProduct(*gin.Context)
}

type invoicesController struct {
	invoicesService invoices.InvoicesService
}

func NewInvoicesController(invoicesService invoices.InvoicesService) InvoicesController {
	return &invoicesController{
		invoicesService: invoicesService,
	}
}

func (i *invoicesController) BuyProduct(ctx *gin.Context) {
	var invoice invoicesMod.Invoice
	if err := ctx.ShouldBindJSON(&invoice); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body")
		ctx.JSON(restErr.Status(), restErr)
		return
	}

	if err := i.invoicesService.BuyProduct(&invoice); err != nil {
		ctx.JSON(err.Status(), err)
		return
	}

	ctx.JSON(http.StatusOK, "Product bought successfully")
}
