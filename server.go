package main

import (
	"github.com/nightlegend/consulapi/conf"
	"github.com/nightlegend/consulapi/router"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

func init() {
	os.Setenv("ENV", "dev")
	log.Info("Init application configure...")
	execDirAbsPath, _ := os.Getwd()
	log.Info("start init env configure")
	env := os.Getenv("ENV")
	log.Info("You load env is:" + env)

	data, err := ioutil.ReadFile(execDirAbsPath + "/conf/" + env + ".conf.yaml")
	if err != nil {
		log.Errorln(err)
	}
	var config *conf.Config
	err = yaml.Unmarshal([]byte(data), &config)
	if err != nil {
		log.Errorln(err)
	}
	log.Info(config.ConsulUrl)

}
func main() {
	router.Start()
}
