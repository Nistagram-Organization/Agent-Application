package main

import (
	"github.com/Nistagram-Organization/Agent-Application/src/application"
	_ "github.com/Nistagram-Organization/Agent-Application/src/datasources/mysql/agent_application_db"
)

func main() {
	application.StartApplication()
}
