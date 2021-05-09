package ping

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	PingController pingControllerInterface = &pingController{}
)

type pingControllerInterface interface {
	Ping(*gin.Context)
}

type pingController struct{}

func (c *pingController) Ping(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}
