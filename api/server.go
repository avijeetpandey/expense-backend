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
	router.POST("/expense", server.addExpense)
	router.DELETE("/expense/:id", server.deleteExpense)
	router.GET("/expense/:id", server.getExpense)

	// CRUD Profile
	router.POST("/profile", server.createProfile)
	router.GET("/profile/:id", server.getProfile)
	router.DELETE("/profile/:id", server.deleteProfile)

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
