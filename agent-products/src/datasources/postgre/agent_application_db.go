package postgre

import (
	"fmt"
	"github.com/Nistagram-Organization/agent-shared/src/datasources"
	"github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

const (
	dsn = "dsn"
)

type postgreSQLClient struct {
	Client *gorm.DB
}

func NewPostgreSqlDatabaseClient() datasources.DatabaseClient {
	return &postgreSQLClient{}
}

func (c *postgreSQLClient) Init() error {
	dataSourceName := os.Getenv(dsn)
	dataSourceName, _ = pq.ParseURL(dataSourceName)
	dataSourceName += " sslmode=require"

	fmt.Println(dataSourceName)
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
