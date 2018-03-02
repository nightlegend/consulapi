package api

import (
	"github.com/hashicorp/consul/api"
	log "github.com/sirupsen/logrus"
	"os"
)

// NewCli is new a consul client.
func NewCli() *api.Client {
	config := api.DefaultConfig()
	log.Info("new consul api...")
	log.Info(os.Getenv("CONSUL_API"))
	log.Info(os.Getenv("ENV"))
	config.Address = os.Getenv("CONSUL_API")
	client, err := api.NewClient(config)
	if err != nil {
		panic(err)
	}
	return client
}
