package services

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"wedding-invitation-api/user/models"
)

type UserController struct {
}

type UseInfo struct {
	Name   string `form:"name" json:"name" `
	Mobile string `form:"mobile" json:"mobile"`
	Number int    `form:"number" json:"number" `
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
	if err := ctx.ShouldBind(&userInfo); err == nil { // 必须绑定同一个结构体吗
		fmt.Println(userInfo)
		models.DB.Create(&userInfo)
		ctx.JSON(http.StatusOK, gin.H{
			"code":        0,
			"success":     true,
			"description": "",
		})
	} else {

	}
	//user := models.User{
	//	Name:   "ijdsf",
	//	Number: 545,
	//	Mobile: "43545",
	//}
	//db.Create(&user)
	//ctx.JSON(http.StatusOK, gin.H{
	//	"code":        0,
	//	"success":     true,
	//	"description": "",
	//})
}
