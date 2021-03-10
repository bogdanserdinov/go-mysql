package handlers

import (
	"awesomeProject/nix/entity"
	"awesomeProject/nix/pkg/db"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func GetOneComment(w http.ResponseWriter,r *http.Request){
	vars := mux.Vars(r)
	id := vars["id"]

	gormDB := db.OpenDataBase()

	var comments entity.Comment
	idStr,err := strconv.Atoi(id)

	if err != nil{
		log.Println("could not convert id to int",err.Error())
	}
	gormDB.Where("ID = ?", idStr).First(&comments)
	json.NewEncoder(w).Encode(&comments)
	xml.NewEncoder(w).Encode(&comments)
}

func GetAllComment(w http.ResponseWriter,r *http.Request){
	gormDB := db.OpenDataBase()
	var comments []entity.Comment
	gormDB.Table("comments").Select("PostID, ID, Name, Email, Body").Scan(&comments)

	json.NewEncoder(w).Encode(&comments)
	xml.NewEncoder(w).Encode(&comments)
}

func CreateComment(w http.ResponseWriter,r *http.Request){
	var newComment entity.Comment
	json.NewDecoder(r.Body).Decode(&newComment)

	gormDB := db.OpenDataBase()
	gormDB.Create(&newComment)

	json.NewEncoder(w).Encode(&newComment)
	xml.NewEncoder(w).Encode(&newComment)
}

func DeleteComment(w http.ResponseWriter,r *http.Request){
	vars := mux.Vars(r)
	id := vars["id"]
	idStr,err := strconv.Atoi(id)
	if err != nil{
		log.Println("could not convert id to int",err.Error())
	}

	var newComment entity.Comment
	gormDB := db.OpenDataBase()
	gormDB.First(&newComment, idStr)
	gormDB.Delete(&entity.Comment{}, idStr)

	json.NewEncoder(w).Encode(&newComment)
	xml.NewEncoder(w).Encode(&newComment)
}

func UpdateComment(w http.ResponseWriter,r *http.Request){
	vars := mux.Vars(r)
	id := vars["id"]

	idStr,err := strconv.Atoi(id)
	if err != nil{
		log.Println("could not convert id to int")
	}

	var newComment entity.Comment
	json.NewDecoder(r.Body).Decode(&newComment)

	gormDB := db.OpenDataBase()
	gormDB.Model(&entity.Comment{}).Where("id = ?",idStr).Updates(entity.Comment{
		PostID: newComment.PostID,
		Name:  newComment.Name,
		Email:  newComment.Email,
		Body:   newComment.Body,
	})

	json.NewEncoder(w).Encode(&newComment)
	xml.NewEncoder(w).Encode(&newComment)
}

