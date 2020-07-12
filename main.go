package main

import (
	"github.com/kondo1018008/DojoAPI/pkg/db"
	"github.com/kondo1018008/DojoAPI/pkg/server/model"
)

func main(){
	db.Init()

	defer db.Close()





}
