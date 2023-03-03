package main

import "fmt"

func main() {
	a, b := 12, 345
	c, d := 1.2, 3.45

	fmt.Printf("%d %d \n", a, b)
	fmt.Printf("%X %X \n", a, b)   //hexadecimal
	fmt.Printf("%#x %#x \n", a, b) //starts with 0x then the hexadecimal number

	fmt.Println()

	fmt.Printf("%.2f %.2f", c, d)

	fmt.Println()
	//Literal vertical bars
	//print in column that is width of 6 characters
	fmt.Printf("|%9d|%9d|\n", a, b)
	//Does the same thing as above but with leading zeros
	fmt.Printf("|%09d|%09d|\n", a, b)
	//left justifying
	fmt.Printf("|%-9d|%-9d|\n", a, b)

	fmt.Println()

	fmt.Printf("|%9f| %9.2f|", c, d)

	fmt.Println()
	fmt.Println()

	//formatting slices
	s := []int{1, 2, 3}

	fmt.Printf("%T \n", s)  //shows type
	fmt.Printf("%v \n", s)  //shoes values
	fmt.Printf("%#v \n", s) //shows []int{1, 2, 3}

	fmt.Println()

	//arrays
	ar := [3]rune{'a', 'b', 'c'}

	fmt.Printf("%T \n", ar)  //shows type -> [3]int32
	fmt.Printf("%v \n", ar)  //shoes integer values of the chars [97 98 99]
	fmt.Printf("%#v \n", ar) // type and values -> [3]int32{97, 98, 99}
	fmt.Printf("%q\n", ar)   //shows the chars themselves -> ['a' 'b' 'c']

}
