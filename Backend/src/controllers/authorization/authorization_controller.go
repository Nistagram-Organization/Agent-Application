package authorization

import (
	"github.com/Nistagram-Organization/Agent-Application/src/model/credentials"
	"github.com/Nistagram-Organization/Agent-Application/src/services/authorization"
	"github.com/Nistagram-Organization/Agent-Application/src/utils/rest_errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthorizationController interface {
	Login(*gin.Context)
}

type authorizationController struct {
	authorizationService authorization.AuthorizationService
}

func NewAuthorizationController(authorizationService authorization.AuthorizationService) AuthorizationController {
	return &authorizationController{
		authorizationService: authorizationService,
	}
}

func (ac *authorizationController) Login(ctx *gin.Context) {
	var credentials credentials.Credentials
	if err := ctx.ShouldBindJSON(&credentials); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body")
		ctx.JSON(restErr.Status(), restErr)
		return
	}

	token, err := ac.authorizationService.Login(credentials)
	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}
	ctx.JSON(http.StatusCreated, token)
}
