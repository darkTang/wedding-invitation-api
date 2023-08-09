package services

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wedding-invitation-api/user/models"
)

type UserController struct {
}

var db = models.DB

func (*UserController) GetAllUsers(ctx *gin.Context) {
	var userList []models.User
	db.Find(&userList)
	data := map[string][]models.User{"list": userList}
	ctx.JSON(http.StatusOK, gin.H{
		"code":        0,
		"success":     true,
		"description": "",
		"data":        data,
	})
}

func (*UserController) AddUser(ctx *gin.Context) {
	var userInfo models.User
	if err := ctx.ShouldBind(&userInfo); err == nil { // 必须绑定model.User才会生效
		models.DB.Create(&userInfo)
		ctx.JSON(http.StatusOK, gin.H{
			"code":        0,
			"success":     true,
			"description": "",
			"data": nil,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code":        1,
			"success":     false,
			"description": err.Error(),
			"data": nil,
		})
	}
}
