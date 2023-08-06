package main

import (
	"periph.io/x/conn/v3/gpio"
	"periph.io/x/host/v3"
	"periph.io/x/host/v3/bcm283x"
)

func main() {
	host.Init()

	bcm283x.GPIO18.Out(gpio.High)
}
