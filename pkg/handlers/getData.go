package handlers

import (
	"awesomeProject/nix/pkg/db"
	"awesomeProject/nix/pkg/operation"
	"net/http"
	"sync"
	"time"
)

func GetDate(w http.ResponseWriter,r *http.Request){
	p := operation.GetPosts(7)

	var mutex = &sync.Mutex{}

	gormDB := db.OpenDataBase()
	for _,value := range p{
		go db.WriteToDBPost(value,gormDB,mutex)
		go operation.GetComment(value)
	}

	time.Sleep(2*time.Second)
	w.Write([]byte("successfully completed writing to db"))
}

