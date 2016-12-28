package CHIP_IO

import (
	"log"
	"strconv"
)

// Modes holds possible states for a pin's direction: either input or output.
var Modes = struct {
	Input  string
	Output string
}{"in", "out"}

// Values holds possible states for a pin's value (in either input or output): high and low.
var Values = struct {
	High string
	Low  string
}{"high", "low"}

type lines struct {
	PWM0      int
	LCD_D13   int
	CSIHSYNC  int
	XIO_P0    int
	AP_EINT3  int
	LCD_D14   int
	CSIVSYNC  int
	XIO_P1    int
	TWI1_SCK  int
	LCD_D15   int
	CSID0     int
	XIO_P2    int
	TWI1_SDA  int
	LCD_D18   int
	CSID1     int
	XIO_P3    int
	TWI2_SCK  int
	LCD_D19   int
	CSID2     int
	XIO_P4    int
	TWI2_SDA  int
	LCD_D20   int
	CSID3     int
	XIO_P5    int
	LCD_D2    int
	LCD_D21   int
	CSID4     int
	XIO_P6    int
	LCD_D3    int
	LCD_D22   int
	CSID5     int
	XIO_P7    int
	LCD_D4    int
	LCD_D23   int
	CSID6     int
	LCD_D5    int
	LCD_CLK   int
	CSID7     int
	LCD_D6    int
	LCD_DE    int
	AP_EINT1  int
	LCD_D7    int
	LCD_HSYNC int
	UART1_TX  int
	LCD_D10   int
	LCD_VSYNC int
	UART1_RX  int
	LCD_D11   int
	CSIPCK    int
	LCD_D12   int
	CSICK     int
}

/*
 The Lines struct contains pin names associated to line numbers for ease of use.
 Be sure to call Setup() first to detect the current kernel's base number for XIO pins.
*/
var Lines lines

/*
 Setup detects the base number for XIO_P[0-7] pins from the current kernel's interface.
 Root privileges are needed to access the "/sys/class/gpio" path!
 Alternatively, a udev rule can be installed: https://bbs.nextthing.co/t/gpio-handling-as-a-normal-user/4311/4
*/
func Setup() {
	checkRootOrGroup()

	labelFile, err := shell("grep -l pcf8574a /sys/class/gpio/*/*label")
	if err != nil {
		log.Fatal("Unable to detect base number for XIO (unable to locate label file): " + err.Error())
	}
	baseFile, err := shell("dirname " + labelFile)
	if err != nil {
		log.Fatal("Unable to detect base number for XIO (unable to locate base file): " + err.Error())
	}
	baseFile += "/base"
	base, err := shell("cat " + baseFile)
	if err != nil {
		log.Fatal("Unable to detect base number for XIO (unable to read base file): " + err.Error())
	}
	baseInt, err := strconv.Atoi(base)
	if err != nil {
		log.Fatal("Unable to detect base number for XIO (unable to parse base file content): " + err.Error())
	}

	Lines = lines{
		PWM0:      34,
		LCD_D13:   109,
		CSIHSYNC:  130,
		XIO_P0:    baseInt,
		AP_EINT3:  35,
		LCD_D14:   110,
		CSIVSYNC:  131,
		XIO_P1:    baseInt + 1,
		TWI1_SCK:  47,
		LCD_D15:   111,
		CSID0:     132,
		XIO_P2:    baseInt + 2,
		TWI1_SDA:  48,
		LCD_D18:   114,
		CSID1:     133,
		XIO_P3:    baseInt + 3,
		TWI2_SCK:  49,
		LCD_D19:   115,
		CSID2:     134,
		XIO_P4:    baseInt + 4,
		TWI2_SDA:  50,
		LCD_D20:   116,
		CSID3:     135,
		XIO_P5:    baseInt + 5,
		LCD_D2:    98,
		LCD_D21:   117,
		CSID4:     136,
		XIO_P6:    baseInt + 6,
		LCD_D3:    99,
		LCD_D22:   118,
		CSID5:     137,
		XIO_P7:    baseInt + 7,
		LCD_D4:    100,
		LCD_D23:   119,
		CSID6:     138,
		LCD_D5:    101,
		LCD_CLK:   120,
		CSID7:     139,
		LCD_D6:    102,
		LCD_DE:    121,
		AP_EINT1:  193,
		LCD_D7:    103,
		LCD_HSYNC: 122,
		UART1_TX:  195,
		LCD_D10:   106,
		LCD_VSYNC: 123,
		UART1_RX:  196,
		LCD_D11:   107,
		CSIPCK:    128,
		LCD_D12:   108,
		CSICK:     129,
	}
}
