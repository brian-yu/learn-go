package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	conns := make(map[string]io.Reader, len(args))
	for _, arg := range args {
		arr := strings.Split(arg, "=")
		city := arr[0]
		addr := arr[1]
		conn := connect(city, addr)
		defer conn.Close()
		conns[city] = conn
	}

	cities := make([]string, len(conns))
	i := 0
	for city := range conns {
		cities[i] = city
		i++
	}

	for {
		for _, city := range cities {
			time := mustReadTime(conns[city])
			fmt.Printf("%s: %s\t", city, time)
		}
		fmt.Println()
	}
}

func connect(city string, addr string) net.Conn {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	return conn
}

func mustReadTime(src io.Reader) string {
	var time string
	_, err := fmt.Fscanln(src, &time)
	if err != nil {
		log.Fatal(err)
	}
	return time
}
