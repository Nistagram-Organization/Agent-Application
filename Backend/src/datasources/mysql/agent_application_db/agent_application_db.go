package agent_application_db

import (
	"fmt"
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

var (
	Client mysqlClientInterface = &mysqlClient{}
)

type mysqlClientInterface interface {
	Init() error
	Migrate(...interface{}) error
	Get(interface{}, uint) *gorm.DB
}

type mysqlClient struct {
	Client *gorm.DB
}

func (c *mysqlClient) Init() error {
	username := os.Getenv(mysqlUsername)
	password := os.Getenv(mysqlPassword)
	host := os.Getenv(mysqlHost)
	schema := os.Getenv(mysqlSchema)

	dataSourceName := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username,
		password,
		host,
		schema,
	)

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

func (c *mysqlClient) Get(entity interface{}, pk uint) *gorm.DB {
	return c.Client.First(&entity, pk)
}
