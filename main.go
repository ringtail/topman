package main

import (
	"flag"
	"io/ioutil"
	"os"
	"encoding/json"
	"github.com/ringtail/topman/seaway"
	log "github.com/Sirupsen/logrus"
	"os/signal"
	"fmt"
)

var (
	config_file    *string
	dingding_token *string
	interval       *int
	debug          *bool
)

func init() {
	config_file = flag.String("config", "topman.conf", "You can specific a config file.")
	dingding_token = flag.String("token", "", "dingbot token")
	interval = flag.Int("interval", 30, "interval time")
	debug = flag.Bool("debug", false, "debug mode")
}

func main() {
	flag.Parse()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)

	if config_file == nil || dingding_token == nil {
		fmt.Println("token must provide.")
		return
	}
	if *debug == true {
		log.SetLevel(log.DebugLevel)
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
	corsairs := make([]seaway.Corsair, 0)

	config := &seaway.Config{
		ScreenshotCorsairs: make([]*seaway.ScreenshotCorsair, 0),
		PingCorsairs:       make([]*seaway.PingCorsair, 0),
		TcpCorsairs:        make([]*seaway.TcpCorsair, 0),
	}

	err = json.Unmarshal(raw, &config)

	if err != nil {
		log.Errorf("Failed to Unmarshal your config file %s", err.Error())
		return
	}

	for _, c := range config.ScreenshotCorsairs {
		corsairs = append(corsairs, c)
	}

	for _, c := range config.PingCorsairs {
		corsairs = append(corsairs, c)
	}

	for _, c := range config.TcpCorsairs {
		corsairs = append(corsairs, c)
	}

	captain := &seaway.Captain{
		PhoneNumber: *dingding_token,
	}

	topman := &seaway.Topman{
		Interval: *interval,
	}

	if err := topman.OnDuty(corsairs, captain); err != nil {
		log.Println("🔭topman has unexpectedly exited or been terminated. Because of %s", err.Error())
		os.Exit(-1)
	}
}
