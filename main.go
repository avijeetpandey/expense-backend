package main

import (
	"expense-service/database"
	"expense-service/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDatabaseConnection()
	router := gin.Default()
	routes.InitRoutes(router)
	router.Run(":9000")
}
