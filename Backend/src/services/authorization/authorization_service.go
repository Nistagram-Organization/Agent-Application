package authorization

import (
	"github.com/Nistagram-Organization/Agent-Application/src/model/credentials"
	credentialsRepo "github.com/Nistagram-Organization/Agent-Application/src/repositories/credentials"
	"github.com/Nistagram-Organization/Agent-Application/src/utils/bcrypt_utils"
	"github.com/Nistagram-Organization/Agent-Application/src/utils/rest_errors"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

const (
	expirationTimeInMinutes = 30
	jwtKey                  = "jwt_key"
)

type AuthorizationService interface {
	Login(credentials.Credentials) (string, rest_errors.RestErr)
	Validate(*credentials.Credentials) rest_errors.RestErr
}

type authorizationService struct {
	credentialsRepository credentialsRepo.CredentialsRepository
}

func NewAuthorizationService(credentialsRepository credentialsRepo.CredentialsRepository) AuthorizationService {
	return &authorizationService{
		credentialsRepository: credentialsRepository,
	}
}

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
	if err := as.Validate(&credentials); err != nil {
		return "", err
	}
	return getNewAccessToken(credentials.Username)
}

func (as *authorizationService) Validate(credentials *credentials.Credentials) rest_errors.RestErr {
	plainPassword := credentials.Password
	dbCredentials, err := as.credentialsRepository.GetByUsername(credentials.Username)
	if err != nil {
		return err
	}

	if !bcrypt_utils.CompareHashAndValue(dbCredentials.Password, plainPassword) {
		return rest_errors.NewUnauthorizedError("invalid password")
	}
	return nil
}
