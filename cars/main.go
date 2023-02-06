package main

import (
	"cars/cars"
	"fmt"
)

func main() {
	//expects 5
	fmt.Println(cars.CalculateWorkingCarsPerMinute(420, 80))
	//expects 3990000
	fmt.Println(cars.CalculateCost(420))
	//expects 336
	fmt.Println(cars.CalculateWorkingCarsPerHour(420, 80))
}
