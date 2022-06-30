package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", "" 			// Initializing and declaring variables using ':=' instead of explicit declaration of type
	for _, arg := range os.Args[1:] { 	// '_' is a blank identifier, 'range' produces a pair of values: the index and the value of the element at that index
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
