package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/offlinebrain/go-jwt-example/auth"
	"github.com/offlinebrain/go-jwt-example/repository/inmem"
	"net/http"
)

type TokenRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func GenerateToken(ctx *gin.Context) {
	var request TokenRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	user, err := inmem.Instance.Get(request.Username)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	if err = user.CheckPassword(request.Password); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		ctx.Abort()
		return
	}

	token, err := auth.GenerateJWT(user.Email, user.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
