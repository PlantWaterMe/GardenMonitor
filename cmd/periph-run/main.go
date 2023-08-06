package main

import (
	"log"

	"periph.io/x/host/v3"
	"periph.io/x/host/v3/allwinner"
	"periph.io/x/host/v3/bcm283x"
)

func main() {
	_, err := host.Init()
	if err != nil {
		panic(err)
	}

	if allwinner.IsH3() {
		log.Println("H3")
	}

	if bcm283x.GPIO18.Read() {
		log.Println("GPIO18 is high")
	} else {
		log.Println("GPIO18 is low")
	}

	bcm283x.GPIO18.Out(true)

	if bcm283x.GPIO1.Read() {
		log.Println("GPIO1 is high")
	} else {
		log.Println("GPIO1 is low")
	}

	bcm283x.GPIO1.Out(true)

}
