package agent_application_db

import (
	"fmt"
	"github.com/Nistagram-Organization/Agent-Application/src/models/delivery_informations"
	"github.com/Nistagram-Organization/Agent-Application/src/models/invoice_items"
	"github.com/Nistagram-Organization/Agent-Application/src/models/invoices"
	"github.com/Nistagram-Organization/Agent-Application/src/models/products"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

const (
	mysqlUsername = "mysql_username"
	mysqlPassword = "mysql_password"
	mysqlHost     = "mysql_host"
	mysqlSchema   = "mysql_schema"
)

var (
	Client   *gorm.DB
	username = os.Getenv(mysqlUsername)
	password = os.Getenv(mysqlPassword)
	host     = os.Getenv(mysqlHost)
	schema   = os.Getenv(mysqlSchema)
)

func connectToDB() {
	dataSourceName := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username,
		password,
		host,
		schema,
	)

	var err error
	Client, err = gorm.Open(mysql.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	log.Println("connected to database")
}

func migrateDB() {
	err := Client.AutoMigrate(
		&products.Product{},
		&invoices.Invoice{},
		&invoice_items.InvoiceItem{},
		&delivery_informations.DeliveryInformation{},
	)

	if err != nil {
		panic(err)
	}

	log.Println("migration successful")
}

func init() {
	connectToDB()
	migrateDB()
}
