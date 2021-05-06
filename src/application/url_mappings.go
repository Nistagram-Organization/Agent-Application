package application

import "github.com/Nistagram-Organization/Agent-Application/src/controllers/ping"

func mapUrls() {
	router.GET("/ping", ping.Ping)
}
