package main

import (
	"fmt"
	"os"
	"time"

	"github.com/fatih/color"
)

func main() {
	if len(os.Args) != 2 {
		color.Red("usage: cpu <string>\n")
		os.Exit(1)
	}

	s := os.Args[1]

	for {
		time.Sleep(1 * time.Second)
		fmt.Println(s)
	}
}
