package main

import (
	"fmt"
	speed "needForSpeed/nfs"
)

func main() {
	fmt.Println("good morning")

	var s = speed.Car{
		Speed:        5,
		BatteryDrain: 3,
		Battery:      100,
		Distance:     1,
	}
	fmt.Println(s.Speed + s.Distance)
	fmt.Println(speed.Drive(s))
}
