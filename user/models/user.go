package models

type User struct { // 默认表名为users
	Id     int    `json:"id"`
	Name   string `json:"name" form:"name"`
	Number int    `json:"number" form:"number"`
	Mobile string `json:"mobile" form:"mobile"`
}

func (User) TableName() string {
	return "user"
}
