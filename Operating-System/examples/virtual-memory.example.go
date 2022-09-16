package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	var p *int

	i := 0
	p = &i

	if p == nil {
		panic("Bag Address")
	}

	fmt.Printf("(%d) address of p: %08x\n", os.Getpid(), p)

	for {
		time.Sleep(1 * time.Second)
		*p = *p + 1
		fmt.Printf("(%d) p: %d\n", os.Getpid(), *p)
	}
}
