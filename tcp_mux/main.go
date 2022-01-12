package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Panic(err)
		}

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	request(conn)
}

func request(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	i := 0
	for scanner.Scan() {
		text := scanner.Text()
		if i == 0 {
			st := strings.Fields(text)
			fmt.Fprintln(os.Stdout, "********URI:", st[1])
			respond(conn, st[1])
		}
		if text == "" {
			break
		}
		i++
	}
}

func respond(conn net.Conn, uri string) {
	body := `<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Document</title>
	</head>
	<body>
		<p> Hello World</p>
	</body>
	</html>`

	something := `<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Document</title>
	</head>
	<body>
		<p> something elsed</p>
	</body>
	</html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")

	if uri == "/" {
		fmt.Fprint(conn, body)
	} else if uri == "/something" {
		fmt.Fprint(conn, something)
	} else {
		fmt.Fprint(conn, "404 Not found")
	}
}
