package utils

import (
	"gopkg.in/mgo.v2"
	"os"
)

func MogConn() *mgo.Session {
	mongoUrl := os.Getenv("MONGO_URL")
	session, err := mgo.Dial("mongodb://" + mongoUrl)
	if err != nil {
		panic(err)

	}
	return session
}
