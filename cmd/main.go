package main

import (
	"awesomeProject/nix/pkg/db"
	"awesomeProject/nix/pkg/operation"
	"sync"
	"time"
)

func main(){
	p := operation.GetPosts(7)

	var mutex = &sync.Mutex{}
	for _,value := range p{
		go db.WriteToDBPost(value,mutex)
		//fmt.Println(value)
		go operation.GetComment(value)
	}

	time.Sleep(2*time.Second)
}
