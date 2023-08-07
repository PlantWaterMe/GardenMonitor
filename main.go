package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/PlantWaterMe/GardenMonitor/sensor"
	"periph.io/x/host/v3"
	"periph.io/x/host/v3/allwinner"
)

func main() {

	err := Init()
	if err != nil {
		log.Fatal(err)
	}

	ds := sensor.New(allwinner.PA16, allwinner.PA1)

	for {
		if ds.Probe() == sensor.NotEmpty {
			fmt.Println("Not Empty")
		} else {
			fmt.Println("Empty")
		}
	}
}

func Init() error {
	// Make sure periph is initialized.
	state, err := host.Init()
	if err != nil {
		return errors.New(fmt.Sprintf("failed to initialize periph: %v", err))
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

	return nil
}
