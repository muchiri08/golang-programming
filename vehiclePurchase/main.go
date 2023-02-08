package main

import (
	"fmt"
	purchase "vehiclePurchase/vp"
)

func main() {
	fmt.Println(purchase.NeedsLicense("truck"))
	fmt.Println(purchase.ChooseVehicle("BMW", "Chevrolet"))
	fmt.Println(purchase.CalculateResellPrice(55000, 10))
}
