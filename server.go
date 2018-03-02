package main

import (
	"github.com/nightlegend/consulapi/conf"
	"github.com/nightlegend/consulapi/router"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

// Init environment
func init() {
	log.Info("Init application configure...")

	var config *conf.Config
	execDirAbsPath, _ := os.Getwd()
	os.Setenv("CONSUL_API", "10.222.49.65:8500")
	env := os.Getenv("ENV")
	if env == "" {
		env = "dev"
	}
	data, err := ioutil.ReadFile(execDirAbsPath + "/conf/" + env + ".conf.yaml")
	if err != nil {
		log.Errorln(err)
	}
	err = yaml.Unmarshal([]byte(data), &config)
	if err != nil {
		log.Errorln(err)
	}

	log.Info("Load environment is ", env)
}

// Main function
func main() {
	router.Start()
}
