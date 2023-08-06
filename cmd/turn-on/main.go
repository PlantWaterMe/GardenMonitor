package main

import (
	"fmt"
	"time"

	"github.com/stianeikeland/go-rpio"
)

func main() {
	err := rpio.Open()
	if err != nil {
		panic(fmt.Sprint("unable to open gpio", err.Error()))
	}
	defer rpio.Close()

	pin := rpio.Pin(16)
	rpio.PinMode(pin, rpio.Output)

	pin.High()
	time.Sleep(5 * time.Second)
	pin.Low()
}
