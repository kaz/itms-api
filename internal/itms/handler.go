package itms

import (
	"fmt"
)

type (
	Instrument string
	Mode       string
)

const (
	InstrumentBath         Instrument = "bath"
	InstrumentFloorHeating Instrument = "floor_heating"

	ModeOn  Mode = "on"
	ModeOff Mode = "off"
)

func SwitchInstrumentMode(instrument Instrument, mode Mode) error {
	switch instrument {
	case InstrumentBath:
		switch mode {
		case ModeOn:
			return turnOnBath()
		case ModeOff:
			return turnOffBath()
		}
	case InstrumentFloorHeating:
		switch mode {
		case ModeOn:
			return turnOnFloorHeating()
		case ModeOff:
			return turnOffFloorHeating()
		}

	}
	return fmt.Errorf("no such instrument/action")
}
