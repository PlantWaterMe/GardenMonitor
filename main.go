package main

import (
	"log"

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

	err = inp.In(gpio.PullDown, gpio.BothEdges)
	if err != nil {
		log.Println(err)
	}

	log.Printf("PA16 is %s", allwinner.PA16.Read())
	log.Printf("PA1 is %s", inp.Read())

	for {
		inp.WaitForEdge(-1)
		log.Printf("PA1 is %s", inp.Read())
		log.Printf("PA1 is %s", inp.FastRead().String())
	}
}
