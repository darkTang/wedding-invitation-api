package controllers

import (
	"github.com/gin-gonic/gin"
	"wedding-invitation-api/user/services"
)

func UserController(r *gin.Engine) {
	userRouter := r.Group("/api")
	user := &services.UserController{}
	{
		userRouter.GET("/queryUserInfo", user.GetAllUsers)
		userRouter.POST("/collectUserInfo", user.AddUser)
	}
}
