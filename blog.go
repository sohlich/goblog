package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sohlich/goblog/blog"
	"github.com/sohlich/goblog/repository"
)

var router *mux.Router

//Main method launching server
func main() {

	//Define rest
	router = mux.NewRouter()
	http.Handle("/", blog.HttpSecurityInterceptor(router))
	router.HandleFunc("/articles/{permalink}", blog.GetPost)
	router.HandleFunc("/register", blog.RegisterFormProcess).Methods("POST")
	router.HandleFunc("/register", blog.RegisterForm).Methods("GET")
		router.HandleFunc("/login", blog.LoginFormProcess).Methods("POST")
	router.HandleFunc("/login", blog.LoginForm).Methods("GET")
	router.HandleFunc("/new", blog.InsertPost).Methods("POST")
	router.HandleFunc("/new", blog.InsertPostForm).Methods("GET")
	router.HandleFunc("/", blog.Index).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", nil)) //http handler handles request first
}

func init() {
	repository.PostRepository()
	repository.UserRepository()
}

func cleanUp() {
	repository.PostRepository().Close()
}
