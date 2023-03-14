package main

import "fmt"

/*
Interfaces are not generic types. They are also implicit, so it can be harder to track down which other
variables an interface applies to. Interfaces are to help reuse code and form a sort of contract between
different types and different functions.
Writing interfaces is tough. It takes experience.
*/

type bot interface { // define a type called bot which is an interface
	getGreeting() string // If another type with function getGreeting() which returns a string
} // ...exists in the program, it is an "honorary" member of type bot.
// ...and can now use function printGreeting
// The functions need to exactly match the input and return types, number, and order for this to apply.

// Concrete types like int, string, struct, etc, can be used to create a value
type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) { // eb and sb are of type bot
	fmt.Println(b.getGreeting())
}

func (eb englishBot) getGreeting() string {
	// Very custom logic for generating an English greeting
	return "Hi There!"
}

func (sb spanishBot) getGreeting() string {
	// Very custom logic for generating a Spanish greeting
	return "Hola!"
}
