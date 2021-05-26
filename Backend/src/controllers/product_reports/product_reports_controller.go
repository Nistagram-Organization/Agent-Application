package product_reports

import (
	"github.com/Nistagram-Organization/Agent-Application/src/services/product_reports"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	ProductReportsController productReportsControllerInterface = &productReportsController{}
)

type productReportsControllerInterface interface {
	GenerateReport(ctx *gin.Context)
}

type productReportsController struct {
}

func (i *productReportsController) GenerateReport(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, product_reports.ProductReportsService.GenerateReport())
}
