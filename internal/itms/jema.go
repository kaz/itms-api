package itms

import (
	"fmt"
)

func jemaInstrument(name string, index int, action string) error {
	if err := jemaControl(action, index); err != nil {
		return fmt.Errorf("turning %s %s: %w", action, name, err)
	}
	return nil

}

func bath(action string) error {
	return jemaInstrument("bath", 1, action)
}
func turnOnBath() error {
	return bath("ON")
}
func turnOffBath() error {
	return bath("OFF")
}

func floorHeating(action string) error {
	return jemaInstrument("floor heating", 2, action)
}
func turnOnFloorHeating() error {
	return floorHeating("ON")
}
func turnOffFloorHeating() error {
	return floorHeating("OFF")
}
