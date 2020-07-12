package handler

import (
	"github.com/gin-gonic/gin"
	db2 "github.com/kondo1018008/DojoAPI/pkg/db"
	"github.com/kondo1018008/DojoAPI/pkg/server/model"
	"net/http"
	"time"
)

func HandleUserGet() gin.HandlerFunc{
	return func(c gin.Context){

	}
}

func HandleUserCreate() gin.HandlerFunc{
	return func(c gin.Context){
		user := model.User{}
		now := time.Now()
		db := db2.GetDB()
		user.CreatedAt = now
		user.UpdatedAt = now

		err := c.BindJSON(&user)
		if err != nil {
			c.String(http.StatusBadRequest, "Request is failed: "+err.Error())
		}
		db.NewRecord(user)
		db.Create(&user)
		if db.NewRecord(user) == false {
			c.JSON(http.StatusOK, user)
		}
	}
}

func HandleUserUpdate() gin.HandlerFunc{
	return func(c gin.Context){

	}
}