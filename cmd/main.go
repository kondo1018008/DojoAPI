package main

import (
	"github.com/kondo1018008/DojoAPI/pkg/db"
	"github.com/kondo1018008/DojoAPI/pkg/server"
)

func main(){
	db.Init()
	server.Init()
	db.Close()
}
