package seaway

import (
	"net"
	"time"
	"fmt"
)

type TcpCorsair struct {
	Corsair
	Name       string `json:"name"`
	Host       string `json:"host"`
	blocked    bool
	blockedMsg string
}

func (tc *TcpCorsair) Wigwag() (spotted bool, err error) {
	_, err = net.DialTimeout("tcp", tc.Host, 3*time.Second)
	if err != nil {
		spotted = true
		tc.blocked = true
		tc.blockedMsg = err.Error()
	}
	return spotted, nil
}

func (tc *TcpCorsair) GetName() string {
	return tc.Name
}

func (tc *TcpCorsair) Msg() (info *CorsairInfo, err error) {
	return &CorsairInfo{
		Msg: fmt.Sprintf("Name: %s, Host: %s, Connection Status: %s \n", tc.Name, tc.Host, tc.blockedMsg),
	}, nil
}
