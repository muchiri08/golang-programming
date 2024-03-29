package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
)

var readValue = make(chan int) //for reading random numbers
var writeData = make(chan int) // for getting new random numbers

func set(newValue int) {
	writeData <- newValue
}

func read() int {
	return <-readValue
}

func monitor() {
	var value int
	for {
		select {
		case newValue := <-writeData:
			value = newValue
			fmt.Printf("%d ", value)
		case readValue <- value:
		}
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please give an integer")
		return
	}
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Going to create %d random numbers.\n\n", n)
	rand.Seed(time.Now().Unix())
	go monitor()

	var w sync.WaitGroup

	for r := 0; r < n; r++ {
		w.Add(1)
		go func() {
			defer w.Done()
			set(rand.Intn(10 * n))
		}()
	}

	w.Wait()
	fmt.Printf("\nLast value: %d\n", read())
}
