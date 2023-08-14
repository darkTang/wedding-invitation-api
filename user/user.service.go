package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wedding-invitation-api/user/database"
	"wedding-invitation-api/user/dto"
	"wedding-invitation-api/user/entities"
)

type Service struct {
}

var db = database.DB

func (Service) GetAllUsers(ctx *gin.Context) {
	var userList []entities.User
	db.Find(&userList)
	data := map[string][]entities.User{"list": userList}
	ctx.JSON(http.StatusOK, gin.H{
		"code":        0,
		"success":     true,
		"description": "",
		"data":        data,
	})
}

func (Service) CreateUser(ctx *gin.Context) {
	var userInfo entities.User
	var userDto dto.CreateUserDto
	if err := ctx.ShouldBind(&userDto); err == nil {
		if err = db.Where("mobile = ?", userDto.Mobile).First(&userInfo).Error; err == nil {
			db.Model(&userInfo).Where("mobile = ?", userDto.Mobile).Updates(userDto)
			//userInfo.Name = userDto.Name
			//userInfo.Number = userDto.Number
			//db.Save(&userInfo)
		} else {
			userInfo = entities.User{
				Name:   userDto.Name,
				Number: userDto.Number,
				Mobile: userDto.Mobile,
			}
			db.Create(&userInfo)
		}
		ctx.JSON(http.StatusCreated, gin.H{
			"code":        0,
			"success":     true,
			"description": "",
			"data":        nil,
		})
	} else {
		ctx.JSON(http.StatusNotFound, gin.H{
			"code":        1,
			"success":     false,
			"description": err.Error(),
			"data":        nil,
		})
	}
}
