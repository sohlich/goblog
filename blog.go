package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/sohlich/blog/logic"
)


//Main method launching server
func main() {
	
	//Define rest
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", blog.Index)
	router.HandleFunc("/new", blog.InsertPost).Methods("POST")
	router.HandleFunc("/articles/{permalink}", blog.GetPost)
	log.Fatal(http.ListenAndServe(":8080", router))
}
