package blog

import(
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/sohlich/goblog/repository"
	"html/template"
)

//Route handlers
func Index(w http.ResponseWriter, r *http.Request) {
	log.Print("Accessing Index page")
	jsonEncoder := json.NewEncoder(w)
	collection := repository.PostRepository().FindAllSortByDate(3)
	jsonEncoder.Encode(collection)
}

func InsertPost(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var input repository.Post
	decoder.Decode(&input)
	repository.PostRepository().Add(&input)
}

func InsertPostForm(w http.ResponseWriter, req *http.Request) {
	generatedTemplate, err := template.ParseFiles("templates/postform.html")
	if err != nil{return}
	generatedTemplate.Execute(w,nil)
}


func GetPost(w http.ResponseWriter, r *http.Request){
	encoder := json.NewEncoder(w)
	variables := mux.Vars(r)
	permalink := variables["permalink"]
	result :=  repository.PostRepository().FindByPermalink(permalink)
	encoder.Encode(result)
}


	
