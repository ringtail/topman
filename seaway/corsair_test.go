package seaway

import (
	"testing"
)

var (
	default_corsair *Corsair
)

func init() {
	default_corsair = &Corsair{
		Ip:        "8.8.8.8",
		Threshold: 30,
	}
}

func TestCorsair(t *testing.T) {
	default_corsair.Wigwag()
}
