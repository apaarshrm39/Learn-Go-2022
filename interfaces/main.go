package main

import "fmt"

type englishBot struct {
}

type spanishBot struct {
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func (eb englishBot) getGreeting() string { // equivalent func (englishBot) getGreeting() (string){ If not using eb
	// Very custom logic
	return "Hello!"
}

func (sb spanishBot) getGreeting() string {
	return "Hola"
}

func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

func printGreeting(sb spanishBot) {
	fmt.Println(sb.getGreeting())
}