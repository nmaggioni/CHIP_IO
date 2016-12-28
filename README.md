# CHIP_IO
###### A basic CHIP GPIO library for Go that interacts with `sysfs`.

## About root privileges

Accessing the path to control GPIO pins requires root privileges.

If you do not want to run your Go programs as root, put the following files in place on your C.H.I.P. and reboot:

###### /usr/local/bin/change_gpio_perms.sh

```bash
#!/bin/sh

# CHANGE PERMISSIONS AND GROUP FOR THE XIO-P* PINS
/bin/chown -R root:dialout /sys/devices/platform/soc\@01c00000/1c2b400.i2c/i2c-2/2-0038/gpio
/bin/chmod -R ug+rw /sys/devices/platform/soc\@01c00000/1c2b400.i2c/i2c-2/2-0038/gpio

# CHANGE PERMISSIONS AND GROUP FOR THE R8 PINS
/bin/chown -R root:dialout /sys/devices/platform/soc@01c00000/1c20800.pinctrl/gpio
/bin/chmod -R ug+rw /sys/devices/platform/soc@01c00000/1c20800.pinctrl/gpio
```

###### /etc/udev/rules.d/98-gpio-group.rules

```text
# UDEV RULE TO CHANGE GPIO GROUP
# INITIALLY WRITTEN BY IOT_STEVE
# UPDATED BY XTACOCOREX
# THIS REQUIRES change_gpio_perms.sh TO BE INSTALLED IN /usr/local/bin/

SUBSYSTEM=="gpio", PROGRAM="/usr/local/bin/change_gpio_perms.sh"
```

Additionally, append this to `/etc/rc.local` right **before `exit 0` at the end**.

###### /etc/rc.local

```text
chgrp -R dialout /sys/class/gpio
chmod -R g+rw /sys/class/gpio
```

---

**Note:** these scripts come from [xtacocorex's](https://bbs.nextthing.co/t/gpio-handling-as-a-normal-currentUser/4311/4) and [iot_steve's](https://bbs.nextthing.co/t/gpio-handling-as-a-normal-user/4311/7) posts on NextThing's BBS, but were adapted to use the `dialout` group, which the default `chip` currentUser is already part of.

## Usage

Using the library is very similar to the excellent [xtacocorex's CHIP_IO Python library](https://github.com/xtacocorex/CHIP_IO).

Make sure to have completed the [steps above](#about-root-privileges) if you want to avoid running your code as root.

### GPIO Setup

Import the library, and call the `Setup` function, that will initialize the rest of the library and check for permissions:

```go
package main

import "github.com/nmaggioni/CHIP_IO"

func main() {
    CHIP_IO.Setup()
}
```

### GPIO Output

Export the desired pin and set its mode to Output, then write a value to it. Remember to unexport the pin when done.

```go
package main

import (
    "fmt"
    "github.com/nmaggioni/CHIP_IO"
)

func main() {
    CHIP_IO.Setup()
	
    CHIP_IO.Export(CHIP_IO.Lines.XIO_P0)
    CHIP_IO.Mode(CHIP_IO.Lines.XIO_P0, CHIP_IO.Modes.Output)
	
    fmt.Println("Setting XIO_P0 to: " + CHIP_IO.Values.High)
    CHIP_IO.SetValue(CHIP_IO.Lines.XIO_P0, CHIP_IO.Values.High)
    fmt.Println("Setting XIO_P0 to: " + CHIP_IO.Values.Low)
    CHIP_IO.SetValue(CHIP_IO.Lines.XIO_P0, CHIP_IO.Values.Low)
    
    CHIP_IO.Unexport(CHIP_IO.Lines.XIO_P0)
}
```

### GPIO Input

Inputs work similarly to outputs:

```go
package main

import (
    "fmt"
    "github.com/nmaggioni/CHIP_IO"
)

func main() {
    CHIP_IO.Setup()
	
    CHIP_IO.Export(CHIP_IO.Lines.XIO_P0)
    CHIP_IO.Mode(CHIP_IO.Lines.XIO_P0, CHIP_IO.Modes.Input)
	
    fmt.Println("XIO_P0 is: " + CHIP_IO.GetValue(CHIP_IO.Lines.XIO_P0))
	
    CHIP_IO.Unexport(CHIP_IO.Lines.XIO_P0)
}
```

### GPIO Cleanup

To release all the used pins and clean up their state:

```go
package main

import "github.com/nmaggioni/CHIP_IO"

func main() {
    CHIP_IO.UnexportAll()
}
```

## Examples

Examples can be found in the [examples directory](https://github.com/nmaggioni/CHIP_IO/tree/master/examples).

## Unsupported features

+ PWM
+ SoftPWM
+ Interrupts
+ LRADC
+ SPI
+ I2C

### Credits

The CHIP_IO Go library was inspired by the fully featured [xtacocorex's CHIP_IO Python library](https://github.com/xtacocorex/CHIP_IO).

If you are looking for PWM, ADC and other advanced functionality not present in this library, refer to his implementation.