package main

import "fmt"

type color map[string]string

func main() {
	colors := map[string]string{
		"red":    "#ff000",
		"yellow": "#dd000",
	}

	color2 := color{
		"hi": "there",
	}

	//var colors map[string]string

	//colors := make(map[string]string)

	colors["white"] = "ffffff"

	// delete keys values in maps

	//delete(colors, "white")
	printMap(colors)
	color2.print()
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Printf("the hex for color %v is %v \n", color, hex)
	}
}

func (c color) print() {
	for key, value := range c {
		fmt.Printf("the hex for color %v is %v \n", key, value)
	}
}
