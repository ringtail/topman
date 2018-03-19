package main

import (
	"fmt"
	"os/exec"
	"log"
)

type CorsairInfo struct {
	Msg string
}

type Corsair struct {
	Ip        string
	Threshold int
	Loss      int
}

func (cs *Corsair) Wigwag() (spotted bool, err error) {
	cmd := exec.Command("mtr", "-n -i 0.3 -c 10 -r", cs.Ip)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Failed to pass packet check,Because of %s", err.Error())
		return
	}
	log.Println(out)
	return
}

func (cs *Corsair) Info() (corsairInfo *CorsairInfo, err error) {
	return &CorsairInfo{
		Msg: fmt.Sprintf("Ip:%s, Loss:%s", cs.Ip, cs.Loss),
	}, nil
}
