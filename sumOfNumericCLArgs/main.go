package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) == 1 {
		fmt.Println("Please give one or more inputs")
		os.Exit(1)
	}

	var err = errors.New("an error")
	args := os.Args
	k := 1
	var n, sum float64

	for err != nil {
		if k >= len(args) {
			fmt.Println("None of the args is a numeric")
			return
		}
		n, err = strconv.ParseFloat(args[k], 64)
		k++
	}

	for i := 1; i < len(args); i++ {
		n, err = strconv.ParseFloat(args[i], 64)
		if err == nil {
			sum += n
		}
	}

	fmt.Println("Sum is: ", sum)

}
