package seaway

import (
	"testing"
)

func TestTcpCorsair(t *testing.T) {
	tc := &TcpCorsair{
		Name: "Golang",
		Host: "127.0.0.1:80",
	}
	spotted, err := tc.Wigwag()
	if err != nil {
		t.Errorf("error: %s", err.Error())
		return
	}
	t.Log(spotted)
}
