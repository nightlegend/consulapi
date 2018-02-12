package api

import (
	"container/list"
	"encoding/json"
	"github.com/hashicorp/consul/api"
	"log"
)

var (
	agent = cli.Agent()
)

// RegisterService a new service.
// args:
// 		[id]: unique value.
// 		[name]:	targets name.
// 		[tags]: define a set of labels.
// 		[address]: metrics source address
// 		[port]	metrics	source port
func RegisterService(id string, name string, tags []string, address string, port int) bool {
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
