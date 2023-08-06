package main

import (
	"log"

	"periph.io/x/conn/v3/driver/driverreg"
	"periph.io/x/conn/v3/gpio"
	"periph.io/x/conn/v3/gpio/gpioreg"
	"periph.io/x/host/v3"
	"periph.io/x/host/v3/allwinner"
)

func main() {
	_, err := host.Init()
	if err != nil {
		panic(err)
	}

	if allwinner.IsH3() {
		log.Println("H3")
	}

	err = allwinner.PA16.Out(true)
	if err != nil {
		log.Println(err)
	}

	inp := allwinner.PA1

	if _, err := driverreg.Init(); err != nil {
		log.Fatal(err)
	}

	// Use gpioreg GPIO pin registry to find a GPIO pin by name.
	p := gpioreg.ByName("GPIO17")
	if p == nil {
		log.Fatal("Failed to find GPIO17")
	}

	// A pin can be read, independent of its state; it doesn't matter if it is
	// set as input or output.
	log.Printf("%s is %s\n", p, p.Read())

	err = inp.In(gpio.PullDown, gpio.RisingEdge)
	if err != nil {
		log.Println(err)
	}

	log.Printf("PA16 is %s", allwinner.PA16.Read())
	log.Printf("GPIO is %s", p.Read())

	for {
		if inp.Read() {
			log.Println("PA1 is high")
		}
	}
}
