package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	// Assuming that both getGreeting will have very different logic
	return "Hi There!"
}
func (spanishBot) getGreeting() string {
	// Assuming that both getGreeting will have very different logic
	return "Hola!"
}
