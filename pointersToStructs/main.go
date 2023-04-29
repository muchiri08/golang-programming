package main

import "fmt"

type aStructure struct {
	Name    string
	Surname string
	Height  int32
}

func createStructPointer(n, sn string, h int32) *aStructure {
	if h > 300 {
		h = 0
	}
	return &aStructure{n, sn, h}
}

func createStruct(n, sn string, h int32) aStructure {
	if h > 300 {
		h = 0
	}

	return aStructure{n, sn, h}
}

func main() {
	s1 := createStructPointer("Mihalis", "Tsoukalos", 123)
	s2 := createStruct("Mihalis", "Tsoukalos", 123)
	fmt.Println((*s1).Name)
	fmt.Println(s2.Name)
	fmt.Println(s1)
	fmt.Println(s2)
}
