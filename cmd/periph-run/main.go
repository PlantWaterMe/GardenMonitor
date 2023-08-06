package main

import (
	"log"

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

}
