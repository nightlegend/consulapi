package api

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"os"
)

// NewCli is new a consul client.
func NewCli() *api.Client {
	config := api.DefaultConfig()
	fmt.Println(os.Getenv("CONSUL_API"))
	config.Address = os.Getenv("CONSUL_API")
	client, err := api.NewClient(config)
	if err != nil {
		panic(err)
	}
	return client
}
