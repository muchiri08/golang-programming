package main

import "fmt"

func main() {
	//two dimension array declaration
	twoD := [4][4]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}

	fmt.Println("Two Dimension Arrays in GO")

	for i := 0; i < len(twoD); i++ {
		row := twoD[i]
		for j := 0; j < len(row); j++ {
			result := fmt.Sprintf("[%d, %d] = %d", i, j, twoD[i][j])
			fmt.Println(result)
		}
	}

}
