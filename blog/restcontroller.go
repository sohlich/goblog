package blog

import(
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/sohlich/goblog/repository"
	"html/template"
	"strings"
	"time"
)


type Page struct{
	Posts []repository.Post
}

//Route handlers
func Index(w http.ResponseWriter, r *http.Request) {
	generatedTemplate, err := template.ParseFiles("templates/index.html")
	if err != nil{log.Fatal("Cant process index template")}
	log.Print("Accessing Index page")
	collection := repository.PostRepository().FindAllSortByDate(3)
	page := Page{collection}
	generatedTemplate.Execute(w,page)
}

func InsertPost(w http.ResponseWriter, req *http.Request) {
	input := &repository.Post{
		Content: req.FormValue("content"),
		Title: req.FormValue("title"),
		Tags: strings.Split(req.FormValue("tags"),";"),
		DateTime: time.Now(),
	}
	repository.PostRepository().Add(input)
	http.Redirect(w, req,"/", 302)
}

func InsertPostForm(w http.ResponseWriter, req *http.Request) {
	log.Print("request for InsertNewPost")
	generatedTemplate, err := template.ParseFiles("templates/postform.html")
	if err != nil{
		log.Fatal("Error in parsing template")
		return
	}
	generatedTemplate.Execute(w,nil)
}

func AdminInterface(w http.ResponseWriter, r *http.Request){
	
}


func GetPost(w http.ResponseWriter, r *http.Request){
	encoder := json.NewEncoder(w)
	variables := mux.Vars(r)
	permalink := variables["permalink"]
	result :=  repository.PostRepository().FindByPermalink(permalink)
	encoder.Encode(result)
}


	
