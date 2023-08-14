package main

import (
	"github.com/gin-gonic/gin"
	"wedding-invitation-api/user"
)

func main() {
	r := gin.Default()
	user.Controller(r)
	err := r.Run()
	if err != nil {
		return
	}
}
