package api

import (
	"encoding/json"
	"fmt"
	"github.com/hashicorp/consul/api"
	"log"
)

// Get a handle to the KV API
var cli = NewCli()
var kv = cli.KV()
var helth = cli.Health()

// Put a key value to kv store.
func Put(key string, value string) bool {
	// PUT a new KV pair
	p := &api.KVPair{Key: key, Value: []byte(value)}
	_, err := kv.Put(p, nil)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

// Get a key from store .
func Get(key string) string {
	// Lookup the pair
	pair, _, err := kv.Get(key, nil)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(pair.Value)
}

// List all kv in kv store.
func ListAllKV() string {
	kvList, _, err := kv.List("", nil)
	if err != nil {
		log.Println(err)
		return ""
	}
	jsonString, _ := json.Marshal(kvList)
	log.Println("code inside......")
	log.Println(string(jsonString))
	return string(jsonString)
}
