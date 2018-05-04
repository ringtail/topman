package seaway

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/sparrc/go-ping"
	"time"
)

type PingCorsair struct {
	Corsair
	Name      string `json:"name"`
	Host      string `json:"host"`
	Threshold int    `json:"threshold"`
	Loss      int
}

func (cs *PingCorsair) Wigwag() (spotted bool, err error) {
	pinger, err := ping.NewPinger(cs.Host)
	if err != nil {
		return
	}
	pinger.Count = 3
	pinger.Timeout = time.Second * 5
	pinger.SetPrivileged(true)
	pinger.Run()                 // blocks until finished
	stats := pinger.Statistics() // get send/receive/rtt stats
	log.Debugf("%s packet loss is %f, current threshold is %d", cs.Name, stats.PacketLoss, cs.Threshold)
	if int(stats.PacketLoss) > cs.Threshold {
		cs.Loss = int(stats.PacketLoss)
		return true, nil
	} else {
		cs.Loss = 0
	}
	return
}

func (cs *PingCorsair) GetName() string {
	return cs.Name
}

func (cs *PingCorsair) Msg() (corsairInfo *CorsairInfo, err error) {
	return &CorsairInfo{
		Msg: fmt.Sprintf("Name: %s, Ip: %s, PacketLoss: %d%%", cs.Name, cs.Host, cs.Loss),
	}, nil
}
