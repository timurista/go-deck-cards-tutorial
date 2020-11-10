package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

type bot interface {
	getGreeting() string
}

func main() {
	eb := englishBot{}
	sp := spanishBot{}
	printGreeting(eb)
	printGreeting(sp)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	return "Hi Custom Logic"
}

func (spanishBot) getGreeting() string {
	return "Hola, es mi logica custuma"
}
