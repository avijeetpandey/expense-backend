package api

import (
	"net/http"

	db "github.com/avijeetpandey/expense_tracker/db/sqlc"
	"github.com/gin-gonic/gin"
)

type CreateExpenseRequest struct {
	Name     string `json:"name" binding:"required"`
	Currency string `json:"currency" binding:"required,oneof=USD INR"`
	Balance  int64  `json:"balance" binding:"required,min=1"`
	Tag      string `json:"tag"`
}

// Add Expense to the database
func (server *Server) addExpense(ctx *gin.Context) {
	var req CreateExpenseRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateExpenseParams{
		Name:     req.Name,
		Currency: req.Currency,
		Balance:  req.Balance,
		Tag:      req.Tag,
	}

	expense, err := server.store.CreateExpense(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, expense)
}

type DeleteExpenseRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

// Delete Expense
func (server *Server) deleteExpense(ctx *gin.Context) {
	var req DeleteExpenseRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.store.DeleteExpense(ctx, req.ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Succesfully deleted",
	})
}

type GetExpenseRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

// Get expense
func (server *Server) getExpense(ctx *gin.Context) {
	var req GetExpenseRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	expense, err := server.store.GetExpense(ctx, req.ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, expense)
}

// health endpoint
func (server *Server) healthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Server is up and running",
	})
}
