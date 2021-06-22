package postgre

import (
	"fmt"
	"github.com/Nistagram-Organization/agent-shared/src/datasources"
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
	var dataSourceName string
	var exists bool

	if dataSourceName, exists = os.LookupEnv("DATABASE_URL"); !exists {
		dataSourceName = os.Getenv(dsn)
	}

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
