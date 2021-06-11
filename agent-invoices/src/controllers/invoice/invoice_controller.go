package invoice

import (
	"github.com/Nistagram-Organization/Agent-Application/agent-invoices/src/services/invoice"
	model "github.com/Nistagram-Organization/agent-shared/src/model/invoice"
	"github.com/Nistagram-Organization/agent-shared/src/utils/rest_error"
	"github.com/gin-gonic/gin"
	"net/http"
)

type InvoicesController interface {
	BuyProduct(*gin.Context)
}

type invoicesController struct {
	invoicesService invoice.InvoicesService
}

func NewInvoicesController(invoicesService invoice.InvoicesService) InvoicesController {
	return &invoicesController{
		invoicesService: invoicesService,
	}
}

func (i *invoicesController) BuyProduct(ctx *gin.Context) {
	var invoice model.Invoice
	if err := ctx.ShouldBindJSON(&invoice); err != nil {
		restErr := rest_error.NewBadRequestError("invalid json body")
		ctx.JSON(restErr.Status(), restErr)
		return
	}

	if err := i.invoicesService.BuyProduct(&invoice); err != nil {
		ctx.JSON(err.Status(), err)
		return
	}

	ctx.JSON(http.StatusOK, "Product bought successfully")
}
