package main

import (
	"time"

	"periph.io/x/conn/v3/gpio"
	"periph.io/x/host/v3"
	"periph.io/x/host/v3/allwinner"
)

func main() {
	host.Init()
	t := time.NewTicker(500 * time.Millisecond)
	for l := gpio.Low; ; l = !l {
		allwinner.PD14.Out(l)
		<-t.C
	}
}