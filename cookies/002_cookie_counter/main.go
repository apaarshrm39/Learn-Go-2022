package main

import (
	"io"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

var counter = make(map[string]int)

func main() {
	counter = make(map[string]int)
	http.HandleFunc("/", cook)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func cook(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("UUID")
	if err != nil {
		u := generateRandomNumber()
		http.SetCookie(w, &http.Cookie{
			Name:  "UUID",
			Value: strconv.Itoa(u),
		})

		counter[strconv.Itoa(u)] = 1
		return
	}

	counter[cookie.Value] = counter[cookie.Value] + 1

	//value := counter[cookie.Value]

	io.WriteString(w, strconv.Itoa(counter[cookie.Value]))

}

func generateRandomNumber() int {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	number := r.Intn(50)
	return number
}
