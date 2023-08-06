package main

import (
	"fmt"
	"log"

	"periph.io/x/conn/v3/driver/driverreg"
	"periph.io/x/conn/v3/gpio/gpioreg"
	"periph.io/x/host/v3"
	"periph.io/x/host/v3/allwinner"
)

func main() {
	// Make sure periph is initialized.
	state, err := host.Init()
	if err != nil {
		log.Fatalf("failed to initialize periph: %v", err)
	}

	// Prints the loaded driver.
	fmt.Printf("Using drivers:\n")
	for _, driver := range state.Loaded {
		fmt.Printf("- %s\n", driver)
	}

	// Prints the driver that were skipped as irrelevant on the platform.
	fmt.Printf("Drivers skipped:\n")
	for _, failure := range state.Skipped {
		fmt.Printf("- %s: %s\n", failure.D, failure.Err)
	}

	// Having drivers failing to load may not require process termination. It
	// is possible to continue to run in partial failure mode.
	fmt.Printf("Drivers failed to load:\n")
	for _, failure := range state.Failed {
		fmt.Printf("- %s: %v\n", failure.D, failure.Err)
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

	/*
		err = inp.In(gpio.PullDown, gpio.RisingEdge)
		if err != nil {
			log.Println(err)
		}

		log.Printf("PA16 is %s", allwinner.PA16.Read())
		log.Printf("GPIO17 is %s", p.Read())
	*/
	for {
		if inp.Read() {
			log.Println("PA1 is high")
		}
		if p.Read() {
			log.Println("GPIO17 is high")
		}
	}
}
