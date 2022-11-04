package controller

import (
	"github.com/offlinebrain/go-jwt-example/repository/inmem"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/offlinebrain/go-jwt-example/model"
)

func RegisterUser(context *gin.Context) {
	var user model.User

	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	if err := user.HashPassword(user.Password); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	if err := inmem.Instance.Save(&user); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	context.JSON(http.StatusCreated, gin.H{"username": user.Username})
}
