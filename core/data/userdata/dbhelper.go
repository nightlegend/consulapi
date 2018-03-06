package userdata

import (
	"github.com/nightlegend/consulapi/core/utils"
	log "github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	account Accounts
)

func FindOne(user *Accounts) bool {
	session := utils.MogConn()
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("consul").C("accounts")
	err := c.Find(bson.M{"username": user.USERNAME, "password": user.PASSWORD}).One(&account)
	if err != nil {
		log.Error(err)
		return false
	}
	return true
}

func Add(user *Accounts) bool {
	session := utils.MogConn()
	defer session.Close()
	c := session.DB("consul").C("accounts")
	session.SetMode(mgo.Monotonic, true)
	err := c.Insert(user)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}
