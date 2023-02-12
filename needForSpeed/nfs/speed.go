package speed

type Car struct {
	Speed        int
	BatteryDrain int
	Battery      int
	Distance     int
}

// NewCar creates a new remote controlled car with full Battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	return Car{
		Speed:        speed,
		BatteryDrain: batteryDrain,
		Battery:      100,
		Distance:     0,
	}
}

type Track struct {
	distance int
}

// NewTrack creates a new track
func NewTrack(distance int) Track {
	return Track{
		distance: distance,
	}
}

// Drive drives the car one time. If there is not enough Battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {
	if car.BatteryDrain > car.Battery {
		return car
	}
	car.Battery -= car.BatteryDrain
	car.Distance += car.Speed

	return car
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	check := track.distance * car.BatteryDrain / car.Speed

	if check > car.Battery {
		return false
	} else {
		return true
	}
}
