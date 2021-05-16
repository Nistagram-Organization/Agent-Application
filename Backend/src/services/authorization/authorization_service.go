package authorization

import (
	"github.com/Nistagram-Organization/Agent-Application/src/models/credentials"
	"github.com/Nistagram-Organization/Agent-Application/src/utils/rest_errors"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

const (
	expirationTimeInMinutes = 30
	jwtKey                  = "jwt_key"
)

var (
	AuthorizationService authorizationServiceInterface = &authorizationService{}
)

type authorizationServiceInterface interface {
	Login(credentials.Credentials) (string, rest_errors.RestErr)
}

type authorizationService struct{}

func getExpiresIn() int64 {
	return time.Now().Add(expirationTimeInMinutes * time.Minute).Unix()
}

func getNewAccessToken(username string) (string, rest_errors.RestErr) {
	claims := &jwt.StandardClaims{
		Subject:   username,
		ExpiresAt: getExpiresIn(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	key := []byte(os.Getenv(jwtKey))
	tokenString, err := token.SignedString(key)
	if err != nil {
		err := rest_errors.NewInternalServerError("Failed to generate access token", err)
		return "", err
	}
	return tokenString, nil
}

func (as *authorizationService) Login(credentials credentials.Credentials) (string, rest_errors.RestErr) {
	if err := credentials.Validate(); err != nil {
		return "", err
	}
	return getNewAccessToken(credentials.Username)
}
