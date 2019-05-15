package main

import (
	"fmt"
)

func main() {

	for x := 0; x < 1000000; x++ {
		sqrt := x * x
		fmt.Printf("\r%d", sqrt)
	}
}
