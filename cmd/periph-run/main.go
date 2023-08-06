package main

import (
	"periph.io/x/conn/v3/gpio"
	"periph.io/x/host/v3"
	"periph.io/x/host/v3/allwinner"
)

func main() {
	host.Init()

	allwinner.PD14.Out(gpio.High)
}
