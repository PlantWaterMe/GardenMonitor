package main

import (
	"log"
	"time"

	"periph.io/x/conn/v3/gpio"
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

	err = inp.In(gpio.PullDown, gpio.RisingEdge)
	if err != nil {
		log.Println(err)
	}

	for {

		log.Printf("PA16 is %s", allwinner.PA16.Read())
		log.Printf("PA1 is %s", inp.Read())

		time.Sleep(5 * time.Second)
		lvl := inp.Read()
		if lvl {
			log.Println("PA1 is high")
		} else {
			log.Println("PA1 is low")
		}
	}
}
