package user

import (
	"github.com/gin-gonic/gin"
)

func Controller(r *gin.Engine) {
	userRouter := r.Group("/api")
	var user Service
	{
		userRouter.GET("/queryUserInfo", user.GetAllUsers)
		userRouter.POST("/collectUserInfo", user.CreateUser)
	}
}
