package blog

import(
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)
	
var postRepository , err = CreatePostRepository();

//Route handlers
func Index(w http.ResponseWriter, r *http.Request) {
	jsonEncoder := json.NewEncoder(w)
	log.Print("Accessing Index page")
	var collection []Post
	postRepository.PostCollection.Find(bson.M{}).Limit(3).All(&collection)
	jsonEncoder.Encode(collection)
}

func InsertPost(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var input Post
	decoder.Decode(&input)
	postRepository.Add(&input)
}

func GetPost(w http.ResponseWriter, r *http.Request){
	encoder := json.NewEncoder(w)
	variables := mux.Vars(r)
	permalink := variables["permalink"]
	result :=  postRepository.FindByPermalink(permalink)
	encoder.Encode(result)
}