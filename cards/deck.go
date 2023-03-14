package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of deck which is a slice of strings
type deck []string

func newDeck() deck { // Function newDeck returns value of type "deck"
	cards := deck{} // instantiates cards variable of type deck
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits { // index variable "i" replaced with underscore
		for _, value := range cardValues { // same as above for var "j"
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// Receiver function. By convention, receiver variable is one or two letters, matching type
func (d deck) print() { // Copy of deck in variable "d". All type "deck" get access to print() method
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) { // Call "deal" function with var type "deck" and var type "int"
	return d[:handSize], d[handSize:] // Return two values of type "deck"
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0664)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		// Option #1 - log error and return a call to newDeck()
		// Option #2 - log the error and exit program
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano()) // Create seed from UnixNano int64
	r := rand.New(source)                           // Use source object from above to create new rand
	for i := range d {                              // "i" is index
		newPosition := r.Intn(len(d) - 1)           // r is rand seeded with source
		d[i], d[newPosition] = d[newPosition], d[i] // Swap whatever is at index "i" with newPosition in slice
	}
}
