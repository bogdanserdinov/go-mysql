package db

import (
	"awesomeProject/nix/entity"
	"database/sql"
	"log"
)

func WriteToDBPost(p []entity.Post){
	db,err := sql.Open("mysql","root:blablabla29032002@/practice")
	if err != nil {
		log.Fatal("failed in opening db : ",err)
	}

	for _,value := range p{
		_,err = db.Exec("insert into practice.post values(?,?,?,?) ",value.UserID,value.ID,value.Title,value.Body)

		if err != nil {
			log.Fatal("failed in executing command : ",err)
		}

		//go getComment(value)
	}
}
