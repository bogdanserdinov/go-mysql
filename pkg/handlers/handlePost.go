package handlers

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/bogdanserdinov/go-mysql/entity"
	"github.com/bogdanserdinov/go-mysql/pkg/db"
	"github.com/labstack/echo"
	"log"
	"net/http"
	"strconv"
)


// @Summary Show a post
// @Tags Post
// @Description get string by ID(Post)
// @ID get-post-by-id
// @Accept  json
// @Produce  json
// @Produce  xml
// @Param id path int true "Post ID"
// @Success 200 {integer} integer 1
// @Router /post/:id [get]
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


// @Summary List posts
// @Description get all posts
// @Tags Post
// @ID get-list-of-posts
// @Accept  json
// @Produce  json
// @Produce xml
// @Success 200 {integer} integer 1
// @Router /post/ [get]
func GetAllPost(e echo.Context) error{
	gormDB := db.OpenDataBase()
	var posts []entity.Post
	gormDB.Table("posts").Select("UserID, ID, Title,Body").Scan(&posts)

	jsonPost,err := json.Marshal(&posts)
	if err != nil{
		log.Println("could not convert post to json",err.Error())
		return e.String(http.StatusNotFound,fmt.Sprint("could not parse json"))
	}
	xmlPost,err := xml.Marshal(&posts)
	if err != nil{
		log.Println("could not convert post to xml",err.Error())
		return e.String(http.StatusNotFound,fmt.Sprint("could not parse xml"))
	}
	return e.String(http.StatusOK,fmt.Sprint(string(jsonPost),"\n",string(xmlPost)))
}

// @Summary Create post
// @Description create new post
// @Tags Post
// @ID create-new-post
// @Accept  json
// @Produce  json
// @Produce xml
// @Success 200 {integer} integer 1
// @Router /post/ [post]
func CreatePost(e echo.Context) error{
	var newPost entity.Post
	json.NewDecoder(e.Request().Body).Decode(&newPost)

	gormDB := db.OpenDataBase()
	gormDB.Create(&newPost)

	jsonPost,err := json.Marshal(&newPost)
	if err != nil{
		log.Println("could not convert post to json",err.Error())
		return e.String(http.StatusNotFound,fmt.Sprint("could not parse json"))
	}
	xmlPost,err := xml.Marshal(&newPost)
	if err != nil{
		log.Println("could not convert post to xml",err.Error())
		return e.String(http.StatusNotFound,fmt.Sprint("could not parse xml"))
	}
	return e.String(http.StatusOK,fmt.Sprint(string(jsonPost),"\n",string(xmlPost)))
}


// @Summary Delete post
// @Description delete post post
// @Tags Post
// @ID delete-post
// @Accept  json
// @Produce  json
// @Produce xml
// @Param id path int true "Post ID"
// @Success 200 {integer} integer 1
// @Router /post/:id [post]
func DeletePost(e echo.Context) error{
	id := e.Param("id")
	idStr,err := strconv.Atoi(id)
	if err != nil{
		log.Println("could not convert id to int",err.Error())
		return e.String(http.StatusNotFound,fmt.Sprint("could not parse json"))
	}

	var newPost entity.Post
	gormDB := db.OpenDataBase()
	gormDB.First(&newPost, idStr)
	gormDB.Delete(&entity.Post{}, idStr)

	jsonPost,err := json.Marshal(&newPost)
	if err != nil{
		log.Println("could not convert post to json",err.Error())
		return e.String(http.StatusNotFound,fmt.Sprint("could not parse json"))
	}

	xmlPost,err := xml.Marshal(&newPost)
	if err != nil{
		log.Println("could not convert post to xml",err.Error())
		return e.String(http.StatusNotFound,fmt.Sprint("could not parse xml"))
	}

	return e.String(http.StatusOK,fmt.Sprint(string(jsonPost),"\n",string(xmlPost)))
}


// @Summary Update post
// @Description update post
// @Tags Post
// @ID update-post
// @Accept  json
// @Produce  json
// @Produce xml
// @Param id path int true "Post ID"
// @Success 200 {integer} integer 1
// @Router /post/:id [put]
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
		return e.String(http.StatusNotFound,fmt.Sprint("could not parse json"))
	}

	xmlPost,err := xml.Marshal(&newPost)
	if err != nil{
		log.Println("could not convert post to xml",err.Error())
		return e.String(http.StatusNotFound,fmt.Sprint("could not parse xml"))
	}

	return e.String(http.StatusOK,fmt.Sprint(string(jsonPost),"\n",string(xmlPost)))
}

