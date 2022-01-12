package main

import (
	"fmt"
	"log"
	"net"
)

func main() {

	client, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Panic(err)
	}

	/* For Reader Client
	bs, err := ioutil.ReadAll(client)
	if err != nil {
		log.Panic(err)
	}
	*/

	fmt.Fprintln(client, "I dialed yoyu")
}
