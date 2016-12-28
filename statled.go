package CHIP_IO

type StatLED struct{}

func (StatLED) Enable() {
	shell("i2cset -y -f 0 0x34 0x93 0x1")
}

func (StatLED) Disable() {
	shell("i2cset -y -f 0 0x34 0x93 0x0")
}
