package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

//Now we have a new custom type called bot here
type bot interface {
	getGreeting() string
	//we said anyone with this fn is also now type of bot
}

func main() {
	eb := englishBot{} //struct
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

//Now you are of bot type you can call this below fn as well
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
	//I only accept bot type
}

func (englishBot) getGreeting() string {
	// very custom logic
	return "Hi There"
}

func (spanishBot) getGreeting() string {
	// custom logic
	return "Hola"
}
