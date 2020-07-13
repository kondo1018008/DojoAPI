package handler

import (
	"fmt"
	//"encoding/json"
	"github.com/gin-gonic/gin"
	db2 "github.com/kondo1018008/DojoAPI/pkg/db"
	"github.com/kondo1018008/DojoAPI/pkg/server/model"
	"log"
	"net/http"
	_ "time"
	"github.com/google/uuid"
)


func HandleUserGet() gin.HandlerFunc{
	return func(c *gin.Context){

		user := model.User{}
		db := db2.GetDB()
		token := c.Request.Header.Get("x-token")
		fmt.Println("token="+ token)
		db.Where("Token = ?", token).First(&user)
		fmt.Println(user)
		c.JSON(http.StatusOK, user)
	}
}

func HandleUserCreate() gin.HandlerFunc{
	return func(c *gin.Context){
		user := model.User{}
		//now := time.Now()
		db := db2.GetDB()


		/*var requestBody authCreateRequest
		json.NewDecoder(c.Request.Body).Decode(&requestBody)*/


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

		user.ID = userID.String()
		user.Token = authToken.String()
		//user.CreatedAt = now
		//user.UpdatedAt = now

		if err := c.BindJSON(&user); err != nil {
			c.String(http.StatusBadRequest, "Request is failed: "+ err.Error())
		}

		fmt.Println(db.NewRecord(user))
		db.Create(&user)
		if db.NewRecord(user) == false{
			fmt.Println(db.NewRecord(user))
			c.JSON(201, gin.H{
				"token":authToken.String(),
			})
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

type authCreateResponse struct {
	Token string `json:"token"`
}