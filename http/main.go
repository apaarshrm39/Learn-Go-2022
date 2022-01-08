package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Some error:", err)
		os.Exit(1)
	}
	byteS := make([]byte, 99999) //make big Slice because read function can't resize slice
	resp.Body.Read(byteS)
	fmt.Println(string(byteS))
}
