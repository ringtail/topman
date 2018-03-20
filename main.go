package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"encoding/json"
	"github.com/ringtail/topman/seaway"
	"os/signal"
	"fmt"
)

var (
	config_file *string
	dingding_token *string
	interval *int
)

func init() {
	config_file = flag.String("config", "topman.conf", "You can specific a config file.")
	dingding_token = flag.String("token","","dingbot token")
	interval = flag.Int("interval",300,"interval time")
}

func main() {
	flag.Parse()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)

	if config_file == nil || dingding_token == nil {
		fmt.Println("token must provide.")
		return
	}
	go Sailing(*config_file)
	<-c
	fmt.Println("kill topman and exit.")
}

func Sailing(config_file string) {
	raw, err := ioutil.ReadFile(config_file)
	if err != nil {
		log.Printf("Failed to read config file, Because of %s", err.Error())
		os.Exit(-1)
	}
	corsairs := make([]*seaway.Corsair, 0)
	json.Unmarshal(raw, &corsairs)

	captain := &seaway.Captain{
		PhoneNumber:*dingding_token,
	}

	topman := &seaway.Topman{
		Interval: *interval,
	}

	if err := topman.OnDuty(corsairs, captain); err != nil {
		log.Println("🔭topman has unexpectedly exited or been terminated. Because of %s", err.Error())
		os.Exit(-1)
	}
}
