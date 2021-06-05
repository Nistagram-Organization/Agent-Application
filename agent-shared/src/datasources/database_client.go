package datasources

import (
	"gorm.io/gorm"
)

type DatabaseClient interface {
	Init() error
	Migrate(...interface{}) error
	GetClient() *gorm.DB
}
