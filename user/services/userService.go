package services

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
}

func (*UserController) GetAllUsers(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":        0,
		"description": "",
		"success":     true,
		"data":        "hello",
	})
}
