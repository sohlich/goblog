package blog
	
import(
	"fmt"
	"time"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	)
	
	
type PostRepository struct {
	 MongoSession *mgo.Session
	 PostCollection *mgo.Collection
}


func CreatePostRepository() (PostRepository,error) {
	
	var repository PostRepository	
	var session, err = mgo.Dial("localhost")
	if err != nil {
		return repository, err
	}
	repository.MongoSession = session
	repository.PostCollection = repository.MongoSession.DB("goBlog").C("posts")
	
	return repository,nil
}


func(repository *PostRepository) Add(post *Post) bool{
	post.Permalink = fmt.Sprintf("perma%d", time.Now().Nanosecond())	
	err := repository.PostCollection.Insert(&post)
	if err != nil {
		return false
	}else{
		return true
	}
}

func(repository *PostRepository) FindByPermalink(permalink string) Post{
	var result Post
	repository.PostCollection.Find(bson.M{"permalink":permalink}).One(&result)
	return result
}

func(repository *PostRepository) Close(){
	repository.MongoSession.Close()
}

	