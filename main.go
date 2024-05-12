package main

import (
	"expense-service/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routes.InitRoutes(router)
	router.Run(":9000")
}
