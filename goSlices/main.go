package main

import "fmt"

func main() {

	intSlice := make([]int, 0, 20)
	intSlice = append(intSlice, 23, 23, 35, 67)

	for index := 0; index < len(intSlice); index++ {
		fmt.Print(intSlice[index], " ")
	}
	fmt.Println("\nCapacity: ", cap(intSlice))
	fmt.Println("length: ", len(intSlice))

	//emptying an existing slice
	//intSlice = nil
	//fmt.Println("length after emptying: ", len(intSlice))
	fmt.Println()
	reSlice := intSlice[2:]
	fmt.Println(reSlice)
	fmt.Println("\nCapacity: ", cap(reSlice))
	fmt.Println("length: ", len(reSlice))

	fmt.Println()
	s1 := make([]int, 5)
	reSlice1 := s1[1:3]
	fmt.Println(s1)
	fmt.Println(reSlice1)
	reSlice1[0] = -100
	reSlice1[1] = 123456
	fmt.Println(s1)
	fmt.Println(reSlice1)

}
