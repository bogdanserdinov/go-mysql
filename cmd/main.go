package main

import (
	"awesomeProject/nix/pkg/db"
	"awesomeProject/nix/pkg/operation"
	"database/sql"
	"log"
	"sync"
	"time"
)

func main(){
	p := operation.GetPosts(7)

	var mutex = &sync.Mutex{}

	database,err := sql.Open("mysql","root:blablabla29032002@/public")
	if err != nil {
		log.Fatal("failed in opening db : ",err)
	}
	defer database.Close()
	for _,value := range p{
		go db.WriteToDBPost(value,database,mutex)
		go operation.GetComment(value)
	}


	time.Sleep(2*time.Second)
}
