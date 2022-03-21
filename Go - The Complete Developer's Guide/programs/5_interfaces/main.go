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

// because both spanishBot and englishBot are both members of type 'bot', this function will work with either
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// because we have a getGreeting() function, the englishBot is an honorary member of type 'bot'
func (eb englishBot) getGreeting() string {
	return "Hi there!"
}

// because we have a getGreeting() function, the spanishBot is an honorary member of type 'bot'
func (sb spanishBot) getGreeting() string {
	return "Hola amigo!"
}
