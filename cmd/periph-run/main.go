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

	bcm283x.GPIO18.Out(true)

}
