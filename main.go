package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"encoding/json"
)

var (
	config_file *string
)

func init() {
	config_file = flag.String("config", "topman.conf", "You can specific a config file.")
}

func main() {
	flag.Parse()
	if config_file == nil {
		return
	}
	Sailing(*config_file)
}

func Sailing(config_file string) {
	raw, err := ioutil.ReadFile(config_file)
	if err != nil {
		log.Printf("Failed to read config file, Because of %s", err.Error())
		os.Exit(-1)
	}
	corsairs := make([]*Corsair, 0)
	json.Unmarshal(raw, &corsairs)

	captain := &Captain{}

	topman := &Topman{}

	if err := topman.OnDuty(corsairs, captain); err != nil {
		log.Println("🔭topman has unexpectedly exited or been terminated. Because of %s", err.Error())
		os.Exit(-1)
	}
}
