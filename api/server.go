package api

import (
	db "github.com/avijeetpandey/expense_tracker/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	// CRUD Expense
	router.POST("/add/expense", server.addExpense)

	// HEALTHCHECK
	router.GET("/health", server.healthCheck)

	server.router = router

	return server
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
