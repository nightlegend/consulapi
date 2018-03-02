package consuldata

import (
	"github.com/hashicorp/consul/api"
	"github.com/nightlegend/consulapi/core/utils"
	log "github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var services []*api.AgentServiceRegistration

func FindAll() []*api.AgentServiceRegistration {
	session := utils.MogConn()
	defer session.Close()
	c := session.DB("consul").C("services")
	err := c.Find(nil).All(&services)
	if err != nil {
		log.Println(err)
	}
	log.Info(len(services))
	return services
}

func Add(service *api.AgentServiceRegistration) bool {
	session := utils.MogConn()
	defer session.Close()
	c := session.DB("consul").C("services")
	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	err := c.Insert(service)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

func Delete(id string) bool {
	session := utils.MogConn()
	defer session.Close()
	c := session.DB("consul").C("services")
	err := c.Remove(bson.M{"id": id})
	if err != nil {
		log.Error(err)
		return false
	}
	return true
}
