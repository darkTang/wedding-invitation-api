package main

import (
	"github.com/gin-gonic/gin"
	"wedding-invitation-api/user/controllers"
)

func main() {
	r := gin.Default()
	controllers.UserController(r)
	r.Run()
}
