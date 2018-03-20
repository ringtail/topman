package seaway

import (
	"fmt"
	"github.com/sparrc/go-ping"
	"time"
)

type CorsairInfo struct {
	Msg string
}

type Corsair struct {
	Name string `json:"name"`
	Ip        string `json:"ip"`
	Threshold int	`json:"threshold"`
	Loss      int
}

func (cs *Corsair) Wigwag() (spotted bool, err error) {
	pinger, err := ping.NewPinger(cs.Ip)
	if err != nil {
		return
	}
	pinger.Count = 3
	pinger.Timeout = time.Second * 5
	//pinger.SetPrivileged(true)
	pinger.Run() // blocks until finished
	stats := pinger.Statistics() // get send/receive/rtt stats
	if int(stats.PacketLoss) > cs.Threshold {
		cs.Loss = int(stats.PacketLoss)
		return true ,nil
	}else{
		cs.Loss = 0
	}
	return
}

func (cs *Corsair) Info() (corsairInfo *CorsairInfo, err error) {
	return &CorsairInfo{
		Msg: fmt.Sprintf("Name: %s, Ip: %s, PacketLoss: %d%%",cs.Name, cs.Ip, cs.Loss),
	}, nil
}
