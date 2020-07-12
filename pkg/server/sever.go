package server

import (
	"github.com/gin-gonic/gin"
	"github.com/kondo1018008/DojoAPI/pkg/server/handler"
)

func Init() {
	r := router()
	r.Run()
}

func router() *gin.Engine {
	r := gin.Default()

	u := r.Group("/user")
	{
		u.POST("/create", handler.HandleUserCreate())
		//u.GET("/get", )
		//u.PUT("/update", )
	}
	return r

}
