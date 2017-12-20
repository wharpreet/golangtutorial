package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

}

func (englishBot) getGreeting() string {
	// Assuming that both getGreeting will have very different logic
	return "Hi There!"
}
func (spanishBot) getGreeting() string {
	// Assuming that both getGreeting will have very different logic
	return "Hola!"
}

func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting)
}

func printGreeting(sb spanishBot) {
	fmt.Println(sb.getGreeting)
}
