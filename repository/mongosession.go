package repository

import (
	"gopkg.in/mgo.v2"
)

var session *mgo.Session = nil

//Creates mongo session singleton
//mgo.v2 lib should be designed for concurrent access
func MongoSession() (*mgo.Session, error) {
	if session == nil {
		var err error
		session, err = mgo.Dial("localhost")
		if err != nil {
			return nil, err
		}
	}
	return session, nil
}

//Closes global session for mongodb
func CloseMongoSession() {
	if session == nil {
		session.Close()
	}
}
