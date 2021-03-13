package main

import (
	_ "awesomeProject/nix/docs"
	"awesomeProject/nix/pkg/handlers"
	"github.com/labstack/echo"
	"github.com/swaggo/echo-swagger"
)

// @title REST API with Echo
// @version 1.0
// @description this is simple REST API which works with MySQL and written using Echo framework.

// @host localhost:8080
// @BasePath /api/

func main(){
	e := echo.New()

	e.GET("/api/get-data",handlers.GetData)

	e.GET("/api/post/",handlers.GetAllPost)
	e.POST("/api/post/",handlers.CreatePost)
	e.GET("/api/post/:id",handlers.GetOnePost)
	e.DELETE("/api/post/:id",handlers.DeletePost)
	e.PUT("/api/post/:id",handlers.UpdatePost)

	e.GET("/api/comment/",handlers.GetAllComment)
	e.POST("/api/comment/",handlers.CreateComment)
	e.GET("/api/comment/:id",handlers.GetOneComment)
	e.DELETE("/api/comment/:id",handlers.DeleteComment)
	e.PUT("/api/comment/:id",handlers.UpdateComment)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
