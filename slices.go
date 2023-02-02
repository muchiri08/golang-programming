package main

import "fmt"

func main(){

	fmt.Println("Slice append")

	slice1 := []int{1,2,3}
	slice2 := append(slice1, 4, 5)

	fmt.Println(slice1, slice2)

	fmt.Println("Slice copy")

	slice3 := []int{1,2,3}
	slice4 := make([]int, 2)
	copy(slice4, slice3)

	fmt.Println(slice3, slice4)

}
