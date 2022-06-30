package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[1:], " ")) // Using 'Join' from the strings package to join together our strings
}
