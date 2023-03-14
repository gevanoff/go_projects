package main

func main() {
	// var card string = "Ace of Spades" // Explicit string variable instantiation
	// card := "Ace of Diamonds // := only used for new variables. Implicit for string
	// card := newCard() // infers type string because of newCard function
	// cards := []string{"Ace of Diamonds", newCard()} // instantiate slice of strings
	// cards := deck{"Ace of Diamonds", newCard()} // instantiate slice of strings

	// cards = append(cards, "Six of Spades") //returns new slice to cards made of cards plus "Six of Spades"
	//for i, card := range cards {
	//	fmt.Println(i, card)
	//}

	cards := newDeck()
	cards.shuffle()
	// hand, remainingCards := deal(cards, 5) // Instantiate two variables of type "deck" with values returned by "deal"

	//cards.print()
	//hand.print()
	//remainingCards.print()
	//fmt.Println(cards.toString())
	//cards.saveToFile("my_cards")
	//cards := newDeckFromFile("my_cards")
	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
