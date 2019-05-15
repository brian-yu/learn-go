package main

import (
	"log"
	"net"
)

func main() {
	for {
		go func() {
			conn, err := net.Dial("tcp", "localhost:8000")
			if err != nil {
				log.Fatal(err)
			}
			conn.Close()
		}()
	}
}
