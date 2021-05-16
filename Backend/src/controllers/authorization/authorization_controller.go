package authorization

import (
	"github.com/Nistagram-Organization/Agent-Application/src/models/credentials"
	"github.com/Nistagram-Organization/Agent-Application/src/services/authorization"
	"github.com/Nistagram-Organization/Agent-Application/src/utils/rest_errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	AuthorizationController authorizationControllerInterface = &authorizationController{}
)

type authorizationControllerInterface interface {
	Login(*gin.Context)
}

type authorizationController struct {
}

func (ac *authorizationController) Login(ctx *gin.Context) {
	var credentials credentials.Credentials
	if err := ctx.ShouldBindJSON(&credentials); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body")
		ctx.JSON(restErr.Status(), restErr)
		return
	}

	token, err := authorization.AuthorizationService.Login(credentials)
	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}
	ctx.JSON(http.StatusCreated, token)
}
