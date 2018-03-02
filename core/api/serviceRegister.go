package api

import (
	"container/list"
	"encoding/json"
	"github.com/hashicorp/consul/api"
	"github.com/nightlegend/consulapi/core/data/constdata"
	"github.com/nightlegend/consulapi/core/data/consuldata"
	log "github.com/sirupsen/logrus"
)

var (
	agent    = cli.Agent()
	services []*api.AgentServiceRegistration
)

// RegisterService a new service.
// args:
// 		[id]: unique value.
// 		[name]:	targets name.
// 		[tags]: define a set of labels.
// 		[address]: metrics source address
// 		[port]	metrics	source port
func RegisterService(id string, name string, tags []string, address string, port int, registerType int) bool {
	service := &api.AgentServiceRegistration{
		ID:      id,
		Name:    name,
		Tags:    tags,
		Address: address,
		Port:    port,
	}
	err := agent.ServiceRegister(service)
	if err != nil {
		log.Println(err)
		return false
	}
	if registerType == 0 {
		flag := consuldata.Add(service)
		if flag {
			log.Info("insert to db successful")
		} else {
			log.Info("insert to db failed")
		}
	}

	return true
}

// ServiceDeRegister: delete a registered service.
// 	args:
// 		[serviceId]: register service id.
// 	return: boolean
func ServiceDeRegister(serviceId string) bool {
	err := agent.ServiceDeregister(serviceId)
	if err != nil {
		log.Println(err)
		return false
	}
	res := consuldata.Delete(serviceId)
	if res {
		log.Info("delete service successful")
	} else {
		log.Error("delete service failed")
	}
	return true
}

// GetAllRegisterService: get all register service.
func GetAllRegisterService() map[string]*api.AgentService {
	services, err := agent.Services()
	if err != nil {
		log.Println(err)
	}
	agentService := new(api.AgentService)
	var agentServiceList = list.New()
	for _, v := range services {
		temp, _ := json.Marshal(v)

		err := json.Unmarshal(temp, agentService)
		if err != nil {
			log.Println(err)
		}
		log.Println(agentService.Service)
		agentServiceList.PushBack(agentService)
	}
	log.Println(agentServiceList.Len())
	return services
}

// ReloadData: reload all services from db
func ReloadData() bool {
	services = consuldata.FindAll()
	for _, v := range services {
		err := RegisterService(v.ID, v.Name, v.Tags, v.Address, v.Port, constdata.RELOAD_DATA_TYPE)
		if err {
			log.Info("reload data successful")
		} else {
			log.Error("reload data failed")
			log.Error(err)
			return false
		}
	}
	return true
}
