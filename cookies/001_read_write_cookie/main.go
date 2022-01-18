package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", writeCookie)
	http.HandleFunc("/read", read)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func writeCookie(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:  "apaar",
		Value: "sharma",
		Path:  "/",
	}

	http.SetCookie(w, &cookie)
	fmt.Println("cookie written to browser")
}

func read(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("apaar")
	if err != nil {
		http.Error(w, err.Error(), 404)
	}
	io.WriteString(w, cookie.String())
}
