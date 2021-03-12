package main

import (
	"awesomeProject/nix/pkg/handlers"
	"github.com/labstack/echo"
)

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

	e.Logger.Fatal(e.Start(":8080"))
}
