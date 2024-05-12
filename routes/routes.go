package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRoutes(g *gin.Engine) {

	g.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "Server is up and running")
	})

}
