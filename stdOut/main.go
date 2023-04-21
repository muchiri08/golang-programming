package main

import (
	"io"
	"os"
)

func main() {
	//For STDOUT
	//myString := ""
	//arguments := os.Args
	//
	//if len(arguments) == 1 {
	//	myString = "Please give me one argument!"
	//} else {
	//	myString = arguments[1]
	//}
	//
	//io.WriteString(os.Stdout, myString)
	//io.WriteString(os.Stdout, "\n")

	//For STDIN
	//var f *os.File
	//f = os.Stdin
	//defer f.Close()
	//
	//scanner := bufio.NewScanner(f)
	//for scanner.Scan() {
	//	fmt.Println(">", scanner.Text())
	//}

	//For STDERR
	myString := ""
	arguments := os.Args

	if len(arguments) == 1 {
		myString = "Please give me one argument!"
	} else {
		myString = arguments[1]
	}

	io.WriteString(os.Stdout, "This is a standard output\n")
	io.WriteString(os.Stderr, myString)
	io.WriteString(os.Stderr, "\n")
}
