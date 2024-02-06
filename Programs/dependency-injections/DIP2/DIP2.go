package main

import "fmt"

type ToyRobot struct {
	PowerSource Battery
}

type Battery interface {
	ProvidePower() string
}

type RegularBattery struct{}

func (b RegularBattery) ProvidePower() string {
	return "Powering toy with regular batteries"
}

type SolarBattery struct{}

func (b SolarBattery) ProvidePower() string {
	return "Powering toy with solar energy"
}

func PowerToyRobot(robot *ToyRobot, battery Battery) {
	robot.PowerSource = battery
	fmt.Println(battery.ProvidePower())

}

func main() {
	// Create a ToyRobot instance.
	robot := &ToyRobot{}

	// Use a RegularBattery to power the ToyRobot.
	PowerToyRobot(robot, RegularBattery{})

	// Change the battery type without modifying PowerToyRobot.
	PowerToyRobot(robot, SolarBattery{})
}
