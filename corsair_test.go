package main

import (
	"testing"
)

var (
	default_corsair *Corsair
)

func init() {
	default_corsair = &Corsair{
		Ip:        "aliyun.com",
		Threshold: 30,
	}
}

func TestCorsair(t *testing.T) {
	default_corsair.Wigwag()
}
