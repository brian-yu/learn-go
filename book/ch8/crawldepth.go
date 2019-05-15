package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/brian-yu/learn/book/ch8/links"
)

var (
	tokens   = make(chan struct{}, 20)
	maxDepth = 0
	seen     = make(map[string]bool)
	seenLock = sync.RWMutex
)

func crawl(url string, depth int) []string {
	fmt.Println(url)

	tokens <- struct{}{} // acquire a token
	list, err := links.Extract(url)
	<-tokens // release token
	if err != nil {
		log.Print(err)
	}

	return list
}

func main() {

	var maxDepth int
	flag.IntVar(&maxDepth, "depth", 3, "depth to crawl")
	flag.IntVar(&maxDepth, "d", 3, "depth to crawl (shorthand)")
	flag.Parse()

	worklist := make(chan []string)
	unseenLinks := make(chan string)

	var n int

	n++
	go func() { worklist <- os.Args[1:] }()

	wg := sync.WaitGroup

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for link := range unseenLinks {
				foundLinks := crawl(link, 1)
				// go func() { worklist <- foundLinks }()
			}
		}()
	}

	wg.Wait()
}
