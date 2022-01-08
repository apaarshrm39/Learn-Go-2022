package main

import "fmt"

type bot interface {
	getGreeting() string
}

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

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

/*
func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}
*/

/* Redundant that's why use interface
func printGreeting(sb spanishBot) {
	fmt.Println(sb.getGreeting())
}
*/
