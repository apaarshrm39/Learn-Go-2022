package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer li.Close()
	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatal(err)
		}
		io.WriteString(conn, "\n Hello from TCp Server")
		fmt.Fprintln(conn, "How is your Day")
		fmt.Fprintf(conn, "%v", "Well I Hope!")
		conn.Close()
	}
}
