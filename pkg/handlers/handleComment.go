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

func GetOneComment(e echo.Context) error{
	id := e.Param("id")

	gormDB := db.OpenDataBase()

	var comments entity.Comment
	idStr,err := strconv.Atoi(id)

	if err != nil{
		log.Println("could not convert id to int",err.Error())
	}
	gormDB.Where("ID = ?", idStr).First(&comments)

	jsonComment,err := json.Marshal(&comments)
	if err != nil{
		log.Println("could not convert comment to json",err.Error())
	}
	xmlComment,err := xml.Marshal(&comments)
	if err != nil{
		log.Println("could not convert comment to xml",err.Error())
	}
	return e.String(http.StatusOK,fmt.Sprint(string(jsonComment),"\n",string(xmlComment)))
}

func GetAllComment(e echo.Context) error{
	gormDB := db.OpenDataBase()
	var comments []entity.Comment
	gormDB.Table("comments").Select("PostID, ID, Name, Email, Body").Scan(&comments)

	jsonComment,err := json.Marshal(&comments)
	if err != nil{
		log.Println("could not convert comment to json",err.Error())
	}
	xmlComment,err := xml.Marshal(&comments)
	if err != nil{
		log.Println("could not convert comment to xml",err.Error())
	}
	return e.String(http.StatusOK,fmt.Sprint(string(jsonComment),"\n",string(xmlComment)))
}

func CreateComment(e echo.Context) error{
	var newComment entity.Comment
	json.NewDecoder(e.Request().Body).Decode(&newComment)

	gormDB := db.OpenDataBase()
	gormDB.Create(&newComment)

	jsonComment,err := json.Marshal(&newComment)
	if err != nil{
		log.Println("could not convert comment to json",err.Error())
	}
	xmlComment,err := xml.Marshal(&newComment)
	if err != nil{
		log.Println("could not convert comment to xml",err.Error())
	}
	return e.String(http.StatusOK,fmt.Sprint(string(jsonComment),"\n",string(xmlComment)))
}

func DeleteComment(e echo.Context) error{
	id := e.Param("id")
	idStr,err := strconv.Atoi(id)
	if err != nil{
		log.Println("could not convert id to int",err.Error())
	}

	var newComment entity.Comment
	gormDB := db.OpenDataBase()
	gormDB.First(&newComment, idStr)
	gormDB.Delete(&entity.Comment{}, idStr)

	jsonComment,err := json.Marshal(&newComment)
	if err != nil{
		log.Println("could not convert comment to json",err.Error())
	}
	xmlComment,err := xml.Marshal(&newComment)
	if err != nil{
		log.Println("could not convert comment to xml",err.Error())
	}
	return e.String(http.StatusOK,fmt.Sprint(string(jsonComment),"\n",string(xmlComment)))
}

func UpdateComment(e echo.Context) error{
	id := e.Param("id")

	idStr,err := strconv.Atoi(id)
	if err != nil{
		log.Println("could not convert id to int")
	}

	var newComment entity.Comment
	json.NewDecoder(e.Request().Body).Decode(&newComment)

	gormDB := db.OpenDataBase()
	gormDB.Model(&entity.Post{}).Where("id = ?",idStr).Updates(entity.Comment{
		PostID: newComment.PostID,
		Name:  	newComment.Name,
		Email:  newComment.Email,
		Body:   newComment.Body,
	})

	jsonComment,err := json.Marshal(&newComment)
	if err != nil{
		log.Println("could not convert comment to json",err.Error())
	}
	xmlComment,err := xml.Marshal(&newComment)
	if err != nil{
		log.Println("could not convert comment to xml",err.Error())
	}
	return e.String(http.StatusOK,fmt.Sprint(string(jsonComment),"\n",string(xmlComment)))
}
