package main

import (
	"github.com/gin-gonic/gin"
	"wedding-invitation-api/user"
)

func main() {
	r := gin.Default()
	user.Controller(r)
	err := r.Run() // 默认8080端口 r.Run(":3000")
	if err != nil {
		return
	}
}
