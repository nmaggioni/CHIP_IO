package main

import (
	"fmt"
	"github.com/nmaggioni/CHIP_IO"
	"time"
)

func main() {
	// Initialize the library
	CHIP_IO.Setup()

	// Export the pin
	CHIP_IO.Export(CHIP_IO.Lines.XIO_P0)
	// Set the pin as an output
	CHIP_IO.Mode(CHIP_IO.Lines.XIO_P0, CHIP_IO.Modes.Output)

	// Blink the pin
	for i := 0; i < 3; i++ {
		enableXIO_P0(true)
		time.Sleep(1 * time.Second)
		enableXIO_P0(false)
		time.Sleep(1 * time.Second)
	}

	// Unexport all the pins - applies `Unexport` to all the previously exported pins.
	CHIP_IO.UnexportAll()
}

func enableXIO_P0(enable bool) {
	if enable {
		fmt.Println("Setting XIO_P0 to: " + CHIP_IO.Values.High)
		// Enable the pin
		CHIP_IO.SetValue(CHIP_IO.Lines.XIO_P0, CHIP_IO.Values.High)
	} else {
		fmt.Println("Setting XIO_P0 to: " + CHIP_IO.Values.Low)
		// Disable the pin
		CHIP_IO.SetValue(CHIP_IO.Lines.XIO_P0, CHIP_IO.Values.Low)
	}
}
