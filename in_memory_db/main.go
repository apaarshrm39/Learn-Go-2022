package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
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

func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)

	io.WriteString(conn, "\r\nIN-MEMORY DATABASE\r\n\r\n"+
		"USE:\r\n"+
		"\tSET key value \r\n"+
		"\tGET key \r\n"+
		"\tDEL key \r\n\r\n"+
		"EXAMPLE:\r\n"+
		"\tSET fav chocolate \r\n"+
		"\tGET fav \r\n\r\n\r\n")

	data := make(map[string]string)
	for scanner.Scan() {
		text := scanner.Text()
		fs := strings.Fields(text)
		if len(fs) < 1 {
			continue
		}

		switch fs[0] {
		case "GET":
			q := fs[1]
			fmt.Fprintf(conn, "value stored in %s is %s", fs[1], data[q])
		case "PUT":
			q := fs[1]
			v := fs[2]
			data[q] = v
		case "DELETE":
			q := fs[1]
			delete(data, q)
		default:
			fmt.Fprint(conn, "invalid")
		}

	}

}
