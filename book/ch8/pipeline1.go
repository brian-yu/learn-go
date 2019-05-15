package main

import (
	"fmt"
)

func square(lo int, hi int, out chan<- int, complete chan<- struct{}) {
	for x := lo; x < hi; x++ {
		s := x * x
		out <- s
	}
	complete <- struct{}{}
}

func print(in <-chan int) {
	for v := range in {
		fmt.Printf("\r%d", v)
	}
}

func main() {

	squares := make(chan int)
	complete := make(chan struct{})

	NUM := 1000000
	NUM_WORKERS := 20
	WORKER_PORTION := NUM / NUM_WORKERS

	for x := 0; x < NUM; x += WORKER_PORTION {
		go square(x, x+WORKER_PORTION, squares, complete)
	}

	go print(squares)

	for i := 0; i < NUM_WORKERS; i++ {
		<-complete
	}
}
