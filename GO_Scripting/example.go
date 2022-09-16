package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

func main() {
	s := "world"

	if len(os.Args) > 1 {
		s = os.Args[1]
	}

	color.Red("We have red")
	fmt.Printf("Hello, %v!", s)
	fmt.Println("")

	if s == "fail" {
		os.Exit(30)
	}
}
