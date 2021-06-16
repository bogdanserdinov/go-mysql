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

// @Summary Show a comment
// @Tags Comment
// @Description get string by ID(Comment)
// @ID get-comment-by-id
// @Accept  json
// @Produce  json
// @Produce  xml
// @Param id path int true "Comment ID"
// @Success 200 {integer} integer 1
// @Router /comment/:id [get]
func GetOneComment(e echo.Context) error {
	id := e.Param("id")

	gormDB := db.OpenDataBase()

	var comments entity.Comment
	idStr, err := strconv.Atoi(id)

	if err != nil {
		log.Println("could not convert id to int", err.Error())
	}
	gormDB.Where("ID = ?", idStr).First(&comments)

	jsonComment, err := json.Marshal(&comments)
	if err != nil {
		log.Println("could not convert comment to json", err.Error())
		return e.String(http.StatusNotFound, fmt.Sprint("could not parse json"))
	}
	xmlComment, err := xml.Marshal(&comments)
	if err != nil {
		log.Println("could not convert comment to xml", err.Error())
		return e.String(http.StatusNotFound, fmt.Sprint("could not parse xml"))
	}
	return e.String(http.StatusOK, fmt.Sprint(string(jsonComment), "\n", string(xmlComment)))
}

// @Summary List comments
// @Description get all comments
// @Tags Comment
// @ID get-list-of-comments
// @Accept  json
// @Produce  json
// @Produce xml
// @Success 200 {integer} integer 1
// @Router /comments/ [get]
func GetAllComment(e echo.Context) error {
	gormDB := db.OpenDataBase()
	var comments []entity.Comment
	gormDB.Table("comments").Select("PostID, ID, Name, Email, Body").Scan(&comments)

	jsonComment, err := json.Marshal(&comments)
	if err != nil {
		log.Println("could not convert comment to json", err.Error())
		return e.String(http.StatusNotFound, fmt.Sprint("could not parse json"))
	}
	xmlComment, err := xml.Marshal(&comments)
	if err != nil {
		log.Println("could not convert comment to xml", err.Error())
		return e.String(http.StatusNotFound, fmt.Sprint("could not parse xml"))
	}
	return e.String(http.StatusOK, fmt.Sprint(string(jsonComment), "\n", string(xmlComment)))
}

// @Summary Create comment
// @Description create new comment
// @Tags Comment
// @ID create-new-comment
// @Accept  json
// @Produce  json
// @Produce xml
// @Success 200 {integer} integer 1
// @Router /comment/ [post]
func CreateComment(e echo.Context) error {
	var newComment entity.Comment
	json.NewDecoder(e.Request().Body).Decode(&newComment)

	gormDB := db.OpenDataBase()
	gormDB.Create(&newComment)

	jsonComment, err := json.Marshal(&newComment)
	if err != nil {
		log.Println("could not convert comment to json", err.Error())
		return e.String(http.StatusNotFound, fmt.Sprint("could not parse json"))
	}
	xmlComment, err := xml.Marshal(&newComment)
	if err != nil {
		log.Println("could not convert comment to xml", err.Error())
		return e.String(http.StatusNotFound, fmt.Sprint("could not parse xml"))
	}
	return e.String(http.StatusOK, fmt.Sprint(string(jsonComment), "\n", string(xmlComment)))
}

// @Summary Delete comment
// @Description delete comment
// @Tags Comment
// @ID delete-comment
// @Accept  json
// @Produce  json
// @Produce xml
// @Param id path int true "Comment ID"
// @Success 200 {integer} integer 1
// @Router /comment/:id [post]
func DeleteComment(e echo.Context) error {
	id := e.Param("id")
	idStr, err := strconv.Atoi(id)
	if err != nil {
		log.Println("could not convert id to int", err.Error())
	}

	var newComment entity.Comment
	gormDB := db.OpenDataBase()
	gormDB.First(&newComment, idStr)
	gormDB.Delete(&entity.Comment{}, idStr)

	jsonComment, err := json.Marshal(&newComment)
	if err != nil {
		log.Println("could not convert comment to json", err.Error())
		return e.String(http.StatusNotFound, fmt.Sprint("could not parse json"))
	}
	xmlComment, err := xml.Marshal(&newComment)
	if err != nil {
		log.Println("could not convert comment to xml", err.Error())
		return e.String(http.StatusNotFound, fmt.Sprint("could not parse xml"))
	}
	return e.String(http.StatusOK, fmt.Sprint(string(jsonComment), "\n", string(xmlComment)))
}

// @Summary Update comment
// @Description update comment
// @Tags Comment
// @ID update-comment
// @Accept  json
// @Produce  json
// @Produce xml
// @Success 200 {integer} integer 1
// @Router /comment/:id [put]
func UpdateComment(e echo.Context) error {
	id := e.Param("id")

	idStr, err := strconv.Atoi(id)
	if err != nil {
		log.Println("could not convert id to int")
	}

	var newComment entity.Comment
	json.NewDecoder(e.Request().Body).Decode(&newComment)

	gormDB := db.OpenDataBase()
	gormDB.Model(&entity.Post{}).Where("id = ?", idStr).Updates(entity.Comment{
		PostID: newComment.PostID,
		Name:   newComment.Name,
		Email:  newComment.Email,
		Body:   newComment.Body,
	})

	jsonComment, err := json.Marshal(&newComment)
	if err != nil {
		log.Println("could not convert comment to json", err.Error())
		return e.String(http.StatusNotFound, fmt.Sprint("could not parse json"))
	}
	xmlComment, err := xml.Marshal(&newComment)
	if err != nil {
		log.Println("could not convert comment to xml", err.Error())
		return e.String(http.StatusNotFound, fmt.Sprint("could not parse xml"))
	}
	return e.String(http.StatusOK, fmt.Sprint(string(jsonComment), "\n", string(xmlComment)))
}
