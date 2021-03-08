package main

import (
	"awesomeProject/nix/pkg/db"
	"awesomeProject/nix/pkg/operation"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"sync"
	"time"
)

func main(){
	p := operation.GetPosts(7)

	var mutex = &sync.Mutex{}
	
	dsn := "root:blablabla29032002@tcp(127.0.0.1:3306)/public?charset=utf8mb4&parseTime=True&loc=Local"
	gormDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed in opening db : ",err)
	}
	for _,value := range p{
		go db.WriteToDBPost(value,gormDB,mutex)
		go operation.GetComment(value)
	}

	time.Sleep(2*time.Second)
	fmt.Println("successfully completed writing to db")
}
