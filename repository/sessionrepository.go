package repository
	
import(
	"gopkg.in/mgo.v2"
	)
	
type User struct{
	username string
	password string
	sessionToken string
}	

	
type userRepository struct {
	 MongoSession *mgo.Session
	 UserCollection *mgo.Collection
}

var userRepositoryInstance *userRepository = nil


func UserRepository() (*userRepository) {
	if(userRepositoryInstance == nil){	
		var session, err = mgo.Dial("localhost")
		if err != nil {
			return userRepositoryInstance
		}
		userRepositoryInstance.MongoSession = session
		userRepositoryInstance.UserCollection = userRepositoryInstance.MongoSession.DB("goBlog").C("user")
	}
	return userRepositoryInstance
}


func(repository *userRepository) Add(user User) bool{
	err := userRepositoryInstance.UserCollection.Insert(&user)
	if(err != nil){
		return false
	}
	return true
}




