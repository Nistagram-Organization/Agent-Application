package application

import (
	"github.com/Nistagram-Organization/Agent-Application/src/datasources/mysql/agent_application_db"
	"github.com/Nistagram-Organization/Agent-Application/src/models/credentials"
	"github.com/Nistagram-Organization/Agent-Application/src/models/delivery_informations"
	"github.com/Nistagram-Organization/Agent-Application/src/models/invoice_items"
	"github.com/Nistagram-Organization/Agent-Application/src/models/invoices"
	"github.com/Nistagram-Organization/Agent-Application/src/models/products"
	"log"
)

func initDbs() {
	var err error
	err = agent_application_db.Client.Init()
	if err != nil {
		panic(err)
	}

	log.Println("Successfully connected to database")

	err = agent_application_db.Client.Migrate(
		&products.Product{},
		&invoices.Invoice{},
		&invoice_items.InvoiceItem{},
		&delivery_informations.DeliveryInformation{},
		&credentials.Credentials{},
	)

	if err != nil {
		panic(err)
	}

	log.Println("Database migration successful")
}
