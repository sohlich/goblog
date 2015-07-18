package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/sohlich/goblog/blog"
	"github.com/sohlich/goblog/repository"
	"code.google.com/p/gcfg"
)

const defaultConfig = `
    [mongo]
    port = 27017
    host= localhost`


var router *mux.Router

type Config struct {
    Port string
	Host string
}

type configFile struct {
    Mongo Config
}





//Main method launching server
func main() {

	Init()
	defer cleanUp()
	
	//Define rest
	router = mux.NewRouter()
	http.Handle("/", blog.HttpSecurityInterceptor(router))
	router.HandleFunc("/articles/{permalink}", blog.GetPost)
	router.HandleFunc("/register", blog.RegisterFormProcess).Methods("POST")
	router.HandleFunc("/register", blog.RegisterForm).Methods("GET")
	router.HandleFunc("/login", blog.LoginFormProcess).Methods("POST")
	router.HandleFunc("/login", blog.LoginForm).Methods("GET")
	router.HandleFunc("/logout", blog.Logout).Methods("GET")
	router.HandleFunc("/new", blog.InsertPost).Methods("POST")
	router.HandleFunc("/new", blog.InsertPostForm).Methods("GET")
	router.HandleFunc("/admin", blog.AdminInterface).Methods("GET")
	router.PathPrefix("/css").Handler(http.FileServer(http.Dir("./static")))
	router.PathPrefix("/js").Handler(http.FileServer(http.Dir("./static")))
	router.HandleFunc("/", blog.Index).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", nil)) //http handler handles request first
}

func Init() {
	config := LoadConfiguration("application.conf")	
	repository.SetupMongo(config.Host,config.Port)
	repository.PostRepository()
	repository.UserRepository()
}

func cleanUp() {
	repository.CloseMongoSession()
}


func LoadConfiguration(cfgFile string) Config {
    var err error
    var cfg configFile
    if cfgFile != "" {
        err = gcfg.ReadFileInto(&cfg, cfgFile)
    } else {
        err = gcfg.ReadStringInto(&cfg, defaultConfig)
    }
	if err != nil {log.Panic(err)}
    return cfg.Mongo
}

