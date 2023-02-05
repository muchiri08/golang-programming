package main

import (
	"fmt"
	"lasagna/lasandra"
)

func main() {

	fmt.Println(lasandra.PreparationTime(2))

	fmt.Println(lasandra.RemainingOvenTime(30))

	fmt.Println(lasandra.ElapsedTime(3, 20))
}
