package repository

import (
	"gopkg.in/mgo.v2"
)

var session *mgo.Session = nil
var host string = "bla"
var port string = "27017"

//Creates mongo session singleton
//mgo.v2 lib should be designed for concurrent access
func MongoSession() (*mgo.Session, error) {
	if session == nil {
		var err error
		session, err = mgo.Dial(host+":"+port)
		if err != nil {
			return nil, err
		}
	}
	return session, nil
}

func SetupMongo(cfghost,cfgport string){
	host = cfghost
	port = cfgport	
}


//Closes global session for mongodb
func CloseMongoSession() {
	if session == nil {
		session.Close()
	}
}
