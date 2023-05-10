package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func lineByLine(filename string) error {
	var err error

	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("error reading file %s\n", err)
			break
		}
		fmt.Print(line)
	}
	return nil
}

func main() {
	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Printf("usage: byLine <file1> [<file2> ...]\n")
		return
	}

	for _, file := range flag.Args() {
		err := lineByLine(file)
		if err != nil {
			fmt.Println(err)
		}
	}
}
