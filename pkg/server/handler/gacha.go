package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	db2 "github.com/kondo1018008/DojoAPI/pkg/db"
	"github.com/kondo1018008/DojoAPI/pkg/server/model"
)

func HandleGachaDraw() gin.HandlerFunc{
	return func(c *gin.Context){
		user := model.User{}
		gacha := model.Gacha{}
		charas := model.UserCharacter{}
		db := db2.GetDB()
		token := c.Request.Header.Get("x-token")
		fmt.Println("token="+ token)
		db.Where("Token = ?", token).First(&user)

	}
}
