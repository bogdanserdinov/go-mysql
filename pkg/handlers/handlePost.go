package handlers

import (
	"awesomeProject/nix/entity"
	"awesomeProject/nix/pkg/db"
	"encoding/json"
	"encoding/xml"
	"fmt"
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

func CreatePost(e echo.Context) error{
	var newPost entity.Post
	json.NewDecoder(e.Request().Body).Decode(&newPost)

	gormDB := db.OpenDataBase()
	gormDB.Create(&newPost)

	jsonPost,err := json.Marshal(&newPost)
	if err != nil{
		log.Println("could not convert post to json",err.Error())
	}
	xmlPost,err := xml.Marshal(&newPost)
	if err != nil{
		log.Println("could not convert post to xml",err.Error())
	}
	return e.String(http.StatusOK,fmt.Sprint(string(jsonPost),"\n",string(xmlPost)))
}

func DeletePost(e echo.Context) error{
	id := e.Param("id")
	idStr,err := strconv.Atoi(id)
	if err != nil{
		log.Println("could not convert id to int",err.Error())
	}

	var newPost entity.Post
	gormDB := db.OpenDataBase()
	gormDB.First(&newPost, idStr)
	gormDB.Delete(&entity.Post{}, idStr)

	jsonPost,err := json.Marshal(&newPost)
	if err != nil{
		log.Println("could not convert post to json",err.Error())
	}

	xmlPost,err := xml.Marshal(&newPost)
	if err != nil{
		log.Println("could not convert post to xml",err.Error())
	}

	return e.String(http.StatusOK,fmt.Sprint(string(jsonPost),"\n",string(xmlPost)))
}

func UpdatePost(e echo.Context) error{
	id := e.Param("id")

	idStr,err := strconv.Atoi(id)
	if err != nil{
		log.Println("could not convert id to int")
	}

	var newPost entity.Post
	json.NewDecoder(e.Request().Body).Decode(&newPost)

	gormDB := db.OpenDataBase()
	gormDB.Model(&entity.Post{}).Where("id = ?",idStr).Updates(entity.Post{
		UserID: newPost.UserID,
		Title:  newPost.Title,
		Body:   newPost.Body,
	})

	jsonPost,err := json.Marshal(&newPost)
	if err != nil{
		log.Println("could not convert post to json",err.Error())
	}

	xmlPost,err := xml.Marshal(&newPost)
	if err != nil{
		log.Println("could not convert post to xml",err.Error())
	}

	return e.String(http.StatusOK,fmt.Sprint(string(jsonPost),"\n",string(xmlPost)))
}

