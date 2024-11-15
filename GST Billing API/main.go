package main

import (
	config "GST_billing_api/Config"
	routes "GST_billing_api/Routes"

	"github.com/gin-gonic/gin"
)

func main() {

	config.Connect_DB()

	router := gin.Default()

	routes.ConnectRouter(router)

	router.Run(":8080")
}
