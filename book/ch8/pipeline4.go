package main

import (
	"fmt"
)

// Generator for count
func count(limit int) <-chan int {
	naturals := make(chan int)
	go func() {
		for x := 0; x < limit; x++ {
			naturals <- x
		}
		close(naturals)
	}()
	return naturals
}

// Generator for square
func square(in <-chan int) <-chan int {
	squares := make(chan int)
	go func() {
		for v := range in {
			squares <- v * v
		}
		close(squares)
	}()
	return squares
}

func print(in <-chan int) {
	for v := range in {
		fmt.Printf("\r%d", v)
	}
}

func main() {
	naturals := count(1000000)
	squares := square(naturals)
	print(squares)
}
