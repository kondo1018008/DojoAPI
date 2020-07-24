package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/kondo1018008/DojoAPI/pkg/server/model"
)

var (
	db *gorm.DB
	err error
)

func Init() {
	DBMS := "mysql"
	USER := "root"
	PASS := "dojo"
	PROTOCOL := "tcp(127.0.0.1:3306)"
	DBNAME := "dojo_api"
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
	db, err = gorm.Open(DBMS, CONNECT)
	if err != nil {
		panic(err.Error())
	}

	autoMigration()

	db.LogMode(true)
}

func GetDB() *gorm.DB {
	return db
}

func Close(){
	if err := db.Close(); err != nil{
		panic(err)
	}
}

func autoMigration() {
	db.Set("gorm:table_options", "ENGINE=InnoDB")
	db.AutoMigrate(&model.User{}, &model.Character{}, &model.UserCharacter{}, &model.Gacha{})
}