package main

import (
	"awesomeProject/nix/pkg/db"
	"awesomeProject/nix/pkg/operation"
	"fmt"
	"sync"
)

func main(){
	p := operation.GetPosts(7)
	go db.WriteToDBPost(p)

	c := operation.GetComment(7)
	var mutex = &sync.Mutex{}

	for _,value := range c{
		go db.WriteToDBComment(value,mutex)
		fmt.Println(value)
	}

}
