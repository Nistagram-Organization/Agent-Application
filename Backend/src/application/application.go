package application

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	router.Use(cors.Default())

	mapUrls()
	initDbs()

	router.Run(":8080")
}
