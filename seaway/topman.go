package seaway

import (
	"time"
	"fmt"
)

const (
	DEFAULT_TIME_TICKER = 30 * time.Second
)

type Topman struct {
	Corsairs []*Corsair
	Captain  *Captain
}

func (tm *Topman) Attention(){
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

func (tm *Topman) OnDuty(corsairs []*Corsair, captain *Captain) (err error) {
	tm.Attention()
	tm.Corsairs = corsairs
	tm.Captain = captain
	ticker := time.NewTicker(DEFAULT_TIME_TICKER)
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

func (tm *Topman) Lookout(corsair *Corsair) (err error) {
	spotted, err := corsair.Wigwag()
	if err != nil {
		return err
	}
	if spotted == true {
		corsairInfo, err := corsair.Info()
		if err != nil {
			return err
		}
		return tm.Report(corsairInfo)
	}
	return
}

func (tm *Topman) LookoutAround() (err error) {
	if len(tm.Corsairs) != 0 {
		for _, c := range tm.Corsairs {
			tm.Lookout(c)
		}
	}
	return
}
