package repository
	
import(
	"log"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	)
	
	
type userRepository struct {
	 MongoSession *mgo.Session
	 UserCollection *mgo.Collection
}

var userRepositoryInstance *userRepository = nil


func UserRepository() (*userRepository) {
	if(userRepositoryInstance == nil){	
		userRepositoryInstance = new(userRepository)
		var session, err = mgo.Dial("localhost")
		if err != nil {
			log.Fatal("Cant initialize repository")
			return userRepositoryInstance
		}
		userRepositoryInstance.MongoSession = session
		userRepositoryInstance.UserCollection = userRepositoryInstance.MongoSession.DB("goBlog").C("user")
	}
	return userRepositoryInstance
}


func(repository *userRepository) Add(user *User) (*User, error){
	err := userRepositoryInstance.UserCollection.Insert(user)
	if(err != nil){
		return nil,err
	}
	return user,nil
}

func(repository *userRepository) FindByUsernameAndPassword(username,password string) (*User, error){
	var user User
	err := userRepositoryInstance.UserCollection.Find(bson.M{"username": username, "password": password}).One(&user)
	if(err != nil){
		return nil,err
	}
	return &user,nil
}

func(repository *userRepository) FindByUsername(username string) (*User, error){
	var user User
	err := userRepositoryInstance.UserCollection.Find(bson.M{"username": username}).One(&user)
	if(err != nil){
		return nil,err
	}
	log.Println(user)
	return &user,nil
}




