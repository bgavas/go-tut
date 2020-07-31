package main

import "fmt"

func main() {

	deck := newDeck()
	deck.print()

	deck.shuffle()
	hand, remainingCards := deal(deck, 5)
	hand.print()
	remainingCards.print()

	fmt.Println("Error:", hand.saveToFile("hand"))

	fmt.Println("Hand:", newDeckFromFile("hand"))

}
