package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	li, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()

	fmt.Fprintln(li, "Hello World")

}
