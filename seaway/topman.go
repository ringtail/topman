package seaway

import (
	"time"
	"fmt"
	log "github.com/Sirupsen/logrus"
)

type Topman struct {
	Interval int
	Corsairs []Corsair
	Captain  *Captain
	IsLock   bool
}

func (tm *Topman) Attention() {
	attention_msg := `

                  ___                     ___          ___          ___              
      ___        /  /\        ___        /  /\        /  /\        /  /\             
     /__/\      /  /::\      /  /\      /  /::|      /  /::\      /  /::|            
     \  \:\    /  /:/\:\    /  /::\    /  /:|:|     /  /:/\:\    /  /:|:|            
      \__\:\  /  /:/  \:\  /  /:/\:\  /  /:/|:|__  /  /::\ \:\  /  /:/|:|__          
      /  /::\/__/:/ \__\:\/  /::\ \:\/__/:/_|::::\/__/:/\:\_\:\/__/:/ |:| /\         
     /  /:/\:\  \:\ /  /:/__/:/\:\_\:\__\/  /~~/:/\__\/  \:\/:/\__\/  |:|/:/         
    /  /:/__\/\  \:\  /:/\__\/  \:\/:/     /  /:/      \__\::/     |  |:/:/          
   /__/:/      \  \:\/:/      \  \::/     /  /:/       /  /:/      |__|::/           
   \__\/        \  \::/        \__\/     /__/:/       /__/:/       /__/:/            
                 \__\/                   \__\/        \__\/        \__\/
  ======================================================================
                     Yes Commander. I am on duty! 
  ======================================================================
	`
	fmt.Println(attention_msg)
}

func (tm *Topman) OnDuty(corsairs []Corsair, captain *Captain) (err error) {
	tm.Attention()
	tm.Corsairs = corsairs
	tm.Captain = captain
	log.Debugf("Topman is onDuty with interval %d\n", tm.Interval)
	tm.LookoutAround()
	ticker := time.NewTicker(time.Second * time.Duration(tm.Interval))
	quit := make(chan struct{})
	for {
		select {
		case <-ticker.C:
			tm.LookoutAround()
		case <-quit:
			ticker.Stop()
			return
		}
	}
	return
}

func (tm *Topman) Report(corsairInfo *CorsairInfo) (err error) {
	return tm.Captain.Dispose(corsairInfo)
}

func (tm *Topman) Lookout(corsair Corsair) (err error) {
	spotted, err := corsair.Wigwag()
	if err != nil {
		return err
	}
	if spotted == true {
		corsairInfo, err := corsair.Msg()
		if err != nil {
			return err
		}
		return tm.Report(corsairInfo)
	}
	return
}

func (tm *Topman) LookoutAround() (err error) {
	log.Debugf("=============================")
	log.Debugf("Topman is looking out around.")
	if tm.IsLock == true {
		log.Warn("Last duty is not finished, please increase the value of interval.")
		return
	}
	tm.IsLock = true
	if len(tm.Corsairs) != 0 {
		for index, _ := range tm.Corsairs {
			c := tm.Corsairs[index]
			err = tm.Lookout(c)
			if err != nil {
				log.Warnf("Topman failed to lookout %s,Because of %s", c.GetName(), err.Error())
			}
		}
	}
	tm.IsLock = false
	return
}
