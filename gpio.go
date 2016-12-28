package CHIP_IO

import (
	"log"
	"path/filepath"
	"regexp"
	"strconv"
)

/*
 Export makes a pin available for use by the user.

 A line number from the Lines struct is to be given. It is common rule to unexport the pin when done.
*/
func Export(line int) {
	shell("echo " + strconv.Itoa(line) + " >/sys/class/gpio/export")
}

/*
 Unexport frees a pin from the user's control.

 A line number from the Lines struct is to be given.
*/
func Unexport(line int) {
	shell("echo " + strconv.Itoa(line) + " >/sys/class/gpio/unexport")
}

// UnexportAll unexports all of the exported pins. See Unexport.
func UnexportAll() {
	matches, err := filepath.Glob("/sys/class/gpio/gpio[0-9]*")
	if err != nil {
		log.Fatal("Unable to unexport all: error in globbing.")
	}
	if len(matches) == 0 {
		return
	}

	for _, match := range matches {
		r, _ := regexp.Compile("^[^0-9]*")
		pinNumber, _ := strconv.Atoi(r.ReplaceAllString(match, ""))
		Unexport(pinNumber)
	}
}

/*
 Mode sets a given pin as an input or as an output.

 A line number from the Lines struct is to be given, and a mode from the Modes struct.
*/
func Mode(line int, mode string) {
	shell("echo '" + mode + "' >/sys/class/gpio/gpio" + strconv.Itoa(line) + "/direction")
}

/*
 SetValue sets a given pin to a given output value (the pin must be set as an output for this to work).

 A line number from the Lines struct is to be given, and a value from the Values struct.
*/
func SetValue(line int, value string) {
	shell("echo '" + value + "' >/sys/class/gpio/gpio" + strconv.Itoa(line) + "/value")
}

/*
 GetValue reads the current value of a given pin (the pin must be set as an input for this to work).

 A line number from the Lines struct is to be given.
 A string is returned, as one of the Values struct's elements.

 If the read value is not mappable, a fatal error is logged.
*/
func GetValue(line int) string {
	valueString, _ := shell("cat /sys/class/gpio/gpio" + strconv.Itoa(line) + "/value")
	switch valueString {
	case "0":
		return Values.Low
		break
	case "1":
		return Values.High
		break
	default:
		log.Fatal("Unexpected value for GPIO read: " + valueString)
	}
	return ""
}

/*
 Invert reverses the value that a pin reads (high becomes low and vice versa).

 A line number from the Lines struct is to be given, and a bool that indicates if the pin must be inverted or not.
*/
func Invert(line int, invert bool) {
	var active_low string
	if invert {
		active_low = "1"
	} else {
		active_low = "0"
	}
	shell("echo " + active_low + " >/sys/class/gpio/gpio" + strconv.Itoa(line) + "/active_low")
}

/*
 IsInverted checks if a pin is inverted.

 A line number from the Lines struct is to be given.
 A bool is returned: true if the pin is actually inverted, false if not.
*/
func IsInverted(line int) bool {
	valueString, _ := shell("cat /sys/class/gpio/gpio" + strconv.Itoa(line) + "/active_low")
	return valueString != "0"
}
