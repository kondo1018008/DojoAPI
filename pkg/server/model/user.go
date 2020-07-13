package model

type User struct{
	ID string
	Name string
	Token string `gorm:"primary_key"`
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