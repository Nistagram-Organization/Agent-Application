package mysql

import (
	"fmt"
	"github.com/Nistagram-Organization/agent-shared/src/datasources"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

const (
	mysqlUsername = "mysql_username"
	mysqlPassword = "mysql_password"
	mysqlHost     = "mysql_host"
	mysqlSchema   = "mysql_schema"
)

type mysqlClient struct {
	Client *gorm.DB
}

func NewMySqlDatabaseClient() datasources.DatabaseClient {
	return &mysqlClient{}
}

func (c *mysqlClient) Init() error {
	var dataSourceName string
	var exists bool

	if dataSourceName, exists = os.LookupEnv("JAWSDB_URL"); !exists {
		username := os.Getenv(mysqlUsername)
		password := os.Getenv(mysqlPassword)
		host := os.Getenv(mysqlHost)
		schema := os.Getenv(mysqlSchema)

		dataSourceName = fmt.Sprintf(
			"%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			username,
			password,
			host,
			schema,
		)
	}
	client, err := gorm.Open(mysql.Open(dataSourceName), &gorm.Config{})

	if err != nil {
		return err
	}

	c.Client = client
	return nil
}

func (c *mysqlClient) Migrate(interfaces ...interface{}) error {
	return c.Client.AutoMigrate(interfaces...)
}

func (c *mysqlClient) GetClient() *gorm.DB {
	return c.Client
}
