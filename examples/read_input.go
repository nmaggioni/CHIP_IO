package main

import (
	"fmt"
	"github.com/nmaggioni/CHIP_IO"
)

func main() {
	// Initialize the library
	CHIP_IO.Setup()

	// Export the pin
	CHIP_IO.Export(CHIP_IO.Lines.XIO_P0)
	// Set the pin as an input
	CHIP_IO.Mode(CHIP_IO.Lines.XIO_P0, CHIP_IO.Modes.Input)

	// Read the pin's value
	fmt.Println("XIO_P0 is: " + CHIP_IO.GetValue(CHIP_IO.Lines.XIO_P0))

	// Unexport the pin
	CHIP_IO.Unexport(CHIP_IO.Lines.XIO_P0)
}
