package sensor

import (
	"os"
	"time"

	"periph.io/x/conn/v3/gpio"
	"periph.io/x/host/v3/allwinner"
)

type Level bool

const (
	Empty    Level = false // Depth sensor shows is empty
	NotEmpty Level = true  // Depth sensor does not show empty
)

type Depth struct {
	PowerPin   gpio.PinIO
	MeasurePin gpio.PinIO
}

func New(PowerPin *allwinner.Pin, MeasurePin *allwinner.Pin) *Depth {

	// set PowerPin to output and High
	PowerPin.Out(true)

	// set MeasurePin to input
	MeasurePin.In(gpio.Float, gpio.BothEdges)

	return &Depth{
		PowerPin:   PowerPin,
		MeasurePin: MeasurePin,
	}
}

func (d *Depth) Probe() Level {
	return d.ProbeFor(5 * time.Millisecond)
}

func (d *Depth) ProbeFor(duration time.Duration) Level {
	debug := os.Getenv("DEBUG")
	if debug == "true" {
		return NotEmpty
	}

	if d.MeasurePin.WaitForEdge(duration) {
		return NotEmpty
	}
	return Empty
}
