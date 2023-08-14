package dto

type CreateUserDto struct {
	Name   string `json:"name" form:"name" binding:"required"`
	Number uint   `json:"number" form:"number" binding:"required"`
	Mobile string `json:"mobile" form:"mobile" binding:"required"`
}
