package main

import "fmt"

//Data types exercise

func main() {
	fmt.Println("32132 x 42452 = ", 32132*42452)
	fmt.Println((true && true) || (false && true) || !(false && false))

	fmt.Println("Enter a number")
	var input float64

	fmt.Scanf("%f", &input)

	output := input * 2

	fmt.Println(output)
}
