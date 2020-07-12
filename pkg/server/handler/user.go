package handler

import (
	"encoding/json"
	//"encoding/json"
	"github.com/gin-gonic/gin"
	db2 "github.com/kondo1018008/DojoAPI/pkg/db"
	"github.com/kondo1018008/DojoAPI/pkg/server/model"
	"log"
	"net/http"
	"time"
	"github.com/google/uuid"
)

/*func HandleUserGet() gin.HandlerFunc{
	return func(c gin.Context){

	}
}*/

func HandleUserCreate() gin.HandlerFunc{
	return func(c *gin.Context){
		user := model.User{}
		now := time.Now()
		db := db2.GetDB()


		var requestBody authCreateRequest
		json.NewDecoder(c.Request.Body).Decode(&requestBody)
		/*if err := c.BindJSON(&user); err != nil {
			c.String(http.StatusBadRequest, "Request is failed: "+ err.Error())
		}*/

		userID, err := uuid.NewRandom()
		if err != nil {
			log.Println(err)
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		authToken, err := uuid.NewRandom()
		if err != nil {
			log.Println(err)
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		user = model.User{ID:userID.String(), Name:requestBody.Name, Token:authToken.String()}
		user.CreatedAt = now
		user.UpdatedAt = now

		db.NewRecord(user)
		db.Create(&user)
		if db.NewRecord(user) == false {
			c.JSON(http.StatusOK, user)
		}
	}
}
/*
func HandleUserUpdate() gin.HandlerFunc{
	return func(c *gin.Context){

	}
}*/


type authCreateRequest struct {
	Name string `json:"name"`
}

/*type authCreateResponse struct {
	Token string `json:"token"`
}*/