package handlers

import (
	"github.com/bogdanserdinov/go-mysql/pkg/db"
	"github.com/bogdanserdinov/go-mysql/pkg/operation"
	"github.com/labstack/echo"
	"net/http"
	"sync"
	"time"
)

func GetData(e echo.Context) error {
	p := operation.GetPosts(7)

	var mutex = &sync.Mutex{}

	gormDB := db.OpenDataBase()
	for _, value := range p {
		go db.WriteToDBPost(value, gormDB, mutex)
		go operation.GetComment(value)
	}

	time.Sleep(2 * time.Second)
	return e.String(http.StatusOK, "successfully completed writing to db")
}
