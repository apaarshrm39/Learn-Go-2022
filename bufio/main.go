package main

import (
	"bufio"
	"fmt"
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
		go handle(conn)
	}
}

// It reads each request
func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		fmt.Fprintf(conn, "I heard you say %s\n,ln", ln)
	}
	defer conn.Close()

	// WE NEVER GET HERE
	// WE HAVE AN OPEN STREAM CONNECTION
	// HOW DOES THE THE ABOVE READER KNOW WHEN IT'S DONE
	fmt.Println("code got here")
}
