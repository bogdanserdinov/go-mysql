package main

import (
	"awesomeProject/nix/pkg/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main(){
	r := mux.NewRouter()

	r.HandleFunc("/api/get/",handlers.GetDate)

	r.HandleFunc("/api/post/",handlers.GetAllPost).Methods("GET")
	r.HandleFunc("/api/post/",handlers.CreatePost).Methods("POST")
	r.HandleFunc("/api/post/{id}",handlers.GetOnePost).Methods("GET")
	r.HandleFunc("/api/post/{id}",handlers.DeletePost).Methods("DELETE")
	r.HandleFunc("/api/post/{id}",handlers.UpdatePost).Methods("PUT")

	r.HandleFunc("/api/comment/",handlers.GetAllComment).Methods("GET")
	r.HandleFunc("/api/comment/",handlers.CreateComment).Methods("POST")
	r.HandleFunc("/api/comment/{id}",handlers.GetOneComment).Methods("GET")
	r.HandleFunc("/api/comment/{id}",handlers.DeleteComment).Methods("DELETE")
	r.HandleFunc("/api/comment/{id}",handlers.UpdateComment).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8080",r))
}
