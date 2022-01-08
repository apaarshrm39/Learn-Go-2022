package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file := os.Args[1]
	read(file)
}

func read(file string) {
	content, err := os.Open(file)
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
	io.Copy(os.Stdout, content)
}
