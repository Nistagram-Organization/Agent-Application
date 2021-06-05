module github.com/Nistagram-Organization/Agent-Application/agent-invoices

go 1.15

replace github.com/Nistagram-Organization/Agent-Application/agent-shared => ../agent-shared

require (
	github.com/Nistagram-Organization/Agent-Application/agent-shared v0.0.0-00010101000000-000000000000
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-gonic/gin v1.7.1
	github.com/stretchr/testify v1.7.0
	gorm.io/driver/mysql v1.1.0
	gorm.io/gorm v1.21.10
)