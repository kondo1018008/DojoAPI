package model

import (
	"github.com/jinzhu/gorm"
)

type User struct{
	gorm.Model
	ID string`gorm:"primary_key"`
	Name string `gorm:"primary_key"`
	Token string
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