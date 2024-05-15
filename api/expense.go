package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// health endpoint
func (server *Server) healthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Server is up and running",
	})
}
