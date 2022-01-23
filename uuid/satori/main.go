package main

import (
	"fmt"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

func main() {
	http.HandleFunc("/", foo)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session")
	if err != nil {
		id := uuid.NewV4()
		cookie := &http.Cookie{
			Name:  "session",
			Value: id.String(),
			// secure : true
			HttpOnly: true,
		}
		http.SetCookie(w, cookie)
	}

	fmt.Println(cookie)
}
