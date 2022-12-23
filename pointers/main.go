package main

import "fmt"

func main() {
	x := 1
	fmt.Println(incr(&x))
}

func incr(p *int) int {
	*p++
	return *p
}
