package main

import (
	"github.com/nmaggioni/CHIP_IO"
	"time"
)

func main() {
	// Declare instance
	var statLed CHIP_IO.StatLED

	// Blink the LED
	for i := 0; i < 3; i++ {
		// Turn it ON
		statLed.Enable()
		time.Sleep(500 * time.Millisecond)
		// Turn it OFF
		statLed.Disable()
		time.Sleep(250 * time.Millisecond)
	}
}
