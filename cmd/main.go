package main

import (
	"awesomeProject/nix/pkg/db"
	"awesomeProject/nix/pkg/handlers"
	"awesomeProject/nix/pkg/operation"
	"fmt"
	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"sync"
	"time"
)

func GetDate(w http.ResponseWriter,r *http.Request){
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

func main(){
	r := mux.NewRouter()

	r.HandleFunc("/api/get/",GetDate)

	r.HandleFunc("/api/post/",handlers.GetAllPost).Methods("GET")
	r.HandleFunc("/api/post/",handlers.CreatePost).Methods("POST")
	r.HandleFunc("/api/post/{id}",handlers.GetOnePost).Methods("GET")
	r.HandleFunc("/api/post/{id}",handlers.DeletePost).Methods("DELETE")
	r.HandleFunc("/api/post/{id}",handlers.UpdatePost).Methods("PUT")

	r.HandleFunc("/api/comment/",handlers.GetAllComment).Methods("GET")
	r.HandleFunc("/api/comment/",handlers.CreateComment).Methods("POST")
	r.HandleFunc("/api/comment/{id}",handlers.GetOneComment).Methods("GET")
	r.HandleFunc("/api/comment/{id}",handlers.DeleteComment).Methods("DELETE")
	r.HandleFunc("/api/comment/{id}",handlers.UpdateComment).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8080",r))

}
