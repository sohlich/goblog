package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/sohlich/goblog/logic"
	"github.com/sohlich/goblog/repository"
)


//Main method launching server
func main() {
	
	//Define rest
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/articles/{permalink}", blog.GetPost)
	router.HandleFunc("/new", blog.InsertPost).Methods("POST")
	router.HandleFunc("/new", blog.InsertPostForm).Methods("GET")
	router.HandleFunc("/", blog.Index)
	log.Fatal(http.ListenAndServe(":8080", router))
}


func init(){
	repository.PostRepository()
//	repository.UserRepository()
}


func cleanUp(){
	defer repository.PostRepository().Close()
}