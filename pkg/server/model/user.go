package model

import (
	"github.com/jinzhu/gorm"
)

type User struct{
	gorm.Model
	Name string `json:"name"`
	Token string `json:"token"`
}
/*


type UserGetResponse struct {
	Name string `json:"name"`
}

func HandleUserCreate() gin.HandlerFunc{

}

func HandleUserUpdate() gin.HandlerFunc{

}

 */