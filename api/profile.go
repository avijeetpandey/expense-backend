package api

import (
	"net/http"

	db "github.com/avijeetpandey/expense_tracker/db/sqlc"
	"github.com/gin-gonic/gin"
)

type CreateProfileRequest struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
	Bio   string `json:"bio" binding:"required"`
}

// function to create a profile of the user
func (server *Server) createProfile(ctx *gin.Context) {
	var req CreateProfileRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateProfileParams{
		Name:  req.Name,
		Email: req.Email,
		Bio:   req.Bio,
	}

	profile, err := server.store.CreateProfile(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, profile)
}

type GetProfileRequest struct {
	ID int64 `uri:"id" json:"id" binding:"required,min=1"`
}

// function to fetch a profile
func (server *Server) getProfile(ctx *gin.Context) {
	var req GetProfileRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	profile, err := server.store.GetProfile(ctx, req.ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, profile)
}

// function to delete profile
func (server *Server) deleteProfile(ctx *gin.Context) {
	var req GetProfileRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.store.DeleteProfile(ctx, req.ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "profile deleted",
	})
}
