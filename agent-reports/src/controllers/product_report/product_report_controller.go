package product_report

import (
	"github.com/Nistagram-Organization/Agent-Application/agent-reports/src/services/product_report"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ProductReportsController interface {
	GenerateReport(ctx *gin.Context)
}

type productReportsController struct {
	productReportsService product_report.ProductReportsService
}

func NewProductReportController(productReportsService product_report.ProductReportsService) ProductReportsController {
	return &productReportsController{
		productReportsService: productReportsService,
	}
}

func (i *productReportsController) GenerateReport(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, i.productReportsService.GenerateReport())
}
