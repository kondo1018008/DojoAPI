package server

import (
	"github.com/gin-gonic/gin"
)


func server(){
	r := gin.Default()

	r.POST("/user/create", )
	r.GET("/user/get", )
	r.PUT("/user/update", )

	r.Run()
}