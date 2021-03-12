package main

import (
	"awesomeProject/nix/pkg/handlers"
	"github.com/gorilla/mux"
	"github.com/labstack/echo"
)

func main(){
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/api/get/",handlers.GetDate)

	/*r.HandleFunc("/api/post/",handlers.GetAllPost).Methods("GET")
	r.HandleFunc("/api/post/",handlers.CreatePost).Methods("POST")
	r.HandleFunc("/api/post/{id}",handlers.GetOnePost).Methods("GET")
	r.HandleFunc("/api/post/{id}",handlers.DeletePost).Methods("DELETE")
	r.HandleFunc("/api/post/{id}",handlers.UpdatePost).Methods("PUT")

	r.HandleFunc("/api/comment/",handlers.GetAllComment).Methods("GET")
	r.HandleFunc("/api/comment/",handlers.CreateComment).Methods("POST")
	r.HandleFunc("/api/comment/{id}",handlers.GetOneComment).Methods("GET")
	r.HandleFunc("/api/comment/{id}",handlers.DeleteComment).Methods("DELETE")
	r.HandleFunc("/api/comment/{id}",handlers.UpdateComment).Methods("PUT")*/

	e := echo.New()

	e.GET("/api/post/",handlers.GetAllPost)
	e.POST("/api/post/",handlers.CreatePost)
	e.GET("/api/post/{id}",handlers.GetOnePost)
	e.DELETE("/api/post/{id}",handlers.DeletePost)
	e.PUT("/api/post/{id}",handlers.UpdatePost)

	e.GET("/api/comment/",handlers.GetAllComment)
	e.POST("/api/comment/",handlers.CreateComment)
	e.GET("/api/comment/{id}",handlers.GetOneComment)
	e.DELETE("/api/comment/{id}",handlers.DeleteComment)
	e.PUT("/api/comment/{id}",handlers.UpdateComment)

	e.Logger.Fatal(e.Start(":8080"))
}
