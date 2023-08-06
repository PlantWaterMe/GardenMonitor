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

	err = allwinner.PA1.In(gpio.Float, gpio.BothEdges)
	if err != nil {
		log.Println(err)
	}

	for {

		log.Printf("PA16 is %s", allwinner.PA16.FastRead())
		log.Printf("PA1 is %s", allwinner.PA1.FastRead())

		time.Sleep(5 * time.Second)
		lvl := allwinner.PA1.FastRead()
		if lvl {
			log.Println("PA1 is high")
		} else {
			log.Println("PA1 is low")
		}
	}
}
