package db

import (
	"awesomeProject/nix/entity"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"sync"
)

//WriteToDBPost - function required for writing data to post entity
func WriteToDBPost(p entity.Post,mutex *sync.Mutex){
	mutex.Lock()
	db,err := sql.Open("mysql","root:blablabla29032002@/public")
	if err != nil {
		log.Fatal("failed in opening db : ",err)
	}
	fmt.Println("writing to db post...")
	fmt.Println(p)

	_,err = db.Exec("insert into public.posts values(?,?,?,?) ",p.UserID,p.ID,p.Title,p.Body)

	if err != nil {
		log.Fatal("failed in executing command : ",err)
	}

	mutex.Unlock()
}
