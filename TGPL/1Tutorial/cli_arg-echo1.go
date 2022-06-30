package main

import (
	"fmt"
	"os"
) 

func main() {
	var s, sep string 			// Declaring string variables 's', 'sep'

	for i := 1; i < len(os.Args); i++ { 	// Loop through all CLI arguments minus the initial

		s += sep + os.Args[i] 		// Store results in 's'
		sep = " " 			// Spacing after the first argument is stored
	}

		fmt.Println(s) 
}
