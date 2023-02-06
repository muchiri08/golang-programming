package main

import (
	"AnnalynInfiltration/annalyn"
	"fmt"
)

func main() {
	fmt.Println(annalyn.CanFastAttack(true))
	fmt.Println(annalyn.CanSpy(false, false, false))
	fmt.Println(annalyn.CanSignalPrisoner(true, true))
	fmt.Println(annalyn.CanFreePrisoner(false, false, true, true))
}
