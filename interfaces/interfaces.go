package main

import "fmt"

// Don't need to specify that the structs implement this because of duck typing
type bot interface {
	getGreeting() string // Basically any type in the file that implement this function, is of type bot (duck typing)
}

type englishBot struct{}

type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreetingEnglish(eb)
	printGreetingSpanish(sb)

	printGreeting(eb)
	printGreeting(sb)
}

func printGreetingEnglish(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

func printGreetingSpanish(sb spanishBot) {
	fmt.Println(sb.getGreeting())
}

// Receiver function for the english bot struct
// func (eb englishBot) getGreeting() string {
// No var name needed for receiver if it is not used
func (englishBot) getGreeting() string {
	return "Hello!"
}


func (spanishBot) getGreeting() string {
	return "Hola!"
}

// Using Interfaces

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
