package product_reports

import (
	"github.com/Nistagram-Organization/Agent-Application/src/services/product_reports"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ProductReportsController interface {
	GenerateReport(ctx *gin.Context)
}

type productReportsController struct {
	productReportsService product_reports.ProductReportsService
}

func NewProductReportsController(productReportsService product_reports.ProductReportsService) ProductReportsController {
	return &productReportsController{
		productReportsService: productReportsService,
	}
}

func (i *productReportsController) GenerateReport(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, i.productReportsService.GenerateReport())
}
