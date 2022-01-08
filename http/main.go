package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type mywriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Some error:", err)
		os.Exit(1)
	}
	/* Condensing this code below
	byteS := make([]byte, 99999) //make big Slice because read function can't resize slice
	resp.Body.Read(byteS)
	fmt.Println(string(byteS))
	*/

	//io.Copy(os.Stdout, resp.Body)
	m := mywriter{}
	io.Copy(m, resp.Body)
}

func (m mywriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("just worte this many bytes:", len(bs))
	return len(bs), nil
}
