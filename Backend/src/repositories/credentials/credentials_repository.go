package credentials

import (
	"fmt"
	"github.com/Nistagram-Organization/Agent-Application/src/datasources/mysql/agent_application_db"
	"github.com/Nistagram-Organization/Agent-Application/src/model/credentials"
	"github.com/Nistagram-Organization/Agent-Application/src/utils/rest_errors"
	"gorm.io/gorm"
)

type CredentialsRepository interface {
	GetByUsername(string) (*credentials.Credentials, rest_errors.RestErr)
}

type credentialsRepository struct {
	db *gorm.DB
}

func NewCredentialsRepository() CredentialsRepository {
	return &credentialsRepository{
		agent_application_db.Client.GetClient(),
	}
}

func (c *credentialsRepository) GetByUsername(username string) (*credentials.Credentials, rest_errors.RestErr) {
	var credentials credentials.Credentials
	if err := c.db.Where("username = ?", username).Take(&credentials).Error; err != nil {
		return nil, rest_errors.NewNotFoundError(fmt.Sprintf("Error when trying to get credentials with username %s", username))
	}
	return &credentials, nil
}
