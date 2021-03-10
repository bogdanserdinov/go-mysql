package handlers

import (
	"awesomeProject/nix/entity"
	"awesomeProject/nix/pkg/db"
	"encoding/json"
	"encoding/xml"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func GetOnePost(w http.ResponseWriter,r *http.Request){
	vars := mux.Vars(r)
	id := vars["id"]

	gormDB := db.OpenDataBase()

	var posts entity.Post
	idStr,err := strconv.Atoi(id)

	if err != nil{
		log.Println("could not convert id to int",err.Error())
	}
	gormDB.Where("ID = ?", idStr).First(&posts)
	json.NewEncoder(w).Encode(&posts)
	xml.NewEncoder(w).Encode(&posts)
}

func GetAllPost(w http.ResponseWriter,r *http.Request){
	gormDB := db.OpenDataBase()
	var posts []entity.Post
	gormDB.Table("posts").Select("UserID, ID, Title,Body").Scan(&posts)

	json.NewEncoder(w).Encode(&posts)
	xml.NewEncoder(w).Encode(&posts)
}

func CreatePost(w http.ResponseWriter,r *http.Request){
	var newPost entity.Post
	json.NewDecoder(r.Body).Decode(&newPost)

	gormDB := db.OpenDataBase()
	gormDB.Create(&newPost)

	json.NewEncoder(w).Encode(&newPost)
	xml.NewEncoder(w).Encode(&newPost)
}

func DeletePost(w http.ResponseWriter,r *http.Request){
	vars := mux.Vars(r)
	id := vars["id"]
	idStr,err := strconv.Atoi(id)
	if err != nil{
		log.Println("could not convert id to int",err.Error())
	}

	var newPost entity.Post
	gormDB := db.OpenDataBase()
	gormDB.First(&newPost, idStr)
	gormDB.Delete(&entity.Post{}, idStr)

	json.NewEncoder(w).Encode(&newPost)
	xml.NewEncoder(w).Encode(&newPost)
}

func UpdatePost(w http.ResponseWriter,r *http.Request){
	vars := mux.Vars(r)
	id := vars["id"]

	idStr,err := strconv.Atoi(id)
	if err != nil{
		log.Println("could not convert id to int")
	}

	var newPost entity.Post
	json.NewDecoder(r.Body).Decode(&newPost)

	gormDB := db.OpenDataBase()
	gormDB.Model(&entity.Post{}).Where("id = ?",idStr).Updates(entity.Post{
		UserID: newPost.UserID,
		Title:  newPost.Title,
		Body:   newPost.Body,
	})

	json.NewEncoder(w).Encode(&newPost)
	xml.NewEncoder(w).Encode(&newPost)
}

