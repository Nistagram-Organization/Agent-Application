package postgre

import (
	"github.com/Nistagram-Organization/agent-shared/src/datasources"
	"github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

const (
	dsn    = "DATABASE_URL"
	heroku = "HEROKU"
)

type postgreSQLClient struct {
	Client *gorm.DB
}

func NewPostgreSqlDatabaseClient() datasources.DatabaseClient {
	return &postgreSQLClient{}
}

func (c *postgreSQLClient) Init() error {
	dataSourceName := os.Getenv(dsn)

	if os.Getenv(heroku) == "true" {
		dataSourceName, _ = pq.ParseURL(dataSourceName)
		dataSourceName = dataSourceName + " sslmode=require"
	}
	client, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dataSourceName,
	}), &gorm.Config{})

	if err != nil {
		return err
	}

	c.Client = client
	return nil
}

func (c *postgreSQLClient) Migrate(interfaces ...interface{}) error {
	return c.Client.AutoMigrate(interfaces...)
}

func (c *postgreSQLClient) GetClient() *gorm.DB {
	return c.Client
}
