package handlers

import (
	"awesomeProject/nix/entity"
	"awesomeProject/nix/pkg/db"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/labstack/echo"
	"log"
	"net/http"
	"strconv"
)

func GetOnePost(e echo.Context) error{
	id := e.Param("id")


	gormDB := db.OpenDataBase()

	var posts entity.Post
	idStr,err := strconv.Atoi(id)

	if err != nil{
		log.Println("could not convert id to int",err.Error())
	}
	gormDB.Where("ID = ?", idStr).First(&posts)

	jsonPost,err := json.Marshal(&posts)
	if err != nil{
		log.Println("could not convert post to json",err.Error())
	}
	xmlPost,err := xml.Marshal(&posts)
	if err != nil{
		log.Println("could not convert post to xml",err.Error())
	}
	return e.String(http.StatusOK,fmt.Sprint(string(jsonPost),"\n",string(xmlPost)))
}

func GetAllPost(e echo.Context) error{
	gormDB := db.OpenDataBase()
	var posts []entity.Post
	gormDB.Table("posts").Select("UserID, ID, Title,Body").Scan(&posts)

	jsonPost,err := json.Marshal(&posts)
	if err != nil{
		log.Println("could not convert post to json",err.Error())
	}
	xmlPost,err := xml.Marshal(&posts)
	if err != nil{
		log.Println("could not convert post to xml",err.Error())
	}
	return e.String(http.StatusOK,fmt.Sprint(string(jsonPost),"\n",string(xmlPost)))
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

