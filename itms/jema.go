package itms

import "fmt"

func jemaInstrument(name string, index int, action string) error {
	if err := jemaControl(action, index); err != nil {
		return fmt.Errorf("turning %s %s: %w", action, name, err)
	}
	return nil

}

func bath(action string) error {
	return jemaInstrument("bath", 1, action)
}
func TurnOnBath() error {
	return bath("ON")
}
func TurnOffBath() error {
	return bath("OFF")
}

func floorHeating(action string) error {
	return jemaInstrument("floor heating", 2, action)
}
func TurnOnFloorHeating() error {
	return floorHeating("ON")
}
func TurnOffFloorHeating() error {
	return floorHeating("OFF")
}
