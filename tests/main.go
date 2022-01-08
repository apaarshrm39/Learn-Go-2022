package main

import (
	"errors"
	"log"
	"os"
)

func main() {
	result, err := divide(100.00, 0)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	log.Println(result)
}

func divide(x, y float32) (float32, error) {
	var result float32
	if y == 0 {
		return result, errors.New("cannot devide by zero")
	}
	result = x / y

	return result, nil
}
