package repository
	
import(
	"log"
	"fmt"
	"time"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	)
		
type postRepository struct {
	 MongoSession *mgo.Session
	 PostCollection *mgo.Collection
}

var postRepositoryInstance *postRepository = nil


//Singleton pattern for repository
func PostRepository() (*postRepository) {
	if(postRepositoryInstance == nil){
		postRepositoryInstance = new(postRepository)
		session, err := mgo.Dial("localhost")
		if err != nil {
			log.Fatal("Cant connect to database")
			return postRepositoryInstance
		}
		postRepositoryInstance.MongoSession = session
		postRepositoryInstance.PostCollection = 
					postRepositoryInstance.MongoSession.DB("goBlog").C("posts")
		log.Print("PostRepsoitory initializes sucessfuly")
	}
	return postRepositoryInstance
}


func(repository *postRepository) Add(post *Post) bool{
	post.Permalink = fmt.Sprintf("perma%d", time.Now().Nanosecond())	
	err := repository.PostCollection.Insert(&post)
	if err != nil {
		return false
	}else{
		return true
	}
}

func(repository *postRepository) FindByPermalink(permalink string) Post{
	var result Post
	repository.PostCollection.Find(bson.M{"permalink":permalink}).One(&result)
	return result
}

func(repository *postRepository) FindAllSortByDate(limit int) []Post{
	var result []Post
	repository.PostCollection.Find(bson.M{}).Sort("datetime").Limit(limit).All(&result)
	return result
}


func(repository *postRepository) Close(){
	repository.MongoSession.Close()
}

	